package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sarweshmaharjan/api-simulator.git/config"
)

var db *sql.DB

func OpenDBConnection() {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Cfg.PostgresUser,
		config.Cfg.PostgresPassword,
		config.Cfg.PostgresHost,
		config.Cfg.PostgresPort,
		config.Cfg.PostgresDB,
	)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	db = conn
}

func Init() error {
	existQuery := `
	SELECT EXISTS (
		SELECT FROM
			information_schema.tables
		WHERE
			table_schema = 'api_simulator_v1' AND
			table_name   = 'direct_transfer_response'
	);`
	var exists bool
	err := db.QueryRow(existQuery).Scan(&exists)
	if err != nil {
		panic(err)
	}

	if !exists {
		createTableSQL := `
		CREATE SCHEMA IF NOT EXISTS api_simulator_v1;
		CREATE TABLE api_simulator_v1.direct_transfer_response (
			id serial primary key,
			createdAt timestamp not null default NOW(),
			requestId varchar(255),
			status varchar(50),
			detailsTransferId varchar(255)
		);
		`
		_, err := db.Exec(createTableSQL)
		if err != nil {
			panic(err)
		}
	}

	return nil
}

func ExecQuery(query string, args ...any) error {
	_, err := db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}
