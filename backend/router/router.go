package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"music/api"
	"music/middleware"
	"net/http"
)

func Router() *gin.Engine {
	ginModel := viper.GetString("gin.model")
	gin.SetMode(ginModel)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hell world")
	})

	user := r.Group("/user")
	{
		user.POST("/login", api.User{}.Login)

		user.Use(middleware.Auth())
		user.POST("/logout", api.User{}.Logout)

		user.POST("/update", func(c *gin.Context) {
			c.String(http.StatusOK, "user update")
		})
	}

	song := r.Group("/song")
	{
		song.POST("/all", api.Song{}.All)
		song.POST("/list", api.Song{}.List)
	}

	album := r.Group("/album")
	{
		album.POST("/list", api.Album{}.List)
	}

	return r
}
