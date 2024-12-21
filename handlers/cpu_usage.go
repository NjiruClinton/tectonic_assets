package handlers

import (
	"github.com/NjiruClinton/tectonic_assets/db"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetCPUUsage(c echo.Context) error {
	query := "SELECT process_name, usage, timestamp FROM cpu_usage ORDER BY timestamp"
	dbConn, err := db.Pgdb()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	rows, err := dbConn.Query(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer rows.Close()
	var cpuUsages []map[string]interface{}
	for rows.Next() {
		var processName string
		var usage float64
		var timestamp string
		err = rows.Scan(&processName, &usage, &timestamp)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		cpuUsage := map[string]interface{}{
			"process_name": processName,
			"usage":        usage,
			"timestamp":    timestamp,
		}
		cpuUsages = append(cpuUsages, cpuUsage)
	}
	return c.JSON(http.StatusOK, cpuUsages)

}
