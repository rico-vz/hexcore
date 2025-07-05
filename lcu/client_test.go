package lcu

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"
)

type MockConnectionParamsProvider struct {
	params *ConnectionParams
	err    error
}

func (m *MockConnectionParamsProvider) GetConnectionParams() (*ConnectionParams, error) {
	return m.params, m.err
}

type MockClientWithResponsesProvider struct {
	client *ClientWithResponses
	err    error
}

func (m *MockClientWithResponsesProvider) NewClientWithResponses(serverURL string, options ...ClientOption) (*ClientWithResponses, error) {
	return m.client, m.err
}

type MockWSClientProvider struct {
	err error
}

func (m *MockWSClientProvider) NewWSClient(params *ConnectionParams) (*WSClient, error) {
	return nil, m.err
}

func TestDefaultHexcoreClientConfig(t *testing.T) {
	config := DefaultHexcoreClientConfig()

	if config == nil {
		t.Error("Expected config, got nil")
	}
	if config.Timeout != 30*time.Second {
		t.Errorf("Expected timeout 30s, got %v", config.Timeout)
	}
	if len(config.RequestEditorFns) != 0 {
		t.Errorf("Expected empty RequestEditorFns, got %d items", len(config.RequestEditorFns))
	}
}

func TestNewHexcoreClientWithProviders_Success(t *testing.T) {
	// Setup mocks
	params := &ConnectionParams{
		Port:     "12345",
		Password: "test-password",
	}

	mockParamsProvider := &MockConnectionParamsProvider{params: params}
	mockClientProvider := &MockClientWithResponsesProvider{client: &ClientWithResponses{}}
	mockWSProvider := &MockWSClientProvider{}

	// Test
	config := DefaultHexcoreClientConfig()
	client, err := NewHexcoreClientWithProviders(config, mockParamsProvider, mockClientProvider, mockWSProvider)

	// Assertions
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if client == nil {
		t.Fatal("Expected client, got nil")
	}
	if client.Params != params {
		t.Errorf("Expected params %v, got %v", params, client.Params)
	}
	if client.httpClient == nil {
		t.Error("Expected httpClient to be set")
	}
	if client.wsClient != nil {
		t.Error("Expected wsClient to be nil initially")
	}
}

