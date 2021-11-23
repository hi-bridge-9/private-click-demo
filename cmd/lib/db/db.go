package db

import (
	"database/sql"
	"fmt"
	"os"
)

var (
	user     = os.Getenv("MYSQL_USER")
	pass     = os.Getenv("MYSQL_PASSWORD")
	ip       = os.Getenv("MYSQL_IP")
	port     = os.Getenv("MYSQL_PORT")
	protocol = os.Getenv("MYSQL_PROTOCOL")
	name     = os.Getenv("MYSQL_DATABASE")
)

func Connect() *sql.DB {
	conf := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8mb4", user, pass, protocol, ip, port, name)
	db, err := sql.Open("mysql", conf)
	if err != nil {
		panic(err.Error())
	}
	return db
}
