package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	// 读取配置文件config
	viper.SetConfigFile("config.yaml")
	// viper.SetConfigName("config")
	// viper.SetConfigType("yaml") // 远程获取配置文件类型，如etcd
	// 配置文件可配置多个目录
	viper.AddConfigPath(".")
	viper.AddConfigPath("./conf")
	// viper.AddConfigPath("$HOME/")

	// 读取配置
	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {            // 处理读取配置文件的错误
		fmt.Printf("viper.ReadInConfig() failed,err: %s \n", err)
		return
	}

	// 实时监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		// 文件修改 触发调用
		fmt.Println("Config file changed reload done:", e.Name)
	})
	return
}
