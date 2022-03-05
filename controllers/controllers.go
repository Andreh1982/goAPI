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
	personalidade, err := database.GetPersons(a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, personalidade)
}

func (a *APIEnv) GetPerson(c *gin.Context) {
	id := c.Params.ByName("id")
	personalidade, exists, err := database.GetPersonByID(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Personalidade não encontrada."})
		return
	}
	c.IndentedJSON(http.StatusOK, personalidade)
}

func (a *APIEnv) CreatePerson(c *gin.Context) {
	personalidade := models.Personalidade{}
	err := c.BindJSON(&personalidade)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := a.DB.Create(&personalidade).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, personalidade)
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
		c.JSON(http.StatusNotFound, gin.H{"Error": "Personalidade não encontrada."})
		return
	}
	err = database.DeletePerson(id, a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": "Personalidade deletada."})
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
		c.JSON(http.StatusNotFound, gin.H{"Error": "Personalidade não existe."})
		return
	}
	updatedPerson := models.Personalidade{}
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
