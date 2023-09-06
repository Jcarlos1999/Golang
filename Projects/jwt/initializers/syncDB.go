package initializers

import "github.com/Jcarlos1999/Golang/Projects/jwt/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}
