package profiler

import (
	"database/sql"
	"fmt"
	"github.com/NjiruClinton/tectonic_assets/db"
	_ "github.com/lib/pq"
	"github.com/shirou/gopsutil/cpu"
	"time"
)

type Profiler struct {
	name      string
	processID int
	db        *sql.DB
	interval  time.Duration
}

func NewProfiler(name string, processID int, db *sql.DB, interval time.Duration) *Profiler {
	return &Profiler{
		name:      name,
		processID: processID,
		db:        db,
		interval:  interval,
	}
}

func (p *Profiler) collectCPUUsage() (float64, error) {
	fmt.Println("Collecting CPU usage...")
	percentages, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	return percentages[0], nil
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
