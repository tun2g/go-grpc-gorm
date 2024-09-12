package database

import (
	"app/src/config"
	"app/src/lib/logger"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var log = logger.NewLogger("Database")

func InitDB() *gorm.DB {
	dataSourceName := fmt.Sprintf("host=%s port=%v user=%s dbname=%s password=%s sslmode=disable",
		config.AppConfiguration.DbHost,
		config.AppConfiguration.DbPort,
		config.AppConfiguration.DbUser,
		config.AppConfiguration.DbName,
		config.AppConfiguration.DbPassword,
	)

	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	log.Info("Connected to database")

	return db
}
