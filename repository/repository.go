package repository

import "adcash/models"

type DatabaseRepo interface {
	NewLoanApplication(models.Loan) error
	IsBlacklisted(personalID int) (bool, error)
}
