package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Initialize() error {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DATABASE")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, port, dbname)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}
