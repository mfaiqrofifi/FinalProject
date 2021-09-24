package configs

import (
	"social_media/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:yuleyek@tcp(127.0.0.1:3306)/social_media_project?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB failed connect")
	}
	migration()
}

func migration() {
	DB.AutoMigrate(&models.UserDatabase{}, &models.StatusUserDB{})
}
