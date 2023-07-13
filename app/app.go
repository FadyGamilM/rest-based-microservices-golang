package app

import (
	"FadyGamilM/banking/app/handlers"
	repository_adapters "FadyGamilM/banking/domain/adapters"
	"FadyGamilM/banking/infra/db"
	service_adapters "FadyGamilM/banking/service/adapters"
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)

func Start_Server() {
	// mux := http.NewServeMux()
	// ! => refacotr to goriall mux
	router := mux.NewRouter()
	router.HandleFunc("/api", readinessHandler)

	//! connecting the ports and adapters together
	// => utilize the mock database repository
	// customer_handler := handlers.CustomerHandler{
	// 	Customer_service: service_adapters.NewCustomerServiceBusinessLogic(repository_adapters.NewCustomerRepoStub()),
	// }

	// => utilize the postgres repository
	dsn := "host=localhost port=1122 user=fady password=fady dbname=bankingdb sslmode=disable timezone=UTC connect_timeout=5"
	dbConn, err := db.ConnectToPostgresInstance(dsn)
	if err != nil {
		panic("Error while connecting to postgres database")
	}
	customer_handler := handlers.CustomerHandler{
		Customer_service: service_adapters.NewCustomerServiceBusinessLogic(repository_adapters.NewCustomerRepoPostgres(dbConn.DB)),
	}

	router.HandleFunc("/customers", customer_handler.GetAllCustomers)
	fmt.Println("Server is up and running on port 5050")
	http.ListenAndServe(":5050", router)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte("api is up"))
}
