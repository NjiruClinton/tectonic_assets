package profiler

import (
	"errors"
	"fmt"
	"github.com/shirou/gopsutil/process"
	"go.uber.org/zap"
	_ "go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

func AttachProcess() *process.Process {
	pid := int32(39698)
	p, err := process.NewProcess(pid)
	if err != nil {
		fmt.Println("Error attaching to process:", err)
		return nil
	}
	// child processes
	childPIDs, err := findChildPIDs(int(pid))
	if err != nil {
		fmt.Println("Error finding child PIDs:", err)
		return nil
	}
	fmt.Println("Child PIDs:", childPIDs)

	fmt.Println("Process ID:", p.Pid)
	return p
}

func interruptProcessTree(logger *zap.Logger, ppid int, sig syscall.Signal) error {
	// Find all descendant PIDs of the given PID & then signal them.
	// Any shell doesn't signal its children when it receives a signal.
	// Children may have their own process groups, so we need to signal them separately.
	children, err := findChildPIDs(ppid)
	if err != nil {
		return err
	}

	children = append(children, ppid)
	uniqueProcess, err := uniqueProcessGroups(children)
	if err != nil {
		logger.Error("failed to find unique process groups", zap.Int("pid", ppid), zap.Error(err))
		uniqueProcess = children
	}

	for _, pid := range uniqueProcess {
		err := syscall.Kill(-pid, sig)
		// ignore the ESRCH error as it means the process is already dead
		var errno syscall.Errno
		if errors.As(err, &errno) && err != nil && !errors.Is(errno, syscall.ESRCH) {
			logger.Error("failed to send signal to process", zap.Int("pid", pid), zap.Error(err))
		}
	}
	return nil
}

func uniqueProcessGroups(pids []int) ([]int, error) {
	uniqueGroups := make(map[int]bool)
	var uniqueGPIDs []int

	for _, pid := range pids {
		pgid, err := getProcessGroupID(pid)
		if err != nil {
			return nil, err
		}
		if !uniqueGroups[pgid] {
			uniqueGroups[pgid] = true
			uniqueGPIDs = append(uniqueGPIDs, pgid)
		}
	}

	return uniqueGPIDs, nil
}

func getProcessGroupID(pid int) (int, error) {
	statusPath := filepath.Join("/proc", strconv.Itoa(pid), "status")
	statusBytes, err := os.ReadFile(statusPath)
	if err != nil {
		return 0, err
	}

	status := string(statusBytes)
	for _, line := range strings.Split(status, "\n") {
		if strings.HasPrefix(line, "NSpgid:") {
			return extractIDFromStatusLine(line), nil
		}
	}

	return 0, nil
}

func extractIDFromStatusLine(line string) int {
	fields := strings.Fields(line)
	if len(fields) == 2 {
		id, err := strconv.Atoi(fields[1])
		if err == nil {
			return id
		}
	}
	return -1
}

func findChildPIDs(parentPID int) ([]int, error) {
	var childPIDs []int

	// Recursive helper function to find all descendants of a given PID.
	var findDescendants func(int)
	findDescendants = func(pid int) {
		procDirs, err := os.ReadDir("/proc")
		if err != nil {
			return
		}

		for _, procDir := range procDirs {
			if !procDir.IsDir() {
				continue
			}

			childPid, err := strconv.Atoi(procDir.Name())
			if err != nil {
				continue
			}

			statusPath := filepath.Join("/proc", procDir.Name(), "status")
			statusBytes, err := os.ReadFile(statusPath)
			if err != nil {
				continue
			}

			status := string(statusBytes)
			for _, line := range strings.Split(status, "\n") {
				if strings.HasPrefix(line, "PPid:") {
					fields := strings.Fields(line)
					if len(fields) == 2 {
						ppid, err := strconv.Atoi(fields[1])
						if err != nil {
							break
						}
						if ppid == pid {
							childPIDs = append(childPIDs, childPid)
							findDescendants(childPid)
						}
					}
					break
				}
			}
		}
	}

	// Start the recursion with the initial parent PID.
	findDescendants(parentPID)

	return childPIDs, nil
}
