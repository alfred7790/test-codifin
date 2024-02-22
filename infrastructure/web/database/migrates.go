package database

import (
	"codifin-challenge/domain/model"
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20240221000000",
			Migrate: func(db *gorm.DB) error {
				return migrateSchema(db)
			},
			Rollback: func(tx *gorm.DB) error {
				return rollbackSchema(db)
			},
		},
	})

	if err := m.Migrate(); err != nil {
		return fmt.Errorf("error running migrations: %w", err)
	}
	return nil
}

func migrateSchema(db *gorm.DB) error {
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

func rollbackSchema(db *gorm.DB) error {
	modelsToDrop := []interface{}{
		&model.ItemCart{},
		&model.ShoppingCart{},
		&model.Product{},
	}

	for _, v := range modelsToDrop {
		if err := db.Migrator().DropTable(v); err != nil {
			return fmt.Errorf("error dropping table: %w", err)
		}
	}

	return nil
}
