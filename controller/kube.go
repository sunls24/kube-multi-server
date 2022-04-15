package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"kube-multi-server/model"
	"kube-multi-server/service"
	"strconv"
)

type kubeController struct {
	service service.KubeService
}

func NewKubeController(s service.KubeService) *kubeController {
	return &kubeController{service: s}
}

func (k *kubeController) ListKube(ctx *gin.Context) {
	list, err := k.service.ListKube()
	if err != nil {
		Error(ctx, err)
		return
	}
	Success(ctx, list)
}

func (k *kubeController) GetKube(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		Error(ctx, errors.Wrap(err, "strconv.Atoi"))
		return
	}
	kube, err := k.service.GetKube(id)
	if err != nil {
		Error(ctx, err)
		return
	}
	Success(ctx, kube)
}

func (k *kubeController) PostKube(ctx *gin.Context) {
	var kube model.Kube
	if err := ctx.BindJSON(&kube); err != nil {
		Error(ctx, err)
		return
	}
	id, err := k.service.AddKube(kube)
	if err != nil {
		Error(ctx, err)
		return
	}
	Success(ctx, map[string]interface{}{"id": id})
}

func (k *kubeController) PutKube(ctx *gin.Context) {
	var kube model.Kube
	if err := ctx.BindJSON(&kube); err != nil {
		Error(ctx, err)
		return
	}

	if err := k.service.UpdateKube(kube); err != nil {
		Error(ctx, err)
		return
	}
	Success(ctx, nil)
}

func (k *kubeController) DeleteKube(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		Error(ctx, errors.Wrap(err, "strconv.Atoi"))
		return
	}
	err = k.service.DeleteKube(id)
	if err != nil {
		Error(ctx, err)
		return
	}
	Success(ctx, nil)
}
