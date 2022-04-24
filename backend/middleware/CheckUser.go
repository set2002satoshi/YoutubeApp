package middleware

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CheckUser() gin.HandlerFunc {
	const SecretKey = "secret"
	return func(c *gin.Context) {
		fmt.Println("動いてる")
		cookie, err := c.Cookie("token")
		token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error){
			return []byte(SecretKey), nil
		})
		fmt.Println(cookie, err)

		if err != nil {
			c.JSON(401, gin.H{"message": "not user"})
			c.Abort()
		}

		c.JSON(200, gin.H{"message": token})

		c.Next()

	    
	}
}
