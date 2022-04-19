package main

import (
	"database/sql"
	"flag"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"kube-multi-server/router"
	"strings"
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
	dbPath := flag.String("db", "/opt/data/data.db", "sqlite file path")
	flag.Parse()

	connect, err := sql.Open("sqlite3", *dbPath)
	if err != nil {
		log.Fatal(errors.Wrap(err, "sql.open"))
	}

	err = createTable(connect)
	if err != nil {
		log.Fatal(errors.Wrap(err, "create SQL"))
	}

	g := gin.Default()
	_ = g.SetTrustedProxies(nil)
	v1 := g.Group("api/v1")
	router.InitCommon(v1)
	router.InitKube(v1, connect)

	log.Info("start http listen addr: ", *addr)
	log.Fatal(g.Run(*addr))
}

const sqlDir = "./sql/"

// 创建数据表，每次启动都会执行，所以SQL中需包含(if not exists)
func createTable(connect *sql.DB) error {
	files, err := ioutil.ReadDir(sqlDir)
	if err != nil {
		return errors.Wrapf(err, "read dir %s", sqlDir)
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		if !strings.Contains(strings.ToLower(f.Name()), ".sql") {
			continue
		}

		path := sqlDir + f.Name()
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return errors.Wrapf(err, "read file %s", path)
		}
		_, err = connect.Exec(string(data))
		if err != nil {
			return errors.Wrapf(err, "exec SQL (%s)", path)
		}
	}
	return nil
}
