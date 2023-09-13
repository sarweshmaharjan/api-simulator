package connection

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type MyDB struct {
	*sql.DB
}

func PrimayConnection() *sql.DB {
	user := "admin"
	password := "admin"
	dbname := "simulator"
	host := "localhost"
	port := 5432
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable", user, password, dbname, host, port)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func (db MyDB) Save(query string) {
	defer db.Close()
	db.Exec(query)
}
