package models

type Tabler interface {
	TableName() string
}

func (OmsLocation) TableName() string {
	return "oms_location"
}

type OmsLocation struct {
	LocId         int    `json:"loc_id"`
	CompanyId     int    `json:"company_id"`
	LocCode       string `json:"loc_code"`
	LocName       string `json:"loc_name"`
	LocStreet     string `json:"loc_street"`
	LocCity       string `json:"loc_city"`
	LocRegion     string `json:"loc_region"`
	LocPostcode   string `json:"loc_postcode"`
	LocCountryId  string `json:"loc_country_id"`
	LocTelephone  string `json:"loc_telephone"`
	LocLong       string `json:"loc_long"`
	LocLat        string `json:"loc_lat"`
	LocZone       string `json:"loc_zone"`
	IsWarehouse   bool   `json:"is_warehouse"`
	UseInFrontend bool   `json:"use_in_frontend"`
	LocNote       string `json:"loc_note"`
	CompanyName   string `json:"company_name"`
}

type LocationJson struct {
	Location OmsLocation `json:"location" binding:"required"`
}
