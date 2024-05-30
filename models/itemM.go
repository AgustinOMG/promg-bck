package models

type Item struct {
	Code         string  `json:"code" binding:"required"`
	Description  string  `json:"description" binding:"required"`
	Currency     string  `json:"currency" binding:"required"`
	Mu           string  `json:"mu" binding:"required"`
	SellingPrice float64 `json:"sellingPrice" binding:"required"`
	BuyingPrice  float64 `json:"buyingPrice" binding:"required"`
	Supplier     string  `json:"supplier" binding:"required"`
	CID          string  `json:"cid" binding:"required"`
	Id           string  `bson:"_id,omitempty"`
}
