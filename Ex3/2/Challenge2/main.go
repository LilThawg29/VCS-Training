package main

import (
	"github.com/challenge2/database"
	"github.com/challenge2/routes"
)

// Khai báo Swagger

// @title           Swagger Example API
// @version         1.0
// @description     Đây là swagger api của challenge 2
// @host localhost:8080
// @BasePath /

func main(){
	database.InitialMigration()
	routes.ServerStart()
}

