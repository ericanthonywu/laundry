package Config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func Db() *gorm.DB {
	//fmt.Sprintf()
	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	//	os.Getenv("DB_HOST"),
	//	os.Getenv("DB_USERNAME"),
	//	os.Getenv("DB_PASSWORD"),
	//	os.Getenv("DB_DATABASE"),
	//	os.Getenv("DB_PORT"))
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		os.Getenv("DB_CONNECTION"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}
	return db
}
