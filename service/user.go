package service

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/model"
	"NFT-BASE-BACK/sdk/service"
	"NFT-BASE-BACK/utils"
	"crypto/md5"
	"fmt"
	emailverifier "github.com/AfterShip/email-verifier"
	"log"
	"strings"
)

var (
	verifier = emailverifier.NewVerifier()
)

func RegisterEmailToken(p model.Person, ReceiverName string) base.ErrCode {
	// check input
	if p.Email == "" {
		return base.EmptyInput
	}
	p.Token = utils.GenEmailToken()
	p.GenTime = utils.GetTimeNow()
	if err := p.StoreEmailToken(); err != nil {
		log.Println(err)
		return base.StoreEmailTokenError
	}
	if err := utils.Email(ReceiverName, p.Email, p.Token); err != nil {
		log.Println(err)
		return base.SendEmailError
	}
	return base.Success
}

func ActivateToken(token string) base.ErrCode {
	// check input
	if token == "" {
		return base.EmptyInput
	}
	_, genTime := model.GetGenTime(token)

	if genTime == "" {
		return base.TokenNotExist
	}
	if _, s := model.GetTokenStatus(token); s == "yes" {
		return base.TokenAlreadyActivated
	}
	if !utils.IsTimeValid(genTime) {
		return base.TokenInvalidError
	}
	if err := model.ActivateEmailToken(token); err != nil {
		log.Println(err)
		return base.ActivateEmailError
	} else {
		return base.Success
	}
}

func CheckEmailToken(p model.Person) base.ErrCode {
	// check input
	if p.Email == "" {
		return base.EmptyInput
	}
	p1, err := model.GetPerson(p.Email)
	if err != nil {
		log.Println(err)
		return base.GetPersonError
	}
	if p1.Activated == "no" {
		return base.TokenNotActivated
	}
	return base.Success
}

func Register(p model.Person) base.ErrCode {

	// check input
	if p.Email == "" {
		return base.EmptyInput
	}
	if p.Passwd == "" {
		return base.EmptyInput
	}
	//p.ActivateEmailToken()

	// check email format
	ret, err := verifier.Verify(p.Email)
	if err != nil {
		return base.EmailFormatError
	}
	if !ret.Syntax.Valid {
		return base.EmailFormatError
	}
	//// check if using ust email
	//b1 := strings.Contains(p.Email, "ust.hk")
	//if b1 == false{
	//	return base.EmailFormatError
	//}

	// check passwd length
	//l3 := strings.Count(p.Passwd, "") - 1
	//if l3 >= 24 || l3 < 8 {
	//	return base.PasswdLengthError
	//}

	// generate UserId
	t1 := strings.Replace(p.Email, "@", "-", -1)
	UserId := strings.Replace(t1, ".", "-", -1)

	p1, _ := model.GetPerson(p.Email)
	log.Println(p.Email, " registering, get sql info: ", p1, " ", UserId)
	if p1.Email == p.Email && p1.Activated == "no" {
		log.Println("delete account,profile,collection of : ", p.Email)
		_ = model.DeleteUser(p.Email)
		_ = model.DeleteProfile(p.Email)
		_ = model.DeleteCollection(UserId)
	} else if p1.Email == p.Email && p1.Activated == "yes" {
		return base.AccountExistError
	}

	//
	//if p1.Email == p.Email {
	//	if p1.Activated == "yes" {
	//		log.Println(base.AccountExistError.String())
	//		return base.AccountExistError
	//	} else {
	//		model.DeleteAccount(p1.Email)
	//	}
	//}

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(p.Passwd))
	Result := Md5Inst.Sum([]byte(""))

	p.Passwd = fmt.Sprintf("%x", Result)

	if err := p.Register(UserId); err != base.Success {
		log.Println(err)
		return base.ServerError
	}

	//e1 := model.UpdateId(p.Email, UserId)
	//if e1 != nil {
	//	log.Println(e1)
	//	return base.ServerError
	//}
	e2 := model.InsertAccount(p.Email, UserId)
	if e2 != nil {
		log.Println(e2)
		return base.InsertProfileError
	}
	if UserId != "admin-unifit-art" {
		tokenInfo, errCode := Transfer(100, "admin-unifit-art", UserId)
		if errCode != base.Success {
			log.Println(string(errCode))
			return errCode
		}
		log.Println("transfer to ", UserId, "Succeed", ", info:", tokenInfo)
	}

	log.Println("Successfully insert into default profile: ", UserId)
	e3 := model.CreatCollection("Default Collection", "https://unift-1312994969.cos.ap-guangzhou.myqcloud.com/unifit/39181658756551_.pic.jpg?q-sign-algorithm=sha1&q-ak=AKIDBD4i9ML5aswlLgmfJisnTt30f6JJ6duu&q-sign-time=1658756853%3B1745156853&q-key-time=1658756853%3B1745156853&q-header-list=host&q-url-param-list=&q-signature=53c49f4266ed89b708bd02efceffddfebe889b77", "https://unift-1312994969.cos.ap-guangzhou.myqcloud.com/unifit/39191658756557_.pic.jpg?q-sign-algorithm=sha1&q-ak=AKIDBD4i9ML5aswlLgmfJisnTt30f6JJ6duu&q-sign-time=1658756879%3B1745156879&q-key-time=1658756879%3B1745156879&q-header-list=host&q-url-param-list=&q-signature=052298646752db6cba304dd39e5ab6d68e08667b", "https://unift-1312994969.cos.ap-guangzhou.myqcloud.com/unifit/39171658756544_.pic.jpg?q-sign-algorithm=sha1&q-ak=AKIDBD4i9ML5aswlLgmfJisnTt30f6JJ6duu&q-sign-time=1658756789%3B1745156789&q-key-time=1658756789%3B1745156789&q-header-list=host&q-url-param-list=&q-signature=731b782a5a946c99a1235a9da90ef5820ef989d8", 0, "Default Collection", UserId, p.Email, utils.GetTimeNow())
	if e3 != nil {
		log.Println(e3)
		return base.CreateCollectionError
	}

	if err := RegisterEmailToken(p, p.Email); err != base.Success {
		log.Println(err)
		return base.InputError
	}

	err3 := service.Enroll(UserId)
	if err3 != nil {
		// TODO: check why this enroll error happens and if it influence transactions
		log.Println("Register to Fabric failed: ", err3, " ", UserId)
		return base.Success
	}

	log.Println("Register Succeed, email:", p.Email)
	return base.Success
}

