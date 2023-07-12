package ports

import (
	"FadyGamilM/banking/domain/core"
)

// secondery port
type CustomerRepository interface {
	GetAll() ([]core.Customer, error)
}
