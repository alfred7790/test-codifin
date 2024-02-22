package database

import (
	"codifin-challenge/domain/model"
	"fmt"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	models := []interface{}{
		&model.Product{},
		&model.ShoppingCart{},
		&model.ItemCart{},
	}

	if err := db.AutoMigrate(models...); err != nil {
		return fmt.Errorf("error auto-migrating schema: %w", err)
	}

	return nil
}
