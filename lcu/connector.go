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

var (
	portRegex  = regexp.MustCompile(`--app-port=([0-9]+)`)
	tokenRegex = regexp.MustCompile(`--remoting-auth-token=([\w-]+)`)
)

func GetConnectionParams() (*ConnectionParams, error) {
	var processName string
	switch runtime.GOOS {
	case "windows":
		processName = "LeagueClientUx.exe"
	case "darwin":
		processName = "LeagueClientUx"
	default:
		return nil, fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	procs, err := process.Processes()
	if err != nil {
		return nil, fmt.Errorf("failed to list processes: %w", err)
	}

	for _, p := range procs {
		name, _ := p.Name()
		if name == processName {
			cmdline, err := p.Cmdline()
			if err != nil {
				continue
			}

			portMatch := portRegex.FindStringSubmatch(cmdline)
			tokenMatch := tokenRegex.FindStringSubmatch(cmdline)

			if len(portMatch) > 1 && len(tokenMatch) > 1 {
				return &ConnectionParams{
					Port:     portMatch[1],
					Password: tokenMatch[1],
				}, nil
			}
		}
	}

	return nil, fmt.Errorf("Client process not found")
}
