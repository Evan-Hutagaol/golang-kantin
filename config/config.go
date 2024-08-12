package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load file")
	}
}

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

func ConnectToDB() *gorm.DB {
	var dbConfig DBConfig = DBConfig{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}
	fmt.Println("dbconfig : ", dbConfig)
	fmt.Println("test")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name)

	var err error
	fmt.Println("dsn : ", dsn)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	// DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		// fmt.Println("failed")
		// panic("Database Connection Error")
		log.Fatal("Database connection error: ", err)
	}
	fmt.Println("Success")

	// sqlDB, err := DB.DB()
	// if err != nil {
	// 	log.Fatal("Failed to get database instance: ", err)
	// }
	// sqlDB.SetMaxOpenConns(50);
	// sqlDB.SetMaxIdleConns(10);
	// sqlDB.SetConnMaxIdleTime(30 * time.Minute)
	return DB
}
