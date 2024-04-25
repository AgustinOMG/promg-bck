package models

type User struct {
	UID        string   `json:"uid,omitempty" validate:"required"`
	Name       string   `json:"name,omitempty" validate:"required"`
	Email      string   `json:"email,omitempty" validate:"required"`
	Company    []string `json:"company,omitempty" validate:"required"`
	Nickname   string   `json:"nickname,omitempty" validate:"required"`
	Telephone  string   `json:"telephone,omitempty" validate:"required"`
	Department string   `json:"department,omitempty" validate:"required"`
	Level      string   `json:"level,omitempty" validate:"required"`
	Photo      string   `json:"photo,omitempty" validate:"required"`
}
