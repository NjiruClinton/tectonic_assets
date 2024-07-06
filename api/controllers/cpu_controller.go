package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func CollectAndSendCPUUsage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Collecting CPU Usage")
	for {
		w.Header().Set("Content-Type", "application/json")
		cpuUsage := getCPUUsage()
		json.NewEncoder(w).Encode(cpuUsage)

		time.Sleep(1 * time.Minute)
	}
}

func getCPUUsage() string {
	return "10%" // Example value
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "home")
}
