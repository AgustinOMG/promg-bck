package models

type Supplier struct {
	CID      string `json:"cid,omitempty" validate:"required"`
	Name     string `json:"name,omitempty" validate:"required"`
	Nickname string `json:"nickname,omitempty" validate:"required"`
	Rfc      string `json:"rfc,omitempty" validate:"required"`
	Street   string `json:"street,omitempty" validate:"required"`
	City     string `json:"city,omitempty" validate:"required"`
	State    string `json:"state,omitempty" validate:"required"`
	PC       string `json:"pc,omitempty" validate:"required"`
	Id       string `bson:"_id,omitempty"`
}
