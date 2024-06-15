package main

import (
	"slow/config"
	"slow/controller"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

func initLogger(conf config.LogConfig) func() {
	logger := &lumberjack.Logger{
		Filename:   conf.Filename,
		MaxSize:    conf.MaxSize,    // 文件大小
		MaxBackups: conf.MaxBackups, // 备份个数
		LocalTime:  conf.LocalTime,  // 本地时间
		Compress:   conf.Compress,   // 压缩
	}
	logrus.SetOutput(logger)
	logrus.SetLevel(logrus.Level(conf.Level)) // iota类型转换
	logrus.SetReportCaller(conf.ReportCaller)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	return func() {
		logger.Close()
	}
}
func main() {

	gconf := config.ExporterConfig{
		Log: config.LogConfig{
			"logs/slow-mysql.log",
			100,
			7,
			true,
			true,
			int(logrus.DebugLevel),
			true,
		},
	}
	// 初始化日志
	closelog := initLogger(gconf.Log)
	defer closelog()

	r := gin.Default()
	r.POST("/post", controller.PostJson)
	r.Run(":9999")
}

// GOOS：目标平台的操作系统（darwin、freebsd、linux、windows）
// GOARCH：目标平台的体系架构（386、amd64、arm）

// go env -w GOARCH=amd64
// go env -w GOOS=windows
// go run .\main.go

// go env -w GOARCH=arm64
// go env -w GOOS=linux
// go build main.go
