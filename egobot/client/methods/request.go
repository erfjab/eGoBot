package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	baseURL        = "https://api.telegram.org/bot"
	defaultTimeout = 30 * time.Second
)

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

func (r *Requester) Request(method string, params interface{}) ([]byte, error) {
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
