package controller   

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/sessions/cookie"


	"github.com/set2002satoshi/YoutubeApp/backend/db"
	"github.com/set2002satoshi/YoutubeApp/backend/model"
	"github.com/set2002satoshi/YoutubeApp/backend/service"
)


func Home(c *gin.Context) {
	userDATA, err := service.ClickUser(c)
	if err != nil {
		mes := fmt.Errorf("not user")
		c.JSON(http.StatusForbidden, gin.H{
			"status": 403,
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




func ListUsers(c *gin.Context) {
	_, err := service.ClickUser(c)
	users := []model.Users{}
	DbEngine := db.OpenDB()

	if err != nil {
		fmt.Println("Error1")
		return
	}

	if reps := DbEngine.Find(&users); reps.Error != nil {
		c.JSON(http.StatusNotFound, reps.Error)
		fmt.Println("Error2")
		return
	}
	fmt.Println("success")
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"users":  users,
	})

}



