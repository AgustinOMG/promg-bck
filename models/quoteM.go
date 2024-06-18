package models

type Quote struct {
	CID           string      `json:"cid,omitempty" validate:"required"`
	Folio         int         `json:"folio,omitempty" validate:"required"`
	Client        Client      `json:"client,omitempty" validate:"required"`
	Terms         string      `json:"terms,omitempty" validate:"required"`
	Date          string      `json:"date,omitempty" validate:"required"`
	Elements      []QuoteItem `json:"elements,omitempty" validate:"required"`
	Serie         string      `json:"serie,omitempty" validate:"required"`
	Description   string      `json:"description,omitempty" validate:"required"`
	PayMethod     string      `json:"paymethod,omitempty" validate:"required"`
	Currency      string      `json:"currency,omitempty" validate:"required"`
	Delivery      int         `json:"delivery,omitempty" validate:"required"`
	DeliveryUnits string      `json:"deliveryunit,omitempty" validate:"required"`
	Title         string      `json:"title,omitempty" validate:"required"`
	SubTotal      float64     `json:"subtotal,omitempty" validate:"required"`
	IVA           float64     `json:"iva,omitempty" validate:"required"`
	Total         float64     `json:"total,omitempty" validate:"required"`
	Comments      string      `json:"comments,omitempty" validate:"required"`
	Seller        string      `json:"seller,omitempty" validate:"required"`
	Regards       string      `json:"regards,omitempty" validate:"required"`
	Id            string      `bson:"_id,omitempty"`
}
