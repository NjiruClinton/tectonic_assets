package profiler

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
	_ "go.uber.org/zap/zapcore"
	"time"
)

func AttachProcess() *process.Process {
	processPath := GetProcessPath()
	p, err := ReattachProcessByName(processPath)
	if err != nil {
		fmt.Println("Error attaching to process:", err)
		return nil
	}
	fmt.Println("Successfully attached to process with PID:", p.Pid)
	return p
}

func GetProcessPath() string {
	return "/Users/admin/Desktop/stuff/clinton/toyl/tmp/main"
}

// logic for re attaching process after it has restarted
// it will have a new pid each time it restarts
func ReattachProcessByName(processPath string) (*process.Process, error) {
	// process is running at path
	procs, err := process.Processes()
	if err != nil {
		return nil, fmt.Errorf("error retrieving processes: %v", err)
	}

	for _, p := range procs {
		if p.Pid < 100 {
			continue
		}
		println("Checking process PID:", p.Pid)
		exe, err := p.Exe()
		if err != nil {
			continue
		}
		if exe == processPath {
			return p, nil
		}
	}

	return nil, fmt.Errorf("process with path %s not found", processPath)
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
