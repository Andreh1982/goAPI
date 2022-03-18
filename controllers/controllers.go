package controllers

import (
	"fmt"
	"net/http"

	"goAPI/database"
	"goAPI/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type APIEnv struct {
	DB *gorm.DB
}

func (a *APIEnv) GetPersons(c *gin.Context) {
	CorsSetup(c)
	person, err := database.GetPersons(a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, person)
}

func (a *APIEnv) GetPerson(c *gin.Context) {
	id := c.Params.ByName("id")
	person, exists, err := database.GetPersonByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Person not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, person)
}

func (a *APIEnv) CreatePerson(c *gin.Context) {
	person := models.Person{}
	err := c.BindJSON(&person)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := a.DB.Create(&person).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, person)
}

func (a *APIEnv) DeletePerson(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := database.GetPersonByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Person not found."})
		return
	}
	err = database.DeletePerson(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": "Person deleted."})
}

func (a *APIEnv) UpdatePerson(c *gin.Context) {
	id := c.Params.ByName("id")
	_, exists, err := database.GetPersonByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Person not exists."})
		return
	}
	updatedPerson := models.Person{}
	err = c.BindJSON(&updatedPerson)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := database.UpdatePerson(a.DB, &updatedPerson); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	a.GetPerson(c)
}

func CorsSetup(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
}
