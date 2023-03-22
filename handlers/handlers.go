package handlers

import (
	"adcash/driver"
	"adcash/models"
	"adcash/repository"
	"adcash/repository/dbrepo"
	"fmt"
	"math"
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

func (m *Repository) Loans(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintln(w, "Error parsing form")
		return
	}

	personalID, err := strconv.Atoi(r.FormValue("personal_id"))
	if err != nil {
		fmt.Fprintln(w, "Personal ID must be a valid number")
		return
	}

	loans, err := m.DB.GetLoans(personalID)
	if err != nil {
		fmt.Fprintln(w, "Error receiving user loans")
		return
	}

	if len(loans) > 0 {
		for _, loan := range loans {
			monthlyPayment := (float64(loan.Amount) * 0.05) / (1 - math.Pow(1 + 0.05, -float64(loan.Term)))
			loanMsg := fmt.Sprintf("Personal ID: %d\nName: %s\nAmount: %d\nTerm: %d months\nMonthly payment:  %.2f\n---------", loan.PersonalID, loan.Name, loan.Amount, loan.Term, monthlyPayment)
			fmt.Fprintln(w, loanMsg)
		}
	} else {
		fmt.Fprintln(w, "User doesn't have any loans yet")
	}
}

func (m *Repository) Apply(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintln(w, "Error parsing form")
		return
	}

	personalID, err := strconv.Atoi(r.FormValue("personal_id"))
	if err != nil {
		fmt.Fprintln(w, "Personal ID must be a valid number")
		return
	}

	name := r.FormValue("name")
	if err != nil || name == "" {
		fmt.Fprintln(w, "Please insert a valid name")
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
		fmt.Fprintln(w, "Loan application rejected, please contact Adcash support")
		return
	}

	loanCount, err := m.DB.LoanCountWithin24Hours(personalID)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	if loanCount > 0 {
		fmt.Fprintln(w, "Loan application already submitted, please try again in 24 hours")
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

	monthlyPayment := (float64(amount) * 0.05) / (1 - math.Pow(1 + 0.05, -float64(term)))

	successMsg := fmt.Sprintf("Thanky you %s! Loan application has been submitted.\nLoan amount: %d\nTerm: %d months\nMonthly payment: %.2f\n---------", name, amount, term, monthlyPayment)
	fmt.Fprintln(w, successMsg)
}
