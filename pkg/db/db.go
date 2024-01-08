package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func InitializeDB(filePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
