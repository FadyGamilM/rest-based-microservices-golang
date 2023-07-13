package handlers

import (
	"FadyGamilM/banking/service/ports"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	Customer_service ports.CustomerService
}

func (ch *CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	customers, err := ch.Customer_service.GetAllCustomers()
	if err != nil {
		log.Printf("Error while fetching the customer from customer service business logic => %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)
}

func (ch *CustomerHandler) GetCustomerById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	req_params := mux.Vars(r)
	customerID, err := strconv.Atoi(req_params["customer_id"])
	if err != nil {
		// Handle the error if the conversion fails
		http.Error(w, "Invalid ID", http.StatusBadRequest)
	}
	customer, err := ch.Customer_service.GetCustomerById(customerID)
	if err != nil {
		log.Printf("Error while fetching the customer from customer service business logic => %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
}
