package model

import (
	"NFT-BASE-BACK/base"
	"log"
)

var (
	//edit user profile
	getProfileByID     = string("select * from accounts where user_id=?;")
	getLogoImageByID   = string("select logo_image from accounts where user_id=?;")
	queryUserId        = string("select user_id from accounts where email=?;")
	updateUserProfile  = string("update accounts set user_name=?,banner_image=?,logo_image=?,poison=?,organization=? where user_id=?;")
	getProfileByKey    = string("select user_id,email,user_name,banner_image,logo_image,poison,organization,created_at from accounts where user_name like concat ('%',?,'%') limit ? offset ?;")
	queryUserName      = string("select user_name from accounts where user_id=?;")
	queryNameByEmail   = string("select user_name from accounts where email=?;")
	updateAccountToken = string("update accounts set token=? where user_id=?;")
)

type UserID struct {
	UserID string
}

type UserProfile struct {
	UserID           string `json:"user_id" db:"user_id"`
	UserEmail        string `json:"email" db:"email"`
	UserName         string `json:"user_name" db:"user_name"`
	BannerImage      string `json:"banner_image" db:"banner_image"`
	LogoImage        string `json:"logo_image" db:"logo_image"`
	Poison           string `json:"poison" db:"poison"`
	Organization     string `json:"organization" db:"organization"`
	RegistrationTime string `json:"registration_time" db:"created_at"`
}

type UserProfileInfo struct {
	UserID           string `json:"user_id" db:"user_id"`
	UserEmail        string `json:"email" db:"email"`
	UserName         string `json:"user_name" db:"user_name"`
	BannerImage      string `json:"banner_image" db:"banner_image"`
	LogoImage        string `json:"logo_image" db:"logo_image"`
	Poison           string `json:"poison" db:"poison"`
	Organization     string `json:"organization" db:"organization"`
	Token            uint64 `json:"token" db:"token"`
	RegistrationTime string `json:"registration_time" db:"created_at"`
}

// update user passwd
func UpdateAccountToken(userId string) error {
	_, e := db.Exec(updateAccountToken, 0, userId)
	if e != nil {
		log.Println(e)
		return e
	}
	return nil
}

func GetUserName(UserId string) (string, error) {
	var g string
	err := db.Get(&g, queryUserName, UserId)
	if err != nil {
		return "", err
	}
	return g, nil
}

func GetLogoImage(UserId string) (string, error) {
	var g string
	err := db.Get(&g, getLogoImageByID, UserId)
	if err != nil {
		return "", err
	}
	return g, nil
}

func GetNameByEmail(email string) (string, error) {
	var g string
	err := db.Get(&g, queryNameByEmail, email)
	if err != nil {
		return "", err
	}
	return g, nil
}

func EditProfile(email, username, organization, poison, logo, banner string) (UserProfileInfo, base.ErrCode) {
	//p := UserProfileInfo{}
	userID, err := GetUserIDByEmail(email)
	if err != nil {
		return UserProfileInfo{}, base.UserIDNotExist
	}
	//e := db.Get(&p, getProfileByID, userID)
	//
	//if e != nil {
	//	log.Println(base.QueryError, base.QueryError.String(), e)
	//	return UserProfileInfo{}, base.QueryError
	//}
	args := []string{}
	str := "update accounts set "
	if username != "" {
		str = str + "user_name=?,"
		args = append(args, username)
	}

	if organization != "" {
		str = str + "organization=?,"
		args = append(args, organization)
	}
	if poison != "" {
		str = str + "poison=?,"
		args = append(args, poison)
	}
	if logo != "" {
		str = str + "logo_image=?,"
		args = append(args, logo)
	}

	if banner != "" {
		str = str + "banner_image=?,"
		args = append(args, banner)
	}
	str = str[:len(str)-1]
	str = str + " " + "where user_id=?;"
	args = append(args, userID)
	params := make([]interface{}, len(args))
	for i, v := range args {
		params[i] = v
	}
	result, e := db.Exec(str, params...)
	if e != nil {
		log.Println(base.UserProfileUpdateError, base.UserProfileUpdateError.String(), e)
		return UserProfileInfo{}, base.UserProfileUpdateError
	}
	rowsAffected, _ := result.RowsAffected()
	lastInsertId, _ := result.LastInsertId()
	log.Println("rowsAffected: ", rowsAffected, "lastInsertId: ", lastInsertId)
	code, userProfileInfo := GetUserInfoByID(userID)
	if code != base.Success {
		log.Println(code)
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
	p := UserProfileInfo{}
	userID, _ := GetUserIDByEmail(email)
	if userID == "" {
		return base.UserIDNotExist, UserProfileInfo{}
	}

	e := db.Get(&p, getProfileByID, userID)
	if e != nil {
		log.Println(base.QueryError, base.QueryError.String(), e)
		return base.QueryError, UserProfileInfo{}
	}

	return base.Success, p
}

func GetUserInfoByID(userID string) (base.ErrCode, UserProfileInfo) {
	p := UserProfileInfo{}

	e := db.Get(&p, getProfileByID, userID)
	if e != nil {
		log.Println(base.QueryError, base.QueryError.String(), e)
		return base.QueryError, UserProfileInfo{}
	}
	return base.Success, p
}

func GetUserListByKey(keyword string, pageNum, pageSize int64) ([]UserProfile, error) {
	offset := (pageNum - 1) * pageSize
	var userInfos []UserProfile
	err := db.Select(&userInfos, getProfileByKey, keyword, pageSize, offset)
	if err != nil {
		log.Println(base.QueryError, base.QueryError.String(), err)
		return []UserProfile{}, err
	}
	return userInfos, nil
}
