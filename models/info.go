package models

type Info struct {
	ID         string
	Name       string
	Birthday   string
	Province   string
	City       string
	Area       string
	Sex        int
	ValidState int
	RealState  int
	Error      string
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
	db.Create(info)
	return
}
