package handlers

import (
	"FadyGamilM/banking/service/ports"
	"encoding/json"
	"net/http"
)

type CustomerHandler struct {
	Customer_service ports.CustomerService
}

func (ch *CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	customers, err := ch.Customer_service.GetAllCustomers()
	if err != nil {
		panic("Error while fetching the customer from customer service business logic")
	}
	json.NewEncoder(w).Encode(customers)
}
