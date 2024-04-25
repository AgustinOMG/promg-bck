package models

type Item struct {
	Code         string `json:"code,omitempty" validate:"required"`
	Description  string `json:"description,omitempty" validate:"required"`
	Currency     string `json:"currency,omitempty" validate:"required"`
	Mu           string `json:"mu,omitempty" validate:"required"`
	SellingPrice string `json:"sellingprice,omitempty" validate:"required"`
	BuyingPrice  string `json:"buyingprice,omitempty" validate:"required"`
	Supplier     string `json:"supplier,omitempty" validate:"required"`
	CID          string `json:"cid,omitempty" validate:"required"`
}
