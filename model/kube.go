package model

import "database/sql"

type Kube struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Config    string `json:"config"`
	Namespace string `json:"namespace"`
}

type KubeList []Kube

func (k *Kube) Scan(row *sql.Row) error {
	if row.Err() != nil {
		return row.Err()
	}
	return row.Scan(&k.Id, &k.Name, &k.Config, &k.Namespace)
}

func (kl *KubeList) ScanList(rows *sql.Rows) error {
	var err error
	for rows.Next() {
		k := Kube{}
		err = rows.Scan(&k.Id, &k.Name, &k.Config, &k.Namespace)
		if err != nil {
			return err
		}
		*kl = append(*kl, k)
	}
	return nil
}
