package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	partner "github.com/sarweshmaharjan/api-simulator.git/data"
	connection "github.com/sarweshmaharjan/api-simulator.git/internal/db"
	"github.com/sarweshmaharjan/api-simulator.git/pkg/utils"
)

func GetDirectTransfer(ctx *gin.Context) {
	directTransfer := partner.GetDirectTransfer()
	sql := connection.PrimayConnection()
	connection.Init(sql)
	// Insert a JSON value into the "simulator" table
	insertSQL := `
		INSERT INTO api_simulator_v1.simulator_responses (response) VALUES ($1);
	`
	responseJSON := utils.ToJSON(directTransfer)
	_, err := sql.Exec(insertSQL, responseJSON)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusCreated, directTransfer)
}
