package models

type QuoteItem struct {
	Code        string  `json:"code,omitempty" validate:"required"`
	Description string  `json:"description,omitempty" validate:"required"`
	MU          string  `json:"mu,omitempty" validate:"required"`
	Amount      int     `json:"amount,omitempty" validate:"required"`
	Price       float32 `json:"price,omitempty" validate:"required"`
	Id          string  `bson:"_id,omitempty"`
}
