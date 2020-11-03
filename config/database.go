package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB singleton
var DB *gorm.DB
var err error

// DBConfig represnets db configration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// buildDBConfig => setup db config
func buildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "chris",
		Password: "chris",
		DBName:   "gin_blog",
	}
	return &dbConfig
}

// getDsn => dsn for postgres
func getDsn(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Port,
	)
}

// connectDB => connect to postgresql db
func connectDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(getDsn(buildDBConfig())), &gorm.Config{})
	return db, err
}

// GetDB => get db instance @singleton
func GetDB() (*gorm.DB, error) {
	if DB == nil {
		DB, err = connectDB()
	}
	return DB, err
}
