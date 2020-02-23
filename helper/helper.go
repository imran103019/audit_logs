package helper
 
import (
   "github.com/labstack/echo"
   "github.com/spf13/viper"
)
 
func ValidateUser(username, password string, c echo.Context) (bool, error) {
    if username == viper.GetString("BASIC_AUTH_USERNAME") && password == viper.GetString("BASIC_AUTH_PASSWORD") {
        return true, nil
    }
    return false, nil
}