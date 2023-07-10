package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

func main() {

	http.HandleFunc("/api", readinessHandler)
	http.HandleFunc("/customers", getAllCustomers)
	http.ListenAndServe(":5050", nil)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte("api is up"))
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	customers := []User{
		{"fady gamil"},
		{"magy magdy"},
	}
	json.NewEncoder(w).Encode(customers)
}
