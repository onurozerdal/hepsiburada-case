package infrastructure

import (
	"database/sql"
)

type SqlHandler struct {
}

var Connection *sql.DB

func NewSqlHandler(connection *sql.DB) *SqlHandler {
	Connection = connection
	return &SqlHandler{}
}

func (db *SqlHandler) QueryRow(query string, args ...interface{}) *sql.Row {
	return Connection.QueryRow(query, args...)
}

type SqlRow struct {
	Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)
	if err != nil {
		return err
	}

	return nil
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	return r.Rows.Close()
}
