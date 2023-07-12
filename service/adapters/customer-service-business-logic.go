package adapters

import (
	"FadyGamilM/banking/domain/core"
	"FadyGamilM/banking/domain/ports"
)

// the business-logic depends on the secondry ports (customer repo interface)
type CustomerServiceBusinessLogic struct {
	Customer_repo ports.CustomerRepository
}

// implement the primary port
func (cs CustomerServiceBusinessLogic) GetAllCustomers() ([]core.Customer, error) {
	// the repo method returns the list of customers and error
	return cs.Customer_repo.GetAll()
}

// Factory Method Pattern
func NewCustomerServiceBusinessLogic(Repo ports.CustomerRepository) CustomerServiceBusinessLogic {
	return CustomerServiceBusinessLogic{
		Customer_repo: Repo,
	}
}
