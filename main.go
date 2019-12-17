package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"react-go/Config"
	entities "react-go/Models"
	"react-go/Web/router"
	_ "github.com/go-sql-driver/mysql"
)

var err error

func main() {

	// Creating a connection to the database
	//Using sql
	Config.DB, err = gorm.Open("mysql",Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("status: ", err)
	}

	defer Config.DB.Close()

	// Migrations
	Config.DB.AutoMigrate(&entities.User{})
	Config.DB.AutoMigrate(&entities.Post{})

	// Setup routes
	r := router.SetupRoutes()
	r.Use(gin.Recovery())
	// running
	_ = r.Run(":9000")

}
