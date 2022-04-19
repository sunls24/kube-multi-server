# kube-multi-server

## 多集群管理工具 server 端

### Run

需要开启 CGO；添加参数 `-db ./data.db` 来指定数据库文件（仅开发过程）

```shell
CGO_ENABLED=1 go run ./main.go -db ./data.db
```

