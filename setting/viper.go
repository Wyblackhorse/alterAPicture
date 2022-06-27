package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init(templateUrl string) (err error) {
	//指定配置文件 如果是 json 就写json
	viper.SetConfigFile(templateUrl)
	//指定文件路径
	viper.AddConfigPath(".")
	//读取配置文件
	err = viper.ReadInConfig()
	//读取配置信息失败
	if err != nil {
		fmt.Println("读取配置失败:", err)
	}
	//监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了!")
	})
	fmt.Println("读取配置成功")
	return err
}
