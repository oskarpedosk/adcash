package main

import (
	"adcash/driver"
	"adcash/handlers"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

const portNumber = ":8080"
const dsn = "../db/adcash.db"

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	fmt.Printf("Starting application on http://localhost%s\n", portNumber)

	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (*driver.DB, error) {
	log.Println("Connecting to database...")
	db, err := driver.ConnectDB(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	log.Println("Connected to database!")

	db.InitDB(dsn)

	repo := handlers.NewRepo(db)
	handlers.NewHandlers(repo)

	return db, nil
}
