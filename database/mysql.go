package database

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "log"
import "os"
import "fmt"
//import "time"

var My *sql.DB // MySQL


func init() {
	var err error
	var dsn string

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName)

	My, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = My.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
