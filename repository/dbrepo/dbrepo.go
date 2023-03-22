package dbrepo

import (
	"adcash/models"
	"adcash/repository"
	"context"
	"database/sql"
	"time"
)

type sqliteDBRepo struct {
	DB *sql.DB
}

func NewSQLiteRepo(conn *sql.DB) repository.DatabaseRepo {
	return &sqliteDBRepo{
		DB: conn,
	}
}

func (m *sqliteDBRepo) GetLoans(personalID int) ([]models.Loan, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
	SELECT
		*
	FROM
		loan_applications
	WHERE
		personal_id = $1
	`

	var loans []models.Loan

	rows, err := m.DB.QueryContext(ctx, query, personalID)
	if err != nil {
		return loans, err
	}
	defer rows.Close()
	for rows.Next() {
		var loan models.Loan
		err := rows.Scan(
			&loan.ID,
			&loan.PersonalID,
			&loan.Name,
			&loan.Amount,
			&loan.Term,
			&loan.CreatedAt,
			&loan.UpdatedAt,
		)
		if err != nil {
			return loans, err
		}
		loans = append(loans, loan)
	}

	return loans, nil
}

func (m *sqliteDBRepo) NewLoanApplication(loan models.Loan) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
	INSERT INTO
		loan_applications
		(personal_id, name, amount, term) 
	VALUES 
		($1, $2, $3, $4)`

	_, err := m.DB.ExecContext(ctx, stmt,
		loan.PersonalID,
		loan.Name,
		loan.Amount,
		loan.Term,
	)

	if err != nil {
		return err
	}

	return nil
}

func (m *sqliteDBRepo) LoanCountWithin24Hours(personalID int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
	SELECT
		COUNT(*)
	FROM
		loan_applications
	WHERE
		created_at >= datetime('now', '-24 hours')
	AND
		personal_id = $1
	`

	var count int

	err := m.DB.QueryRowContext(ctx, query, personalID).Scan(&count)
	if err != nil {
		return count, err
	}

	return count, nil
}

func (m *sqliteDBRepo) IsBlacklisted(personalID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
	SELECT
		1
	FROM
		blacklist
	WHERE
		personal_id = $1
	`
	blacklisted := 0

	row, err := m.DB.QueryContext(ctx, query, personalID)
	if err != nil {
		return true, err
	}

	defer row.Close()
	for row.Next() {
		err := row.Scan(&blacklisted)
		if err != nil {
			return true, err
		}
	}

	if blacklisted != 0 {
		return true, nil
	}

	return false, nil
}
