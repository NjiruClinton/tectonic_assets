package main

import (
	"fmt"
	"github.com/NjiruClinton/tectonic_assets/api/routes"
	"github.com/NjiruClinton/tectonic_assets/db"
	"github.com/NjiruClinton/tectonic_assets/profiler"
	"github.com/NjiruClinton/tectonic_assets/tests"
	"github.com/NjiruClinton/tectonic_assets/timetool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"net/http"
	"text/template"
	"time"
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

type Count struct {
	Count int
}

func customRenderer() {
	e := echo.New()

	count := Count{Count: 0}

	e.Renderer = newTemplate()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		count.Count++
		return c.Render(200, "index.html", count)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func main() {
	mux := routes.SetupRoutes()
	http.Handle("/", mux)
	timeTool := timetool.NewTime()
	timeTool.Start()

	dbConn, err := db.Pgdb()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer dbConn.Close()
	prof := profiler.NewProfiler("example_process", 1, dbConn, 5*time.Second)
	go prof.Start()
	customRenderer()

	tests.TestCPU()
	timeTool.Stop()
	jsonData, err := timeTool.ToJSON()
	if err != nil {
		fmt.Println("Error formatting data:", err)
		return
	}
	fmt.Println("Profiling Data:", jsonData)
}
