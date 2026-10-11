package stream

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

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
			if !s.readFrame() {
				return
			}
		}
	}
}

func (s *Stream) readFrame() bool {
	if s.conn == nil {
		return false
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
		return false
	}

	s.dispatchFrame(message)
	return true
}

func (s *Stream) dispatchFrame(frame []byte) {
	decoder := json.NewDecoder(bytes.NewReader(frame))
	for decoder.More() {
		var rawMsg map[string]any
		if err := decoder.Decode(&rawMsg); err != nil {
			s.messageChannel <- Message{
				Error:     fmt.Errorf("json decode error: %w", err),
				Timestamp: time.Now(),
			}
			return
		}

		s.messageChannel <- Message{
			Type:      messageType(rawMsg),
			Data:      rawMsg,
			Timestamp: time.Now(),
		}
	}
}

func messageType(rawMsg map[string]any) string {
	if event, ok := rawMsg["event"].(string); ok {
		return event
	}
	if stream, ok := rawMsg["stream"].(string); ok {
		return stream
	}
	return ""
}
