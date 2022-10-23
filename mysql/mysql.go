package mysql

import (
	"blog/tools"
	"database/sql"
	"log"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

type Infor struct {
	UserName     string
	Password     string
	Addr         string
	DatabaseName string
}

var conn *sql.DB

// 连接Mysql
func Connect(info *Infor) {
	url := tools.PingString([]string{
		info.UserName,
		":",
		info.Password,
		"@tcp(",
		info.Addr,
		")/",
		info.DatabaseName,
		"?charset=utf8&parseTime=True",
	})
	db, err := sql.Open("mysql", url)
	if err != nil {
		log.Fatalf("Failed to connect to mysql, err:%v\n", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to mysql, err:%v\n", err.Error())
	}
	conn = db
	log.Println("Connect to mysql successfuly!")
}

// 获取Mysql连接
func GetDB() *sql.DB {
	return conn
}
