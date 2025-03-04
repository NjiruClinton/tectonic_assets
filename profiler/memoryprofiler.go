package profiler

import (
	"database/sql"
	"fmt"
	"github.com/NjiruClinton/tectonic_assets/db"
	"github.com/shirou/gopsutil/process"
	"time"
)

type MemoryProfiler struct {
	name      string
	processID int
	db        *sql.DB
	interval  time.Duration
	process   *process.Process
}

func NewMemoryProfiler(name string, processID int, db *sql.DB, interval time.Duration) (*MemoryProfiler, error) {
	p, err := process.NewProcess(int32(processID))
	if err != nil {
		return nil, fmt.Errorf("error attaching to process: %v", err)
	}
	return &MemoryProfiler{
		name:      name,
		processID: processID,
		db:        db,
		interval:  interval,
		process:   p,
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
		usage, err := p.collectMemoryUsage()
		if err != nil {
			fmt.Println("Error collecting memory usage:", err)
			continue
		}
		err = p.storeMemoryUsage(usage)
		if err != nil {
			fmt.Println("Error storing memory usage:", err)
		}
		time.Sleep(p.interval)
	}
}
