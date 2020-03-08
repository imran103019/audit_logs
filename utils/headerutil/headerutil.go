
package headerutil

import (
	"github.com/labstack/echo"
	"strconv"
	log "github.com/sirupsen/logrus"
)


func GetConsumerIDFromHeader(c echo.Context) uint64 {
	consumerId, err := strconv.ParseUint(c.Request().Header.Get("Consumer-Id"), 10, 32)
	if(err != nil) {
		log.Error(err.Error)
		return 1
	}
	return consumerId
}