package core

import (
	"encoding"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/erfjab/egobot/models"
)

var (
	ErrCallbackDataKeyCount          = errors.New("callback data key/value count mismatch")
	ErrCallbackDataInvalidSeparator  = errors.New("callback data separator cannot be empty")
	ErrCallbackDataContainsSeparator = errors.New("callback data value contains separator")
	ErrCallbackDataPrefixInvalid     = errors.New("callback data prefix contains separator")
	ErrCallbackDataModelInvalid      = errors.New("callback data model must be a struct or pointer to struct")
	ErrCallbackDataMarkerMissing     = errors.New("callback data marker not found (embed tools.CallbackData)")
)

const callbackDataContextKey = "callback_data"

const (
	callbackMarkerPkgPath = "github.com/erfjab/egobot/tools"
	callbackMarkerName    = "CallbackData"
)

// CallbackData helps build and parse callback data strings.
// Format: prefix + separator + values
// Example: NewCallbackData("order", "action", "id").Pack("view", "42") => "order:view:42"
type CallbackData struct {
	Prefix    string
	Separator string
	Keys      []string
}

// NewCallbackData creates a new CallbackData instance with default separator ":".
func NewCallbackData(prefix string, keys ...string) *CallbackData {
	return &CallbackData{
		Prefix:    prefix,
		Separator: ":",
		Keys:      keys,
	}
}

// WithSeparator sets a custom separator (ignored if empty).
func (c *CallbackData) WithSeparator(separator string) *CallbackData {
	if separator != "" {
		c.Separator = separator
	}
	return c
}

// Pack builds callback data from values in the same order as Keys.
func (c *CallbackData) Pack(values ...string) (string, error) {
	if c.Separator == "" {
		return "", ErrCallbackDataInvalidSeparator
	}
	if strings.Contains(c.Prefix, c.Separator) {
		return "", ErrCallbackDataPrefixInvalid
	}
	if len(values) != len(c.Keys) {
		return "", ErrCallbackDataKeyCount
	}
	for _, value := range values {
		if strings.Contains(value, c.Separator) {
			return "", fmt.Errorf("%w: %q", ErrCallbackDataContainsSeparator, value)
		}
	}
	if len(values) == 0 {
		return c.Prefix, nil
	}
	return c.Prefix + c.Separator + strings.Join(values, c.Separator), nil
}

// PackStruct builds callback data from an exported struct's fields.
// The field names must match Keys in order. Supports string, bool, ints, and types
// implementing fmt.Stringer or encoding.TextMarshaler.
func (c *CallbackData) PackStruct(v interface{}) (string, error) {
	val, ok := structValue(v)
	if !ok {
		return "", ErrCallbackDataKeyCount
	}

	values := make([]string, 0, len(c.Keys))
	for _, key := range c.Keys {
		field := val.FieldByName(key)
		if !field.IsValid() {
			return "", fmt.Errorf("callback data field not found: %s", key)
		}
		if !field.CanInterface() {
			return "", fmt.Errorf("callback data field not exported: %s", key)
		}
		values = append(values, formatCallbackValue(field.Interface()))
	}
	return c.Pack(values...)
}

// Parse parses callback data and returns a map of key/value pairs.
// Returns false if prefix or part count does not match.
func (c *CallbackData) Parse(data string) (map[string]string, bool) {
	if c.Separator == "" {
		return nil, false
	}
	if len(c.Keys) == 0 {
		if data == c.Prefix {
			return map[string]string{}, true
		}
		return nil, false
	}
	prefix := c.Prefix + c.Separator
	if !strings.HasPrefix(data, prefix) {
		return nil, false
	}
	payload := strings.TrimPrefix(data, prefix)
	parts := strings.Split(payload, c.Separator)
	if len(parts) != len(c.Keys) {
		return nil, false
	}
	result := make(map[string]string, len(c.Keys))
	for i, key := range c.Keys {
		result[key] = parts[i]
	}
	return result, true
}

// ParseToStruct parses callback data into an exported struct's fields.
// The field names must match Keys in order. Supports string, bool, ints, and
// types implementing encoding.TextUnmarshaler.
func (c *CallbackData) ParseToStruct(data string, v interface{}) bool {
	val, ok := structTarget(v)
	if !ok {
		return false
	}

	payload, ok := c.Parse(data)
	if !ok {
		return false
	}
	for _, key := range c.Keys {
		field := val.FieldByName(key)
		if !field.IsValid() || !field.CanSet() {
			return false
		}
		if !setCallbackField(field, payload[key]) {
			return false
		}
	}
	return true
}

// Filter creates a filter for callback queries that match this CallbackData.
func (c *CallbackData) Filter() FilterFunc {
	return func(update *models.Update) bool {
		if update.CallbackQuery == nil {
			return false
		}
		_, ok := c.Parse(update.CallbackQuery.Data)
		return ok
	}
}

