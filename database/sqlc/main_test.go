package database

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:hyacinth@localhost:5432/banking?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	connection, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot Connect to Postgres: ", err)
	}

	log.Default()

	testQueries = New(connection)

	os.Exit(m.Run())
}
