package db

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	once       sync.Once
	connection *sql.DB
)

// GetConnection only return 1 connection (singletion)
func GetConnection() *sql.DB {
	once.Do(func() {
		var err error
		connStr := "user=admin password=admin dbname=test sslmode=disable"
		connection, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
		}
	})
	return connection
}
