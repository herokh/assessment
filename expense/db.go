package expense

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(databaseUrl string) {
	var err error
	db, err = sql.Open("postgres", databaseUrl)

	if err != nil {
		log.Fatal("Failed to connect database", err)
	}

	err = createExpensesTable()
	if err != nil {
		log.Fatal("Failed to create expenses table", err)
	}
}

func createExpensesTable() error {
	createTable := `
		CREATE TABLE IF NOT EXISTS expenses (
			id SERIAL PRIMARY KEY,
			title TEXT,
			amount FLOAT,
			note TEXT,
			tags TEXT[]
		);
	`
	_, err := db.Exec(createTable)
	if err != nil {
		return err
	}
	return nil
}
