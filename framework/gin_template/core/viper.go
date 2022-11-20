package core

import (
	"flag"
	"fmt"
	"gin_template/core/internal"
	"gin_template/global"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

// 优先级: 命令行 > 环境变量 > 默认值
func Viper(path ...string) *viper.Viper {
	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()

		if config == "" {
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDebugFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\\n", gin.EnvGinMode, internal.ConfigDebugFile)
				case gin.TestMode:
					config = internal.ConfigTestFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\\n", gin.EnvGinMode, internal.ConfigTestFile)
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\\n", gin.EnvGinMode, internal.ConfigReleaseFile)
				default:
				}
			}
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper 传递的值,config的路径为%s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err.Error()))
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("config file changed: %s \n", in.Name)
		if err := v.Unmarshal(&global.GLA_CONFIG); err != nil {
			fmt.Println(err)
		}

		fmt.Printf("config:%+v \n", global.GLA_CONFIG)
	})
	if err := v.Unmarshal(&global.GLA_CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
