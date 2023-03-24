package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func CreatePostgresDatabase() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "postgres", "test_db")
	return sql.Open("postgres", psqlconn)
}
