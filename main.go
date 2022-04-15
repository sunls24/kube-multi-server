package main

import (
	"database/sql"
	"flag"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"kube-multi-server/router"
)

func init() {
	// 日志打印格式
	format := new(log.TextFormatter)
	format.FullTimestamp = true
	format.TimestampFormat = "06-01-02 15:04:05"
	log.SetFormatter(format)
}

func main() {
	addr := flag.String("addr", ":8088", "http service address")
	flag.Parse()

	connect, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(errors.Wrap(err, "sql.open"))
	}

	g := gin.Default()
	_ = g.SetTrustedProxies(nil)
	v1 := g.Group("api/v1")
	router.InitCommon(v1)
	router.InitKube(v1, connect)

	log.Info("start http listen addr: ", *addr)
	log.Fatal(g.Run(*addr))
}
