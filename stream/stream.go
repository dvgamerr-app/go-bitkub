package stream

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	BaseURL         = "wss://api.bitkub.com/websocket-api"
	DefaultPingTime = 30 * time.Second
	DefaultTimeout  = 60 * time.Second
)

type Stream struct {
	conn              *websocket.Conn
	url               string
	config            StreamConfig
	mu                sync.RWMutex
	ctx               context.Context
	cancel            context.CancelFunc
	messageChannel    chan Message
	isConnected       bool
	lastPong          time.Time
	reconnectAttempts int
}

func New(config *StreamConfig) *Stream {
	if config == nil {
		config = &StreamConfig{
			ReconnectInterval: 5 * time.Second,
			MaxReconnect:      10,
			PingInterval:      DefaultPingTime,
			ReadTimeout:       DefaultTimeout,
			MessageBuffer:     100,
		}
	} else {
		if config.PingInterval <= 0 {
			config.PingInterval = DefaultPingTime
		}
		if config.ReadTimeout <= 0 {
			config.ReadTimeout = DefaultTimeout
		}
		if config.ReconnectInterval <= 0 {
			config.ReconnectInterval = 5 * time.Second
		}
		if config.MessageBuffer <= 0 {
			config.MessageBuffer = 100
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &Stream{
		config:         *config,
		ctx:            ctx,
		cancel:         cancel,
		messageChannel: make(chan Message, config.MessageBuffer),
		lastPong:       time.Now(),
	}
}

func (s *Stream) ConnectMarket(streams ...string) error {
	streamNames := ""
	for i, stream := range streams {
		if i > 0 {
			streamNames += ","
		}
		streamNames += stream
	}

	s.url = fmt.Sprintf("%s/%s", BaseURL, streamNames)
	return s.connect()
}

func (s *Stream) ConnectOrderBook(symbolID int) error {
	s.url = fmt.Sprintf("%s/orderbook/%d", BaseURL, symbolID)
	return s.connect()
}

func (s *Stream) connect() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.isConnected {
		return fmt.Errorf("already connected")
	}

	dialer := websocket.DefaultDialer
	dialer.HandshakeTimeout = 10 * time.Second

	conn, _, err := dialer.Dial(s.url, nil)
	if err != nil {
		return fmt.Errorf("websocket dial failed: %w", err)
	}

	s.conn = conn
	s.isConnected = true
	s.reconnectAttempts = 0
	s.lastPong = time.Now()

	conn.SetPongHandler(func(string) error {
		s.lastPong = time.Now()
		return nil
	})

	go s.readMessages()
	go s.pingPong()

	return nil
}

func (s *Stream) readMessages() {
	defer func() {
		s.mu.Lock()
		s.isConnected = false
		s.mu.Unlock()
	}()

	for {
		select {
		case <-s.ctx.Done():
			return
		default:
			if s.conn == nil {
				return
			}

			s.conn.SetReadDeadline(time.Now().Add(s.config.ReadTimeout))
			_, message, err := s.conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					s.messageChannel <- Message{
						Error:     err,
						Timestamp: time.Now(),
					}
				}

				s.handleReconnect()
				return
			}

			decoder := json.NewDecoder(bytes.NewReader(message))
			for decoder.More() {
				var rawMsg map[string]any
				if err := decoder.Decode(&rawMsg); err != nil {
					s.messageChannel <- Message{
						Error:     fmt.Errorf("json decode error: %w", err),
						Timestamp: time.Now(),
					}
					break
				}

				msg := Message{
					Data:      rawMsg,
					Timestamp: time.Now(),
				}

				if event, ok := rawMsg["event"].(string); ok {
					msg.Type = event
				} else if stream, ok := rawMsg["stream"].(string); ok {
					msg.Type = stream
				}

				s.messageChannel <- msg
			}
		}
	}
}

func (s *Stream) pingPong() {
	ticker := time.NewTicker(s.config.PingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-s.ctx.Done():
			return
		case <-ticker.C:
			s.mu.RLock()
			conn := s.conn
			isConnected := s.isConnected
			s.mu.RUnlock()

			if !isConnected || conn == nil {
				continue
			}

			if time.Since(s.lastPong) > s.config.ReadTimeout {
				s.messageChannel <- Message{
					Error:     fmt.Errorf("connection timeout: no pong received"),
					Timestamp: time.Now(),
				}
				s.handleReconnect()
				return
			}

			if err := conn.WriteControl(websocket.PingMessage, nil, time.Now().Add(time.Second)); err != nil {
				s.messageChannel <- Message{
					Error:     fmt.Errorf("ping failed: %w", err),
					Timestamp: time.Now(),
				}
				s.handleReconnect()
				return
			}
		}
	}
}

func (s *Stream) handleReconnect() {
	s.mu.Lock()
	if !s.isConnected {
		s.mu.Unlock()
		return
	}

	if s.conn != nil {
		s.conn.Close()
		s.conn = nil
	}
	s.isConnected = false
	s.reconnectAttempts++
	s.mu.Unlock()

	if s.config.MaxReconnect > 0 && s.reconnectAttempts >= s.config.MaxReconnect {
		s.messageChannel <- Message{
			Error:     fmt.Errorf("max reconnect attempts reached (%d)", s.config.MaxReconnect),
			Timestamp: time.Now(),
		}
		return
	}

	time.Sleep(s.config.ReconnectInterval)

	if err := s.connect(); err != nil {
		s.messageChannel <- Message{
			Error:     fmt.Errorf("reconnect failed (attempt %d): %w", s.reconnectAttempts, err),
			Timestamp: time.Now(),
		}
		go s.handleReconnect()
	} else {
		s.messageChannel <- Message{
			Type:      "reconnected",
			Timestamp: time.Now(),
		}
	}
}

func (s *Stream) Messages() <-chan Message {
	return s.messageChannel
}

func (s *Stream) Close() error {
	s.cancel()

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.conn != nil {
		s.conn.WriteControl(websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
			time.Now().Add(time.Second))
		s.conn.Close()
		s.conn = nil
	}

	s.isConnected = false

	go func() {
		time.Sleep(100 * time.Millisecond)
		defer recover()
		close(s.messageChannel)
	}()

	return nil
}

func (s *Stream) IsConnected() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.isConnected
}
