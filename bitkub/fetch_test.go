package bitkub

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
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
		timeSynced = false
		timeOffsetMs = 0
	})
	timeSynced = false
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

func TestFetchSecureStatusErrorDoesNotLeakCredentials(t *testing.T) {
	resetFetchState(t)

	apiKey = "api-key-value"
	secretKey = "secret-key-value"
	apiBitkubTime = &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		return response("2000000000000"), nil
	})}
	apiBitkub = &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		resp := response("{}")
		resp.StatusCode = http.StatusUnauthorized
		return resp, nil
	})}

	var result ResponseAPI
	err := FetchSecure(http.MethodPost, "/secure", nil, &result)
	if err == nil {
		t.Fatalf("expected status error")
	}
	if strings.Contains(err.Error(), apiKey) || strings.Contains(err.Error(), "X-Btk-Sign") {
		t.Fatalf("error leaks credentials: %v", err)
	}
}

func TestFetchSecureRequiresCredentials(t *testing.T) {
	resetFetchState(t)

	apiKey = ""
	secretKey = ""
	apiBitkub = &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		t.Fatalf("secure request without credentials reached the network")
		return nil, nil
	})}

	var result ResponseAPI
	if err := FetchSecure(http.MethodPost, "/secure", nil, &result); !errors.Is(err, ErrMissingCredentials) {
		t.Fatalf("expected missing credentials error, got %v", err)
	}
}

func TestFetchSecureAcceptsForeignPayload(t *testing.T) {
	resetFetchState(t)

	apiKey = "api-key"
	secretKey = "secret-key"
	apiBitkubTime = &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		return response("2000000000000"), nil
	})}
	apiBitkub = &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		return response(`{"value":"ok"}`), nil
	})}

	var result map[string]string
	if err := FetchSecure(http.MethodGet, "/secure", nil, &result); err != nil {
		t.Fatalf("FetchSecure returned error: %v", err)
	}
	if err := FetchSecureV4(http.MethodGet, "/secure", nil, &result); err != nil {
		t.Fatalf("FetchSecureV4 returned error: %v", err)
	}
}

func TestFetchNonSecureRejectsErrorStatus(t *testing.T) {
	resetFetchState(t)

	apiBitkub = &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		resp := response("<html>bad gateway</html>")
		resp.StatusCode = http.StatusBadGateway
		return resp, nil
	})}

	var result any
	err := FetchNonSecure(http.MethodGet, "/public", nil, &result)
	if err == nil || !strings.Contains(err.Error(), "HTTP 502") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestFetchGetWithoutBodyOmitsContentType(t *testing.T) {
	resetFetchState(t)

	apiBitkub = &http.Client{Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
		if req.Header.Get("Content-Type") != "" {
			t.Fatalf("unexpected content type on bodyless request")
		}
		return response(`{}`), nil
	})}

	var result any
	if err := FetchNonSecure(http.MethodGet, "/public", nil, &result); err != nil {
		t.Fatalf("FetchNonSecure returned error: %v", err)
	}
}

func TestServerTimeFailureIsNotCached(t *testing.T) {
	resetFetchState(t)

	calls := 0
	apiBitkubTime = &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		calls++
		if calls <= serverTimeAttempts {
			return nil, errors.New("offline")
		}
		return response("2000000000000"), nil
	})}

	if _, err := GetServerTime(); err == nil {
		t.Fatalf("expected server time failure")
	}
	if _, err := GetServerTime(); err != nil {
		t.Fatalf("server time did not recover: %v", err)
	}
}

func TestServerTimeHonorsContextCancellation(t *testing.T) {
	resetFetchState(t)

	apiBitkubTime = &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		return nil, errors.New("offline")
	})}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := GetServerTimeWithContext(ctx); !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context cancellation, got %v", err)
	}
}
