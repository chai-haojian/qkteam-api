package main

import (
	"fmt"
	"qkteam-api/config"
	"qkteam-api/dao/mysql"
	"qkteam-api/logger"
	snowflake "qkteam-api/pkg"
	"qkteam-api/routes"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	//加载配置
	if err := config.Init(); err != nil {
		fmt.Printf("Init config faild, err:%v\n", err)
	}
	//初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("Init logger faild, err:%v\n", err)
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")
	//初始化MySQL连接
	if err := mysql.Init(); err != nil {
		fmt.Printf("Init mysql faild, err:%v\n", err)
	}
	defer mysql.Close()
	//初始化雪花算法
	if err := snowflake.Init(viper.GetString("app.start_time"), viper.GetInt64("app.machine_id")); err != nil {
		fmt.Printf("Init snowflake failed,err:%v", err)
		return
	}
	//注册路由
	r := routes.Setup()

	r.Run()
}
