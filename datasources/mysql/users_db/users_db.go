package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	//Client ... represents the mysql client
	Client *sql.DB

	username = os.Getenv("mysql_users_username")
	password = os.Getenv("mysql_users_password")
	host     = os.Getenv("mysql_users_host")
	schema   = os.Getenv("mysql_users_schema")
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, schema)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("DB connection success!")
}
