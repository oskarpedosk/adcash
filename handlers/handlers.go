package handlers

import (
	"adcash/driver"
	"adcash/helpers"
	"adcash/models"
	"adcash/repository"
	"adcash/repository/dbrepo"
	"encoding/json"
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
		helpers.ServerError(w, err)
		return
	}

	if !r.URL.Query().Has("personal_id") {
		helpers.ClientError(w, "please enter personal id", 400)
		return
	}

	personalID, err := strconv.Atoi(r.URL.Query().Get("personal_id"))
	if err != nil {
		helpers.ClientError(w, "personal id must be a valid number", 400)
		return
	}

	loans, err := m.DB.GetLoans(personalID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	var loansData []map[string]interface{}

	if len(loans) > 0 {
		for _, loan := range loans {
			monthlyPayment := math.Round((float64(loan.Amount)*0.05)/(1-math.Pow(1+0.05, -float64(loan.Term)))*100) / 100
			data := map[string]interface{}{
				"personal_id":     loan.PersonalID,
				"name":            loan.Name,
				"amount":          loan.Amount,
				"term":            loan.Term,
				"monthly_payment": monthlyPayment,
			}

			loansData = append(loansData, data)
		}
	} else {
		helpers.ClientError(w, "user doesnt have any loans yet", 400)
		return
	}

	jsonData, err := json.Marshal(loansData)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jsonData)
}

func (m *Repository) Apply(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	personalID, err := strconv.Atoi(r.FormValue("personal_id"))
	if err != nil {
		helpers.ClientError(w, "personal id must be a valid number", 400)
		return
	}

	name := r.FormValue("name")
	if err != nil || name == "" {
		helpers.ClientError(w, "name cant be empty", 400)
		return
	}

	amount, err := strconv.Atoi(r.FormValue("amount"))
	if err != nil {
		helpers.ClientError(w, "loan amount must be a valid number", 400)
		return
	}

	term, err := strconv.Atoi(r.FormValue("term"))
	if err != nil {
		helpers.ClientError(w, "loan term must be a valid number", 400)
		return
	}

	blacklisted, err := m.DB.IsBlacklisted(personalID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	if blacklisted {
		helpers.ClientError(w, "loan application rejected, please contact Adcash support", 400)
		return
	}

	loanCount, err := m.DB.LoanCountWithin24Hours(personalID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	if loanCount > 0 {
		helpers.ClientError(w, "loan application already submitted, please try again in 24 hours", 400)
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
		helpers.ServerError(w, err)
		return
	}

	monthlyPayment := math.Round((float64(amount)*0.05)/(1-math.Pow(1+0.05, -float64(term)))*100) / 100

	data := map[string]interface{}{
		"personal_id":     loan.PersonalID,
		"name":            loan.Name,
		"amount":          loan.Amount,
		"term":            loan.Term,
		"monthly_payment": monthlyPayment,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(jsonData)
}
