package models

type User struct {
	UID        string   `json:"uid" binding:"required"`
	Name       string   `json:"name" binding:"required"`
	Email      string   `json:"email" binding:"required"`
	Company    []string `json:"company" binding:"required"`
	Nickname   string   `json:"nickname" binding:"required"`
	Telephone  string   `json:"telephone" binding:"required"`
	Department string   `json:"department" binding:"required"`
	Level      string   `json:"level" binding:"required"`
	Photo      string   `json:"photo" binding:"required"`
	Status     string   `json:"status" binding:"required"`
	Id         string   `bson:"_id,omitempty"`
}
