package migration

import (
	"002_sample-go-app/migration/entities"
	"gorm.io/gorm"
)

func CreateTables(db *gorm.DB) error {
	if err := db.AutoMigrate(entities.User{}); err != nil {
		return err
	}

	return nil
}
