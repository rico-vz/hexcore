package lcu

import (
	"fmt"
	"regexp"
	"runtime"

	"github.com/shirou/gopsutil/v4/process"
)

type ConnectionParams struct {
	Port     string
	Password string
}

type ProcessInfo struct {
	Name    string
	Cmdline string
}

type ProcessProvider interface {
	GetProcesses() ([]ProcessInfo, error)
}

type DefaultProcessProvider struct{}

func (p *DefaultProcessProvider) GetProcesses() ([]ProcessInfo, error) {
	procs, err := process.Processes()
	if err != nil {
		return nil, err
	}

	var processInfos []ProcessInfo
	for _, proc := range procs {
		name, err := proc.Name()
		if err != nil {
			continue // Skip processes we can't get name for
		}

		cmdline, err := proc.Cmdline()
		if err != nil {
			continue // Skip processes we can't get cmdline for
		}

		processInfos = append(processInfos, ProcessInfo{
			Name:    name,
			Cmdline: cmdline,
		})
	}

	return processInfos, nil
}

type OSProvider interface {
	GetOS() string
}

type DefaultOSProvider struct{}

func (o *DefaultOSProvider) GetOS() string {
	return runtime.GOOS
}

var (
	portRegex  = regexp.MustCompile(`--app-port=([0-9]+)`)
	tokenRegex = regexp.MustCompile(`--remoting-auth-token=([\w-]+)`)
)

func GetConnectionParams() (*ConnectionParams, error) {
	return GetConnectionParamsWithProviders(&DefaultProcessProvider{}, &DefaultOSProvider{})
}

func GetConnectionParamsWithProviders(procProvider ProcessProvider, osProvider OSProvider) (*ConnectionParams, error) {
	processName, err := GetProcessNameForOS(osProvider.GetOS())
	if err != nil {
		return nil, err
	}

	processes, err := procProvider.GetProcesses()
	if err != nil {
		return nil, fmt.Errorf("failed to list processes: %w", err)
	}

	for _, proc := range processes {
		if proc.Name == processName {
			params, err := ParseConnectionParams(proc.Cmdline)
			if err == nil {
				return params, nil
			}
		}
	}

	return nil, fmt.Errorf("Client process not found")
}

func ParseConnectionParams(cmdline string) (*ConnectionParams, error) {
	portMatch := portRegex.FindStringSubmatch(cmdline)
	tokenMatch := tokenRegex.FindStringSubmatch(cmdline)

	if len(portMatch) > 1 && len(tokenMatch) > 1 {
		return &ConnectionParams{
			Port:     portMatch[1],
			Password: tokenMatch[1],
		}, nil
	}

	return nil, fmt.Errorf("invalid command line format")
}

func GetProcessNameForOS(osName string) (string, error) {
	switch osName {
	case "windows":
		return "LeagueClientUx.exe", nil
	case "darwin":
		return "LeagueClientUx", nil
	default:
		return "", fmt.Errorf("unsupported OS: %s", osName)
	}
}
