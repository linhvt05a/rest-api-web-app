package database
import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"example/rest-api/models"
  )
  
  func Init(url string) *gorm.DB {
    db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&User)

    return db
}