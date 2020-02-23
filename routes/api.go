package route

import (
	"github.com/labstack/echo"
	"github.com/imran103019/audit_logs/api/controllers"
	"github.com/labstack/echo/middleware"
	"github.com/imran103019/audit_logs/helper"
	"github.com/spf13/viper"
)

func Init() *echo.Echo {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	activity := e.Group("/api/v1")
	if(viper.GetBool("ENABLE_BASIC_AUTH")) {
		activity.Use(middleware.BasicAuth(helper.ValidateUser))
	}
	activity.POST("/logs",       controllers.StoreActivity())
	activity.PATCH("/logs/:id",  controllers.UpdateActivity())
	activity.GET("/logs",        controllers.GetActivities())
	activity.DELETE("/logs/:id", controllers.DeleteActivity())


	return e
}