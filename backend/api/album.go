package api

import (
	"github.com/gin-gonic/gin"
	orm "music/database"
	"music/model"
	"time"
)

type Album struct {
	//Restful
}

type AlbumList struct {
	Paging
	SingerId int `json:"singer_id" binding:"omitempty,gte=1"`
}

func (a Album) List(c *gin.Context) {
	var param AlbumList
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
	singerId := param.SingerId

	var album []model.Album
	db := orm.Db

	if singerId > 0 {
		db = db.Where("singer_id = ?", singerId)
	}

	db.Offset(offset).Limit(limit).Find(&album)

	timeTemplate := TimeTemplate("Y-m-d H:i")
	for i := range album {
		album[i].Pubdate = time.Unix(album[i].Pubtime, 0).Format(timeTemplate)
	}

	Response(0, album, "success", c)
}
