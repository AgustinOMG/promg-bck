package models

type Quote struct {
	CID           string `json:"cid,omitempty" validate:"required"`
	Folio         string `json:"folio,omitempty" validate:"required"`
	Client        Client `json:"client,omitempty" validate:"required"`
	Terms         string `json:"terms,omitempty" validate:"required"`
	Date          string `json:"date,omitempty" validate:"required"`
	Elements      []Item `json:"elements,omitempty" validate:"required"`
	Serie         string `json:"serie,omitempty" validate:"required"`
	Heading       string `json:"heading,omitempty" validate:"required"`
	Description   string `json:"description,omitempty" validate:"required"`
	PayMethod     string `json:"paymethod,omitempty" validate:"required"`
	Currency      string `json:"currency,omitempty" validate:"required"`
	delivery      string `json:"delivery,omitempty" validate:"required"`
	DeliveryUnits string `json:"deliveryunits,omitempty" validate:"required"`
	Title         string `json:"title,omitempty" validate:"required"`
	SubTotal      string `json:"subtotal,omitempty" validate:"required"`
	IVA           string `json:"iva,omitempty" validate:"required"`
	Total         string `json:"total,omitempty" validate:"required"`
	Comments      string `json:"comments,omitempty" validate:"required"`
	Seller        string `json:"seller,omitempty" validate:"required"`
	Regards       string `json:"regards,omitempty" validate:"required"`
}
