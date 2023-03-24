package main

import (
	"context"

	"github.com/jackc/pgtype"
	"github.com/weinandt/go-base-projects/postgres/jsonb/postgres"
)

type Metadata struct {
	Location string
}

type User struct {
	ID       string
	Metadata Metadata
}

type CreateUserParams struct {
	metadata pgtype.JSONB
}

const createQuery = `
	INSERT INTO users ( metadata) VALUES ( $1 )`

func main() {
	db, err := postgres.CreateDBPool()
	if err != nil {
		panic(err)
	}

	// Inserting happy path
	var normalMetadata pgtype.JSONB
	normalMetadata.Set(Metadata{
		Location: "testLocation",
	})
	_ = db.QueryRowContext(context.Background(), createQuery, normalMetadata)

	// Inserting default struct
	var emptyMetadata pgtype.JSONB
	emptyMetadata.Set(Metadata{})
	_ = db.QueryRowContext(context.Background(), createQuery, emptyMetadata)

	// Inserting golang nil.
	_ = db.QueryRowContext(context.Background(), createQuery, nil)

	// Inserting array.
	var metadataArray pgtype.JSONB
	metadataArray.Set([]Metadata{
		Metadata{
			Location: "first one",
		},
		Metadata{
			Location: "second one",
		},
	})
	_ = db.QueryRowContext(context.Background(), createQuery, metadataArray)

	// Inserting array of pointers
	var metadataArrayPointers pgtype.JSONB
	metadataArrayPointers.Set([]*Metadata{
		&Metadata{
			Location: "first pointer one",
		},
		&Metadata{
			Location: "second pointer one",
		},
	})
	_ = db.QueryRowContext(context.Background(), createQuery, metadataArrayPointers)

	// Inserting JSON null. Note this is not the same as a postgres null
	var jsonNull pgtype.JSONB
	jsonNull.Set([]byte("null"))

	_ = db.QueryRowContext(context.Background(), createQuery, jsonNull)
}
