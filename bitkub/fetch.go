package bitkub

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

func generateSignature(payload string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(payload))
	return hex.EncodeToString(h.Sum(nil))
}

func FetchSecure(method string, path string, reqBody any, resPayload any) error {
	resp, err := fetch(true, method, path, reqBody)
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

func FetchNonSecure(method string, path string, reqBody any, resPayload any) error {
	resp, err := fetch(true, method, path, reqBody)
	if err != nil {
		return fmt.Errorf("decoding response: %+v", err)
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&resPayload); err != nil {
		return fmt.Errorf("decoding response: %+v", err)
	}

	return nil
}

func fetch(secure bool, method string, path string, reqBody any) (*http.Response, error) {
	var payload []byte = nil

	serverTime, err := getServerTime()
	if err != nil {
		return nil, fmt.Errorf("server time: %+v", err)
	}

	if reqBody != nil {
		payload, err = json.Marshal(reqBody)
		if err != nil {
			return nil, fmt.Errorf("marshaling json: %+v", err)
		}
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", BASE_URL, path), bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("creating request: %+v", err)
	}

	// Generate timestamp and signature

	signaturePayload := fmt.Sprintf(`%s%s%s`, serverTime, req.Method, req.URL.Path)
	if req.URL.RawQuery != "" {
		signaturePayload += fmt.Sprintf(`?%s`, req.URL.RawQuery)
	}
	signature := generateSignature(signaturePayload)

	// Set the required headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	if secure {
		req.Header.Set("X-BTK-TIMESTAMP", serverTime)
		req.Header.Set("X-BTK-APIKEY", apiKey)
		req.Header.Set("X-BTK-SIGN", signature)
	}

	// Make the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Warn().Int("status", resp.StatusCode).Str("method", method).Str("path", path).Err(err).Stack().Send()
		return nil, fmt.Errorf("making request: %+v", err)
	}

	log.Debug().Int("status", resp.StatusCode).Str("method", method).Str("path", path).Send()
	return resp, nil
}

func getServerTime() (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s%s", BASE_URL, "/v3/servertime"))
	if err != nil {
		return "0", err
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return "0", err
	}

	return string(result), nil
}
