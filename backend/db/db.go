package db


import (
	"log"
	"gorm.io/driver/mysql"
    "gorm.io/gorm"

	"github.com/set2002satoshi/YoutubeApp/backend/model"
)

func DbInit() {
	DE := OpenDB()
	DE.AutoMigrate(&model.Users{})
}






func OpenDB() *gorm.DB {
	dsn := "docker:pass@tcp(mySQL:3306)/main?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("OpenDB failed:", err)
	}
	return db
}





