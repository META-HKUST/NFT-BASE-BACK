package model

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/config"
	"fmt"
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
	Token     string `db:"token"`
	Activated string `db:"activated"`
	GenTime   string `db:"genTime"`
}

// db variable
var db *sqlx.DB

// mysql sentences
var (
	// these three are related to account email and passwd
	insert = string("insert into userlogin(email,passwd) values(?,?);")
	query  = string("select email,passwd,token,activated,genTime from userlogin where email=?;")
	update = string("update userlogin set passwd=? where email=?;")

	// email activation
	updateToken      = string("update userlogin set token=? where email=?;")
	updateactivated  = string("update userlogin set activated=? where email=?;")
	updategenTime    = string("update userlogin set genTime=? where email=?;")
	activateToken    = string("update userlogin set activated=? where token=?;")
	queryGentime     = string("select genTime from userlogin where token=?;")
	queryTokenStatus = string("select activated from userlogin where token=?;")
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
		fmt.Println(err)
		return Person{}, err
	}
	return a, nil
}

// store token but do not acitvate it
func (p Person) StoreEmailToken() error {
	r1, e1 := db.Exec(updateToken, p.Token, p.Email)
	if e1 != nil {
		log.Println(e1)
		return e1
	}
	r2, e2 := db.Exec(updateactivated, "no", p.Email)
	if e2 != nil {
		log.Println(e2)
		return e2
	}
	r3, e3 := db.Exec(updategenTime, p.GenTime, p.Email)
	if e3 != nil {
		log.Println(e3)
		return e3
	}
	log.Println("Store email token succeeded, rowsAffected: ", r1.RowsAffected, r2.RowsAffected, r3.RowsAffected)
	return nil
}

func ActivateEmailToken(token string) error {
	r1, e := db.Exec(activateToken, "yes", token)
	if e != nil {
		return e
	}
	log.Println("Activate email token succeeded, rowsAffected: ", r1.RowsAffected)
	return nil
}

func GetGenTime(token string) string {
	var g string
	err := db.Get(&g, queryGentime, token)
	if err != nil {
		log.Println(err)
	}
	return g
}

func GetTokenStatus(token string) string {
	var g string
	err := db.Get(&g, queryTokenStatus, token)
	if err != nil {
		log.Println(err)
	}
	return g
}
