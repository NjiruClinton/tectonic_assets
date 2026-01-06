package profiler

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
	_ "go.uber.org/zap/zapcore"
)

func AttachProcess() *process.Process {
	// process path
	processPath := "/Users/admin/Desktop/stuff/clinton/toyl/tmp/main"
	p, err := ReattachProcessByName(processPath)
	if err != nil {
		fmt.Println("Error attaching to process:", err)
		return nil
	}
	fmt.Println("Successfully attached to process with PID:", p.Pid)
	return p
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
