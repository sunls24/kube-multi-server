package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"kube-multi-server/controller"
	"kube-multi-server/database"
	"kube-multi-server/service"
)

const base = "/kube"
const id = "/:id"

func InitKube(r *gin.RouterGroup, connect *sql.DB) {
	db := database.NewKubeDB(connect)
	svc := service.NewKubeService(db)
	kube := controller.NewKubeController(svc)

	r.GET(base, kube.ListKube)
	r.GET(base+id, kube.GetKube)
	r.POST(base, kube.PostKube)
	r.PATCH(base, kube.PutKube)
	r.DELETE(base+id, kube.DeleteKube)
}
