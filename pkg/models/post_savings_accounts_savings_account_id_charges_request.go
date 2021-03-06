package models

// PostSavingsAccountsSavingsAccountIdChargesRequest PostSavingsAccountsSavingsAccountIdChargesRequest
type PostSavingsAccountsSavingsAccountIdChargesRequest struct {
	ChargeId   int32   `json:"chargeId,omitempty"`
	Locale     string  `json:"locale,omitempty"`
	Amount     float32 `json:"amount,omitempty"`
	DateFormat string  `json:"dateFormat,omitempty"`
	DueDate    string  `json:"dueDate,omitempty"`
}
