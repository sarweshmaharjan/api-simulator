package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type MyDB struct {
	*sql.DB
}

func PrimayConnection() *sql.DB {
	connStr := "postgres://admin:admin@localhost:5432/simulator?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}

func Init(db *sql.DB) error {

	existQuery := `
	SELECT EXISTS (
		SELECT FROM
			information_schema.tables
		WHERE
			table_schema = 'api_simulator_v1' AND
			table_name   = 'simulator_responses'
	);`
	var exists bool
	err := db.QueryRow(existQuery).Scan(&exists)
	if err != nil {
		panic(err)
	}

	if !exists {
		createTableSQL := `
		CREATE SCHEMA IF NOT EXISTS api_simulator_v1;
		CREATE TABLE api_simulator_v1.simulator_responses (
			id serial primary key,
			response json,
			created_at timestamp not null default NOW()
		);
		`
		_, err := db.Exec(createTableSQL)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
