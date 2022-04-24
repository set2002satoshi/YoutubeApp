package service

import (
	// "os"
	"fmt"
// 
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"


	// "github.com/set2002satoshi/GoGinProcess/model"
	// "github.com/set2002satoshi/GoGinProcess/db"
)

const SecretKey = "secret"


func ClickUser(c *gin.Context) (userID string, err error) {
	 
	cookie, err := c.Cookie("jwt")
	// DbEngine := db.OpenDB()

	// fmt.Println(os.Getenv("key"))

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		fmt.Errorf("not user")
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims)


	return claims.Issuer, nil

}


// func comparePassword(hashedPwd string, plainPassword []byte) bool {
// 	byteHash := []byte(hashedPwd)
// 	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
// 	if err != nil {
// 		log.Println(err)
// 		return false
// 	}
// 	return true
// }