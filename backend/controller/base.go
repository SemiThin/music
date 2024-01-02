package controller

import (
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
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Response(code int, data interface{}, msg string, c *gin.Context) {
	json := &JsonStruct{Code: code, Data: data, Msg: msg}
	c.JSON(http.StatusOK, json)
}
