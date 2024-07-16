package initializers

import "gihub.com/Jcarlos1999/Golang/Projects/floricultura/internal/models"

func Sync() {
	DB.AutoMigrate(&models.Flower{}, models.Order{})
	// DB.AutoMigrate(&models.User{}, &models.CreditCard{})
}
