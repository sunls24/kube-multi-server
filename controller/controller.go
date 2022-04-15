package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	statusSuccess = 0
	statusFail    = 1
)

type body struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, body{Status: statusSuccess, Data: data})
}

func Error(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, body{Status: statusFail, Msg: err.Error()})
}
