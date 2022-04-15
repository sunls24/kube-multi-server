package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Healthy(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}
