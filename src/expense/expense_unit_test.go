package expense_test

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/herokh/assessment/expense"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var sqlMock sqlmock.Sqlmock
var sqlDb *sql.DB

func TestGetAll(t *testing.T) {
	teardown := setupTest(t)
	defer teardown(t)

	rows := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).AddRow(
		1, "title", 100, "note", pq.Array([]string{"tag1", "tag2"}))
	sqlMock.ExpectPrepare("SELECT id, title, amount, note, tags FROM expenses").ExpectQuery().WillReturnRows(rows)

	_, err := expense.GetExpenses(sqlDb)

	assert.Nil(t, err)
}
func TestGetOne(t *testing.T) {
	teardown := setupTest(t)
	defer teardown(t)

	rows := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).AddRow(
		1, "title", 100, "note", pq.Array([]string{"tag1", "tag2"}))
	sqlMock.ExpectPrepare(regexp.QuoteMeta("SELECT id, title, amount, note, tags FROM expenses WHERE id = $1")).ExpectQuery().WithArgs("1").WillReturnRows(rows)

	_, err := expense.GetExpense(sqlDb, "1")

	assert.Nil(t, err)
}
func TestCreate(t *testing.T) {
	teardown := setupTest(t)
	defer teardown(t)

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	sqlMock.ExpectQuery(regexp.QuoteMeta("INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3, $4) RETURNING id")).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(rows)

	err := expense.CreateExpense(sqlDb, &expense.Expense{})

	assert.Nil(t, err)
}
func TestUpdate(t *testing.T) {
	teardown := setupTest(t)
	defer teardown(t)

	rows := sqlmock.NewRows([]string{"id", "title", "amount", "note", "tags"}).AddRow(
		1, "title", 100, "note", pq.Array([]string{"tag1", "tag2"}))
	sqlMock.ExpectPrepare(regexp.QuoteMeta("SELECT id, title, amount, note, tags FROM expenses WHERE id = $1")).ExpectQuery().WithArgs("1").WillReturnRows(rows)

	sqlMock.ExpectExec(regexp.QuoteMeta("UPDATE expenses SET title = $1, amount = $2, note = $3, tags = $4 WHERE id = $5")).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))

	err := expense.UpdateExpense(sqlDb, "1", &expense.Expense{})

	assert.Nil(t, err)
}

func setupTest(tb testing.TB) func(tb testing.TB) {
	db, mock, err := sqlmock.New()
	if err != nil {
		tb.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlDb = db
	sqlMock = mock

	return func(tb testing.TB) {
		db.Close()
	}
}
