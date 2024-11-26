package profiler

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
)

func AttachProcess() *process.Process {
	pid := int32(50920)
	p, err := process.NewProcess(pid)
	if err != nil {
		fmt.Println("Error attaching to process:", err)
		return nil
	}
	fmt.Println("Process ID:", p.Pid)
	return p
}
