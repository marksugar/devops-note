package main

import (
	"fmt"

	"tgbot/dao/mysql"
	"tgbot/logger"
	"tgbot/service"
	"tgbot/settings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	// 1,加载配置文件
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed,err:%v\n", err)
		return
	}
	// 2,初始化日志
	if err := logger.Init(viper.GetString("app.mode")); err != nil {
		fmt.Printf("init logger failed,err:%v\n", err)
		return
	}
	// 延迟注册追加缓冲区日志
	defer zap.L().Sync()
	zap.L().Info("logger init success ...")
	// 3,初始化数据库
	if err := mysql.Init(); err != nil {
		fmt.Printf("init logger failed,err:%v\n", err)
		return
	}
	defer mysql.Close()

	//
	go service.TgbotService.TgCallbackInit()
	zap.L().Info("TgCallbackInit init success ...")

	// go func() {
	// 	// for now := range time.Tick(12 * time.Hour) { // 10秒运行一次
	// 	for now := range time.Tick(1 * time.Minute) { // 10秒运行一次
	// 		zap.L().Info(fmt.Sprintf("同步消息: %s", now))
	// 		err := service.SendMessage(7136576847, "我是一个测试") // 推送消息 ，id为群的id或者对话的个人id
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 	}
	// }()

	fmt.Println("ok")
	select {}
}
