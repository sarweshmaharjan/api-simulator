package api

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {
	router.POST("/v1/direct_transfer", GetDirectTransfer)
	router.GET("/v1/prescription/:prescription_token/fill/:claim_id/claims", GetClaimsDetails)
	router.GET("/v1/prescription/:prescription_token", GetPrescriptionDetails)
	router.POST("/v1/insurance", GetInsuranceDetails)
	router.POST("/v1/copay_request", GetCopayRequest)
	router.POST("/v1/fill_request", GetFillRequest)
}
