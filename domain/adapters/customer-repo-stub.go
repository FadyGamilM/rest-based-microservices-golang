package adapters

import (
	"FadyGamilM/banking/domain/core"
	"time"
)

// stub adapter type
type CustomerRepoStub struct {
	Customers []core.Customer
}

// Implement the secondery port
func (cr CustomerRepoStub) GetAll() ([]core.Customer, error) {
	return cr.Customers, nil
}

// Factory Method pattern
func NewCustomerRepoStub() CustomerRepoStub {
	return CustomerRepoStub{Customers: []core.Customer{
		{
			ID:        1,
			Name:      "Fady Gamil",
			BirthDate: time.Date(1999, time.March, 7, 0, 0, 0, 0, time.UTC),
			City:      "Cairo",
			ZipCode:   "11566",
		},
		{
			ID:        2,
			Name:      "Ahmed Rushdi",
			BirthDate: time.Date(2000, time.January, 3, 0, 0, 0, 0, time.UTC),
			City:      "Cairo",
			ZipCode:   "11689",
		},
	}}
}
