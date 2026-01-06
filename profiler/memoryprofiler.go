package profiler

import (
	"database/sql"
	"fmt"
	"github.com/NjiruClinton/tectonic_assets/db"
	"github.com/shirou/gopsutil/process"
	"time"
)

type MemoryProfiler struct {
	name        string
	processID   int
	processPath string
	db          *sql.DB
	interval    time.Duration
	process     *process.Process
}

func NewMemoryProfiler(name string, processID int, processPath string, db *sql.DB, interval time.Duration) (*MemoryProfiler, error) {
	p, err := process.NewProcess(int32(processID))
	if err != nil {
		return nil, fmt.Errorf("error attaching to process: %v", err)
	}
	return &MemoryProfiler{
		name:        name,
		processID:   processID,
		processPath: processPath,
		db:          db,
		interval:    interval,
		process:     p,
	}, nil
}

func (p *MemoryProfiler) collectMemoryUsage() (float32, error) {
	fmt.Println("Collecting memory usage for process ID:", p.processID)
	memInfo, err := p.process.MemoryInfo()
	fmt.Println("Memory info:", memInfo)
	if err != nil {
		return 0, err
	}
	return float32(memInfo.RSS), nil
}

func (p *MemoryProfiler) storeMemoryUsage(usage float32) error {
	fmt.Println("Storing memory usage:", usage)
	query := `INSERT INTO memory_usage (process_name, usage) VALUES ($1, $2)`
	_, err := db.ExecuteQuery(p.db, query, p.name, usage)
	return err
}

func (p *MemoryProfiler) Start() {
	for {
		if !IsProcessRunning(p.process) {
			fmt.Printf("Memory Profiler: Process %s (PID: %d) is not running or terminated\n", p.name, p.processID)
			p.process = WaitAndReattachProcess(p.processPath, 2*time.Second)
			p.processID = int(p.process.Pid)
			fmt.Printf("Memory Profiler: Successfully reattached to process %s (new PID: %d)\n", p.name, p.processID)
		}

		usage, err := p.collectMemoryUsage()
		if err != nil {
			fmt.Printf("Error collecting memory usage for %s: %v\n", p.name, err)
			if !IsProcessRunning(p.process) {
				continue
			}
			time.Sleep(p.interval)
			continue
		}
		err = p.storeMemoryUsage(usage)
		if err != nil {
			fmt.Println("Error storing memory usage:", err)
		}
		time.Sleep(p.interval)
	}
}
