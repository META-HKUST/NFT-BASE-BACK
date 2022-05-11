package entity

type User struct {
	Campus      int    `json:"campus" example:"1 mean CWB and 2 mean GZ"`
	Email       string `json:"email" example:"Sam@ust.hk"`
	Passwd      string `json:"passwd" example:"123"`
	BannerImage string `json:"bannerimage" example:"/home/yezzi/bannerimage"`
	AvatarImage string `json:"avatarimage" example:"/home/yezzi/bannerimage"`
	UserName    string `json:"username" example:"Sam"`
	Id          string `json:"id" example:"1001"`
	Certificate string `json:"certificate" example:"/home/yezzi/certificate_yezzi"`
}
