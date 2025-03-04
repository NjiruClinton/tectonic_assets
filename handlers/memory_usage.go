package handlers

import (
	"github.com/NjiruClinton/tectonic_assets/db"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetMemoryUsage(c echo.Context) error {
	query := "SELECT process_name, usage, timestamp FROM memory_usage ORDER BY timestamp"
	dbConn, err := db.Pgdb()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	rows, err := dbConn.Query(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer rows.Close()
	var memoryUsages []map[string]interface{}
	for rows.Next() {
		var processName string
		var usage float32
		var timestamp string
		err = rows.Scan(&processName, &usage, &timestamp)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		memoryUsage := map[string]interface{}{
			"process_name": processName,
			"usage":        usage,
			"timestamp":    timestamp,
		}
		memoryUsages = append(memoryUsages, memoryUsage)
	}
	return c.JSON(http.StatusOK, memoryUsages)
}
