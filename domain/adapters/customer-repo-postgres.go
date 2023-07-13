package adapters

import (
	"FadyGamilM/banking/domain/core"
	"FadyGamilM/banking/infra/db"
	"context"
	"database/sql"
	"log"
)

// the repository implementation using postgresql
type CustomerRepoPostgres struct {
	db_conn_pool *sql.DB
}

// implement the secondry port
func (crp *CustomerRepoPostgres) GetAll() ([]core.Customer, error) {
	/*
		1. define a context with time out = 3 seconds to limit my database operations
		2. call cancel on this context to release all used resources for this context
		3. define a query
		4. execute the query using db.QueryWityContext() method and pass the context
		5. handle if there is an error due to the query execution
		6. map the result of the query to my domain core model using the rows.Next and rows.Scan methods
		7. return the result
	*/

	ctx, cancel := context.WithTimeout(context.Background(), db.DbTimeOut)
	defer cancel()

	select_all_query := `SELECT id, name, city, zip_code, date_of_birth FROM customers`

	rows, err := crp.db_conn_pool.QueryContext(ctx, select_all_query)
	if err != nil {
		panic("Error while fetching data from database | customers table")
	}

	defer rows.Close()

	all_customers := make([]core.Customer, 0)
	for rows.Next() {
		var customer core.Customer
		err := rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.City,
			&customer.ZipCode,
			&customer.BirthDate,
		)
		if err != nil {
			panic("Error while mapping the db row to customer golang type")
		}
		all_customers = append(all_customers, customer)
	}

	return all_customers, nil
}

func (crp *CustomerRepoPostgres) GetById(customerID int) (*core.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.DbTimeOut)
	defer cancel()

	select_by_id_query := `
        SELECT * FROM customers WHERE id = $1
    `

	row := crp.db_conn_pool.QueryRowContext(ctx, select_by_id_query, customerID)

	var customer core.Customer
	err := row.Scan(
		&customer.ID,
		&customer.Name,
		&customer.City,
		&customer.ZipCode,
		&customer.BirthDate,
	)
	if err != nil {
		log.Printf("Error while fetching customer by id from postgres repo => %v", err)
		return nil, err
	}
	return &customer, nil
}

// Factory method pattern
func NewCustomerRepoPostgres(conn_pool *sql.DB) *CustomerRepoPostgres {
	return &CustomerRepoPostgres{db_conn_pool: conn_pool}
}
