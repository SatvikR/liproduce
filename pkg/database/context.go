package database

import (
	"context"
	"fmt"

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
	// TODO switch this to environment variables
	db = pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "postgres",
		Database: "liproduce",
	})

	db.AddQueryHook(dbLogger{})
}

// GetConnection will return the connection object to run queries on
func GetConnection() *pg.DB {
	return db
}
