package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func NewDataBase(host, port, user, password, name string, retries int) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, name)

	var err error
	for i := 0; i < retries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
		if err == nil {
			db = db.Debug()

			log.Printf("connection with database on %s:%s was stablished, enjoy it!...", host, port)
			return db, nil
		}

		log.Printf("%d remaining attempts to connect with the database: %s\n", retries-i-1, err.Error())
		time.Sleep(time.Second * 3)
	}

	return nil, fmt.Errorf("failed to establish connection to database after %d attempts", retries)
}
