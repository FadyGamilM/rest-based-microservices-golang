package adapters

import (
	"FadyGamilM/banking/domain/core"
	"time"
)

// stub adapter type
type CustomerRepoStub struct {
	customers []core.Customer
}

// Implement the secondery port
func (c CustomerRepoStub) GetAll() ([]core.Customer, error) {
	return c.customers, nil
}

// Factory Method pattern
func NewCustomerRepoStub() CustomerRepoStub {
	return CustomerRepoStub{customers: []core.Customer{
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
