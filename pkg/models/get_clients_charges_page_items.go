package models

// GetClientsChargesPageItems struct for GetClientsChargesPageItems
type GetClientsChargesPageItems struct {
	Id                    int32                          `json:"id,omitempty"`
	ClientId              int32                          `json:"clientId,omitempty"`
	ChargeId              int32                          `json:"chargeId,omitempty"`
	Name                  string                         `json:"name,omitempty"`
	ChargeTimeType        GetClientChargeTimeType        `json:"chargeTimeType,omitempty"`
	DueDate               string                         `json:"dueDate,omitempty"`
	ChargeCalculationType GetClientChargeCalculationType `json:"chargeCalculationType,omitempty"`
	Currency              GetClientChargeCurrency        `json:"currency,omitempty"`
	Amount                float32                        `json:"amount,omitempty"`
	AmountPaid            float32                        `json:"amountPaid,omitempty"`
	AmountWaived          float32                        `json:"amountWaived,omitempty"`
	AmountWrittenOff      float32                        `json:"amountWrittenOff,omitempty"`
	AmountOutstanding     float32                        `json:"amountOutstanding,omitempty"`
	Penalty               bool                           `json:"penalty,omitempty"`
	IsActive              bool                           `json:"isActive,omitempty"`
	IsPaid                bool                           `json:"isPaid,omitempty"`
	IsWaived              bool                           `json:"isWaived,omitempty"`
}
