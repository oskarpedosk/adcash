package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

func ConnectDB(dsn string) (*DB, error) {
	db, err := NewDatabase(dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot ping database!")
	}

	dbConn.SQL = db

	return dbConn, nil
}

func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Initiate database with dummy data
func (m *DB) InitDB(dsn string) error {
	_, err := m.SQL.Query("select * from loan_applications;")
	if err != nil {
		err := m.PopulateDB("tables.sql")
		if err != nil {
			fmt.Println(err)
			return err
		}
		err = m.PopulateDB("dummy.sql")
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *DB) PopulateDB(filename string) error {
	data, err := os.ReadFile("../db/" + filename)
	if err != nil {
		return err
	}
	_, err = m.SQL.Exec(string(data))
	if err != nil {
		return err
	}
	return nil
}
