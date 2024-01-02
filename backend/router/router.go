package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"music/controller"
	"net/http"
)

func Router() *gin.Engine {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viperErr := viper.ReadInConfig()
	if viperErr != nil {
		fmt.Printf("read database failed: %v", viperErr)
	}

	appDebug := viper.Get("APP_DEBUG")
	if appDebug == false {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hell world")
	})

	song := r.Group("/song")
	{
		song.POST("/all", controller.Song{}.All)
	}

	return r
}
