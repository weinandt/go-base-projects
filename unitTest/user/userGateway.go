package user

import (
	"database/sql"
	"fmt"
)

type UserGateway interface {
	createUser(name string) (*User, error)
}

type PostgresUserGateway struct {
	db *sql.DB
}

func NewPostgresUserGateway(db *sql.DB) *PostgresUserGateway {
	return &PostgresUserGateway{
		db: db,
	}
}

func (pug *PostgresUserGateway) createUser(name string) (*User, error) {
	var id string
	err := pug.db.QueryRow("INSERT INTO users (name) VALUES ($1) RETURNING id", name).Scan(&id)
	if err != nil {
		fmt.Println("problem quering postgres", err)

		return nil, err
	}

	return &User{
		id:   id,
		name: name,
	}, nil
}
