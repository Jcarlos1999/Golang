package main

import (
	"context"
	"log"
	"reflect"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	tutorial "tutorial.sqlc.dev/app/internal/db"
)

func run() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "host=localhost user=postgres password=990932 dbname=postgres port=5432 sslmode=verify-full")
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := tutorial.New(conn)

	// create an author
	insertedAuthor, err := queries.CreateAuthor(ctx, tutorial.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  pgtype.Text{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		return err
	}
	log.Println(insertedAuthor)

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
