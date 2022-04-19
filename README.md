# kube-multi-server

## 多集群管理工具 server 端

### Run

需要开启 CGO；添加参数 `-db ./data.db` 来指定数据库文件（仅开发时）

```shell
CGO_ENABLED=1 go run ./main.go -db ./data.db
```

### Build

```shell
docker build -t kube-multi-server:v0.1 .
```

### Deploy

```shell
kubectl create -f deploy.yaml
```
