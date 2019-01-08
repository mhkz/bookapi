package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 用户密码
	Slogan   string `json:"slogan"`   // 签名
	Gravatar string `json:"gravatar"` // 头像
}

func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}

	return false
}
