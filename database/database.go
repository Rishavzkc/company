package database

import (
	"company/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabaseConnection(uri string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to create a connection to database")
	}

	if err := db.AutoMigrate(&model.Company{}); err != nil {
		return nil, fmt.Errorf("failed to create company database")
	}

	return db, nil
}

func CloseDatabaseConnection(db *gorm.DB) error {
	dbSQL, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to fetching sql DB object for closing")
	}

	if err := dbSQL.Close(); err != nil {
		return fmt.Errorf("failed to close connection from database")
	}
	return nil
}
