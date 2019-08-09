package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err error
)

func init(){
	openConn()
}

func openConn() *sql.DB{
	dbConn, err := sql.Open("mysql", "root:123!@#@tcp(localhost:3306)/video_server?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	return dbConn
}