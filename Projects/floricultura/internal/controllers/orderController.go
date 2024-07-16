package controllers

import (
	"errors"
	"net/http"

	"gihub.com/Jcarlos1999/Golang/Projects/floricultura/internal/initializers"
	"gihub.com/Jcarlos1999/Golang/Projects/floricultura/internal/models"
	"github.com/gin-gonic/gin"
)

var orderBody struct {
	Items      []models.Flower
	FinalPrice float64
}

func CreateOrder(c *gin.Context) {

	if c.Bind(&orderBody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Errpr": "failed to read body",
		})
		return
	}

	err := verifyFlowers(orderBody.Items)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	order, err := models.NewOrder(orderBody.Items, orderBody.FinalPrice)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to create order",
		})
		return
	}
	result := initializers.DB.FirstOrCreate(&order)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Success": "Data successfully created",
	})

}

func verifyFlowers(items []models.Flower) error {
	var flower models.Flower
	for _, f := range items {
		initializers.DB.Find(&flower, "name = ?", f.Name)
		if flower.ID == 0 {
			return errors.New("Flower " + f.Name + " dont exist")
		}
	}

	return nil
}
