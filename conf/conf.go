package conf

import (
	"database/sql"
	"log"
)

//DB -
var DB *sql.DB

//LoadDBConfig -
func LoadDBConfig() {
	db, err := sql.Open("mysql", "USER:PWD@tcp(DB-URL:3306)/DB-NAME?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
