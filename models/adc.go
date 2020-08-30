package models

type AdministrativeDivisionsCode struct {
	Value    string    `json:"value"`
	Label    string    `json:"label"`
	Children [0]string `json:"children"`
	Loading  bool      `json:"loading"`
	CodeType string    `json:"code_type"`
}

func GetProvince() (adcs []AdministrativeDivisionsCode) {
	db.Select("distinct province_code value, province_name label, 'city' code_type").Find(&adcs)
	return
}
func GetCity(code string) (adcs []AdministrativeDivisionsCode) {
	db.Select("distinct city_code value, city_name label, 'area' code_type").Where("province_code = ? and city_name is not null", code).Find(&adcs)
	return
}
func GetArea(code string) (adcs []AdministrativeDivisionsCode) {
	db.Select("distinct area_code value, area_name label").Where("city_code = ? and area_name is not null", code).Find(&adcs)
	return
}
