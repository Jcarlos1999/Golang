package controllers

import (
	"fmt"
	"net/http"

	"gihub.com/Jcarlos1999/Golang/Projects/floricultura/internal/initializers"
	"gihub.com/Jcarlos1999/Golang/Projects/floricultura/internal/models"
	"github.com/gin-gonic/gin"
)

var bodyFlower struct {
	Name        string
	Price       float64
	HarvestDate string
}

func CreateFlower(c *gin.Context) {

	if c.Bind(&bodyFlower) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	flower, err := models.NewFlower(bodyFlower.Name, bodyFlower.Price, bodyFlower.HarvestDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to create flower",
		})
		return
	}
	result := initializers.DB.Create(&flower)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to insert flower data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Sucess": "Data successfully inserted",
	})

}

func FindFlowers(c *gin.Context) {

	if c.Bind(&bodyFlower) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	fmt.Print(bodyFlower.Name)
	var flower models.Flower
	initializers.DB.First(&flower, "name = ?", bodyFlower.Name)
	if flower.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Flower" + bodyFlower.Name + " dont exist",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Flower": flower,
	})

}
