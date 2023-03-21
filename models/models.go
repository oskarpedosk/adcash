package models

import "time"

type Loan struct {
	ID         int
	PersonalID int
	Name       string
	Amount     int
	Term       int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
