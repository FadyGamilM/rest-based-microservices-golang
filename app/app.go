package app

import (
	"FadyGamilM/banking/app/handlers"
	repository_adapters "FadyGamilM/banking/domain/adapters"
	service_adapters "FadyGamilM/banking/service/adapters"

	"net/http"

	"github.com/gorilla/mux"
)

func Start_Server() {
	// mux := http.NewServeMux()
	// ! => refacotr to goriall mux
	router := mux.NewRouter()
	router.HandleFunc("/api", readinessHandler)

	// connecting the ports and adapters together
	customer_handler := handlers.CustomerHandler{
		Customer_service: service_adapters.NewCustomerServiceBusinessLogic(repository_adapters.NewCustomerRepoStub()),
	}
	router.HandleFunc("/customers", customer_handler.GetAllCustomers)
	http.ListenAndServe(":5050", router)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte("api is up"))
}
