package repository

import "adcash/models"

type DatabaseRepo interface {
	GetLoans(personalID int) ([]models.Loan, error)
	NewLoanApplication(models.Loan) error
	IsBlacklisted(personalID int) (bool, error)
	LoanCountWithin24Hours(personalID int) (int, error)
}
