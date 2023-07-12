package core

import "time"

type Customer struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"date_of_birth"`
	City      string    `json:"city"`
	ZipCode   string    `json:"zip_code"`
}
