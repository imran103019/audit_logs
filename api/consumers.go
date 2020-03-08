package api

import (
	"github.com/imran103019/audit_logs/conn"
	"github.com/imran103019/audit_logs/dtos"
	repos "github.com/imran103019/audit_logs/repos"
	"github.com/imran103019/audit_logs/utils/errorutil"
	"github.com/imran103019/audit_logs/utils/filterutil"
	"github.com/imran103019/audit_logs/utils/paginationutil"
	"github.com/imran103019/audit_logs/utils/cacheutil"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"fmt"
	"github.com/spf13/viper"
)

var consumerRepo = repos.NewConsumerRepo(conn.Default())

func GetConsumers() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		page := paginationutil.NewPage()
		pageLimit, err := strconv.ParseUint(viper.GetString("MAX_PAGE_LIMIT"), 10, 32)
		if err == nil {
			page.Limit = uint(pageLimit)
		}

		pageCurrent, err := strconv.ParseUint(c.QueryParam("page"), 10, 32)
		if err == nil {
			page.Current = uint(pageCurrent)
		}
		pgntr := paginationutil.NewPaginator(page.PageLimit(), page.Current)

		f := filterutil.ConsumerFilterElements{
			AppName:      "",
			FromDate:     "",
			ToDate:       "",
		}

		resp, err := consumerRepo.GetConsumers(f, pgntr)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, errorutil.NewCustomError(http.StatusInternalServerError,
				errorutil.ErrSomethingWentWrong))
		}

		return c.JSON(http.StatusOK, dtos.Response{
			Data:       resp,
			Pagination: pgntr,
		})
	}
}


func StoreConsumer() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		consumerReq := &dtos.ConsumerRequest{}
		if err := c.Bind(consumerReq); err != nil {
			return err
		}
		errConsumer := consumerRepo.SetConsumer(consumerReq)

		if errConsumer != nil {
			return c.JSON(http.StatusInternalServerError, errorutil.NewCustomError(http.StatusInternalServerError,
				errorutil.ErrSomethingWentWrong))
		}

		return c.JSON(http.StatusCreated, dtos.SuccessResponse{
			Success: true,
		})
	}
}

func UpdateConsumer() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		token := c.Param("token")
		consumerReq := &dtos.ConsumerRequest{}
		if err := c.Bind(consumerReq); err != nil {
			return err
		}
		errConsumer := consumerRepo.UpdateConsumer(consumerReq, token)

		if errConsumer != nil {
			return c.JSON(http.StatusInternalServerError, errorutil.NewCustomError(http.StatusInternalServerError,
				errorutil.ErrSomethingWentWrong))
		}

		return c.JSON(http.StatusOK, dtos.SuccessResponse{
			Success: true,
		})
	}
}

func DeleteConsumer() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		token := c.Param("token")
		errConsumer := consumerRepo.DeleteConsumer(token)

		if errConsumer != nil {
			return c.JSON(http.StatusInternalServerError, errorutil.NewCustomError(http.StatusInternalServerError,
				errorutil.ErrSomethingWentWrong))
		}

        key := fmt.Sprintf("consumer_%s",token)
        if err := cacheutil.ClearCache(key); err != nil {
        	fmt.Println(err)
        }

		return c.JSON(http.StatusOK, dtos.SuccessResponse{
			Success: true,
		})
	}
}

func SetAllConsumerAppTokenCache() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		go consumerRepo.SetAllConsumerAppTokenCache()
		return c.JSON(http.StatusOK, dtos.SuccessResponse{
			Success: true,
		})
	}
}


