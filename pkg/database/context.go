package database

import "github.com/go-pg/pg/v10"

var db *pg.DB // go-pg database object

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
}

// GetConnection will return the connection object to run queries on
func GetConnection() *pg.DB {
	return db
}
