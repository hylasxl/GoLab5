package db

import (
	"gorm.io/gorm"
	"server/models"
)

func MigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&models.Class{},
		&models.Student{},
	); err != nil {
		return err
	}
	print("Migrate successfully\n")
	return nil
}
