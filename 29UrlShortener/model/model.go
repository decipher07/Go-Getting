package model

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Goly struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Redirect string `json:"redirect" gorm:"non null"`
	Goly     string `json:"goly" gorm:"unique;not null"`
	Clicked  uint64 `json:"clicked"`
	Random   bool   `json:"random"`
}

var db *gorm.DB

func Setup() {
	/* Database URL for GORM */
	dsn := "host=127.0.0.1 user=dj password=godisgreat dbname=test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	/* ORM Migrations - Similar to Prisma */
	err = db.AutoMigrate(&Goly{})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database Connected Successfully")
}
