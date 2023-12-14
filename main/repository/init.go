package repository

import (
	c "github.com/blue-farid/WebMidExam/config"
	"github.com/blue-farid/WebMidExam/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitDB() {
	var err error
	dsn := "host=" + c.Cfg.Database.Host + " user=" + c.Cfg.Database.Username + " dbname=" + c.Cfg.Database.DBName + " sslmode=disable" + " password=" +
		c.Cfg.Database.Password

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	mErr := db.AutoMigrate(&model.Basket{}, &model.User{})
	if mErr != nil {
		log.Fatal("Error while initialize the database!", mErr)
	}
}
