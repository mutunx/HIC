package models

type Info struct {
	Name       string `json:"name"`
	Sex        string `json:"sex"`
	Birthday   string `json:"birthday"`
	Province   string `json:"province"`
	City       string `json:"city"`
	ID         string `json:"id"`
	Area       string `json:"area"`
	ValidState string `json:"validState"`
}

func GetInfos() (infos []Info, total int) {
	db.Find(&infos)
	total = len(infos)
	return
}

func GetInfo(queryMap map[string]interface{}) (infos []Info) {
	db.Where(queryMap).Find(&infos)
	return
}

func EditInfo(IDCard string, info interface{}) (count int64, err error) {
	db.Model(&Info{}).Where("id = ?", IDCard).Update(info)
	return
}

func AddInfo(info *Info) (infoId interface{}, err error) {
	db.Create(&Info{})
	return
}
