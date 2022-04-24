package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	// "github.com/gin-contrib/sessions"
	// "github.com/gin-contrib/sessions/cookie"
	"golang.org/x/crypto/bcrypt"

	"github.com/set2002satoshi/YoutubeApp/backend/db"
	"github.com/set2002satoshi/YoutubeApp/backend/model"
	"github.com/set2002satoshi/YoutubeApp/backend/service"
)

const SecretKey = "secret"

func ListUsers(c *gin.Context) {
	userDATA, err := service.ClickUser(c)
	users := []model.Users{}
	DbEngine := db.OpenDB()

	if err != nil {
		fmt.Println("Error1")
		return
	}

	if reps := DbEngine.Where("ID =?", userDATA).Find(&users); reps.Error != nil {
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

func CreateUser(c *gin.Context) {

	DbEngine := db.OpenDB()
	var Cuser model.CreateUsers

	if err := c.BindJSON(&Cuser); err != nil {
		c.JSON(204, err)
		return
	}

	pass, _ := bcrypt.GenerateFromPassword([]byte(Cuser.Password), 14)

	user := model.Users{
		Name:     &Cuser.Name,
		Email:    &Cuser.Email,
		Password: pass,
	}

	if reps := DbEngine.Create(&user); reps.Error != nil {
		c.JSON(204, reps.Error)
	} else {
		c.JSON(200, gin.H{
			"status": "ok",
			"user":   user,
		})
	}
}

func Login(c *gin.Context) {
	DbEngine := db.OpenDB()
	var CatchUser model.LoginUser
	users := model.Users{}

	if c.Request.Method == "OPTION" {
		err := "Method is OPTION"
		c.JSON(200, err)
		fmt.Println(err)
		return
	}

	//データの受け取り
	if err := c.BindJSON(&CatchUser); err != nil {
		c.JSON(204, err)

		return
	}

	// 受け取ったデータの検索
	if result := DbEngine.Where("Email =?", &CatchUser.Email).First(&users); result.Error != nil {
		c.JSON(204, gin.H{"message": "not found"})
		fmt.Println("err")
		return
	}

	// 取得したデータのpasswordが正しいのか？検証
	err := bcrypt.CompareHashAndPassword(users.Password, []byte(CatchUser.Password))
	if err != nil {
		c.JSON(204, err)
		fmt.Println(err)
		return
	}

	// tokenの生成するオプション
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(users.ID)),
		ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.JSON(204, err)
		return
	}

	fmt.Println("set cookie")
	c.SetCookie("jwt", token, 3600, "/", "localhost", false, false)

	c.JSON(200, token)

}

func Logout(c *gin.Context) {
	_, err := service.ClickUser(c)
	if err != nil {
		mes := fmt.Errorf("not user")
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": mes,
		})
		return
	}
	c.SetCookie("jwt", "not token", 0, "/", "localhost", false, false)

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func MethodOptionHandling(c *gin.Context, p *gin.LogFormatterParams) {
	if p.Method == "OPTION" {
		err := "Method is OPTION"
		c.JSON(200, err)
	}

}
