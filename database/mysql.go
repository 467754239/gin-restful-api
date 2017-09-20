package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var SqlDB *sqlx.DB

func init() {
	var err error
	SqlDB, err = sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := SqlDB.Ping(); err != nil {
		log.Fatal(err.Error())
	}
}