// NewCallbackDataFromStruct builds a CallbackData definition from a struct model.
// The model must embed tools.CallbackData (anonymous field) to opt in.
// Optional tags on the embedded marker field:
//   - prefix:"admin" to set prefix (default: derived from struct name)
//   - separator:"|" to set separator (default: ":")
//
// Exported fields (excluding embedded fields) are used as callback keys in order.
func NewCallbackDataFromStruct(model interface{}) (*CallbackData, error) {
	t, ok := structTypeOf(model)
	if !ok {
		return nil, ErrCallbackDataModelInvalid
	}

	var (
		hasMarker bool
		prefix    string
		separator string
		keys      []string
	)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Anonymous && field.Type.Name() == callbackMarkerName && field.Type.PkgPath() == callbackMarkerPkgPath {
			hasMarker = true
			if tagPrefix := field.Tag.Get("prefix"); tagPrefix != "" {
				prefix = tagPrefix
			}
			if tagSeparator := field.Tag.Get("separator"); tagSeparator != "" {
				separator = tagSeparator
			}
			continue
		}

		if field.Anonymous {
			continue
		}
		if !field.IsExported() {
			continue
		}
		if field.Tag.Get("callback") == "-" {
			continue
		}

		keys = append(keys, field.Name)
	}

	if !hasMarker {
		return nil, ErrCallbackDataMarkerMissing
	}
	if prefix == "" {
		prefix = defaultCallbackPrefix(t.Name())
	}

	cb := NewCallbackData(prefix, keys...)
	if separator != "" {
		cb.WithSeparator(separator)
	}

	return cb, nil
}

// GetCallbackStruct reads parsed callback payload from Context.
// Use with OnCallbackStruct registrations.
func GetCallbackStruct[T any](ctx *Context) (T, bool) {
	var zero T
	if ctx == nil {
		return zero, false
	}

	if loadCallbackData(&zero, ctx.Get(callbackDataContextKey)) {
		return zero, true
	}

	return zero, false
}

func newCallbackStructValue(model interface{}) (interface{}, error) {
	t, ok := structTypeOf(model)
	if !ok {
		return nil, ErrCallbackDataModelInvalid
	}

	return reflect.New(t).Interface(), nil
}

func defaultCallbackPrefix(typeName string) string {
	name := strings.ToLower(strings.TrimSpace(typeName))
	if name == "" {
		return "callback"
	}

	name = strings.TrimSuffix(name, "callback")
	name = strings.TrimSuffix(name, "cb")
	if name == "" {
		return "callback"
	}

	return name
}

func formatCallbackValue(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case fmt.Stringer:
		return val.String()
	case encoding.TextMarshaler:
		if b, err := val.MarshalText(); err == nil {
			return string(b)
		}
	}
	return fmt.Sprint(v)
}

func setCallbackField(field reflect.Value, raw string) bool {
	if !field.CanSet() {
		return false
	}

	if field.CanAddr() {
		addr := field.Addr()
		if addr.CanInterface() {
			if unmarshaler, ok := addr.Interface().(encoding.TextUnmarshaler); ok {
				return unmarshaler.UnmarshalText([]byte(raw)) == nil
			}
		}
	}

	switch field.Kind() {
	case reflect.String:
		field.SetString(raw)
		return true
	case reflect.Bool:
		val, err := strconv.ParseBool(raw)
		if err != nil {
			return false
		}
		field.SetBool(val)
		return true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			return false
		}
		field.SetInt(val)
		return true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val, err := strconv.ParseUint(raw, 10, 64)
		if err != nil {
			return false
		}
		field.SetUint(val)
		return true
	case reflect.Float32, reflect.Float64:
		val, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			return false
		}
		field.SetFloat(val)
		return true
	default:
		return false
	}
}

func structTypeOf(v interface{}) (reflect.Type, bool) {
	if v == nil {
		return nil, false
	}
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, false
	}
	return t, true
}

func structValue(v interface{}) (reflect.Value, bool) {
	if v == nil {
		return reflect.Value{}, false
	}
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return reflect.Value{}, false
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return reflect.Value{}, false
	}
	return rv, true
}

func structTarget(v interface{}) (reflect.Value, bool) {
	if v == nil {
		return reflect.Value{}, false
	}
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return reflect.Value{}, false
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return reflect.Value{}, false
	}
	return rv, true
}

func loadCallbackData(out interface{}, stored interface{}) bool {
	if out == nil || stored == nil {
		return false
	}

	outVal := reflect.ValueOf(out)
	if outVal.Kind() != reflect.Ptr || outVal.IsNil() {
		return false
	}

	target := outVal.Elem()
	storedVal := reflect.ValueOf(stored)

	if storedVal.Type().AssignableTo(target.Type()) {
		target.Set(storedVal)
		return true
	}

	if storedVal.Kind() == reflect.Ptr && !storedVal.IsNil() && storedVal.Elem().Type().AssignableTo(target.Type()) {
		target.Set(storedVal.Elem())
		return true
	}

	return false
}

func callbackStructPatternMatches(payload interface{}, pattern interface{}) bool {
	pv, ok := structValue(payload)
	if !ok {
		return false
	}
	ev, ok := structValue(pattern)
	if !ok {
		return false
	}

	if pv.Type() != ev.Type() {
		return false
	}

	t := ev.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.IsExported() || f.Anonymous || f.Tag.Get("match") == "-" {
			continue
		}

		expectedField := ev.Field(i)
		if expectedField.IsZero() {
			continue
		}

		if !reflect.DeepEqual(expectedField.Interface(), pv.Field(i).Interface()) {
			return false
		}
	}

	return true
}
