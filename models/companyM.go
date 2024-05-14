package models

type Company struct {
	Cid     string `json:"cid,omitempty" validate:"required"`
	Name    string `json:"name,omitempty" validate:"required"`
	Rfc     string `json:"rfc,omitempty" validate:"required"`
	Street  string `json:"street,omitempty" validate:"required"`
	City    string `json:"city,omitempty" validate:"required"`
	State   string `json:"state,omitempty" validate:"required"`
	PC      string `json:"pc,omitempty" validate:"required"`
	LogoUri string `json:"logouir,omitempty" validate:"required"`
	Conf    Conf   `json:"conf,omitempty" validate:"required"`
}

type Conf struct {
	QSerie      string `json:"qserie,omitempty" validate:"required"`
	QFolio      string `json:"qfolio,omitempty" validate:"required"`
	QConditions string `json:"qconditions,omitempty" validate:"required"`
}

type NewCompany struct {
	Name    string `json:"name,omitempty" validate:"required"`
	Rfc     string `json:"rfc,omitempty" validate:"required"`
	Street  string `json:"street,omitempty" validate:"required"`
	City    string `json:"city,omitempty" validate:"required"`
	State   string `json:"state,omitempty" validate:"required"`
	PC      string `json:"pc,omitempty" validate:"required"`
	LogoUri string `json:"logouir,omitempty" validate:"required"`
	Conf    Conf   `json:"conf,omitempty" validate:"required"`
}
