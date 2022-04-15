package service

import (
	"github.com/pkg/errors"
	"kube-multi-server/database"
	"kube-multi-server/model"
)

type KubeService interface {
	AddKube(kube model.Kube) (int64, error)
	GetKube(id int) (model.Kube, error)
	ListKube() (model.KubeList, error)
	UpdateKube(kube model.Kube) error
	DeleteKube(id int) error
}

func NewKubeService(db database.KubeDB) KubeService {
	return &kubeService{db: db}
}

type kubeService struct {
	db database.KubeDB
}

func (k *kubeService) AddKube(kube model.Kube) (int64, error) {
	return k.db.AddKube(kube)
}

func (k *kubeService) GetKube(id int) (model.Kube, error) {
	list, err := k.db.ListKube(id)
	if err != nil {
		return model.Kube{}, err
	}
	if len(list) == 0 {
		return model.Kube{}, errors.New("not found kube")
	}
	return list[0], nil
}

func (k *kubeService) ListKube() (model.KubeList, error) {
	return k.db.ListKube(0)
}

func (k *kubeService) UpdateKube(kube model.Kube) error {
	return k.db.UpdateKube(kube)
}

func (k *kubeService) DeleteKube(id int) error {
	return k.db.DeleteKube(id)
}
