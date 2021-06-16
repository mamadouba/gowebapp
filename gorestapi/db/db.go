package db

import (
	"gorestapi/config"
	"gorestapi/logger"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Connect(c *config.Config) {
	logger.Info("Connect to database")
	mdb, err := gorm.Open("postgres", c.GetConnectStr())
	if err != nil {
		logger.Panic(err.Error())
	}
	DB = mdb
}