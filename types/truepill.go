package types

type DirectTransferResponse struct {
	RequestID string                 `json:"request_id" bson:"requestId"`
	Status    string                 `json:"status" bson:"status"`
	TimeStamp float64                `json:"timestamp" bson:"createdAt"`
	Details   *DirectTransferDetails `json:"details" bson:"details"`
}

type DirectTransferDetails struct {
	DirectTransferID string `json:"direct_transfer_id" bson:"directTransferId"`
}

type CallbackRequest struct {
	RequestID    string                 `json:"request_id" bson:"requestId"`
	Timestamp    float64                `json:"timestamp" bson:"timestamp"`
	Status       string                 `json:"status" bson:"status"`
	CallbackType string                 `json:"callback_type" bson:"callbackType"`
	Details      map[string]interface{} `json:"details" bson:"details"`
}

type NotifyRx struct {
	PrescriptionToken         string `json:"prescription_token" bson:"prescriptionToken"`
	PatientToken              string `json:"patient_token" bson:"patientToken"`
	TransferPrescriptionToken string `json:"transfer_prescription_token" bson:"transferPrescriptionToken"`
	MedicationName            string `json:"medication_name" bson:"medicationName"`
	Prescriber                string `json:"prescriber" bson:"prescriber"`
	Location                  string `json:"location" bson:"location"`
}

type DirectTransferWebhook struct {
	PrescriptionToken   string `json:"prescription_token" bson:"prescriptionToken"`
	PatientToken        string `json:"patient_token" bson:"patientToken"`
	Metadata            string `json:"metadata" bson:"metadata"`
	Message             string `json:"message" bson:"message"`
	DirectTransferToken string `json:"direct_transfer_token" bson:"directTransferToken"`
}

type PrescriptionResponse struct {
	Prescription *Prescription `json:"prescription" bson:"prescription"`
}
type Prescription struct {
	PrescriptionToken      string             `json:"prescription_token" bson:"prescriptionToken"`
	RxNumber               string             `json:"rx_number" bson:"rxNumber"`
	Fillable               bool               `json:"fillable" bson:"fillable"`
	PrescribedBrandName    string             `json:"prescribed_brand_name" bson:"prescribedBrandName"`
	PrescribedGenericName  string             `json:"prescribed_generic_name" bson:"prescribedGenericName"`
	PrescribedNDC          string             `json:"prescribed_ndc" bson:"prescribedNDC"`
	PrescribedQuantity     int                `json:"prescribed_quantity" bson:"prescribedQuantity"`
	PrescribedWrittenName  string             `json:"prescribed_written_name" bson:"prescribedWrittenName"`
	PrescribedDrugStrength string             `json:"prescribed_drug_strength" bson:"prescribedDrugStrength"`
	MedicationSIG          string             `json:"medication_sig" bson:"medicationSIG"`
	DaysSupply             string             `json:"days_supply" bson:"daysSupply"`
	DateWritten            string             `json:"date_written" bson:"dateWritten"`
	ExpirationDate         string             `json:"expiration_date" bson:"expirationDate"`
	RefillsRemaining       int                `json:"refills_remaining" bson:"refillsRemaining"`
	QuantityRemaining      int                `json:"quantity_remaining" bson:"quantityRemaining"`
	IsRefill               int               `json:"is_refill" bson:"isRefill"`
	LastFilledDate         string             `json:"last_filled_date" bson:"lastFilledDate"`
	DateFilledUTC          string             `json:"date_filled_utc" bson:"dateFilledUTC"`
	NumberOfRefillsAllowed int                `json:"number_of_refills_allowed" bson:"numberOfRefillsAllowed"`
	Prescriber             string             `json:"prescriber" bson:"prescriber"`
	PrescriberNPI          string             `json:"prescriber_npi" bson:"prescriberNPI"`
	PrescriberOrderNumber  string             `json:"prescriber_order_number" bson:"prescriberOrderNumber"`
	Notes                  string             `json:"notes" bson:"notes"`
	ICDCodes               *ICDCodes          `json:"icd_codes" bson:"icdCodes"`
	Origin                 string             `json:"origin" bson:"origin"`
	IsDAW                  bool               `json:"is_daw" bson:"isDAW"`
	DEASchedule            int                `json:"dea_schedule" bson:"deaSchedule"`
	OriginalPrescribedNDC  string             `json:"original_prescribed_ndc" bson:"originalPrescribedNDC"`
	PrescribedQuantityUnit string             `json:"prescribed_quantity_unit" bson:"prescribedQuantityUnit"`
	PrescriberAddress      *PrescriberAddress `json:"prescriber_address" bson:"prescriberAddress"`
}

type PrescriberAddress struct {
	Name    string `json:"name" bson:"name"`
	Company string `json:"company" bson:"company"`
	Street1 string `json:"street1" bson:"street1"`
	Street2 string `json:"street2" bson:"street2"`
	City    string `json:"city" bson:"city"`
	State   string `json:"state" bson:"state"`
	Zip     string `json:"zip" bson:"zip"`
	Country string `json:"country" bson:"country"`
	Phone   string `json:"phone" bson:"phone"`
	Email   string `json:"email" bson:"email"`
}

type ICDCodes struct {
	ICD10 []string `json:"icd10" bson:"icd10"`
	ICD9  []string `json:"icd9" bson:"icd9"`
}

type InsuranceResponse struct {
	RequestID string            `json:"request_id" bson:"requestId"`
	Status    string            `json:"status" bson:"status"`
	TimeStamp float64           `json:"timestamp" bson:"timestamp"`
	Details   *InsuranceDetails `json:"details" bson:"details"`
}
// InsuranceDetails represents the data structure for response details for create insurance from truepill
type InsuranceDetails struct {
	InsuranceToken string `json:"insurance_token" bson:"insuranceToken"`
}

type CopayResponse struct {
	RequestID string                `json:"request_id" bson:"requestId"`
	Status    string                `json:"status" bson:"status"`
	TimeStamp float64               `json:"timestamp" bson:"timestamp"`
	Details   *CopayResponseDetails `json:"details" bson:"details"`
}

type CopayResponseDetails struct {
	CopayRequestToken              string                           `json:"copay_request_token" bson:"copayRequestToken" `
	CopayRequestPrescriptionTokens []*CopayRequestPrescriptionToken `json:"copay_request_prescription_tokens" bson:"copayRequestPrescriptionTokens"`
}

// CopayRequestPrescriptionToken represents the tokens of copay request
type CopayRequestPrescriptionToken struct {
	PrescriptionToken      string `json:"prescription_token" bson:"prescriptionToken"`
	CopayPrescriptionToken string `json:"copay_request_prescription_token" bson:"copayRequestPrescriptionToken"`
	Status                 string `json:"status" bson:"status"`
}