package model

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	Email      string `json:"email" db:"email"`
	Passwd     string `json:"passwd" db:"passwd"`
	Token      string `json:"emailToken" db:"emailToken"`
	GenTime    string `json:"genTime" db:"genTime"`
	Activated  string `json:"activated" db:"activated"`
	VerifyCode string `json:"verify_code" db:"verify_code"`
	CodeTime   string `json:"codeTime" db:"codeTime"`
	UserId     string `json:"user_id" db:"user_id"`
}

// db variable
var db *sqlx.DB

// mysql sentences
var (
	// these three are related to account email and passwd
	insert        = string("insert into login(email,passwd,user_id) values(?,?,?);")
	query         = string("select email,activated from login where email=?;")
	queryPasswd   = string("select email,passwd from login where email = ?")
	update        = string("update login set passwd=? where email=?;")
	updateUserId  = string("update login set user_id=? where email=?;")
	deleteUser    = string("DELETE FROM login WHERE email=?;")
	deleteProfile = string("DELETE FROM accounts WHERE email=?;")

	// email activation
	updateToken      = string("update login set emailToken=? where email=?;")
	updateactivated  = string("update login set activated=? where email=?;")
	updategenTime    = string("update login set genTime=? where email=?;")
	activateToken    = string("update login set activated=? where emailToken=?;")
	queryGentime     = string("select genTime from login where emailToken=?;")
	queryCodetime    = string("select codeTime from login where email=?;")
	queryTokenStatus = string("select activated from login where email=?;")

	updateVerifyCode  = string("update login set verify_code=? , codeTime=? where email=?;")
	getVerifyCode     = string("select verify_code from login where email=?;")
	updateResetPasswd = string("update login set passwd=? where email=?;")

	insertAccount = string("insert into accounts(user_id,email,user_name,banner_image,logo_image,poison,organization,token) values(?,?,?,?,?,?,?,?);")
)

// 连接池设为最大100，空闲最大20，可以调整
func InitDB(config config.Config) {
	var err error
	db, err = sqlx.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Println(base.OpenSqlError, base.OpenSqlError.String(), err)
		return
	}
	err2 := db.Ping()
	if err2 != nil {
		log.Println(base.ConnectSqlError, base.ConnectSqlError.String(), err2)
		return
	}
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(20)
}

func DeleteUser(email string) error {
	_, e1 := db.Exec(deleteUser, email)
	if e1 != nil {
		log.Println(e1)
		return e1
	}
	return nil
}
func DeleteProfile(email string) error {
	_, e1 := db.Exec(deleteProfile, email)
	if e1 != nil {
		log.Println(e1)
		return e1
	}
	return nil
}

func InsertAccount(email string, Id string) error {
	num := 0
	if email == "contact@unifit.art" {
		num = 100000000
	}
	_, e1 := db.Exec(insertAccount, Id, email, email, "https://unift-1312994969.cos.ap-guangzhou.myqcloud.com/unifit/39171658756544_.pic.jpg?q-sign-algorithm=sha1&q-ak=AKIDBD4i9ML5aswlLgmfJisnTt30f6JJ6duu&q-sign-time=1658756789%3B1745156789&q-key-time=1658756789%3B1745156789&q-header-list=host&q-url-param-list=&q-signature=731b782a5a946c99a1235a9da90ef5820ef989d8", "https://unift-1312994969.cos.ap-guangzhou.myqcloud.com/unifit/39231658756870_.pic.jpg?q-sign-algorithm=sha1&q-ak=AKIDBD4i9ML5aswlLgmfJisnTt30f6JJ6duu&q-sign-time=1658756912%3B1745156912&q-key-time=1658756912%3B1745156912&q-header-list=host&q-url-param-list=&q-signature=573b8e673eac69188bc0615b887eeb9843c67301", "not set up", "not set up", num)
	if e1 != nil {
		return e1
	}

	return nil
}

// first examin if the account exists and then insert
func (p Person) Register(UserId string) base.ErrCode {

	_, e := db.Exec(insert, p.Email, p.Passwd, UserId)
	if e != nil {
		log.Println(base.InsertError, base.InsertError.String(), e)
		return base.InsertError
	}

	return base.Success
}

// check the password according to the account
// user account login
func (p Person) Login() base.ErrCode {

	var p1 Person
	e := db.Get(&p1, queryPasswd, p.Email)

	if e != nil {
		log.Println(base.UserDoNotExist, base.UserDoNotExist.String(), e)
		return base.UserDoNotExist
	}

	if p1.Passwd != p.Passwd {
		return base.WrongLoginError
	}
	return base.Success
}

// update user passwd
func (p Person) Update(newpasswd string) base.ErrCode {

	var p1 Person
	e := db.Get(&p1, query, p.Email)
	if e != nil {
		log.Println(base.QueryError, base.QueryError.String(), e)
		return base.QueryError
	}
	if p1.Passwd != p.Passwd {
		return base.WrongLoginError
	}
	result, e := db.Exec(update, newpasswd, p.Email)
	if e != nil {
		log.Println(base.PasswdUpdateError, base.PasswdUpdateError.String(), e)
		return base.PasswdUpdateError
	}
	rowsAffected, _ := result.RowsAffected()
	lastInsertId, _ := result.LastInsertId()
	log.Println("rowsAffected: ", rowsAffected, "lastInsertId: ", lastInsertId)
	return base.Success
}

func GetPerson(email string) (Person, error) {
	var a Person
	err := db.Get(&a, query, email)
	if err != nil {
		log.Println(err)
		return Person{}, err
	}
	return a, nil
}

// store token but do not acitvate it
func (p Person) StoreEmailToken() error {
	_, e1 := db.Exec(updateToken, p.Token, p.Email)
	if e1 != nil {
		log.Println(e1)
		return e1
	}
	_, e2 := db.Exec(updateactivated, "no", p.Email)
	if e2 != nil {
		log.Println(e2)
		return e2
	}
	_, e3 := db.Exec(updategenTime, p.GenTime, p.Email)
	if e3 != nil {
		log.Println(e3)
		return e3
	}
	return nil
}

func ActivateEmailToken(token string) error {
	_, e := db.Exec(activateToken, "yes", token)
	if e != nil {
		return e
	}
	return nil
}

func GetGenTime(token string) (error, string) {
	var g string
	err := db.Get(&g, queryGentime, token)

	if err != nil {
		return err, ""
	}
	return nil, g
}
func GetCodeTime(email string) (error, string) {
	var g string
	err := db.Get(&g, queryCodetime, email)

	if err != nil {
		return err, ""
	}
	return nil, g
}

func GetTokenStatus(token string) (error, string) {
	var g string
	err := db.Get(&g, queryTokenStatus, token)
	if err != nil {
		return err, ""
	}
	return nil, g
}

func UpdateVerifyCode(email string, code string, time string) error {

	r1, e1 := db.Exec(updateVerifyCode, code, time, email)
	if e1 != nil {
		log.Println(e1)
		return e1
	}
	log.Println(r1)
	return nil
}

func GetVerifyCode(email string) (string, error) {
	var code string
	err := db.Get(&code, getVerifyCode, email)
	if err != nil {
		return "", err
	}
	return code, nil
}

func ResetUpdate(email string, passwd string) error {
	r1, e1 := db.Exec(updateResetPasswd, passwd, email)
	if e1 != nil {
		log.Println(e1)
		return e1
	}
	log.Println(r1)
	return nil
}

func UpdateId(email string, Id string) error {
	_, e1 := db.Exec(updateUserId, Id, email)
	if e1 != nil {
		log.Println(e1)
		return e1
	}
	return nil
}
