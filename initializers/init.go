package initializers

import (
	"log"
	"github.com/AvinFajarF/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conntection() {

	var errDB error

	dsn := "host=localhost user=root password=root dbname=authsimple port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	DB, errDB = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if errDB != nil {
		log.Fatalln("Connection to database error")
	}

	DB.AutoMigrate(&model.User{})
}
