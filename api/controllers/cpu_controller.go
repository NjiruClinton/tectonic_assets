package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CPUUsageData struct {
	CPUUsage  string `json:"cpuUsage"`
	Timestamp string `json:"timestamp"`
}

func CollectAndSendCPUUsage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Collecting CPU Usage")
	var data CPUUsageData
	body, err := ioutil.ReadAll(r.Body) //TODO: ReadAll is deprecated
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

	//for {
	//	w.Header().Set("Content-Type", "application/json")
	//	cpuUsage := getCPUUsage()
	//	json.NewEncoder(w).Encode(cpuUsage)
	//
	//	time.Sleep(1 * time.Minute)
	//}

}

func getCPUUsage() string {
	return "10%" // Example value
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "home")
}
