package models

// GetClientsPageItems struct for GetClientsPageItems
type GetClientsPageItems struct {
	Id              int32                              `json:"id,omitempty"`
	OfficeId        int32                              `json:"officeId,omitempty"`
	OfficeName      string                             `json:"officeName,omitempty"`
	Type            GetClientsClientIdTransactionsType `json:"type,omitempty"`
	Date            string                             `json:"date,omitempty"`
	Currency        GetClientTransactionsCurrency      `json:"currency,omitempty"`
	Amount          float64                            `json:"amount,omitempty"`
	SubmittedOnDate string                             `json:"submittedOnDate,omitempty"`
	Reversed        bool                               `json:"reversed,omitempty"`
}
