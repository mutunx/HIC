package models

func SearchInfo(keyword string) interface{} {
	var searchResult []Info
	db.Model(&Info{}).Where("id like ? or name = ? or sex = ? or birthday = ? or province = ? or city = ? or area = ?", keyword, keyword, keyword, keyword, keyword, keyword, keyword).Find(&searchResult)
	return searchResult
}

func SearchInfoBySpan(sTime, eTime string) interface{} {
	var searchResult []Info
	db.Model(&Info{}).Where("birthday between ? and ?", sTime, eTime).Find(&searchResult)
	return searchResult
}
