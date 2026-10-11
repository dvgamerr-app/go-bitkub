package bitkub

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/dvgamerr-app/go-bitkub/utils"
)

const (
	BASE_URL           = "https://api.bitkub.com"
	serverTimeAttempts = 3
)

type Config struct {
	APITimeout      time.Duration
	ServerTimeout   time.Duration
	MaxIdleConns    int
	IdleConnTimeout time.Duration
}

var (
	apiKey        string
	secretKey     string
	credentialsMu sync.RWMutex
	timeOffsetMu  sync.Mutex
	timeOffsetMs  int64
	timeSynced    bool
	config        = Config{
		APITimeout:      30 * time.Second,
		ServerTimeout:   5 * time.Second,
		MaxIdleConns:    100,
		IdleConnTimeout: 90 * time.Second,
	}
)

type ResponseAPI struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
	Result  any    `json:"result"`
}

type StatusResponse struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// CheckResponseError checks if the response has an error and returns an error if it does
func (r *ResponseAPI) CheckResponseError() error {
	if r.Error != 0 {
		errMsg, exists := ErrorCode[r.Error]
		if !exists {
			errMsg = "Unknown error"
		}
		if r.Message != "" {
			return fmt.Errorf("[error %d] %s: %s", r.Error, errMsg, r.Message)
		}
		return fmt.Errorf("[error %d] %s", r.Error, errMsg)
	}
	return nil
}

type ResponseAPIV4 struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func SetConfig(cfg Config) {
	if cfg.APITimeout > 0 {
		config.APITimeout = cfg.APITimeout
	}
	if cfg.ServerTimeout > 0 {
		config.ServerTimeout = cfg.ServerTimeout
	}
	if cfg.MaxIdleConns > 0 {
		config.MaxIdleConns = cfg.MaxIdleConns
	}
	if cfg.IdleConnTimeout > 0 {
		config.IdleConnTimeout = cfg.IdleConnTimeout
	}
	initHTTPClients()
}

func GetConfig() Config {
	return config
}

func credentials() (string, string) {
	credentialsMu.RLock()
	defer credentialsMu.RUnlock()
	return apiKey, secretKey
}

func Initlizer(key ...string) error {
	credentialsMu.Lock()
	defer credentialsMu.Unlock()

	if len(key) >= 2 {
		apiKey = key[0]
		secretKey = key[1]
	}

	if apiKey == "" || secretKey == "" {
		if err := utils.CheckEnvVars("BTK_APIKEY", "BTK_SECRET"); err != nil {
			return err
		}

		apiKey = os.Getenv("BTK_APIKEY")
		secretKey = os.Getenv("BTK_SECRET")
	}
	return nil
}

func GetStatus() ([]StatusResponse, error) {
	var result []StatusResponse

	if err := FetchNonSecure("GET", "/api/status", nil, &result); err != nil {
		return nil, fmt.Errorf("get status: %w", err)
	}

	return result, nil
}

func GetServerTime() (string, error) {
	return GetServerTimeWithContext(context.Background())
}

func GetServerTimeWithContext(ctx context.Context) (string, error) {
	timeOffsetMu.Lock()
	defer timeOffsetMu.Unlock()

	if !timeSynced {
		offset, err := fetchServerTimeOffset(ctx)
		if err != nil {
			return "0", err
		}
		timeOffsetMs = offset
		timeSynced = true
	}

	return fmt.Sprintf("%d", time.Now().UnixMilli()+timeOffsetMs), nil
}

func fetchServerTimeOffset(ctx context.Context) (int64, error) {
	var lastErr error
	for attempt := 1; attempt <= serverTimeAttempts; attempt++ {
		if attempt > 1 {
			select {
			case <-ctx.Done():
				return 0, fmt.Errorf("server time: %w", ctx.Err())
			case <-time.After(time.Duration(attempt-1) * 500 * time.Millisecond):
			}
		}

		serverTimeMs, err := requestServerTime(ctx)
		if err == nil {
			return serverTimeMs - time.Now().UnixMilli(), nil
		}
		lastErr = err
		if ctx.Err() != nil {
			return 0, fmt.Errorf("server time: %w", ctx.Err())
		}
	}
	return 0, fmt.Errorf("server time after %d attempts: %w", serverTimeAttempts, lastErr)
}

func requestServerTime(ctx context.Context) (int64, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, BASE_URL+"/api/v3/servertime", nil)
	if err != nil {
		return 0, err
	}

	resp, err := apiBitkubTime.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	return parseServerTime(strings.Trim(string(body), "\" \n\r"))
}

func parseServerTime(value string) (int64, error) {
	if parsed, err := time.Parse("2006-01-02T15:04:05.999Z07:00", value); err == nil {
		return parsed.UnixMilli(), nil
	}
	if millis, err := strconv.ParseInt(value, 10, 64); err == nil && millis > 0 {
		return millis, nil
	}
	return 0, fmt.Errorf("parse time %q", value)
}
