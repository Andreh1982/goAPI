package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"goAPI/database"
	"goAPI/models"
	"goAPI/shared"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	fmt.Println("Home Page")
}

func TodasPersonalidades(c *gin.Context) {
	var p []models.Personalidade
	database.DB.Find(&p)
	c.IndentedJSON(http.StatusOK, p)
	shared.ZapLogCustom([]string{"Retornando todas as entradas"}, "info")
	size := "Total de entradas: " + strconv.Itoa(len(p))
	shared.ZapLogCustom([]string{size}, "info")

}

func RetornaUmaPersonalidade(c *gin.Context) {
	var personalidade models.Personalidade

	if err := database.DB.Where("Id = ?", c.Param("id")).First(&personalidade).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Entrada não localizada."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": personalidade})
	shared.ZapLogCustom([]string{personalidade.Nome}, "info")
	shared.ZapLogCustom([]string{personalidade.Historia}, "info")

}

func CriaUmaNovaPersonalidade(c *gin.Context) {
	var novaPersonalidade models.Personalidade

	shared.ZapLogCustom([]string{"Criando nova entrada..."}, "info")

	if err := c.ShouldBindJSON(&novaPersonalidade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Requisição falhou ao inserir uma nova entrada.": err.Error()})
		return
	}

	personalidade := models.Personalidade{Nome: novaPersonalidade.Nome, Historia: novaPersonalidade.Historia}
	database.DB.Create(&personalidade)

	c.JSON(http.StatusOK, gin.H{"data": personalidade})
	shared.ZapLogCustom([]string{novaPersonalidade.Nome}, "info")
	shared.ZapLogCustom([]string{novaPersonalidade.Historia}, "info")
}

func DeletaUmaPersonalidade(c *gin.Context) {

	var personalidade models.Personalidade

	// Get model if exist
	if err := database.DB.Where("id = ?", c.Param("id")).First(&personalidade).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Entrada não localizada!"})
		return
	}

	database.DB.Delete(&personalidade)

	c.JSON(http.StatusOK, gin.H{"Entrada deletada com êxito!": true})

	shared.ZapLogCustom([]string{"[DEL] Deletando entrada..."}, "info")
	shared.ZapLogCustom([]string{personalidade.Nome}, "info")
}

func EditaPersonalidade(c *gin.Context) {

	var personalidade models.Personalidade

	if err := database.DB.Where("id = ?", c.Param("id")).First(&personalidade).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Entrada não localizada!"})
		return
	}

	var updatePersonalidade models.Personalidade
	if err := c.ShouldBindJSON(&updatePersonalidade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Erro ao atualizar entrada.": err.Error()})
		return
	}
	database.DB.Model(&personalidade).Updates(updatePersonalidade)

	c.JSON(http.StatusOK, gin.H{"Entrada atualizada com sucesso.": personalidade})

}
