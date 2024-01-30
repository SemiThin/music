package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	orm "music/database"
	"music/model"
	"time"
)

type User struct {
	//Restful
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CustomerClaims struct {
	Username string `json:"username"`
	Gender   string `json:"gender"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
}

func (u User) Login(c *gin.Context) {
	var param UserLogin
	err := c.ShouldBind(&param)
	if err != nil {
		Response(100, "", err.Error(), c)
		return
	}

	username := param.Username
	password := MD5(param.Password)

	var user model.User
	db := orm.Db

	where := &UserLogin{Username: username, Password: password}
	result := db.Where(where).First(&user)
	if result.RowsAffected == 0 {
		Response(404, "", err.Error(), c)
		return
	}

	token, err := generateToken(user)
	if err != nil {
		Response(403, token, "generateToken fail", c)
		return
	}
	Response(0, token, "success", c)
}

func (u User) Logout(c *gin.Context) {
	uid, bool := c.Get("uid")
	fmt.Println("uid:", uid)
	fmt.Println("exists:", bool)
	Response(0, "", "success", c)
}

func generateToken(user model.User) (string, error) {
	jwtSignKey := viper.GetString("gin.jwt")
	var signingKey = []byte(jwtSignKey)
	fmt.Println(signingKey)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.Id,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 设置过期时间为 24 小时
	})

	token, err := claims.SignedString(signingKey)
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT: %v", err)
	}

	return token, nil
}
