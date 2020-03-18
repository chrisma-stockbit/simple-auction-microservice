package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetDbConnection() (*gorm.DB, error) {
	username := "myuser"
	password := "password"
	host := "localhost"
	port := 3306
	dbname := "simple_auction"
	return openConnection(username, password, host, port, dbname)
}

//openDbConnection open connection to mysql database
func openConnection(username string, password string, host string, port int, dbname string) (*gorm.DB, error) {
	connString := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)
	return gorm.Open("mysql", connString)
}
