package bitkub

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/dvgamerr-app/go-bitkub/utils"
)

const BASE_URL = "https://api.bitkub.com"

type Config struct {
	APITimeout      time.Duration
	ServerTimeout   time.Duration
	MaxIdleConns    int
	IdleConnTimeout time.Duration
}

var (
	apiKey         string
	secretKey      string
	timeOffsetMs   int64
	timeOffsetOnce sync.Once
	timeOffsetErr  error
	config         = Config{
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

func Initlizer(key ...string) error {
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
	timeOffsetOnce.Do(func() {
		var lastErr error
		for attempt := 1; attempt <= 3; attempt++ {
			if attempt > 1 {
				backoff := time.Duration(attempt-1) * 500 * time.Millisecond
				time.Sleep(backoff)
			}

			resp, err := apiBitkubTime.Get(fmt.Sprintf("%s%s", BASE_URL, "/api/v3/servertime"))
			if err != nil {
				lastErr = err
				if attempt < 3 {
					continue
				}
				timeOffsetErr = fmt.Errorf("(%d) %v", attempt, lastErr)
				return
			}

			result, err := io.ReadAll(resp.Body)
			resp.Body.Close()

			if err != nil {
				lastErr = err
				if attempt < 3 {
					continue
				}
				timeOffsetErr = fmt.Errorf("(%d) %v", attempt, lastErr)
				return
			}

			timeStr := string(result)
			timeStr = strings.Trim(timeStr, "\" \n\r")

			serverTime, err := time.Parse("2006-01-02T15:04:05.999Z07:00", timeStr)
			if err != nil {
				serverTimeMs := int64(0)
				fmt.Sscanf(timeStr, "%d", &serverTimeMs)
				if serverTimeMs > 0 {
					timeOffsetMs = serverTimeMs - time.Now().UnixMilli()
					return
				}
				lastErr = err
				if attempt < 3 {
					continue
				}
				timeOffsetErr = fmt.Errorf("parse time (%d) %v", attempt, lastErr)
				return
			}

			timeOffsetMs = serverTime.UnixMilli() - time.Now().UnixMilli()
			return
		}

		timeOffsetErr = fmt.Errorf("(retry) %v", lastErr)
	})

	if timeOffsetErr != nil {
		return "0", timeOffsetErr
	}

	currentTime := time.Now().UnixMilli() + timeOffsetMs
	return fmt.Sprintf("%d", currentTime), nil
}
