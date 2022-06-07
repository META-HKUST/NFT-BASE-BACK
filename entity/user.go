package entity

type Account struct {
	BannerImage string `json:"bannerimage" `
	AvatarImage string `json:"avatarimage" `
	UserId      string `json:"id"`
	Poison      string `json:"poison"`
	Campus      string `json:"campus"`
	Token       uint64 `json:"token"`
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
