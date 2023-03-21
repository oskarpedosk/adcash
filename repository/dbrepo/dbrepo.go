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

func (m *sqliteDBRepo) IsBlacklisted(personalID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
	SELECT
		1
	FROM
		loan_applications
	WHERE
		personal_id = $1
	`
	id := 0

	row, err := m.DB.QueryContext(ctx, query, personalID)
	if err != nil {
		return false, err
	}

	defer row.Close()
	for row.Next() {
		err := row.Scan(
			&id,
		)
		if err != nil {
			return false, err
		}
	}

	if id != 0 {
		return true, nil
	}

	return false, nil
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
