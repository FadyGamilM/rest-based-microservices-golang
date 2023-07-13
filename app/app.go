package app

import (
	"FadyGamilM/banking/app/handlers"
	repository_adapters "FadyGamilM/banking/domain/adapters"
	"FadyGamilM/banking/infra/db"
	service_adapters "FadyGamilM/banking/service/adapters"
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/mux"
)

func Start_Server() {
	// mux := http.NewServeMux()
	// ! => refacotr to goriall mux
	router := mux.NewRouter()

	//! connecting the ports and adapters together
	//* => utilize the mock database repository
	// customer_handler := handlers.CustomerHandler{
	// 	Customer_service: service_adapters.NewCustomerServiceBusinessLogic(repository_adapters.NewCustomerRepoStub()),
	// }

	//* => utilize the postgres repository
	dsn := "host=localhost port=1122 user=fady password=fady dbname=bankingdb sslmode=disable timezone=UTC connect_timeout=5"
	dbConn, err := db.ConnectToPostgresInstance(dsn)
	if err != nil {
		log.Fatalln("Database Connection Error !!")
	}
	customer_handler := handlers.CustomerHandler{
		Customer_service: service_adapters.NewCustomerServiceBusinessLogic(repository_adapters.NewCustomerRepoPostgres(dbConn.DB)),
	}

	router.HandleFunc("/customers", customer_handler.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customer_handler.GetCustomerById).Methods(http.MethodGet)
	fmt.Println("Server is up and running on port 5050")
	http.ListenAndServe(":5050", router)
}
