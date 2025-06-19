package main

import (
	_ "gorm.io/driver/mysql"
	_ "gorm.io/gorm"
	"spleet-api/internal/routes"
)

func loadDb() {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
}

func loadRouter() {
	router := routes.SetupRouter()
	err := router.Run()
	if err != nil {
		panic(err)
	}
}

func main() {
	loadDb()
	loadRouter()
}
