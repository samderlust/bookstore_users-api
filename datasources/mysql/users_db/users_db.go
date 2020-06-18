package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

//UserDB
var (
	Client   *sql.DB
	username = os.Getenv("MYSQL_USERNAME")
	password = os.Getenv("MYSQL_PASSWORD")
	host     = os.Getenv("MYSQL_HOST")
	db       = os.Getenv("MYSQL_DB_USER")
)

func init() {

	fmt.Printf("user: %s, pass: %s, host: %s, db: %s", username,
		password,
		host,
		db)

	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?",
		username,
		password,
		host,
		db)

	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database connected")
}
