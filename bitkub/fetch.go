package bitkub

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog"
)

var (
	apiBitkub     *http.Client
	apiBitkubTime *http.Client
	httpTransport *http.Transport
)

func init() {
	initHTTPClients()
}

func initHTTPClients() {
	cfg := GetConfig()

	httpTransport = &http.Transport{
		MaxIdleConns:        cfg.MaxIdleConns,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     cfg.IdleConnTimeout,
		DisableKeepAlives:   false,
	}

	apiBitkub = &http.Client{
		Transport: httpTransport,
		Timeout:   cfg.APITimeout,
	}

	apiBitkubTime = &http.Client{
		Transport: httpTransport,
		Timeout:   cfg.ServerTimeout,
	}
}

func generateSignature(secret string, payload string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(payload))
	return hex.EncodeToString(h.Sum(nil))
}

type ResponseSecure interface {
	GetError() int
}

func statusError(resp *http.Response, body string) error {
	message, exists := ErrorCode[resp.StatusCode]
	if !exists {
		message = http.StatusText(resp.StatusCode)
	}
	if body != "" {
		return fmt.Errorf("HTTP %d: %s - %s", resp.StatusCode, message, body)
	}
	return fmt.Errorf("HTTP %d: %s", resp.StatusCode, message)
}

func FetchSecure(method string, path string, reqBody any, resPayload any) error {
	return FetchSecureWithContext(context.Background(), method, path, reqBody, resPayload)
}

func FetchSecureWithContext(ctx context.Context, method string, path string, reqBody any, resPayload any) error {
	resp, err := fetchWithContext(ctx, true, method, path, reqBody)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return statusError(resp, "")
	}

	if err = json.NewDecoder(resp.Body).Decode(resPayload); err != nil {
		return fmt.Errorf("decoding response: %w", err)
	}

	if res, ok := resPayload.(*ResponseAPI); ok {
		return res.CheckResponseError()
	}

	return nil
}

func FetchSecureV4(method string, path string, reqBody any, resPayload any) error {
	return FetchSecureV4WithContext(context.Background(), method, path, reqBody, resPayload)
}

func FetchSecureV4WithContext(ctx context.Context, method string, path string, reqBody any, resPayload any) error {
	body, err := fetchBody(ctx, method, path, reqBody)
	if err != nil {
		return err
	}

	return decodeV4(body, resPayload)
}

func fetchBody(ctx context.Context, method string, path string, reqBody any) ([]byte, error) {
	resp, err := fetchWithContext(ctx, true, method, path, reqBody)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}

	if resp.StatusCode > 299 {
		return nil, statusError(resp, string(body))
	}

	return body, nil
}

func decodeV4(body []byte, resPayload any) error {
	if err := json.Unmarshal(body, resPayload); err != nil {
		return fmt.Errorf("decoding response: %w", err)
	}

	res, ok := resPayload.(*ResponseAPIV4)
	if !ok || res.Code == "0" {
		return nil
	}

	if errDesc, exists := ErrorCodeV4[res.Code]; exists {
		return fmt.Errorf("[%s] %s: %s", res.Code, errDesc, res.Message)
	}
	return fmt.Errorf("[%s] %s", res.Code, res.Message)
}

func FetchNonSecure(method string, path string, reqBody any, resPayload any) error {
	return FetchNonSecureWithContext(context.Background(), method, path, reqBody, resPayload)
}

func FetchNonSecureWithContext(ctx context.Context, method string, path string, reqBody any, resPayload any) error {
	resp, err := fetchWithContext(ctx, false, method, path, reqBody)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return statusError(resp, "")
	}

	if err = json.NewDecoder(resp.Body).Decode(resPayload); err != nil {
		return fmt.Errorf("decoding response: %w", err)
	}

	if res, ok := resPayload.(ResponseSecure); ok && res.GetError() != 0 {
		errMsg, exists := ErrorCode[res.GetError()]
		if !exists {
			errMsg = "Unknown error"
		}
		return fmt.Errorf("[error %d] %s", res.GetError(), errMsg)
	}

	return nil
}

var ErrMissingCredentials = errors.New("missing API credentials: set BTK_APIKEY and BTK_SECRET or pass --key and --secret")

func fetchWithContext(ctx context.Context, secure bool, method string, path string, reqBody any) (*http.Response, error) {
	serverTime := ""
	if secure {
		if key, secret := credentials(); key == "" || secret == "" {
			return nil, ErrMissingCredentials
		}
		var err error
		serverTime, err = GetServerTimeWithContext(ctx)
		if err != nil {
			return nil, fmt.Errorf("server time: %w", err)
		}
	}

	var payload []byte
	var body io.Reader
	if reqBody != nil {
		var err error
		payload, err = json.Marshal(reqBody)
		if err != nil {
			return nil, fmt.Errorf("marshaling json: %w", err)
		}
		body = bytes.NewReader(payload)
	}

	req, err := http.NewRequestWithContext(ctx, method, BASE_URL+path, body)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if secure {
		key, secret := credentials()
		signaturePayload := serverTime + req.Method + req.URL.Path
		if req.URL.RawQuery != "" {
			signaturePayload += "?" + req.URL.RawQuery
		}
		signaturePayload += string(payload)

		req.Header.Set("X-BTK-TIMESTAMP", serverTime)
		req.Header.Set("X-BTK-APIKEY", key)
		req.Header.Set("X-BTK-SIGN", generateSignature(secret, signaturePayload))
	}

	resp, err := apiBitkub.Do(req)
	if err != nil {
		logger.Warn().Str("method", method).Str("path", path).Err(err).Stack().Send()
		return nil, fmt.Errorf("making request: %w", err)
	}

	logger.Debug().Int("status", resp.StatusCode).Str("method", method).Str("path", path).Send()
	return resp, nil
}

var logger = zerolog.Nop()

func SetLogger(l zerolog.Logger) {
	logger = l
}
