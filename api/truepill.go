package api

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sarweshmaharjan/api-simulator.git/common"
	"github.com/sarweshmaharjan/api-simulator.git/storage"
	"github.com/sarweshmaharjan/api-simulator.git/types"
)

const (
	RequestID                 = "5a9a79a446f5d554212346"
	PrescriptionToken         = "z3q2jrs312356"
	PatientToken              = "4526d90a12356"
	InsuranceToken            = "skhsyq83rkd3uht691235"
	DirectTransferID          = "c8a73a55dc61325"
	TransferPrescriptionToken = "z3q2jr2312365"
	DirectTransferToken       = "dff0a9812412635"
	CopayPrescriptionToken    = "an7hj7gp63yjh6gpw1235"
	CopayRequestToken         = "ef6p8y6wz7x3hp6ng1235"
)

func GetDirectTransfer(ctx *gin.Context) {
	directTransfer := &types.DirectTransferResponse{
		RequestID: RequestID,
		Status:    "success",
		TimeStamp: float64(time.Now().Unix()),
		Details: &types.DirectTransferDetails{
			DirectTransferID: DirectTransferID,
		},
	}
	sql := storage.PrimayConnection()
	storage.Init(sql)
	// Insert a JSON value into the "simulator" table
	insertSQL := `
		INSERT INTO api_simulator_v1.direct_transfer_response (requestId,status,detailsTransferId) VALUES ($1,$2,$3);
	`
	_, err := sql.Exec(insertSQL, directTransfer.RequestID, directTransfer.Status, directTransfer.Details.DirectTransferID)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusCreated, directTransfer)
	go func() {
		SendNotifyRxWebhook(ctx)
		time.Sleep(5 * time.Second)
		SendDirectTransferWebhook(ctx)
	}()
}

func SendNotifyRxWebhook(ctx *gin.Context) {
	notifyRx := &types.NotifyRx{
		PrescriptionToken:         PrescriptionToken,
		PatientToken:              PatientToken,
		TransferPrescriptionToken: TransferPrescriptionToken,
		MedicationName:            "BONJESTA 20-20 MG",
		Prescriber:                "Dr. Strange",
		Location:                  "Hayward, CA",
	}

	notifyRxMap := map[string]interface{}{
		"prescription_token":          notifyRx.PrescriptionToken,
		"patient_token":               notifyRx.PatientToken,
		"transfer_prescription_token": notifyRx.TransferPrescriptionToken,
		"medication_name":             notifyRx.MedicationName,
		"prescriber":                  notifyRx.Prescriber,
		"location":                    notifyRx.Location,
	}
	callbackResponse := &types.CallbackRequest{
		Timestamp:    float64(time.Now().Unix()),
		CallbackType: "NOTIFY_RX",
		Details:      notifyRxMap,
	}

	webHookURL := "http://localhost:8888/api/v1/truepill/callback"

	callbackJson := common.ToJSON(callbackResponse)
	resp, err := http.Post(webHookURL, "application/json", bytes.NewBuffer(callbackJson))
	if err != nil {
		return
	}
	defer resp.Body.Close()

}

func SendDirectTransferWebhook(ctx *gin.Context) {
	notifyRx := &types.DirectTransferWebhook{
		PrescriptionToken:   PrescriptionToken,
		PatientToken:        PatientToken,
		Metadata:            TransferPrescriptionToken,
		DirectTransferToken: DirectTransferToken,
		Message:             "Direct Transfer Accepted.",
	}

	notifyRxMap := map[string]interface{}{
		"prescription_token":    notifyRx.PrescriptionToken,
		"patient_token":         notifyRx.PatientToken,
		"metadata":              notifyRx.Metadata,
		"message":               notifyRx.Message,
		"direct_transfer_token": notifyRx.DirectTransferToken,
	}
	callbackResponse := &types.CallbackRequest{
		RequestID:    RequestID,
		Status:       "success",
		Timestamp:    float64(time.Now().Unix()),
		CallbackType: "DIRECT_TRANSFER",
		Details:      notifyRxMap,
	}

	webHookURL := "http://localhost:8888/api/v1/truepill/callback"

	callbackJson := common.ToJSON(callbackResponse)
	resp, err := http.Post(webHookURL, "application/json", bytes.NewBuffer(callbackJson))
	if err != nil {
		return
	}
	defer resp.Body.Close()
}

