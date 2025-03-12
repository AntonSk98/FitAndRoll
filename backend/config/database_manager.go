package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DatabaseManager manages the database connection and provides helper functions.
type DatabaseManager struct {
	DB *gorm.DB
}

// NewDatabaseManager initializes and returns a new DatabaseManager with an SQLite connection.
// It logs a fatal error if the connection cannot be established.
func NewDatabaseManager() *DatabaseManager {
	// Open the SQLite database connection
	db, err := gorm.Open(sqlite.Open("fit_and_roll.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	log.Println("Database connection established")

	return &DatabaseManager{DB: db}
}

// Migrate automatically migrates the given database models.
// It logs a fatal error if migration fails.
//
// Example usage:
//
//	dbManager.Migrate(&User{}, &Order{})
func (dbManager *DatabaseManager) Migrate(dst ...interface{}) {
	err := dbManager.DB.AutoMigrate(dst...)
	if err != nil {
		log.Fatalf("Error creating a database strucure for one of the tables: %v. Error: %v", dst, err)
	}
}

// Paginate returns a GORM scope function that applies pagination based on the given page and size.
// If page <= 0, it defaults to page 1. If size < 0, it defaults to 10.
//
// Example usage:
//
//	db.Scopes(dbManager.Paginate(2, 20)).Find(&users)
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

// Transactional executes the given function within a database transaction.
// If the function returns an error, the transaction is rolled back; otherwise, it commits.
//
// Example usage:
//
//	err := dbManager.Transactional(func(tx *gorm.DB) error {
//	    return tx.Create(&User{Name: "John"}).Error
//	})
func (dbManager *DatabaseManager) Transactional(txFunc func(tx *gorm.DB) error) error {
	return dbManager.DB.Transaction(func(tx *gorm.DB) error {
		if err := txFunc(tx); err != nil {
			return err
		}

		return nil
	})
}
