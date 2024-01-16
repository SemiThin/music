package api

import (
	"github.com/gin-gonic/gin"
	orm "music/database"
	"music/model"
)

type Song struct {
	//Restful
}

type SongList struct {
	Paging
	AlbumId int `json:"album_id" binding:"omitempty,gte=1"`
}

func (s Song) All(c *gin.Context) {

	var song []model.Song
	orm.Db.Preload("Album").Find(&song)

	Response(0, song, "success", c)
}

func (s Song) List(c *gin.Context) {

	var param SongList
	err := c.ShouldBind(&param)
	if err != nil {
		Response(100, "", err.Error(), c)
		return
	}

	if param.Page == 0 {
		param.Page = 1
	}

	if param.Size == 0 {
		param.Size = 20
	}

	offset := param.Page - 1
	limit := param.Size
	albumId := param.AlbumId

	var song []model.Song
	db := orm.Db

	if albumId > 0 {
		db = db.Where("album_id = ?", albumId)
	}

	db.Preload("Album").Offset(offset).Limit(limit).Find(&song)

	Response(0, song, "success", c)
}
