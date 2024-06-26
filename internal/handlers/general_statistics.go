package handlers

import (
	"encoding/json"
	"fmt"
	"librestories/services"
	"net/http"
	"time"
)

func GetGeneralStatisticsHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		w.Write([]byte("Invalid method!\n"))
		return
	}
	general_statics, err := services.DisplayGeneralStatistics()
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(general_statics)
	fmt.Printf("%s %s %s\n", r.Method, r.RequestURI, time.Since(start))
}
