package models

type Client struct {
	CID      string `json:"cid" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Rfc      string `json:"rfc" binding:"required"`
	Street   string `json:"street" binding:"required"`
	City     string `json:"city" binding:"required"`
	State    string `json:"state" binding:"required"`
	PC       string `json:"pc" binding:"required"`
}
