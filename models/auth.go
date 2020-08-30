package models

type Auth struct {
	ID        string `gorm:"primary_key" json:"id"`
	LoginName string `json:"login_name"`
	Password  string `json:"password"`
}

func CheckAuth(loginName, password string) bool {
	var auth Auth
	db.Where(Auth{LoginName: loginName, Password: password}).First(&auth)
	if auth.ID != "" {
		return true
	}
	return false
}
