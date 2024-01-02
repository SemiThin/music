package controller

import (
	"github.com/gin-gonic/gin"
	orm "music/database"
	"music/model"
)

type Song struct {
	//Restful
}

func (s Song) All(c *gin.Context) {

	var song []model.Song
	orm.Db.Preload("Album").Find(&song)

	Response(0, song, "success", c)
}
