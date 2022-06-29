package model

import (
	"NFT-BASE-BACK/base"
	"log"
)

var (
	//edit user profile
	getProfileByID    = string("select * from accounts where user_id=?;")
	queryUserId       = string("select userId from login where email=?;")
	updateUserProfile = string("update accounts set user_name=?,banner_image=?,logo_image=?,poison=?,organization=? where user_id=?;")
)

type UserID struct {
	UserID string
}

type UserProfile struct {
	UserID           string `json:"user_id" db:"user_id"`
	UserName         string `json:"user_name" db:"user_name"`
	BannerImage      string `json:"banner_image" db:"banner_image"`
	LogoImage        string `json:"logo_image" db:"logo_image"`
	Poison           string `json:"poison" db:"poison"`
	Organization     string `json:"organization" db:"organization"`
	Token            uint64 `json:"token" db:"token"`
	RegistrationTime string `json:"registration_time" db:"created_at"`
}

type UserProfileInfo struct {
	UserId           string `json:"user_id" `
	UserEmail        string `json:"user_email" `
	UserName         string `json:"user_name" `
	BannerImage      string `json:"banner_image" `
	LogoImage        string `json:"logo_image"`
	Poison           string `json:"poison" `
	Organization     string `json:"organization" `
	RegistrationTime string `json:"registration_time" `
}

func EditProfile(email, username, organization, poison, logo, banner string) (UserProfileInfo, base.ErrCode) {
	p := UserProfile{}
	userID, err := GetUserIDByEmail(email)
	if err != nil {

		return UserProfileInfo{}, base.UserIDNotExist
	}
	e := db.Get(&p, getProfileByID, userID)

	if e != nil {
		log.Println(base.QueryError, base.QueryError.String(), e)
		return UserProfileInfo{}, base.QueryError
	}
	if username != "" {
		p.UserName = username
	}
	if organization != "" {
		p.Organization = organization
	}
	if poison != "" {
		p.Poison = poison
	}
	if logo != "" {
		p.LogoImage = logo
	}
	if banner != "" {
		p.BannerImage = banner
	}
	result, e := db.Exec(updateUserProfile, p.UserName, p.BannerImage, p.LogoImage, p.Poison, p.Organization, userID)
	if e != nil {
		log.Println(base.UserProfileUpdateError, base.UserProfileUpdateError.String(), e)
		return UserProfileInfo{}, base.UserProfileUpdateError
	}
	rowsAffected, _ := result.RowsAffected()
	lastInsertId, _ := result.LastInsertId()
	log.Println("rowsAffected: ", rowsAffected, "lastInsertId: ", lastInsertId)
	code, userProfileInfo := GetUserInfoByID(userID)
	if code != base.Success {
		return UserProfileInfo{}, code
	}

	return userProfileInfo, base.Success
}

func GetUserIDByEmail(email string) (string, error) {
	var g string
	err := db.Get(&g, queryUserId, email)
	if err != nil {
		return "", err
	}
	return g, nil
}

func GetUserInfoEmail(email string) (base.ErrCode, UserProfileInfo) {
	p := UserProfile{}
	userID, _ := GetUserIDByEmail(email)
	if userID == "" {
		return base.UserIDNotExist, UserProfileInfo{}
	}

	e := db.Get(&p, getProfileByID, userID)
	if e != nil {
		log.Println(base.QueryError, base.QueryError.String(), e)
		return base.QueryError, UserProfileInfo{}
	}

	resp := UserProfileInfo{
		p.UserID,
		email,
		p.UserName,
		p.BannerImage,
		p.LogoImage,
		p.Poison,
		p.Organization,
		p.RegistrationTime,
	}
	return base.Success, resp
}

func GetUserInfoByID(userID string) (base.ErrCode, UserProfileInfo) {
	p := UserProfile{}

	e := db.Get(&p, getProfileByID, userID)
	if e != nil {
		log.Println(base.QueryError, base.QueryError.String(), e)
		return base.QueryError, UserProfileInfo{}
	}

	resp := UserProfileInfo{
		p.UserID,
		"",
		p.UserName,
		p.BannerImage,
		p.LogoImage,
		p.Poison,
		p.Organization,
		p.RegistrationTime,
	}
	return base.Success, resp
}
