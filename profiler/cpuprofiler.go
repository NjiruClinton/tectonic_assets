package profiler

import (
	"database/sql"
	"fmt"
	"github.com/NjiruClinton/tectonic_assets/db"
	"github.com/shirou/gopsutil/process"
	"time"
)

type Profiler struct {
	name      string
	processID int
	db        *sql.DB
	interval  time.Duration
	process   *process.Process
}

func NewProfiler(name string, processID int, db *sql.DB, interval time.Duration) (*Profiler, error) {
	p, err := process.NewProcess(int32(processID))
	if err != nil {
		return nil, fmt.Errorf("error attaching to process: %v", err)
	}
	return &Profiler{
		name:      name,
		processID: processID,
		db:        db,
		interval:  interval,
		process:   p,
	}, nil
}

func (p *Profiler) collectCPUUsage() (float64, error) {
	fmt.Println("Collecting CPU usage for process ID:", p.processID)
	percentages, err := p.process.CPUPercent()
	if err != nil {
		return 0, err
	}
	return percentages, nil
}

func (p *Profiler) storeCPUUsage(usage float64) error {
	fmt.Println("Storing CPU usage:", usage)
	query := `INSERT INTO cpu_usage (process_name, usage) VALUES ($1, $2)`
	_, err := db.ExecuteQuery(p.db, query, p.name, usage)
	return err
}

func (p *Profiler) Start() {
	for {
		usage, err := p.collectCPUUsage()
		if err != nil {
			fmt.Println("Error collecting CPU usage:", err)
			continue
		}
		err = p.storeCPUUsage(usage)
		if err != nil {
			fmt.Println("Error storing CPU usage:", err)
		}
		time.Sleep(p.interval)
	}
}
