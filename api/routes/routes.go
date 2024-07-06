package routes

import (
	"github.com/NjiruClinton/tectonic_assets/api/controllers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.RootHandler).Methods("GET")
	router.HandleFunc("/api/collect", controllers.CollectAndSendCPUUsage).Methods("POST")

	return router
}
