package handlers

import (
	"adcash/driver"
	"adcash/models"
	"adcash/repository"
	"adcash/repository/dbrepo"
	"fmt"
	"net/http"
	"strconv"
)

var Repo *Repository

type Repository struct {
	DB repository.DatabaseRepo
}

func NewRepo(db *driver.DB) *Repository {
	return &Repository{
		DB: dbrepo.NewSQLiteRepo(db.SQL),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Instructions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "Hello, World!")
}

func (m *Repository) Loans(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (m *Repository) PostApply(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	personalID, err := strconv.Atoi(r.FormValue("personal_id"))
	if err != nil {
		fmt.Fprintln(w, "Personal ID must be a valid number")
		return
	}

	name := r.FormValue("name")
	if err != nil {
		fmt.Fprintln(w, "Invalid name")
		return
	}

	amount, err := strconv.Atoi(r.FormValue("amount"))
	if err != nil {
		fmt.Fprintln(w, "Loan amount must be a valid number")
		return
	}

	term, err := strconv.Atoi(r.FormValue("term"))
	if err != nil {
		fmt.Fprintln(w, "Loan term must be a valid number")
		return
	}
	
	blacklisted, err := m.DB.IsBlacklisted(personalID)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	if blacklisted {
		fmt.Fprintln(w, "Loan application rejected, please contact Adcash")
		return
	}

	loan := models.Loan{
		PersonalID: personalID,
		Name:       name,
		Amount:     amount,
		Term:       term,
	}

	err = m.DB.NewLoanApplication(loan)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	
	fmt.Fprintln(w, loan)
}
