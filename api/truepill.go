package api

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sarweshmaharjan/api-simulator.git/common"
	"github.com/sarweshmaharjan/api-simulator.git/storage"
	"github.com/sarweshmaharjan/api-simulator.git/types"
)

func GetDirectTransfer(ctx *gin.Context) {
	directTransfer := &types.DirectTransferResponse{
		RequestID: "5a9a79a446f5d554",
		Status:    "success",
		TimeStamp: float64(time.Now().Unix()),
		Details: &types.DirectTransferDetails{
			DirectTransferID: "c8a73a55dc",
		},
	}
	sql := storage.PrimayConnection()
	storage.Init(sql)
	// Insert a JSON value into the "simulator" table
	insertSQL := `
		INSERT INTO api_simulator_v1.simulator_responses (response) VALUES ($1);
	`
	responseJSON := common.ToJSON(directTransfer)
	_, err := sql.Exec(insertSQL, responseJSON)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusCreated, directTransfer)
}
