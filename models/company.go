package models

type CompanyTabler interface {
	TableName() string
}

func (OmsCompany) TableName() string {
	return "oms_company"
}

type OmsCompany struct {
	CompanyId        int    `json:"company_id"`
	CompanyCode      string `json:"company_code"`
	CompanyName      string `json:"company_name"`
	Email            string `json:"email"`
	NoTelephone      string `json:"no_telephone"`
	CompanyStreet    string `json:"company_street"`
	CompanyCountryId string `json:"company_country_id"`
	CompanyRegion    string `json:"company_region"`
	CompanyCity      string `json:"company_city"`
}

type CompanyJson struct {
	Company OmsCompany `json:"company" binding:"required"`
}
