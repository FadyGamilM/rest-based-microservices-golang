package ports

import "FadyGamilM/banking/domain/core"

type CustomerService interface {
	GetAllCustomers() ([]core.Customer, error)
	GetCustomerById(int) (*core.Customer, error)
}
