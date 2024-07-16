package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Jcarlos1999/Golang/Projects/jwt/initializers"
	"github.com/Jcarlos1999/Golang/Projects/jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {

	// Get the cookie off req
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	// decode /validade it

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// find the user with token sub
		var user models.User

		initializers.DB.First(&user, "id = ?", claims["sub"])
		fmt.Println(claims["sub"])
		if user.ID == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// attach to req

		c.Set("user", user)

		// continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