func GetPrescriptionDetails(ctx *gin.Context) {
	prescription := &types.Prescription{
		DateWritten:            "2023-09-17T17:35:40.083Z",
		ExpirationDate:         "2023-12-30T17:35:40.083Z",
		DaysSupply:             "84",
		IsRefill:               0,
		LastFilledDate:         "2023-08-30T17:35:40.083Z",
		MedicationSIG:          "Take one Tablet by mouth at the same time daily",
		NumberOfRefillsAllowed: 0,
		Origin:                 "Electronic",
		PrescribedBrandName:    "BONJESTA 20-20 MG",
		PrescribedDrugStrength: "20-20 MG",
		PrescribedGenericName:  "TAKE 1 TABLET DAILY|TOME 1 TABLETA AL DIA",
		PrescribedNDC:          "55494012060",
		PrescribedQuantity:     84,
		PrescribedWrittenName:  "Lutera-28 Tablet",
		Prescriber:             "Dr. Strange",
		PrescriberAddress: &types.PrescriberAddress{
			Name:    "Dr. Strange",
			Company: "",
			Street1: "12345 Avengers Rd",
			Street2: "",
			City:    "San Francisco",
			State:   "CA",
			Zip:     "94402",
			Country: "US",
			Phone:   "(800) 888-8888",
			Email:   "dr.strange@avengers.com",
		},
		PrescriberNPI:         "1639349256",
		PrescriberOrderNumber: "11111111",
		QuantityRemaining:     0,
		RefillsRemaining:      0,
		RxNumber:              "1966785",
		PrescriptionToken:     PrescriptionToken,
		Notes:                 "prescription notes",
		ICDCodes: &types.ICDCodes{
			ICD10: []string{"Z30.09"},
			ICD9:  []string{"N94.6"},
		},
		IsDAW:                  false,
		Fillable:               true,
		DEASchedule:            0,
		OriginalPrescribedNDC:  "43547041110",
		DateFilledUTC:          "2020-07-30T20:08:55.000Z",
		PrescribedQuantityUnit: "EA",
	}

	resp := &types.PrescriptionResponse{
		Prescription: prescription,
	}

	ctx.JSON(http.StatusCreated, resp)
}

func GetInsuranceDetails(ctx *gin.Context) {
	insuranceResponse := &types.InsuranceResponse{
		RequestID: RequestID,
		Status:    "success",
		TimeStamp: float64(time.Now().Unix()),
		Details: &types.InsuranceDetails{
			InsuranceToken: InsuranceToken,
		},
	}

	ctx.JSON(http.StatusCreated, insuranceResponse)
}

func GetCopayRequest(ctx *gin.Context) {
	copayRxToken := make([]*types.CopayRequestPrescriptionToken, 0)
	copayPrescriptionToken := &types.CopayRequestPrescriptionToken{
		PrescriptionToken:      PrescriptionToken,
		CopayPrescriptionToken: CopayPrescriptionToken,
		Status:                 "pending",
	}
	copayRxToken = append(copayRxToken, copayPrescriptionToken)
	insuranceResponse := &types.CopayResponse{
		RequestID: RequestID,
		Status:    "success",
		TimeStamp: float64(time.Now().Unix()),
		Details: &types.CopayResponseDetails{
			CopayRequestToken:              CopayRequestToken,
			CopayRequestPrescriptionTokens: copayRxToken,
		},
	}

	ctx.JSON(http.StatusCreated, insuranceResponse)
}

// func SendCopayWebhook(ctx *gin.Context) {
// 	copayCallback := ""
// 	callbackResponse := &types.CallbackRequest{
// 		RequestID:    RequestID,
// 		Status:       "success",
// 		Timestamp:    float64(time.Now().Unix()),
// 		CallbackType: "COPAY",
// 		Details:      notifyRxMap,
// 	}

// 	webHookURL := "http://localhost:8888/api/v1/truepill/callback"

// 	callbackJson := common.ToJSON(callbackResponse)
// 	resp, err := http.Post(webHookURL, "application/json", bytes.NewBuffer(callbackJson))
// 	if err != nil {
// 		return
// 	}
// 	defer resp.Body.Close()
// }
