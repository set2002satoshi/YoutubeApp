package controller   

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/set2002satoshi/YoutubeApp/backend/db"
	"github.com/set2002satoshi/YoutubeApp/backend/model"
	"github.com/set2002satoshi/YoutubeApp/backend/service"
)


func Home(c *gin.Context) {
	userDATA, err := service.ClickUser(c)
	if err != nil {
		mes := fmt.Errorf("not user")
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"message": mes,
		})
		return 
	}
	
	DbEngine := db.OpenDB()

	var user model.Users

	DbEngine.Select("name").Where("ID =?", userDATA).Find(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"user": user,
	})


}



func Test(c *gin.Context) {
	
	DbEngine := db.OpenDB()

	var user model.Users

	DbEngine.Select("name").Find(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"user": user,
	})


}



