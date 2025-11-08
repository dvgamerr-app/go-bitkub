package bitkub

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/rs/zerolog/log"
)

var (
	apiBitkub     *http.Client
	apiBitkubTime *http.Client
	httpTransport *http.Transport
	bufferPool    = sync.Pool{
		New: func() any {
			return new(bytes.Buffer)
		},
	}
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

func generateSignature(payload string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(payload))
	return hex.EncodeToString(h.Sum(nil))
}

// ResponseSecure interface for non-secure API responses
type ResponseSecure interface {
	GetError() int
}

func FetchSecure(method string, path string, reqBody any, resPayload any) error {
	return FetchSecureWithContext(context.Background(), method, path, reqBody, resPayload)
}

func FetchSecureWithContext(ctx context.Context, method string, path string, reqBody any, resPayload any) error {
	resp, err := fetchWithContext(ctx, true, method, path, reqBody)
	if err != nil {
		return fmt.Errorf("decoding response: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return fmt.Errorf("%s : %+v", ErrorCode[resp.StatusCode], resp.Request)
	}

	if err = json.NewDecoder(resp.Body).Decode(&resPayload); err != nil {
		return fmt.Errorf("decoding response: %+v", err)
	}

	res := resPayload.(*ResponseAPI)

	if res.Error != 0 {
		return fmt.Errorf("%s : %+v", ErrorCode[res.Error], res)
	}

	return nil
}

func FetchSecureV4(method string, path string, reqBody any, resPayload any) error {
	return FetchSecureV4WithContext(context.Background(), method, path, reqBody, resPayload)
}

func FetchSecureV4WithContext(ctx context.Context, method string, path string, reqBody any, resPayload any) error {
	resp, err := fetchWithContext(ctx, true, method, path, reqBody)
	if err != nil {
		return fmt.Errorf("decoding response: %+v", err)
	}
	defer resp.Body.Close()

	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufferPool.Put(buf)

	if _, err := io.Copy(buf, resp.Body); err != nil {
		return fmt.Errorf("reading response body: %+v", err)
	}

	bodyBytes := buf.Bytes()

	if resp.StatusCode > 299 {
		bodyString := string(bodyBytes)
		log.Error().Str("response_body", bodyString).Msg("API Error Response")
		return fmt.Errorf("HTTP %d : %s - %s", resp.StatusCode, resp.Status, bodyString)
	}

	if err = json.Unmarshal(bodyBytes, &resPayload); err != nil {
		return fmt.Errorf("decoding response: %+v", err)
	}

	res := resPayload.(*ResponseAPIV4)

	if res.Code != "0" {
		errDesc, exists := ErrorCodeV4[res.Code]
		if exists {
			return fmt.Errorf("[%s] %s: %s", res.Code, errDesc, res.Message)
		}
		return fmt.Errorf("[%s] %s", res.Code, res.Message)
	}

	return nil
}

func FetchNonSecure(method string, path string, reqBody any, resPayload any) error {
	return FetchNonSecureWithContext(context.Background(), method, path, reqBody, resPayload)
}

func FetchNonSecureWithContext(ctx context.Context, method string, path string, reqBody any, resPayload any) error {
	resp, err := fetchWithContext(ctx, false, method, path, reqBody)
	if err != nil {
		return fmt.Errorf("decoding response: %+v", err)
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&resPayload); err != nil {
		return fmt.Errorf("decoding response: %+v", err)
	}

	res, ok := resPayload.(ResponseSecure)
	if ok && res.GetError() != 0 {
		errMsg, exists := ErrorCode[res.GetError()]
		if !exists {
			errMsg = "Unknown error"
		}
		return fmt.Errorf("[error %d] %s", res.GetError(), errMsg)
	}

	return nil
}

func fetchWithContext(ctx context.Context, secure bool, method string, path string, reqBody any) (*http.Response, error) {
	var payload []byte = nil

	serverTime, err := GetServerTime()
	if err != nil {
		return nil, fmt.Errorf("server time: %+v", err)
	}

	if reqBody != nil {
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Reset()
		defer bufferPool.Put(buf)

		if err := json.NewEncoder(buf).Encode(reqBody); err != nil {
			return nil, fmt.Errorf("marshaling json: %+v", err)
		}
		payload = make([]byte, buf.Len())
		copy(payload, buf.Bytes())
	}

	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s%s", BASE_URL, path), bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("creating request: %+v", err)
	}

	signaturePayload := fmt.Sprintf(`%s%s%s`, serverTime, req.Method, req.URL.Path)
	if req.URL.RawQuery != "" {
		signaturePayload += fmt.Sprintf(`?%s`, req.URL.RawQuery)
	}
	if len(payload) > 0 {
		signaturePayload += string(payload)
	}
	signature := generateSignature(signaturePayload)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	if secure {
		req.Header.Set("X-BTK-TIMESTAMP", serverTime)
		req.Header.Set("X-BTK-APIKEY", apiKey)
		req.Header.Set("X-BTK-SIGN", signature)
	}

	resp, err := apiBitkub.Do(req)
	if err != nil {
		log.Warn().Int("status", resp.StatusCode).Str("method", method).Str("path", path).Err(err).Stack().Send()
		return nil, fmt.Errorf("making request: %+v", err)
	}

	log.Debug().Int("status", resp.StatusCode).Str("method", method).Str("path", path).Send()
	return resp, nil
}
