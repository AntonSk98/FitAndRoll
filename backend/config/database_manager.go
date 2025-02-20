package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseManager struct {
	DB *gorm.DB
}

func NewDatabaseManager() *DatabaseManager {
	// Open the SQLite database connection
	db, err := gorm.Open(sqlite.Open("fit_and_roll.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	log.Println("Database connection established")

	return &DatabaseManager{DB: db}
}

func (dbManager *DatabaseManager) Migrate(dst ...interface{}) {
	err := dbManager.DB.AutoMigrate(dst...)
	if err != nil {
		log.Fatalf("Error creating a database strucure for one of the tables: %v. Error: %v", dst, err)
	}
}

func (dbManager *DatabaseManager) Paginate(page int, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		if size < 0 {
			size = 10
		}

		offset := (page - 1) * size

		return db.Offset(offset).Limit(size)
	}
}

func (dbManager *DatabaseManager) Transactional(txFunc func(tx *gorm.DB) error) error {
	return dbManager.DB.Transaction(func(tx *gorm.DB) error {
		if err := txFunc(tx); err != nil {
			return err
		}

		return nil
	})
}
