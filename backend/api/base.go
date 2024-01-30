package api

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Restful interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Show(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type JsonStruct struct {
	Code int    `json:"code"`
	Data any    `json:"data,omitempty"`
	Msg  string `json:"msg"`
}

type Paging struct {
	Page int `json:"page" binding:"omitempty,gte=1"`
	Size int `json:"size" binding:"omitempty,gte=1,lte=20"`
}

func Response(code int, data any, msg string, c *gin.Context) {
	json := &JsonStruct{Code: code, Data: data, Msg: msg}
	c.JSON(http.StatusOK, json)
}

func TimeTemplate(format string) string {
	var timeTemplate string
	switch format {
	case "Y-m-d":
		timeTemplate = "2006-01-02"
	case "Y-m-d H:i":
		timeTemplate = "2006-01-02 15:04"
	case "Y-m-d H:i:s":
		timeTemplate = "2006-01-02 15:04:05"
	case "H:i:s":
		timeTemplate = "15:04:05"
	default:
		timeTemplate = format
	}

	return timeTemplate
}

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}
