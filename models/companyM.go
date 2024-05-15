package models

type Company struct {
	Cid    string `json:"cid" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Rfc    string `json:"rfc" binding:"required"`
	Street string `json:"street" binding:"required"`
	City   string `json:"city" binding:"required"`
	State  string `json:"state" binding:"required"`
	PC     string `json:"pc" binding:"required"`
	Logo   string `json:"logo" binding:"required"`
	Conf   Conf   `json:"conf" binding:"required"`
}

type Conf struct {
	QSerie      string `json:"qserie" binding:"required"`
	QFolio      int    `json:"qfolio" binding:"required"`
	QConditions string `json:"qconditions" binding:"required"`
	PSerie      string `json:"pserie" binding:"required"`
	PFolio      int    `json:"pfolio" binding:"required"`
	PConditions string `json:"pconditions" binding:"required"`
	FSerie      string `json:"fserie" binding:"required"`
	FFolio      int    `json:"ffolio" binding:"required"`
}

type NewCompany struct {
	Name   string `json:"name" binding:"required"`
	Rfc    string `json:"rfc" binding:"required"`
	Street string `json:"street" binding:"required"`
	City   string `json:"city" binding:"required"`
	State  string `json:"state" binding:"required"`
	PC     string `json:"pc" binding:"required"`
	Logo   string `json:"logo" binding:"required"`
	Conf   Conf   `json:"conf" binding:"required"`
}
