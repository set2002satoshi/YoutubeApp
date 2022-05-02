package routes

import (

	"time"
	// "github.com/gin-contrib/sessions"
	// "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/set2002satoshi/YoutubeApp/backend/controller"
	// "github.com/set2002satoshi/GoGinProcess/middleware"
)

func SetUpRoute() {

	router := gin.Default()


	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIOMS",
		},

		AllowHeaders: []string{
			"*",
		},

		AllowCredentials: true,

		MaxAge: 24 * time.Hour,
	}))

	router.GET("/", controller.Home)

	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// router.Use(cors.New(config))


	app := router.Group("/app")
	{
		app.GET("/list", controller.ListUsers)
	}
	
	user := router.Group("/user")
	{
		user.POST("/Cuser", controller.CreateUser)
		user.POST("/login", controller.Login)
		user.OPTIONS("/login", controller.Login)
		user.GET("/logout", controller.Logout)

	}

	router.Run(":8080")

}
