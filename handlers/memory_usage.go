package handlers

import (
	"database/sql"
	"fmt"
	"github.com/NjiruClinton/tectonic_assets/db"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func GetMemoryUsage(c echo.Context) error {
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")
	processFilter := c.QueryParam("process")

	if startDate == "" {
		startDate = c.FormValue("start_date")
	}
	if endDate == "" {
		endDate = c.FormValue("end_date")
	}
	if processFilter == "" {
		processFilter = c.FormValue("process")
	}

	query := "SELECT process_name, usage, timestamp FROM memory_usage WHERE 1=1"
	args := []interface{}{}
	argIndex := 1

	if startDate != "" {
		if parsedTime, err := time.Parse("2006-01-02T15:04", startDate); err == nil {
			query += " AND timestamp >= $" + fmt.Sprintf("%d", argIndex)
			args = append(args, parsedTime)
			argIndex++
		}
	} else {
		query += " AND timestamp >= NOW() - INTERVAL '5 hours'"
	}
	if endDate != "" {
		if parsedTime, err := time.Parse("2006-01-02T15:04", endDate); err == nil {
			query += " AND timestamp <= $" + fmt.Sprintf("%d", argIndex)
			args = append(args, parsedTime)
			argIndex++
		}
	}
	if processFilter != "" {
		query += " AND process_name = $" + fmt.Sprintf("%d", argIndex)
		args = append(args, processFilter)
	}

	query += " ORDER BY timestamp ASC"

	dbConn, err := db.Pgdb()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer dbConn.Close()

	var rows *sql.Rows
	if len(args) > 0 {
		rows, err = dbConn.Query(query, args...)
	} else {
		rows, err = dbConn.Query(query)
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer rows.Close()
	var memoryUsages []map[string]interface{}
	for rows.Next() {
		var processName string
		var usage float32
		var timestamp time.Time
		err = rows.Scan(&processName, &usage, &timestamp)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		memoryUsage := map[string]interface{}{
			"process_name": processName,
			"usage":        usage,
			"timestamp":    timestamp.Format(time.RFC3339),
		}
		memoryUsages = append(memoryUsages, memoryUsage)
	}
	return c.JSON(http.StatusOK, memoryUsages)
}

func GetMemoryProcesses(c echo.Context) error {
	query := "SELECT DISTINCT process_name FROM memory_usage ORDER BY process_name"

	dbConn, err := db.Pgdb()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer dbConn.Close()

	rows, err := dbConn.Query(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer rows.Close()

	var processes []string
	for rows.Next() {
		var processName string
		err = rows.Scan(&processName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		processes = append(processes, processName)
	}

	return c.JSON(http.StatusOK, processes)
}
