package app

import (
	"FadyGamilM/banking/app/handlers"
	"net/http"
)

func Start_Server() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api", readinessHandler)
	mux.HandleFunc("/customers", handlers.GetAllCustomers)
	http.ListenAndServe(":5050", mux)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte("api is up"))
}
