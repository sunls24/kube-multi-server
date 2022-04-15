package database

import (
	"database/sql"
	"github.com/pkg/errors"
	"kube-multi-server/model"
	"kube-multi-server/util"
)

type KubeDB interface {
	AddKube(kube model.Kube) (int64, error)
	ListKube(id int) (model.KubeList, error)
	UpdateKube(kube model.Kube) error
	DeleteKube(id int) error
}

func NewKubeDB(connect *sql.DB) KubeDB {
	return &kubeDB{connect: connect}
}

type kubeDB struct {
	connect *sql.DB
}

func (k *kubeDB) AddKube(kube model.Kube) (int64, error) {
	result, err := k.connect.Exec("INSERT INTO kube (name, config, namespace) VALUES (?, ?, ?)", kube.Name, kube.Config, kube.Namespace)
	if err != nil {
		return 0, errors.Wrap(err, "Exec INSERT")
	}
	id, err := result.LastInsertId()
	return id, errors.Wrap(err, "LastInsertId")
}

func (k *kubeDB) ListKube(id int) (model.KubeList, error) {
	var rows *sql.Rows
	var err error
	if id <= 0 {
		rows, err = k.connect.Query("SELECT * FROM kube")
	} else {
		rows, err = k.connect.Query("SELECT * FROM kube WHERE id = ?", id)
	}

	if err != nil {
		return nil, errors.Wrap(err, "Query")
	}
	list := model.KubeList{}
	err = list.ScanList(rows)
	return list, errors.Wrap(err, "ScanList")
}

func (k *kubeDB) UpdateKube(kube model.Kube) error {
	if util.IsZero(kube.Id) {
		return errors.New("cannot update，id has zero value")
	}

	_, err := k.connect.Exec("UPDATE kube SET name = ?, namespace = ?, config = ? WHERE id = ?",
		kube.Name, kube.Namespace, kube.Config, kube.Id)
	return errors.Wrap(err, "Exec UPDATE")
}

func (k *kubeDB) DeleteKube(id int) error {
	if util.IsZero(id) {
		return errors.New("cannot delete，id has zero value")
	}
	_, err := k.connect.Exec("DELETE FROM kube WHERE id = ?", id)
	return errors.Wrap(err, "Exec DELETE")
}
