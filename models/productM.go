package models

type Product struct {
	CID         string `json:"cid,omitempty" validate:"required"`
	Code        string `json:"code,omitempty" validate:"required"`
	Description string `json:"description,omitempty" validate:"required"`
	Currency    string `json:"currency,omitempty" validate:"required"`
	Price       string `json:"price,omitempty" validate:"required"`
	Items       []Item `json:"items,omitempty" validate:"required"`
	Id          string `bson:"_id,omitempty"`
}
