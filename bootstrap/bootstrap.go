package bootstrap

import (
	"github.com/imran103019/audit_logs/conn"
	"github.com/spf13/viper"
	"github.com/imran103019/audit_logs/routes"

)

func Run() {
	conn.Connect()
	conn.ConnectRedis()
	// conn.Ping()
	apiRouter := route.Init()
	apiRouter.Logger.Fatal(apiRouter.Start(viper.GetString("APP_PORT")))
}

