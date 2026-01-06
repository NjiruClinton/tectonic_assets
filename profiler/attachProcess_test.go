package profiler

import (
	"github.com/shirou/gopsutil/process"
	"testing"
)

func TestAttachProcess(t *testing.T) {
	// Mock process ID for testing
	pid := int32(48507)

	// Check if the process exists
	exists, err := process.PidExists(pid)
	if err != nil {
		t.Fatalf("Error checking if PID %d exists: %v", pid, err)
	}
	if !exists {
		t.Fatalf("Process with PID %d does not exist", pid)
	}

	// Call the AttachProcess function
	p := AttachProcess()

	// Check if the process is attached successfully
	if p == nil {
		t.Errorf("Failed to attach to process with PID %d", pid)
	} else {
		t.Logf("Successfully attached to process with PID %d", pid)
	}

	// Check if the process ID matches
	if p.Pid != pid {
		t.Errorf("Expected PID %d, but got %d", pid, p.Pid)
	}
}
