package route

import (
	"github.com/labstack/echo"
	"github.com/imran103019/audit_logs/api"
	"github.com/imran103019/audit_logs/dummy"
	"github.com/labstack/echo/middleware"
	"github.com/imran103019/audit_logs/helper"
	customMiddleware "github.com/imran103019/audit_logs/middleware"
	"github.com/spf13/viper"
)

func Init() *echo.Echo {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	route := e.Group("/api/v1")
	route.Use(customMiddleware.Authorization())
	route.POST("/logs",       api.StoreActivity())
	route.PATCH("/logs/:id",  api.UpdateActivity())
	route.GET("/logs",        api.GetActivities())

	consumerRoute := e.Group("/api/v1")
	if(viper.GetBool("ENABLE_BASIC_AUTH")) {
		consumerRoute.Use(middleware.BasicAuth(helper.ValidateUser))
	}

	consumerRoute.POST("/consumers",          api.StoreConsumer())
	consumerRoute.PATCH("/consumers/:token",  api.UpdateConsumer())
	consumerRoute.GET("/consumers",           api.GetConsumers())
	consumerRoute.DELETE("/consumers/:token", api.DeleteConsumer())
	consumerRoute.POST("/consumers/cache",    api.SetAllConsumerAppTokenCache())
    
	return e
}