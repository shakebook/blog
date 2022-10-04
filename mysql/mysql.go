package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

type Information struct {
	UserName     string
	Password     string
	Addr         string
	DatabaseName string
}

var conn *sql.DB

// 连接Mysql
func Connect(info *Information) {
	fmt.Println("Connect to mysql ing...")
	db, err := sql.Open("mysql", genURL(info))
	if err != nil {
		log.Fatalf("Failed to connect to mysql, err:%v\n", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to mysql, err:%v\n", err.Error())
	}
	conn = db
	fmt.Println("Connect to mysql successfuly!")
}

// 拼接地址
func genURL(info *Information) string {
	var b strings.Builder
	b.WriteString(info.UserName)
	b.WriteString(":")
	b.WriteString(info.Password)
	b.WriteString("@tcp(")
	b.WriteString(info.Addr)
	b.WriteString(")/")
	b.WriteString(info.DatabaseName)
	b.WriteString("?charset=utf8&parseTime=True")
	return b.String()
}

// 获取Mysql连接
func GetDB() *sql.DB {
	return conn
}
