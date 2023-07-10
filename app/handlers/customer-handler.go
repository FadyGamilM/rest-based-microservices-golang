package handlers

import (
	"FadyGamilM/banking/app/types"
	"encoding/json"
	"net/http"
)

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	customers := []types.User{
		{"fady gamil"},
		{"magy magdy"},
	}
	json.NewEncoder(w).Encode(customers)
}
