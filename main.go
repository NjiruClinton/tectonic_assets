package main

import (
	"fmt"
	"io"
	"net/http"
	"text/template"

	"github.com/NjiruClinton/tectonic_assets/api/controllers"
	"github.com/NjiruClinton/tectonic_assets/config"
	"github.com/NjiruClinton/tectonic_assets/db"
	"github.com/NjiruClinton/tectonic_assets/handlers"
	"github.com/NjiruClinton/tectonic_assets/profiler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	tmpl *template.Template
}

func newTemplate() *Template {
	return &Template{
		tmpl: template.Must(template.ParseGlob("views/*.html")),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.tmpl.ExecuteTemplate(w, name, data)
}

func newServer() *echo.Echo {
	e := echo.New()

	e.Renderer = newTemplate()
	e.Use(middleware.Logger())

	e.Static("/assets", "views")
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index.html", nil)
	})
	e.GET("/cpu_usage", handlers.GetCPUUsage)
	e.POST("/cpu_usage", handlers.GetCPUUsage)
	e.GET("/memory_usage", handlers.GetMemoryUsage)
	e.POST("/memory_usage", handlers.GetMemoryUsage)
	e.GET("/cpu_processes", handlers.GetCPUProcesses)
	e.GET("/memory_processes", handlers.GetMemoryProcesses)
	e.POST("/api/collect", echo.WrapHandler(http.HandlerFunc(controllers.CollectAndSendCPUUsage)))

	return e
}

func startProfilers(cfg config.Config) {
	if cfg.TargetProcessPath == "" {
		fmt.Println("Profiling disabled: set TARGET_PROCESS_PATH to monitor a process")
		return
	}

	dbConn, err := db.Pgdb()
	if err != nil {
		fmt.Println("Profiling disabled: database unavailable:", err)
		return
	}

	attachedProcess, err := profiler.AttachProcess(cfg.TargetProcessPath)
	if err != nil {
		_ = dbConn.Close()
		fmt.Println("Profiling disabled:", err)
		return
	}

	cpuProfiler, err := profiler.NewProfiler(cfg.TargetName, int(attachedProcess.Pid), cfg.TargetProcessPath, dbConn, cfg.ProfileInterval)
	if err != nil {
		_ = dbConn.Close()
		fmt.Println("Profiling disabled: create CPU profiler:", err)
		return
	}

	memoryProfiler, err := profiler.NewMemoryProfiler(cfg.TargetName, int(attachedProcess.Pid), cfg.TargetProcessPath, dbConn, cfg.ProfileInterval)
	if err != nil {
		_ = dbConn.Close()
		fmt.Println("Profiling disabled: create memory profiler:", err)
		return
	}

	go cpuProfiler.Start()
	go memoryProfiler.Start()
	fmt.Printf("Profiling %s every %s\n", cfg.TargetName, cfg.ProfileInterval)
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println("Invalid configuration:", err)
		return
	}

	server := newServer()
	go startProfilers(cfg)
	server.Logger.Fatal(server.Start(cfg.ServerAddress))
}
