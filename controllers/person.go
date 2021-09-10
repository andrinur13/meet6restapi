package controllers

import (
	"meet6restapi/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (Conn *DBConn) CreatePerson(c *gin.Context) {
	var person structs.Person

	person.FirstName = c.PostForm("first_name")
	person.LastName = c.PostForm("last_name")

	Conn.DB.Create(&person)

	result := gin.H{
		"result": person,
	}

	c.JSON(http.StatusOK, result)
}

func (Conn *DBConn) GetPersonByID(c *gin.Context) {
	var person structs.Person
	var result = gin.H{}

	id := c.Param("id")

	err := Conn.DB.Where("id = ?", id).First(&person).Error

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": person,
		}
	}

	c.JSON(http.StatusOK, result)

}

func (Conn *DBConn) GetPersons(c *gin.Context) {
	var person []structs.Person
	var result = gin.H{}

	Conn.DB.Find(&person)

	if len(person) <= 0 {
		result = gin.H{
			"result": nil,
		}
	} else {
		result = gin.H{
			"result": person,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (Conn *DBConn) UpdatePerson(c *gin.Context) {

	id := c.Query("id")

	FirstName := c.PostForm("first_name")
	LastName := c.PostForm("last_name")

	var (
		person    structs.Person
		newPerson structs.Person
		result    gin.H
	)

	err := Conn.DB.First(&person, id).Error

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
		c.JSON(http.StatusNotFound, result)
		return
	}

	newPerson.FirstName = FirstName
	newPerson.LastName = LastName

	err = Conn.DB.Model(&person).Updates(newPerson).Error

	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	result = gin.H{
		"result": "update data sukses",
	}

	c.JSON(http.StatusOK, result)
}

func (Conn *DBConn) DeletePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	id := c.Param("id")

	err := Conn.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
		c.JSON(http.StatusNotFound, result)
		return
	}

	err = Conn.DB.Delete(&person).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	result = gin.H{
		"result": "sukes deleted",
	}
	c.JSON(http.StatusOK, result)
	return
}
