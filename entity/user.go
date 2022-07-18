package entity

type Account struct {
	BannerImage string `json:"banner_image" db:"banner_image"`
	LogoImage   string `json:"logo_image" db:"logo_image"`
	UserId      string `json:"user_id" db:"user_id"`
	Poison      string `json:"poison" db:"poison"`
	Campus      string `json:"campus" db:"campus"`
	Token       uint64 `json:"token" db:"token"`
}

type Person struct {
	Email  string `json:"email" db:"email"`
	Passwd string `json:"passwd" db:"passwd"`
	Name   string `json:"name" db:"name"`
	Activate
}

type Activate struct {
	Token     string `db:"token" json:"token"`
	Activated string `db:"activated" json:"activated"`
	GenTime   string `db:"genTime" json:"genTime"`
}

type RePasswd struct {
	Email      string `json:"email"`
	VerifyCode string `json:"verifyCode"`
}
