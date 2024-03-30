package main

import (
	"PostBank_Virtual_Banking/app"
	"PostBank_Virtual_Banking/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func createDataBaseIfNotExist() {
	dsn := "root:root@tcp(localhost:13306)/mysql"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Create the database if it doesn't exist
	db.Exec("CREATE DATABASE IF NOT EXISTS PostBank_Virtual")

}
func main() {

	createDataBaseIfNotExist()

	dsn := "root:root@tcp(127.0.0.1:13306)/PostBank_Virtual?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	err = db.AutoMigrate(&domain.Proceeding{}, &domain.Gama{})
	if err != nil {
		panic(err.Error())
	}

	app.StartApplication(db)
}
