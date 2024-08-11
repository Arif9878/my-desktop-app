package internal

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Item struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"type:text"`
}

func InitDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	// Automatically migrate your schema, to keep it up to date.
	db.AutoMigrate(&Item{})

	return db
}
