package main

import (
	"fmt"
	"github.com/NjiruClinton/tectonic_assets/api/routes"
	"github.com/NjiruClinton/tectonic_assets/db"
	"github.com/NjiruClinton/tectonic_assets/tests"
	"github.com/NjiruClinton/tectonic_assets/timetool"
	"net/http"
)

func main() {
	mux := routes.SetupRoutes()
	http.Handle("/", mux)

	timeTool := timetool.NewTime()
	timeTool.Start()

	// workload here

	db.Pgdb()
	fmt.Println("database connected... ")

	port := 8080
	fmt.Printf("Server running at http://localhost:%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

	tests.TestCPU()
	timeTool.Stop()
	jsonData, err := timeTool.ToJSON()
	if err != nil {
		fmt.Println("Error formatting data:", err)
		return
	}
	fmt.Println("Profiling Data:", jsonData)
}
