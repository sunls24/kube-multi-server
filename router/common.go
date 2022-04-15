package router

import (
	"github.com/gin-gonic/gin"
	"kube-multi-server/controller"
)

func InitCommon(r *gin.RouterGroup) {
	r.GET("/healthy", controller.Healthy)
}
