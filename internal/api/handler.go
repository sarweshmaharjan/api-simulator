package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type DirectTransferResponse struct {
	RequestID string                 `json:"request_id" bson:"requestId"`
	Status    string                 `json:"status" bson:"status"`
	TimeStamp float64                `json:"timestamp" bson:"timestamp"`
	Details   *DirectTransferDetails `json:"details" bson:"details"`
}

type DirectTransferDetails struct {
	DirectTransferID string `json:"direct_transfer_id" bson:"directTransferId"`
}

func GetDirectTransfer(ctx *gin.Context) {
	directTransfer := DirectTransferResponse{
		RequestID: "5a9a79a446f5d554",
		Status:    "success",
		TimeStamp: float64(time.Now().Unix()),
		Details: &DirectTransferDetails{
			DirectTransferID: "c8a73a55dc",
		},
	}
	ctx.JSON(http.StatusCreated, directTransfer)
}
