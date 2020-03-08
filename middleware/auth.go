package middleware

import (  
    "fmt"
    "github.com/imran103019/audit_logs/repos"
    "github.com/imran103019/audit_logs/conn"
    "time"
    "github.com/imran103019/audit_logs/utils/cacheutil"
    "github.com/spf13/viper"
    "encoding/json"
    "github.com/imran103019/audit_logs/dtos"
    "github.com/labstack/echo"
    "net/http"
    "strconv"
    log "github.com/sirupsen/logrus"
)

func Authorization() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            token := c.Request().Header.Get("App-Key")
            key := fmt.Sprintf("consumer_%s",token)
            consumer, err := cacheutil.GetCache(key)
            
            if(err != nil) {
                log.Error("Error while fetching consumer from cache: ", err.Error)
            }

            if(consumer != "") {

                var cacheConsumer dtos.ConsumerResponse 
                if err := json.Unmarshal([]byte(consumer), &cacheConsumer); err != nil {
                    log.Error("Error while unmarshal consumer cache: ", err.Error())
                }
                c.Request().Header.Set("Consumer-Id", strconv.Itoa(cacheConsumer.Id))
                return next(c)
            }

            repo := repos.NewConsumerRepo(conn.Default())
            resp, err := repo.GetConsumerByToken(token)

            if(err != nil) {
                log.Error("Error while getting consumer from DB: ", err.Error())
                return c.JSON(http.StatusUnauthorized, dtos.UnauthorizedResponse{
                    Message:"Unauthorized",
                })
            }    

            respByte, err := json.Marshal(resp)
            if err := cacheutil.SetCache(key, string(respByte),
                viper.GetDuration("GLOBAL_CACHE_DURATION")*time.Second); err != nil {
                log.Error("Failed to set cache for consumers:", err.Error())
            }
            if(resp.Id > 0) {
                c.Request().Header.Set("Consumer-Id", strconv.Itoa(resp.Id))
                return next(c)
            }

            return c.JSON(http.StatusUnauthorized, dtos.UnauthorizedResponse{
                Message:"Unauthorized",
            })

        }
    }
}
