package app

import (
	"FadyGamilM/banking/app/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func Start_Server() {
	// mux := http.NewServeMux()
	// ! => refacotr to goriall mux
	router := mux.NewRouter()
	router.HandleFunc("/api", readinessHandler)
	router.HandleFunc("/customers", handlers.GetAllCustomers)
	http.ListenAndServe(":5050", router)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte("api is up"))
}
