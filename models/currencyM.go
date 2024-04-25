package models

type Currency struct {
	MXN string `json:"mxn,omitempty" validate:"required"`
	USD string `json:"USD,omitempty" validate:"required"`
	EUR string `json:"eur,omitempty" validate:"required"`
}
