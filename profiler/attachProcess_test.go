package profiler

import (
	"os"
	"testing"
)

func TestProcessMonitoring(t *testing.T) {
	if os.Getenv("RUN_PROCESS_INTEGRATION_TEST") == "" {
		t.Skip("set RUN_PROCESS_INTEGRATION_TEST=1 to scan operating-system processes")
	}

	processPath, err := os.Executable()
	if err != nil {
		t.Fatalf("os.Executable() error = %v", err)
	}

	attached, err := AttachProcess(processPath)
	if err != nil {
		t.Fatalf("AttachProcess() error = %v", err)
	}
	if !IsProcessRunning(attached) {
		t.Fatal("IsProcessRunning() = false, want true for the test process")
	}
}

func TestReattachProcessRequiresPath(t *testing.T) {
	if _, err := ReattachProcessByName(""); err == nil {
		t.Fatal("ReattachProcessByName() error = nil, want missing path error")
	}
}

func TestIsProcessRunningRejectsNil(t *testing.T) {
	if IsProcessRunning(nil) {
		t.Fatal("IsProcessRunning(nil) = true, want false")
	}
}
