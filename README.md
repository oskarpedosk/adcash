# Consumer loan app
A headless web service with possibilites to view user loans and apply for a new loan.

## Instructions
Install necessary dependancies.
chi router:
```go get -u github.com/go-chi/chi/v5```
go-sqlite3
```go get -u github.com/mattn/go-sqlite3```

## Run application
Navigate to cmd directory
```cd cmd```
Run application
```go run .```

Make post requests to the application from *another* terminal window.
Apply for a loan
```curl -X POST localhost:8080/apply -d 'personal_id=5&name=Super+Mario&amount=5000&term=24'```
View user loans
```curl -X POST localhost:8080/loans -d 'personal_id=2'```