package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/erfjab/egobot/models"
)

const (
	baseURL        = "https://api.telegram.org/bot"
	defaultTimeout = 30 * time.Second
)

// inputFileType is the reflect.Type for models.InputFile, used for fast type checks.
var inputFileType = reflect.TypeOf(models.InputFile{})

type Requester struct {
	Token      string
	HTTPClient *http.Client
}

func NewRequester(token string) *Requester {
	return &Requester{
		Token: token,
		HTTPClient: &http.Client{
			Timeout: defaultTimeout,
		},
	}
}

// Request sends a Telegram Bot API call. It automatically switches to
// multipart/form-data when params contains an InputFile with raw Data bytes;
// otherwise it encodes the body as JSON.
func (r *Requester) Request(method string, params interface{}) ([]byte, error) {
	if params != nil && hasUploads(params) {
		return r.requestMultipart(method, params)
	}
	return r.requestJSON(method, params)
}

// requestJSON encodes params as JSON and POSTs it to the Telegram API.
func (r *Requester) requestJSON(method string, params interface{}) ([]byte, error) {
	url := fmt.Sprintf("%s%s/%s", baseURL, r.Token, method)

	var body []byte
	var err error

	if params != nil {
		body, err = json.Marshal(params)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal params: %w", err)
		}
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	return r.do(req)
}

// requestMultipart builds a multipart/form-data body from params and POSTs it.
// Fields that are InputFile with Data are written as file parts; everything
// else is written as plain form fields (complex types are JSON-encoded).
func (r *Requester) requestMultipart(method string, params interface{}) ([]byte, error) {
	url := fmt.Sprintf("%s%s/%s", baseURL, r.Token, method)

	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)

	if err := writeMultipartFields(w, params); err != nil {
		return nil, fmt.Errorf("failed to build multipart form: %w", err)
	}
	if err := w.Close(); err != nil {
		return nil, fmt.Errorf("failed to close multipart writer: %w", err)
	}

	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	return r.do(req)
}

// do executes the HTTP request and returns the raw response body.
func (r *Requester) do(req *http.Request) ([]byte, error) {
	resp, err := r.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("telegram API error: %s", string(respBody))
	}

	return respBody, nil
}

func (r *Requester) ParseResponse(respBody []byte, target interface{}) error {
	var apiResp struct {
		Ok          bool            `json:"ok"`
		Result      json.RawMessage `json:"result,omitempty"`
		ErrorCode   int             `json:"error_code,omitempty"`
		Description string          `json:"description,omitempty"`
	}

	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if !apiResp.Ok {
		return fmt.Errorf("telegram API error [%d]: %s", apiResp.ErrorCode, apiResp.Description)
	}

	if target != nil && len(apiResp.Result) > 0 {
		if err := json.Unmarshal(apiResp.Result, target); err != nil {
			return fmt.Errorf("failed to unmarshal result: %w", err)
		}
	}

	return nil
}

// hasUploads reports whether params (a struct or pointer to struct) contains at
// least one InputFile field whose Data slice is non-empty.
func hasUploads(params interface{}) bool {
	v := reflect.ValueOf(params)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return false
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return false
	}
	for i := 0; i < v.NumField(); i++ {
		if f, ok := resolveInputFile(v.Field(i)); ok && len(f.Data) > 0 {
			return true
		}
	}
	return false
}

// writeMultipartFields writes every non-zero struct field of params into w.
// InputFile fields with Data are written as file parts (CreateFormFile).
// All other fields are written as plain text form fields; complex types
// (slices, maps, nested structs) are JSON-encoded first.
func writeMultipartFields(w *multipart.Writer, params interface{}) error {
	v := reflect.ValueOf(params)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("params must be a struct, got %s", v.Kind())
	}
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// Resolve the json field name and omitempty flag.
		tag := field.Tag.Get("json")
		if tag == "" || tag == "-" {
			continue
		}
		parts := strings.SplitN(tag, ",", 2)
		name := parts[0]
		omitempty := len(parts) == 2 && strings.Contains(parts[1], "omitempty")

		if omitempty && value.IsZero() {
			continue
		}

		// Handle InputFile (raw upload).
		if f, ok := resolveInputFile(value); ok {
			if len(f.Data) > 0 {
				filename := f.Name
				if filename == "" {
					filename = name
				}
				part, err := w.CreateFormFile(name, filename)
				if err != nil {
					return fmt.Errorf("multipart: CreateFormFile %q: %w", name, err)
				}
				if _, err := part.Write(f.Data); err != nil {
					return fmt.Errorf("multipart: write file %q: %w", name, err)
				}
			} else if f.FileID != "" {
				if err := w.WriteField(name, f.FileID); err != nil {
					return err
				}
			} else if f.URL != "" {
				if err := w.WriteField(name, f.URL); err != nil {
					return err
				}
			}
			continue
		}

		// For interface{} fields that haven't been resolved as InputFile above,
		// unwrap the interface before deciding how to serialise.
		actual := value
		if actual.Kind() == reflect.Interface {
			if actual.IsNil() {
				continue
			}
			actual = actual.Elem()
		}
		if actual.Kind() == reflect.Ptr {
			if actual.IsNil() {
				continue
			}
			actual = actual.Elem()
		}

		str, err := scalarToString(actual)
		if err != nil {
			return fmt.Errorf("multipart: field %q: %w", name, err)
		}
		if str == "" && omitempty {
			continue
		}
		if err := w.WriteField(name, str); err != nil {
			return err
		}
	}
	return nil
}

// resolveInputFile attempts to extract an models.InputFile from v, handling
// pointer and interface indirections.
func resolveInputFile(v reflect.Value) (models.InputFile, bool) {
	// Unwrap interface.
	if v.Kind() == reflect.Interface {
		if v.IsNil() {
			return models.InputFile{}, false
		}
		v = v.Elem()
	}
	// Unwrap pointer.
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return models.InputFile{}, false
		}
		v = v.Elem()
	}
	if v.Type() == inputFileType {
		return v.Interface().(models.InputFile), true
	}
	return models.InputFile{}, false
}

// scalarToString converts a reflect.Value to its string representation.
// Primitive kinds are formatted directly; everything else is JSON-encoded.
func scalarToString(v reflect.Value) (string, error) {
	switch v.Kind() {
	case reflect.String:
		return v.String(), nil
	case reflect.Bool:
		return strconv.FormatBool(v.Bool()), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10), nil
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64), nil
	default:
		data, err := json.Marshal(v.Interface())
		if err != nil {
			return "", err
		}
		return string(data), nil
	}
}
