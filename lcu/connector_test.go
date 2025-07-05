package lcu

import (
	"errors"
	"testing"
)

type MockProcessProvider struct {
	processes []ProcessInfo
	err       error
}

func (m *MockProcessProvider) GetProcesses() ([]ProcessInfo, error) {
	return m.processes, m.err
}

type MockOSProvider struct {
	osName string
}

func (m *MockOSProvider) GetOS() string {
	return m.osName
}

func TestGetConnectionParamsWithProviders_Success(t *testing.T) {
	mockProcProvider := &MockProcessProvider{
		processes: []ProcessInfo{
			{
				Name:    "LeagueClientUx.exe",
				Cmdline: "LeagueClientUx.exe --app-port=12345 --remoting-auth-token=abc-123-def",
			},
		},
	}
	mockOSProvider := &MockOSProvider{osName: "windows"}

	result, err := GetConnectionParamsWithProviders(mockProcProvider, mockOSProvider)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("Expected result, got nil")
	}
	if result.Port != "12345" {
		t.Errorf("Expected port '12345', got '%s'", result.Port)
	}
	if result.Password != "abc-123-def" {
		t.Errorf("Expected password 'abc-123-def', got '%s'", result.Password)
	}
}

func TestGetConnectionParamsWithProviders_UnsupportedOS(t *testing.T) {
	mockProcProvider := &MockProcessProvider{}
	mockOSProvider := &MockOSProvider{osName: "linux"}

	result, err := GetConnectionParamsWithProviders(mockProcProvider, mockOSProvider)

	if err == nil {
		t.Error("Expected error for unsupported OS, got nil")
	}
	if result != nil {
		t.Error("Expected nil result for unsupported OS")
	}
	expectedMsg := "unsupported OS: linux"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestGetConnectionParamsWithProviders_ProcessListError(t *testing.T) {
	mockProcProvider := &MockProcessProvider{
		err: errors.New("process error"),
	}
	mockOSProvider := &MockOSProvider{osName: "windows"}

	result, err := GetConnectionParamsWithProviders(mockProcProvider, mockOSProvider)

	if err == nil {
		t.Error("Expected error when process listing fails, got nil")
	}
	if result != nil {
		t.Error("Expected nil result when process listing fails")
	}
	expectedMsg := "failed to list processes: process error"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestGetConnectionParamsWithProviders_ProcessNotFound(t *testing.T) {
	mockProcProvider := &MockProcessProvider{
		processes: []ProcessInfo{
			{
				Name:    "SomeOtherProcess.exe",
				Cmdline: "SomeOtherProcess.exe --some-args",
			},
		},
	}
	mockOSProvider := &MockOSProvider{osName: "windows"}

	result, err := GetConnectionParamsWithProviders(mockProcProvider, mockOSProvider)

	if err == nil {
		t.Error("Expected error when process not found, got nil")
	}
	if result != nil {
		t.Error("Expected nil result when process not found")
	}
	expectedMsg := "Client process not found"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestGetConnectionParamsWithProviders_InvalidCommandLine(t *testing.T) {
	mockProcProvider := &MockProcessProvider{
		processes: []ProcessInfo{
			{
				Name:    "LeagueClientUx.exe",
				Cmdline: "LeagueClientUx.exe --some-other-args",
			},
		},
	}
	mockOSProvider := &MockOSProvider{osName: "windows"}

	result, err := GetConnectionParamsWithProviders(mockProcProvider, mockOSProvider)

	if err == nil {
		t.Error("Expected error when command line is invalid, got nil")
	}
	if result != nil {
		t.Error("Expected nil result when command line is invalid")
	}
	expectedMsg := "Client process not found"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestGetConnectionParamsWithProviders_MultipleProcesses(t *testing.T) {
	mockProcProvider := &MockProcessProvider{
		processes: []ProcessInfo{
			{
				Name:    "SomeOtherProcess.exe",
				Cmdline: "SomeOtherProcess.exe --some-args",
			},
			{
				Name:    "LeagueClientUx.exe",
				Cmdline: "LeagueClientUx.exe --app-port=9999 --remoting-auth-token=xyz-456-ghi",
			},
			{
				Name:    "AnotherProcess.exe",
				Cmdline: "AnotherProcess.exe --different-args",
			},
		},
	}
	mockOSProvider := &MockOSProvider{osName: "windows"}

	result, err := GetConnectionParamsWithProviders(mockProcProvider, mockOSProvider)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("Expected result, got nil")
	}
	if result.Port != "9999" {
		t.Errorf("Expected port '9999', got '%s'", result.Port)
	}
	if result.Password != "xyz-456-ghi" {
		t.Errorf("Expected password 'xyz-456-ghi', got '%s'", result.Password)
	}
}

func TestParseConnectionParams_Success(t *testing.T) {
	testCases := []struct {
		name          string
		cmdline       string
		expectedPort  string
		expectedToken string
	}{
		{
			name:          "Standard format",
			cmdline:       "LeagueClientUx.exe --app-port=12345 --remoting-auth-token=abc-123-def",
			expectedPort:  "12345",
			expectedToken: "abc-123-def",
		},
		{
			name:          "Different order",
			cmdline:       "LeagueClientUx.exe --remoting-auth-token=xyz-789-ghi --app-port=54321",
			expectedPort:  "54321",
			expectedToken: "xyz-789-ghi",
		},
		{
			name:          "With additional args",
			cmdline:       "LeagueClientUx.exe --some-arg=value --app-port=11111 --another-arg --remoting-auth-token=test-token-123",
			expectedPort:  "11111",
			expectedToken: "test-token-123",
		},
		{
			name:          "Token with underscores",
			cmdline:       "LeagueClientUx.exe --app-port=8888 --remoting-auth-token=token_with_underscores",
			expectedPort:  "8888",
			expectedToken: "token_with_underscores",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseConnectionParams(tc.cmdline)

			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if result == nil {
				t.Fatal("Expected result, got nil")
			}
			if result.Port != tc.expectedPort {
				t.Errorf("Expected port '%s', got '%s'", tc.expectedPort, result.Port)
			}
			if result.Password != tc.expectedToken {
				t.Errorf("Expected password '%s', got '%s'", tc.expectedToken, result.Password)
			}
		})
	}
}

func TestParseConnectionParams_Failure(t *testing.T) {
	testCases := []struct {
		name    string
		cmdline string
	}{
		{
			name:    "Missing port",
			cmdline: "LeagueClientUx.exe --remoting-auth-token=abc-123-def",
		},
		{
			name:    "Missing token",
			cmdline: "LeagueClientUx.exe --app-port=12345",
		},
		{
			name:    "Invalid port format",
			cmdline: "LeagueClientUx.exe --app-port=abc --remoting-auth-token=abc-123-def",
		},
		{
			name:    "Empty port",
			cmdline: "LeagueClientUx.exe --app-port= --remoting-auth-token=abc-123-def",
		},
		{
			name:    "Empty token",
			cmdline: "LeagueClientUx.exe --app-port=12345 --remoting-auth-token=",
		},
		{
			name:    "No arguments",
			cmdline: "LeagueClientUx.exe",
		},
		{
			name:    "Wrong argument format",
			cmdline: "LeagueClientUx.exe --port=12345 --token=abc-123-def",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseConnectionParams(tc.cmdline)

			if err == nil {
				t.Error("Expected error for invalid command line, got nil")
			}
			if result != nil {
				t.Error("Expected nil result for invalid command line")
			}
			expectedMsg := "invalid command line format"
			if err.Error() != expectedMsg {
				t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
			}
		})
	}
}

func TestGetProcessNameForOS_Success(t *testing.T) {
	testCases := []struct {
		osName       string
		expectedName string
	}{
		{
			osName:       "windows",
			expectedName: "LeagueClientUx.exe",
		},
		{
			osName:       "darwin",
			expectedName: "LeagueClientUx",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.osName, func(t *testing.T) {
			result, err := GetProcessNameForOS(tc.osName)

			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if result != tc.expectedName {
				t.Errorf("Expected process name '%s', got '%s'", tc.expectedName, result)
			}
		})
	}
}

func TestGetProcessNameForOS_UnsupportedOS(t *testing.T) {
	testCases := []string{"linux", "freebsd", "openbsd", "netbsd", "solaris", "plan9", "js"}

	for _, osName := range testCases {
		t.Run(osName, func(t *testing.T) {
			result, err := GetProcessNameForOS(osName)

			if err == nil {
				t.Error("Expected error for unsupported OS, got nil")
			}
			if result != "" {
				t.Errorf("Expected empty string for unsupported OS, got '%s'", result)
			}
			expectedMsg := "unsupported OS: " + osName
			if err.Error() != expectedMsg {
				t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
			}
		})
	}
}

func TestGetConnectionParams_Integration(t *testing.T) {
	// This test verifies that the default implementation works
	// It will fail if League Client is not running, but that's expected
	result, err := GetConnectionParams()

	if err != nil {
		t.Logf("GetConnectionParams returned error (expected if League Client not running): %v", err)
	}
	if result != nil {
		t.Logf("GetConnectionParams returned: Port=%s, Password=%s", result.Port, result.Password)
		if result.Port == "" {
			t.Error("Port should not be empty when result is returned")
		}
		if result.Password == "" {
			t.Error("Password should not be empty when result is returned")
		}
	}
}

// Benchmark tests
func BenchmarkParseConnectionParams(b *testing.B) {
	cmdline := "LeagueClientUx.exe --app-port=12345 --remoting-auth-token=abc-123-def"

	for i := 0; i < b.N; i++ {
		_, _ = ParseConnectionParams(cmdline)
	}
}

func BenchmarkGetProcessNameForOS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GetProcessNameForOS("windows")
	}
}
