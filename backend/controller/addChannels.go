package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	// "github.com/set2002satoshi/YoutubeApp/backend/db"
	// "github.com/set2002satoshi/YoutubeApp/backend/model"
	"github.com/set2002satoshi/YoutubeApp/backend/service"
)


func SearchChannels(c *gin.Context) {
	// _, err := service.ClickUser(c)
	// if err != nil {
	// 	mes := fmt.Errorf("not user")
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"status": 401,
	// 		"message": mes,
	// 	})
	// 	return
	// }


	ChannelsList, err := service.SearchChannels(c.Param("Keyword"))
	if err != nil {
		mes := fmt.Errorf("not channels")
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"message": mes,
		})
	} 

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"ChannelsList": ChannelsList,
	})


}