func Login(p model.Person) (base.ErrCode, string, string) {
	// check input
	if p.Email == "" {
		return base.EmptyInput, "", ""
	}
	if p.Passwd == "" {
		return base.EmptyInput, "", ""
	}

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(p.Passwd))
	Result := Md5Inst.Sum([]byte(""))

	p.Passwd = fmt.Sprintf("%x", Result)

	if err := p.Login(); err != base.Success {
		return err, "", ""
	}
	token, err := utils.GenToken(p)
	if err != nil {
		return base.ServerError, "", ""
	}
	if err := CheckEmailToken(p); err != base.Success {
		return err, "", ""
	}
	t1 := strings.Replace(p.Email, "@", "-", -1)
	UserId := strings.Replace(t1, ".", "-", -1)
	return base.Success, token, UserId
}

func ForgetPasswd(email string) base.ErrCode {
	// check input
	if email == "" {
		return base.EmptyInput
	}
	p := model.Person{
		Email: email,
	}
	if err := CheckEmailToken(p); err != base.Success {
		return base.TokenNotActivated
	}
	p, err1 := model.GetPerson(email)
	if err1 != nil {
		return base.GetPersonError
	}
	if p.Activated == "no" {
		return base.TokenNotActivated
	}

	c := utils.GenVerifyCode()
	time := utils.GetTimeNow()
	err := model.UpdateVerifyCode(email, c, time)
	if err != nil {
		return base.ServerError
	}
	reName, _ := model.GetNameByEmail(email)
	if reName == "" {
		reName = "Sir/Madam"
	}
	log.Println(" sending reset email to: ", reName)
	err = utils.ResetEmail(reName, email, c)
	if err != nil {
		return base.ResetEmailError
	}
	return base.Success
}

func ResetPasswd(email string, code string, pd string) base.ErrCode {
	// check input
	if email == "" {
		return base.EmptyInput
	}
	c1, err1 := model.GetVerifyCode(email)
	if err1 != nil {
		return base.InputError
	}
	log.Println("Reset Passwd: Get verify code: ", c1)
	err2, cdTime := model.GetCodeTime(email)
	if err2 != nil {
		return base.ServerError
	}
	if utils.IsTimeValid(cdTime) == false {
		return base.WrongVerifyCode
	}

	// TODO: Set the expire time of verify code
	if c1 != code {
		return base.WrongVerifyCode
	}

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(pd))
	Result := Md5Inst.Sum([]byte(""))

	pd = fmt.Sprintf("%x", Result)

	err := model.ResetUpdate(email, pd)
	if err != nil {
		return base.InputError
	}
	return base.Success
}
