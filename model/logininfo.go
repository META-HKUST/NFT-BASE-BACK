package model

import (
	"NFT-BASE-BACK/base"
	"NFT-BASE-BACK/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type Error struct {
	Num int
	error
}
type Person struct {
	Email  string
	Passwd string
}

// db variable
var db *sqlx.DB

// mysql sentences
var (
	insert = string("insert into logininfo(email,passwd) values(?,?);")
	query  = string("select email,passwd from logininfo where email=?;")
	update = string("update logininfo set passwd=? where email=?;")
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
	return
}

// first examin if the account exists and then insert
func (p Person) Register() base.ErrCode {

	var a Person
	db.Get(&a, query, p.Email)
	if (a != Person{}) {
		return base.AccountExistError
	}

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
	log.Println("rowsAffected: ", result.RowsAffected, "lastInsertId: ", result.LastInsertId)
	return base.Success
}
