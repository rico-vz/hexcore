package lcu

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/coder/websocket"
)

// Message types
const (
	MessageTypeWelcome     = 0
	MessageTypePrefix      = 1
	MessageTypeCall        = 2
	MessageTypeCallResult  = 3
	MessageTypeCallError   = 4
	MessageTypeSubscribe   = 5
	MessageTypeUnsubscribe = 6
	MessageTypePublish     = 7
	MessageTypeEvent       = 8
)

type EventHandler func(payload interface{})

type WSClient struct {
	conn          *websocket.Conn
	params        *ConnectionParams
	session       string
	handlers      map[string][]EventHandler
	handlersMutex sync.RWMutex
	ctx           context.Context
	cancel        context.CancelFunc
	closed        bool
	closedMutex   sync.RWMutex
}

func NewWSClient(params *ConnectionParams) (*WSClient, error) {
	if params == nil {
		var err error
		params, err = GetConnectionParams()
		if err != nil {
			return nil, fmt.Errorf("failed to get connection params: %w", err)
		}
	}

	ctx, cancel := context.WithCancel(context.Background())

	client := &WSClient{
		params:   params,
		handlers: make(map[string][]EventHandler),
		ctx:      ctx,
		cancel:   cancel,
	}

	if err := client.connect(); err != nil {
		cancel()
		return nil, fmt.Errorf("failed to connect: %w", err)
	}

	go client.messageLoop()

	return client, nil
}

func (c *WSClient) connect() error {
	wsURL := fmt.Sprintf("wss://riot:%s@127.0.0.1:%s/", c.params.Password, c.params.Port)

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, //gosec:disable G402 -- Uses self-signed certificates
			},
		},
	}

	conn, _, err := websocket.Dial(c.ctx, wsURL, &websocket.DialOptions{
		HTTPClient:   httpClient,
		HTTPHeader:   http.Header{},
		Subprotocols: []string{"wamp"},
	})
	if err != nil {
		return fmt.Errorf("failed to dial websocket: %w", err)
	}

	conn.SetReadLimit(100 * 1024 * 1024) // fix "Read limited" error

	c.conn = conn
	return nil
}

func (c *WSClient) Subscribe(topic string, handler EventHandler) error {
	c.handlersMutex.Lock()
	defer c.handlersMutex.Unlock()

	if c.isClosed() {
		return fmt.Errorf("client is closed")
	}

	c.handlers[topic] = append(c.handlers[topic], handler)

	message := []interface{}{MessageTypeSubscribe, topic}
	return c.sendMessage(message)
}

func (c *WSClient) Unsubscribe(topic string, handler EventHandler) error {
	c.handlersMutex.Lock()
	defer c.handlersMutex.Unlock()

	if c.isClosed() {
		return fmt.Errorf("client is closed")
	}

	handlers := c.handlers[topic]
	for i, h := range handlers {
		if fmt.Sprintf("%p", h) == fmt.Sprintf("%p", handler) {
			c.handlers[topic] = append(handlers[:i], handlers[i+1:]...)
			break
		}
	}

	if len(c.handlers[topic]) == 0 {
		delete(c.handlers, topic)
		message := []interface{}{MessageTypeUnsubscribe, topic}
		return c.sendMessage(message)
	}

	return nil
}

func (c *WSClient) UnsubscribeAll(topic string) error {
	c.handlersMutex.Lock()
	defer c.handlersMutex.Unlock()

	if c.isClosed() {
		return fmt.Errorf("client is closed")
	}

	if _, exists := c.handlers[topic]; exists {
		delete(c.handlers, topic)
		message := []interface{}{MessageTypeUnsubscribe, topic}
		return c.sendMessage(message)
	}

	return nil
}

func (c *WSClient) sendMessage(message []interface{}) error {
	if c.isClosed() {
		return fmt.Errorf("client is closed")
	}

	data, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	return c.conn.Write(c.ctx, websocket.MessageText, data)
}

func (c *WSClient) messageLoop() {
	defer c.close()

	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			_, data, err := c.conn.Read(c.ctx)
			if err != nil {
				if !c.isClosed() {
					log.Printf("WebSocket read error: %v", err)
				}
				return
			}

			if err := c.handleMessage(data); err != nil {
				log.Printf("Error handling message: %v", err)
			}
		}
	}
}

func (c *WSClient) handleMessage(data []byte) error {
	var message []interface{}
	if err := json.Unmarshal(data, &message); err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}

	if len(message) == 0 {
		return fmt.Errorf("empty message")
	}

	messageType, ok := message[0].(float64)
	if !ok {
		return fmt.Errorf("invalid message type")
	}

	switch int(messageType) {
	case MessageTypeWelcome:
		if len(message) > 1 {
			if session, ok := message[1].(string); ok {
				c.session = session
			}
		}
	case MessageTypeCallResult:
		log.Printf("Unknown call result: %v", message[1:])
	case MessageTypeCallError:
		log.Printf("Unknown call error: %v", message[1:])
	case MessageTypeEvent:
		if len(message) >= 3 {
			topic, ok := message[1].(string)
			if !ok {
				return fmt.Errorf("invalid topic in event message")
			}

			payload := message[2]
			c.emitEvent(topic, payload)
		}
	default:
		log.Printf("Unknown message type %d: %v", int(messageType), message[1:])
	}

	return nil
}

func (c *WSClient) emitEvent(topic string, payload interface{}) {
	c.handlersMutex.RLock()
	handlers := make([]EventHandler, len(c.handlers[topic]))
	copy(handlers, c.handlers[topic])
	c.handlersMutex.RUnlock()

	for _, handler := range handlers {
		go handler(payload)
	}
}

func (c *WSClient) Close() error {
	c.closedMutex.Lock()
	defer c.closedMutex.Unlock()

	if c.closed {
		return nil
	}

	c.closed = true
	c.cancel()

	if c.conn != nil {
		return c.conn.Close(websocket.StatusNormalClosure, "")
	}

	return nil
}

func (c *WSClient) close() {
	c.closedMutex.Lock()
	defer c.closedMutex.Unlock()

	if !c.closed {
		c.closed = true
		c.cancel()
	}
}

func (c *WSClient) isClosed() bool {
	c.closedMutex.RLock()
	defer c.closedMutex.RUnlock()
	return c.closed
}

func (c *WSClient) Session() string {
	return c.session
}

func (c *WSClient) IsConnected() bool {
	return !c.isClosed() && c.conn != nil
}
