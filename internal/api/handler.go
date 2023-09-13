package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	partner "github.com/sarweshmaharjan/api-simulator.git/data"
	connection "github.com/sarweshmaharjan/api-simulator.git/internal/db"
)

func GetDirectTransfer(ctx *gin.Context) {
	directTransfer := partner.GetDirectTransfer()
	sql := connection.PrimayConnection()

	exists, err := tableExists(sql, "simulator")
	if err != nil {
		log.Fatal(err)
	}

	if !exists {
		// Create the "simulator" table with a "response" column of type JSON
		createTableSQL := `
			CREATE TABLE simulator (
				id SERIAL PRIMARY KEY,
				response JSON
			);
		`
		_, err := sql.Exec(createTableSQL)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Insert a JSON value into the "simulator" table
	insertSQL := `
		INSERT INTO simulator (response) VALUES ($1);
	`
	responseJSON := `{"key": "value"}`
	_, err = sql.Exec(insertSQL, responseJSON)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusCreated, directTransfer)
}

// tableExists checks if a table with the given name exists in the database.
func tableExists(db *sql.DB, tableName string) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM information_schema.tables
			WHERE table_name = $1
		);
	`
	var exists bool
	err := db.QueryRow(query, tableName).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
