package ports

import "FadyGamilM/banking/domain/core"

type CustomerService interface {
	GetAllCustomers() ([]core.Customer, error)
}
