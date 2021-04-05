package database

import (
	"context"
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
)

var db *pg.DB // go-pg database object

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	query, err := q.FormattedQuery()
	if err != nil {
		return err
	}
	fmt.Println(string(query))
	return nil
}

// CreateConnection connects to the postgres database.
// Stores db variables in this file
func CreateConnection() {

	db = pg.Connect(&pg.Options{
		Addr:     os.Getenv("DB_ADDR"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	})

	if os.Getenv("ENVIRONMENT") != "production" {
		db.AddQueryHook(dbLogger{})
	}
}

// GetConnection will return the connection object to run queries on
func GetConnection() *pg.DB {
	return db
}
