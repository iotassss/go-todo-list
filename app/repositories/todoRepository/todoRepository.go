package todoRepository

import (
	"database/sql"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Repository struct {
	DB *sql.DB
}

func Connect(dataSourceName string) *Repository {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &Repository{DB: db}
}

func (r *Repository) Close() {
	r.DB.Close()
}

func (r *Repository) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return r.DB.Query(query, args...)
}

func (r *Repository) QueryRow(query string, args ...interface{}) *sql.Row {
	return r.DB.QueryRow(query, args...)
}

func (r *Repository) Exec(query string, args ...interface{}) (sql.Result, error) {
	return r.DB.Exec(query, args...)
}
