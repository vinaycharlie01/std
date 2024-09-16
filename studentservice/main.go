package main

import (
	"log"
	"myapp/studentservice/model"
	"myapp/studentservice/router"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "user:user_password@tcp(127.0.0.1:3306)/mydatabase?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(model.NEWStudnetTable())
	r := gin.Default()

	router.RegisterRouter(r, db)

	r.Run()

}
