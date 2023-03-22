# Consumer Loan App

The Consumer Loan App is a headless web service that allows users to view their loans and apply for a new loan.

## Installation

Before running the Consumer Loan App, ensure you have the following dependencies installed:

* [chi router](https://github.com/go-chi/chi)
go get -u github.com/go-chi/chi/v5

* [go-sqlite3](https://github.com/mattn/go-sqlite3)
go get -u github.com/mattn/go-sqlite3


## Usage

1. Navigate to the `cmd` directory of the project using the terminal.
```
cd cmd
```
2. Run the application using the following command:
```
go run .
```
3. Use Postmaster, Insomnia or another terminal window to make GET and POST requests to the application:
* To view a user's loans, run the following command in another terminal window:
    ```
    curl -X GET localhost:8080/api/loans -d 'personal_id=2'
    ```
* To apply for a loan, run the following command in another terminal window:
    ```
    curl -X POST localhost:8080/api/apply -d 'personal_id=5&name=Super+Mario&amount=5000&term=24'
    ```

## Dependencies

* Built using Go version 1.20
* Uses [sqlite3](https://github.com/mattn/go-sqlite3)
* Uses the [chi router](https://github.com/go-chi/chi)
