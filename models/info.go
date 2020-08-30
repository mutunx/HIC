package models

type Info struct {
	Name      string `json:"name"`
	Sex       string `json:"sex"`
	Nation    string `json:"nation"`
	Birthday  string `json:"birthday"`
	Hometown  string `json:"hometown"`
	Career    string `json:"career"`
	IDCard    string `json:"IDCard"`
	Unit      string `json:"unit"`
	Living    string `json:"living"`
	Father    string `json:"father"`
	Mother    string `json:"mother"`
	Education string `json:"education"`
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
