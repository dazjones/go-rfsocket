package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/dazjones/go-rfsocket/conf"
	"fmt"
	"log"
)

var db gorm.DB

func Init(c *conf.Configuration) {
	cs := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.MySQLUser, c.MySQLPass, c.MySQLHost, c.MySQLPort, c.MySQLName)

	con, err := gorm.Open("mysql", cs)

	if err != nil {
		log.Fatalln(err)
	}
	db = *con

	db.AutoMigrate(
		&Transmission{},
	)
}
