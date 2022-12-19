package expense

import (
	"database/sql"

	"github.com/lib/pq"
)

func GetExpenses(db *sql.DB) ([]Expense, error) {
	var expenses = []Expense{}
	stmt, err := db.Prepare("SELECT id, title, amount, note, tags FROM expenses")
	if err != nil {
		return expenses, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return expenses, err
	}
	for rows.Next() {
		var expense Expense
		err := rows.Scan(&expense.Id, &expense.Title, &expense.Amount, &expense.Note, pq.Array(&expense.Tags))
		if err != nil {
			return expenses, err
		}
		expenses = append(expenses, expense)
	}

	return expenses, nil
}

func GetExpense(db *sql.DB, id string) (Expense, error) {
	expense := Expense{}
	stmt, err := db.Prepare("SELECT id, title, amount, note, tags FROM expenses WHERE id = $1")
	if err != nil {
		return expense, err
	}
	row := stmt.QueryRow(id)

	err = row.Scan(&expense.Id, &expense.Title, &expense.Amount, &expense.Note, pq.Array(&expense.Tags))
	if err != nil {
		return expense, err
	}
	return expense, nil
}

func CreateExpense(db *sql.DB, expense *Expense) error {
	row := db.QueryRow("INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3, $4) RETURNING id",
		expense.Title, expense.Amount, expense.Note, pq.Array(expense.Tags))
	err := row.Scan(&expense.Id)

	if err != nil {
		return err
	}
	return nil
}

func UpdateExpense(db *sql.DB, id string, expense *Expense) error {
	_, err := GetExpense(db, id)
	if err != nil {
		return err
	}
	row, err := db.Exec("UPDATE expenses SET title = $1, amount = $2, note = $3, tags = $4 WHERE id = $5",
		expense.Title, expense.Amount, expense.Note, pq.Array(expense.Tags), id)
	if err != nil {
		return err
	}
	_, err = row.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
