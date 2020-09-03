package models

func SearchInfo(keyword string) interface{} {
	var searchResult []Info
	db.Model(&Info{}).Where("id = %s or name = %s or sex = %s or birthday = %s or province = %s or city = %s or area = %s").Find(&searchResult)
	return searchResult
}

func SearchInfoBySpan(sTime, eTime string) interface{} {
	var searchResult []Info
	db.Model(&Info{}).Where("birthday between %s and %s", sTime, eTime).Find(&searchResult)
	return searchResult
}
