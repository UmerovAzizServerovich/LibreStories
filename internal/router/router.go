package router

import (
	"librestories/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/general_statistics", handlers.GetGeneralStatisticsHandler)
}
