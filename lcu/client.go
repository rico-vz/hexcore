package lcu

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

type HexcoreClient struct {
	*ClientWithResponses
	Params     *ConnectionParams
	httpClient *http.Client
	wsClient   *WSClient
	wsProvider WSClientProvider
}

type HexcoreClientConfig struct {
	Timeout          time.Duration
	RequestEditorFns []RequestEditorFn
}

type ConnectionParamsProvider interface {
	GetConnectionParams() (*ConnectionParams, error)
}

type DefaultConnectionParamsProvider struct{}

func (p *DefaultConnectionParamsProvider) GetConnectionParams() (*ConnectionParams, error) {
	return GetConnectionParams()
}

type ClientWithResponsesProvider interface {
	NewClientWithResponses(serverURL string, options ...ClientOption) (*ClientWithResponses, error)
}

type DefaultClientWithResponsesProvider struct{}

func (p *DefaultClientWithResponsesProvider) NewClientWithResponses(serverURL string, options ...ClientOption) (*ClientWithResponses, error) {
	return NewClientWithResponses(serverURL, options...)
}

type WSClientProvider interface {
	NewWSClient(params *ConnectionParams) (*WSClient, error)
}

type DefaultWSClientProvider struct{}

func (p *DefaultWSClientProvider) NewWSClient(params *ConnectionParams) (*WSClient, error) {
	return NewWSClient(params)
}

func DefaultHexcoreClientConfig() *HexcoreClientConfig {
	return &HexcoreClientConfig{
		Timeout: 30 * time.Second,
	}
}

func NewHexcoreClient() (*HexcoreClient, error) {
	return NewHexcoreClientWithConfig(DefaultHexcoreClientConfig())
}

func NewHexcoreClientWithConfig(config *HexcoreClientConfig) (*HexcoreClient, error) {
	return NewHexcoreClientWithProviders(
		config,
		&DefaultConnectionParamsProvider{},
		&DefaultClientWithResponsesProvider{},
		&DefaultWSClientProvider{},
	)
}

func NewHexcoreClientWithProviders(
	config *HexcoreClientConfig,
	paramsProvider ConnectionParamsProvider,
	clientProvider ClientWithResponsesProvider,
	wsProvider WSClientProvider,
) (*HexcoreClient, error) {
	params, err := paramsProvider.GetConnectionParams()
	if err != nil {
		return nil, fmt.Errorf("could not get LCU connection params: %w", err)
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			MinVersion:         tls.VersionTLS12,
		},
		MaxIdleConns:        10,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     30 * time.Second,
	}

	httpClient := &http.Client{
		Transport: transport,
		Timeout:   config.Timeout,
	}

	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte("riot:"+params.Password))
	serverURL := fmt.Sprintf("https://127.0.0.1:%s", params.Port)

	options := []ClientOption{
		WithHTTPClient(httpClient),
	}

	authEditor := func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", authHeader)
		return nil
	}
	options = append(options, WithRequestEditorFn(authEditor))

	for _, editor := range config.RequestEditorFns {
		options = append(options, WithRequestEditorFn(editor))
	}

	clientWithResponses, err := clientProvider.NewClientWithResponses(serverURL, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	return &HexcoreClient{
		ClientWithResponses: clientWithResponses,
		Params:              params,
		httpClient:          httpClient,
		wsClient:            nil,
		wsProvider:          wsProvider,
	}, nil
}

func (c *HexcoreClient) Close() error {
	if c.wsClient != nil {
		c.wsClient.Close()
	}
	if c.httpClient != nil {
		c.httpClient.CloseIdleConnections()
	}
	return nil
}

func (c *HexcoreClient) GetHTTPClient() *http.Client {
	return c.httpClient
}

func (c *HexcoreClient) GetWSClient() (*WSClient, error) {
	if c.wsClient == nil {
		var err error
		c.wsClient, err = c.wsProvider.NewWSClient(c.Params)
		if err != nil {
			return nil, fmt.Errorf("failed to create WebSocket client: %w", err)
		}
	}
	return c.wsClient, nil
}

func (c *HexcoreClient) Subscribe(topic string, handler EventHandler) error {
	wsClient, err := c.GetWSClient()
	if err != nil {
		return err
	}
	return wsClient.Subscribe(topic, handler)
}

func (c *HexcoreClient) Unsubscribe(topic string, handler EventHandler) error {
	if c.wsClient == nil {
		return fmt.Errorf("WebSocket client not initialized")
	}
	return c.wsClient.Unsubscribe(topic, handler)
}

func (c *HexcoreClient) UnsubscribeAll(topic string) error {
	if c.wsClient == nil {
		return fmt.Errorf("WebSocket client not initialized")
	}
	return c.wsClient.UnsubscribeAll(topic)
}
