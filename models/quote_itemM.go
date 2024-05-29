package models

type QuoteItem struct {
	CID         string `json:"cid,omitempty" validate:"required"`
	Code        string `json:"code,omitempty" validate:"required"`
	Description string `json:"description,omitempty" validate:"required"`
	MU          string `json:"mu,omitempty" validate:"required"`
	Amount      string `json:"amount,omitempty" validate:"required"`
	Price       string `json:"price,omitempty" validate:"required"`
	Id          string `bson:"_id,omitempty"`
}
