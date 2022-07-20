package model

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	Email  string `db:"email"`
	Passwd string `db:"passwd"`
	Activate
}

type Activate struct {
	Token      string `db:"emailToken"`
	GenTime    string `db:"genTime"`
	Activated  string `db:"activated"`
	VerifyCode string `db:"verify_code"`
	CodeTime   string `db:"codeTime"`
	UserId     string `db:"userId"`
}

// db variable
var db *sqlx.DB

// mysql sentences
var (
	// these three are related to account email and passwd
	insert       = string("insert into login(email,passwd) values(?,?);")
	query        = string("select email,passwd from login where email=?;")
	update       = string("update login set passwd=? where email=?;")
	updateUserId = string("update login set user_id=? where email=?;")

	// email activation
	updateToken      = string("update login set emailToken=? where email=?;")
	updateactivated  = string("update login set activated=? where email=?;")
	updategenTime    = string("update login set genTime=? where email=?;")
	activateToken    = string("update login set activated=? where emailToken=?;")
	queryGentime     = string("select genTime from login where emailToken=?;")
	queryTokenStatus = string("select activated from login where email=?;")

	updateVerifyCode  = string("update login set verify_code=? where email=?;")
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

func InsertAccount(email string, Id string) error {
	r1, e1 := db.Exec(insertAccount, Id, email, email, "https://lh3.googleusercontent.com/khLUIXJ0H0HDHvbaMPg-yhHKnTHPS7TVg4q_sHj9YNaLci-6tEw5K8UElAwZ0ov7ABBpnj0eNw9S1I637mpLqcuSdxrZYGuAaXU_rg=s0", "https://lh3.googleusercontent.com/ydFS04Bvu6ek2JcqwEsiIavNCMfFhforiZ24EcBYMurNvj1QgCDj1Ol1V9tZuH8Kfomqd83Umtr7iFe_FQZ3ptF5_0uQwiG_nYbP1hQ=h600", "not set up", "not set up", 100)
	if e1 != nil {
		log.Println(e1)
		return e1
	}
	log.Println(r1)
	return nil
}

// first examin if the account exists and then insert
func (p Person) Register() base.ErrCode {

	result, e := db.Exec(insert, p.Email, p.Passwd)
	if e != nil {
		log.Println(base.InsertError, base.InsertError.String(), e)
		return base.InsertError
	}

	rowsAffected, _ := result.RowsAffected()
	lastInsertId, _ := result.LastInsertId()
	log.Println("rowsAffected: ", rowsAffected, "lastInsertId: ", lastInsertId)
	return base.Success
}

// check the password according to the account
// user account login
func (p Person) Login() base.ErrCode {

	var p1 Person
	e := db.Get(&p1, query, p.Email)
	if e != nil {
		log.Println(base.QueryError, base.QueryError.String(), e)
		return base.QueryError
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

func GetTokenStatus(token string) (error, string) {
	var g string
	err := db.Get(&g, queryTokenStatus, token)
	if err != nil {
		return err, ""
	}
	return nil, g
}

func UpdateVerifyCode(email string, code string) error {
	r1, e1 := db.Exec(updateVerifyCode, code, email)
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
	r1, e1 := db.Exec(updateUserId, Id, email)
	if e1 != nil {
		log.Println(e1)
		return e1
	}
	log.Println(r1)
	return nil
}
