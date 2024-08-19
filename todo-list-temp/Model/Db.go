package model

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:root@/todo?charset=utf8")
	if err != nil {
		log.Fatalln("Mysql Connect Error")
	}
	err = CreateTable()
	if err != nil {
		log.Fatalln("Create Table Error")
	}
}
