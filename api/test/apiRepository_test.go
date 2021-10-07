package test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/onurozerdal/hepsiburada-case/api/infrastructure"
	"github.com/onurozerdal/hepsiburada-case/api/repository"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

type SqlHandler struct {
	connection *sql.DB
}

func TestBrowsingHistories(t *testing.T) {
	db, mock, err := sqlmock.New()

	sqlHandler := infrastructure.NewSqlHandler(db)
	repo := repository.NewApiRepository(*sqlHandler)

	userId := "user-100"

	query := "select h.product_id as productId from history h where h.user_id ='user-100' order by h.date desc limit 10"

	rows := sqlmock.NewRows([]string{"productId"}).
		AddRow("product-100")

	mock.ExpectQuery(query).WillReturnRows(rows)

	r, err := repo.BrowsingHistories(userId)
	assert.NotEmpty(t, r)
	assert.NoError(t, err)
}

func TestCheckBrowsingHistories(t *testing.T) {
	db, mock, err := sqlmock.New()

	sqlHandler := infrastructure.NewSqlHandler(db)
	repo := repository.NewApiRepository(*sqlHandler)

	userId := "user-100"

	query := "select count(*) as count from history h where h.user_id = 'user-100'"

	rows := sqlmock.NewRows([]string{"count"}).
		AddRow("1")

	mock.ExpectQuery(query).WillReturnRows(rows)

	r := repo.CheckBrowsingHistories(userId)
	assert.True(t, r)
	assert.NoError(t, err)
}

func TestBestseller10Products(t *testing.T) {
	db, mock, err := sqlmock.New()

	sqlHandler := infrastructure.NewSqlHandler(db)
	repo := repository.NewApiRepository(*sqlHandler)

	query := "select b.product_id as productId from bestseller b order by b.quantity desc limit 10"

	rows := sqlmock.NewRows([]string{"productId"}).
		AddRow("product-100")

	mock.ExpectQuery(query).WillReturnRows(rows)

	r, err := repo.Bestseller10Products()
	assert.NotEmpty(t, r)
	assert.NoError(t, err)
}
