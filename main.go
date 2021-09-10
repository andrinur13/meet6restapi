package main

import (
	"fmt"
	"meet6restapi/config"
	"meet6restapi/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := gin.Default()

	db := config.InitDB()
	DBConn := &controllers.DBConn{DB: db}

	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	fmt.Println("Server running on port 8080")

	router.POST("/person", DBConn.CreatePerson)
	router.GET("/persons", DBConn.GetPersons)
	router.GET("/person/:id", DBConn.GetPersonByID)
	router.PUT("/person", DBConn.UpdatePerson)
	router.DELETE("/person/:id", DBConn.DeletePerson)

	router.Run(":8080")
}
