package api

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {
	router.POST("/v1/direct_transfer", GetDirectTransfer)
}
