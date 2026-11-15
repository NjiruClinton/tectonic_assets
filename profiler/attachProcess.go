package profiler

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/shirou/gopsutil/process"
)

func AttachProcess(processPath string) (*process.Process, error) {
	p, err := ReattachProcessByName(processPath)
	if err != nil {
		return nil, fmt.Errorf("attach to process: %w", err)
	}
	fmt.Println("Successfully attached to process with PID:", p.Pid)
	return p, nil
}

// logic for re attaching process after it has restarted
// it will have a new pid each time it restarts
func ReattachProcessByName(processPath string) (*process.Process, error) {
	processPath = strings.TrimSpace(processPath)
	if processPath == "" {
		return nil, fmt.Errorf("process path is required")
	}
	wantedPath := canonicalPath(processPath)

	// process is running at path
	procs, err := process.Processes()
	if err != nil {
		return nil, fmt.Errorf("retrieve processes: %w", err)
	}

	for _, p := range procs {
		if p.Pid < 100 {
			continue
		}
		exe, err := p.Exe()
		if err != nil {
			continue
		}
		if canonicalPath(exe) == wantedPath {
			return p, nil
		}
	}

	return nil, fmt.Errorf("process with path %s not found", processPath)
}

func canonicalPath(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		absPath = path
	}
	absPath = filepath.Clean(absPath)

	resolvedPath, err := filepath.EvalSymlinks(absPath)
	if err == nil {
		return filepath.Clean(resolvedPath)
	}
	return absPath
}

func IsProcessRunning(p *process.Process) bool {
	if p == nil {
		return false
	}
	running, err := p.IsRunning()
	if err != nil {
		return false
	}
	return running
}

func WaitAndReattachProcess(processPath string, checkInterval time.Duration) *process.Process {
	fmt.Printf("Process not found or terminated. Waiting for process at %s to restart...\n", processPath)

	for {
		p, err := ReattachProcessByName(processPath)
		if err == nil {
			fmt.Printf("Process restarted! Successfully reattached to process with PID: %d\n", p.Pid)
			return p
		}

		fmt.Printf("Process still not available, checking again in %v...\n", checkInterval)
		time.Sleep(checkInterval)
	}
}
