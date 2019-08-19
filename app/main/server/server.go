package server

import (
	"battery-anlysis-platform/app/main/conf"
	"battery-anlysis-platform/app/main/dao"
	"battery-anlysis-platform/app/main/middleware"
	"battery-anlysis-platform/pkg/config"
	"github.com/gin-gonic/gin"
)

func Start() error {
	confParams := &conf.Params{}
	config.Load("conf/app.ini", "main", confParams)

	// 初始化数据库连接
	dao.InitMySQL(confParams.MysqlUri)

	gin.SetMode(confParams.RunMode)
	r := gin.Default()
	r.Use(middleware.Session(confParams.SecretKey))
	register(r)

	return r.Run(confParams.HttpAddr)
}