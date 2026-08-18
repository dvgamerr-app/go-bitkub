package bitkub

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (fn roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return fn(req)
}

func response(body string) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func resetFetchState(t *testing.T) {
	previousAPI := apiBitkub
	previousTimeAPI := apiBitkubTime
	previousAPIKey := apiKey
	previousSecretKey := secretKey
	t.Cleanup(func() {
		apiBitkub = previousAPI
		apiBitkubTime = previousTimeAPI
		apiKey = previousAPIKey
		secretKey = previousSecretKey
		timeOffsetOnce = sync.Once{}
		timeOffsetErr = nil
		timeOffsetMs = 0
	})
	timeOffsetOnce = sync.Once{}
	timeOffsetErr = nil
	timeOffsetMs = 0
}

func TestFetchNonSecureSkipsAuthentication(t *testing.T) {
	resetFetchState(t)

	timeRequests := 0
	apiBitkubTime = &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		timeRequests++
		return response("2000000000000"), nil
	})}

	apiBitkub = &http.Client{Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
		if req.Header.Get("X-BTK-TIMESTAMP") != "" || req.Header.Get("X-BTK-APIKEY") != "" || req.Header.Get("X-BTK-SIGN") != "" {
			t.Fatalf("non-secure request contains authentication headers")
		}
		return response(`{"value":"ok"}`), nil
	})}

	var result struct {
		Value string `json:"value"`
	}
	if err := FetchNonSecure(http.MethodGet, "/public", nil, &result); err != nil {
		t.Fatalf("FetchNonSecure returned error: %v", err)
	}
	if result.Value != "ok" {
		t.Fatalf("unexpected response value: %q", result.Value)
	}
	if timeRequests != 0 {
		t.Fatalf("non-secure request fetched server time %d times", timeRequests)
	}
}

func TestFetchSecureAddsAuthentication(t *testing.T) {
	resetFetchState(t)

	apiKey = "api-key"
	secretKey = "secret-key"
	timeRequests := 0
	apiBitkubTime = &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		timeRequests++
		return response("2000000000000"), nil
	})}

	apiBitkub = &http.Client{Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
		if req.Header.Get("X-BTK-TIMESTAMP") == "" || req.Header.Get("X-BTK-APIKEY") != apiKey || req.Header.Get("X-BTK-SIGN") == "" {
			t.Fatalf("secure request is missing authentication headers")
		}
		return response(`{"error":0,"result":"ok"}`), nil
	})}

	var result ResponseAPI
	if err := FetchSecure(http.MethodPost, "/secure", nil, &result); err != nil {
		t.Fatalf("FetchSecure returned error: %v", err)
	}
	if timeRequests != 1 {
		t.Fatalf("secure request fetched server time %d times", timeRequests)
	}
}

func TestFetchTransportErrorDoesNotPanic(t *testing.T) {
	resetFetchState(t)

	apiBitkub = &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		return nil, errors.New("transport unavailable")
	})}

	var result any
	err := FetchNonSecure(http.MethodGet, "/public", nil, &result)
	if err == nil || !strings.Contains(err.Error(), "transport unavailable") {
		t.Fatalf("unexpected transport error: %v", err)
	}
}
