package db

import (
	"database/sql"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	once       sync.Once
	dbInstance *gorm.DB // Add this line
	sqlDB      *sql.DB  // Add this line
)

// InitDB initializes the database connection with pooling using GORM.
func InitDB() (*gorm.DB, error) {
	var err error

	once.Do(func() {
		dbInstance, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		// Set connection pool settings here if needed
		sqlDB, err := dbInstance.DB()
		if err != nil {
			panic("failed to get database instance")
		}
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
	})

	return dbInstance, err
}

// GetDB returns the initialized database instance.
func GetDB() (*gorm.DB, error) {
	return InitDB()
}

func CloseDB() error {
	return sqlDB.Close()
}
