package types

type DirectTransferResponse struct {
	RequestID string                 `json:"request_id" bson:"requestId"`
	Status    string                 `json:"status" bson:"status"`
	TimeStamp float64                `json:"timestamp" bson:"timestamp"`
	Details   *DirectTransferDetails `json:"details" bson:"details"`
}

type DirectTransferDetails struct {
	DirectTransferID string `json:"direct_transfer_id" bson:"directTransferId"`
}
