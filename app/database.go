package app

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB singleton
var DB *gorm.DB
var err error

// Database configuration for database
type Database struct {
	Type     string
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}


// buildDatabase => setup db config
func buildDatabase() *Database {
	dbConfig := Database{
		Host:     "localhost",
		Port:     5432,
		User:     "chris",
		Password: "chris",
		Name:     "gin_blog",
	}
	return &dbConfig
}

// getDsn => dsn for postgres
func getDsn(dbConfig *Database) string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Name,
		dbConfig.Port,
	)
}

// connectDB => connect to postgresql db
func connectDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(getDsn(buildDatabase())), &gorm.Config{})
	return db, err
}

// InitDB => get db instance @singleton
func InitDB() {
	if DB == nil {
		DB, err = connectDB()
	}
}