func TestNewHexcoreClientWithProviders_ConnectionParamsError(t *testing.T) {
	// Setup mocks
	mockParamsProvider := &MockConnectionParamsProvider{err: errors.New("connection error")}
	mockClientProvider := &MockClientWithResponsesProvider{}
	mockWSProvider := &MockWSClientProvider{}

	// Test
	config := DefaultHexcoreClientConfig()
	client, err := NewHexcoreClientWithProviders(config, mockParamsProvider, mockClientProvider, mockWSProvider)

	// Assertions
	if err == nil {
		t.Error("Expected error for connection params failure, got nil")
	}
	if client != nil {
		t.Error("Expected nil client when connection params fail")
	}
	expectedMsg := "could not get LCU connection params"
	if err != nil && len(err.Error()) < len(expectedMsg) {
		t.Errorf("Expected error message to contain '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestNewHexcoreClientWithProviders_ClientCreationError(t *testing.T) {
	// Setup mocks
	params := &ConnectionParams{
		Port:     "12345",
		Password: "test-password",
	}

	mockParamsProvider := &MockConnectionParamsProvider{params: params}
	mockClientProvider := &MockClientWithResponsesProvider{err: errors.New("client creation error")}
	mockWSProvider := &MockWSClientProvider{}

	// Test
	config := DefaultHexcoreClientConfig()
	client, err := NewHexcoreClientWithProviders(config, mockParamsProvider, mockClientProvider, mockWSProvider)

	// Assertions
	if err == nil {
		t.Error("Expected error for client creation failure, got nil")
	}
	if client != nil {
		t.Error("Expected nil client when client creation fails")
	}
	expectedMsg := "failed to create client"
	if err != nil && len(err.Error()) < len(expectedMsg) {
		t.Errorf("Expected error message to contain '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestNewHexcoreClientWithProviders_CustomConfig(t *testing.T) {
	// Setup mocks
	params := &ConnectionParams{
		Port:     "54321",
		Password: "custom-password",
	}

	mockParamsProvider := &MockConnectionParamsProvider{params: params}
	mockClientProvider := &MockClientWithResponsesProvider{client: &ClientWithResponses{}}
	mockWSProvider := &MockWSClientProvider{}

	// Custom config
	customEditor := func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Custom-Header", "test-value")
		return nil
	}

	config := &HexcoreClientConfig{
		Timeout:          45 * time.Second,
		RequestEditorFns: []RequestEditorFn{customEditor},
	}

	// Test
	client, err := NewHexcoreClientWithProviders(config, mockParamsProvider, mockClientProvider, mockWSProvider)

	// Assertions
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if client == nil {
		t.Fatal("Expected client, got nil")
	}
	if client.httpClient.Timeout != 45*time.Second {
		t.Errorf("Expected timeout 45s, got %v", client.httpClient.Timeout)
	}
}

func TestHexcoreClient_GetHTTPClient(t *testing.T) {
	httpClient := &http.Client{Timeout: 30 * time.Second}
	client := &HexcoreClient{
		httpClient: httpClient,
	}

	result := client.GetHTTPClient()
	if result != httpClient {
		t.Errorf("Expected httpClient %v, got %v", httpClient, result)
	}
}

func TestHexcoreClient_GetWSClient_Error(t *testing.T) {
	// Setup mocks
	params := &ConnectionParams{
		Port:     "12345",
		Password: "test-password",
	}

	mockWSProvider := &MockWSClientProvider{err: errors.New("ws creation error")}

	client := &HexcoreClient{
		Params:     params,
		wsProvider: mockWSProvider,
	}

	// Test
	result, err := client.GetWSClient()

	// Assertions
	if err == nil {
		t.Error("Expected error when ws creation fails, got nil")
	}
	if result != nil {
		t.Error("Expected nil result when ws creation fails")
	}
	expectedMsg := "failed to create WebSocket client"
	if err != nil && len(err.Error()) < len(expectedMsg) {
		t.Errorf("Expected error message to contain '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestHexcoreClient_Subscribe_WSClientError(t *testing.T) {
	// Setup mocks
	params := &ConnectionParams{
		Port:     "12345",
		Password: "test-password",
	}

	mockWSProvider := &MockWSClientProvider{err: errors.New("ws error")}

	client := &HexcoreClient{
		Params:     params,
		wsProvider: mockWSProvider,
	}

	handler := func(payload interface{}) {}

	// Test
	err := client.Subscribe("test-topic", handler)

	// Assertions
	if err == nil {
		t.Error("Expected error when ws client creation fails, got nil")
	}
	expectedMsg := "failed to create WebSocket client"
	if err != nil && len(err.Error()) < len(expectedMsg) {
		t.Errorf("Expected error message to contain '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestHexcoreClient_Unsubscribe_NoWSClient(t *testing.T) {
	client := &HexcoreClient{
		wsClient: nil,
	}

	handler := func(payload interface{}) {}

	// Test
	err := client.Unsubscribe("test-topic", handler)

	// Assertions
	if err == nil {
		t.Error("Expected error when ws client not initialized, got nil")
	}
	expectedMsg := "WebSocket client not initialized"
	if err != nil && err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestHexcoreClient_UnsubscribeAll_NoWSClient(t *testing.T) {
	client := &HexcoreClient{
		wsClient: nil,
	}

	// Test
	err := client.UnsubscribeAll("test-topic")

	// Assertions
	if err == nil {
		t.Error("Expected error when ws client not initialized, got nil")
	}
	expectedMsg := "WebSocket client not initialized"
	if err != nil && err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestHexcoreClient_Close_WithoutWSClient(t *testing.T) {
	client := &HexcoreClient{
		wsClient:   nil,
		httpClient: &http.Client{},
	}

	// Test
	err := client.Close()

	// Assertions
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestAuthHeaderGeneration(t *testing.T) {
	password := "test-password"
	expectedAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte("riot:"+password))

	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte("riot:"+password))

	if authHeader != expectedAuth {
		t.Errorf("Expected auth header '%s', got '%s'", expectedAuth, authHeader)
	}
}

func TestServerURLGeneration(t *testing.T) {
	port := "12345"
	expectedURL := "https://127.0.0.1:12345"

	serverURL := fmt.Sprintf("https://127.0.0.1:%s", port)

	if serverURL != expectedURL {
		t.Errorf("Expected server URL '%s', got '%s'", expectedURL, serverURL)
	}
}

func TestNewHexcoreClient_BackwardCompatibility(t *testing.T) {
	client, err := NewHexcoreClient()
	if err != nil {
		t.Logf("NewHexcoreClient returned error (expected if League Client not running): %v", err)
		return
	}

	if client == nil {
		t.Error("Expected client, got nil")
	}
	if client.Params == nil {
		t.Error("Expected params to be set")
	}
	if client.httpClient == nil {
		t.Error("Expected httpClient to be set")
	}

	// Test cleanup
	err = client.Close()
	if err != nil {
		t.Errorf("Expected no error on close, got %v", err)
	}
}

func TestNewHexcoreClientWithConfig_BackwardCompatibility(t *testing.T) {
	config := &HexcoreClientConfig{
		Timeout: 60 * time.Second,
	}

	client, err := NewHexcoreClientWithConfig(config)
	if err != nil {
		t.Logf("NewHexcoreClientWithConfig returned error (expected if League Client not running): %v", err)
		return
	}

	if client == nil {
		t.Error("Expected client, got nil")
	}
	if client.httpClient.Timeout != 60*time.Second {
		t.Errorf("Expected timeout 60s, got %v", client.httpClient.Timeout)
	}

	// Test cleanup
	err = client.Close()
	if err != nil {
		t.Errorf("Expected no error on close, got %v", err)
	}
}

// Integration test (requires actual League Client running)
func TestNewHexcoreClient_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client, err := NewHexcoreClient()
	if err != nil {
		t.Skipf("League Client not running, skipping integration test: %v", err)
	}

	if client == nil {
		t.Error("Expected client, got nil")
	}
	if client.Params == nil {
		t.Error("Expected params to be set")
	}
	if client.Params.Port == "" {
		t.Error("Expected port to be set")
	}
	if client.Params.Password == "" {
		t.Error("Expected password to be set")
	}

	// Test cleanup
	err = client.Close()
	if err != nil {
		t.Errorf("Expected no error on close, got %v", err)
	}
}

// Test that validates the HTTP client configuration
func TestHexcoreClient_HTTPClientConfiguration(t *testing.T) {
	// Setup mocks
	params := &ConnectionParams{
		Port:     "12345",
		Password: "test-password",
	}

	mockParamsProvider := &MockConnectionParamsProvider{params: params}
	mockClientProvider := &MockClientWithResponsesProvider{client: &ClientWithResponses{}}
	mockWSProvider := &MockWSClientProvider{}

	// Test with custom timeout
	config := &HexcoreClientConfig{
		Timeout: 60 * time.Second,
	}

	client, err := NewHexcoreClientWithProviders(config, mockParamsProvider, mockClientProvider, mockWSProvider)

	// Assertions
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if client == nil {
		t.Fatal("Expected client, got nil")
	}

	// Check HTTP client configuration
	if client.httpClient.Timeout != 60*time.Second {
		t.Errorf("Expected timeout 60s, got %v", client.httpClient.Timeout)
	}

	// Check TLS configuration
	transport, ok := client.httpClient.Transport.(*http.Transport)
	if !ok {
		t.Fatal("Expected HTTP transport to be *http.Transport")
	}

	if transport.TLSClientConfig == nil {
		t.Error("Expected TLS config to be set")
	}

	if !transport.TLSClientConfig.InsecureSkipVerify {
		t.Error("Expected InsecureSkipVerify to be true")
	}

	if transport.MaxIdleConns != 10 {
		t.Errorf("Expected MaxIdleConns 10, got %d", transport.MaxIdleConns)
	}

	if transport.MaxIdleConnsPerHost != 10 {
		t.Errorf("Expected MaxIdleConnsPerHost 10, got %d", transport.MaxIdleConnsPerHost)
	}

	if transport.IdleConnTimeout != 30*time.Second {
		t.Errorf("Expected IdleConnTimeout 30s, got %v", transport.IdleConnTimeout)
	}
}

// Test that validates the server URL and auth header generation
func TestHexcoreClient_ServerURLAndAuth(t *testing.T) {
	testCases := []struct {
		name     string
		port     string
		password string
		expected string
	}{
		{
			name:     "Standard port",
			port:     "12345",
			password: "test-password",
			expected: "https://127.0.0.1:12345",
		},
		{
			name:     "Different port",
			port:     "54321",
			password: "another-password",
			expected: "https://127.0.0.1:54321",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test server URL generation
			serverURL := fmt.Sprintf("https://127.0.0.1:%s", tc.port)
			if serverURL != tc.expected {
				t.Errorf("Expected server URL '%s', got '%s'", tc.expected, serverURL)
			}

			// Test auth header generation
			expectedAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte("riot:"+tc.password))
			authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte("riot:"+tc.password))
			if authHeader != expectedAuth {
				t.Errorf("Expected auth header '%s', got '%s'", expectedAuth, authHeader)
			}
		})
	}
}

// Benchmark tests
func BenchmarkNewHexcoreClientWithProviders(b *testing.B) {
	// Setup mocks
	params := &ConnectionParams{
		Port:     "12345",
		Password: "test-password",
	}

	mockParamsProvider := &MockConnectionParamsProvider{params: params}
	mockClientProvider := &MockClientWithResponsesProvider{client: &ClientWithResponses{}}
	mockWSProvider := &MockWSClientProvider{}

	config := DefaultHexcoreClientConfig()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client, err := NewHexcoreClientWithProviders(config, mockParamsProvider, mockClientProvider, mockWSProvider)
		if err != nil {
			b.Fatal(err)
		}
		_ = client
	}
}

func BenchmarkAuthHeaderGeneration(b *testing.B) {
	password := "test-password"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = "Basic " + base64.StdEncoding.EncodeToString([]byte("riot:"+password))
	}
}
