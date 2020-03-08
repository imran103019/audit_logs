package api

import (
	"github.com/imran103019/audit_logs/conn"
	"github.com/imran103019/audit_logs/dtos"
	repos "github.com/imran103019/audit_logs/repos"
	"github.com/imran103019/audit_logs/utils/errorutil"
	"github.com/imran103019/audit_logs/utils/filterutil"
	"github.com/imran103019/audit_logs/utils/headerutil"
	"github.com/imran103019/audit_logs/utils/paginationutil"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"fmt"
)

var activityRepo = repos.NewActivityRepo(conn.Default())

func GetActivities() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		page := paginationutil.NewPage()
		pageLimit, err := strconv.ParseUint(c.QueryParam("per_page"), 10, 32)
		if err == nil {
			page.Limit = uint(pageLimit)
		}

		pageCurrent, err := strconv.ParseUint(c.QueryParam("page"), 10, 32)
		if err == nil {
			page.Current = uint(pageCurrent)
		}
		pgntr := paginationutil.NewPaginator(page.PageLimit(), page.Current)

		consumerId := headerutil.GetConsumerIDFromHeader(c)

		f := filterutil.ActivityFilterElements{
			Source:       c.QueryParam("source"),
			EntityName:   c.QueryParam("entity_name"),
			EntityID:     c.QueryParam("entity_id"),
			Type:         c.QueryParam("type"),
			ActionBy:     c.QueryParam("action_by"),
			FromDate:     c.QueryParam("from_date"),
			ToDate:       c.QueryParam("to_date"),
			ConsumerID:   consumerId,
		}

		resp, err := activityRepo.GetActivities(f, pgntr)

		if err != nil {

			fmt.Print(err)
			return c.JSON(http.StatusInternalServerError, errorutil.NewCustomError(http.StatusInternalServerError,
				errorutil.ErrSomethingWentWrong))
		}

		return c.JSON(http.StatusOK, dtos.Response{
			Data:       resp,
			Pagination: pgntr,
		})
	}
}


func StoreActivity() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		activityReq := &dtos.ActivityRequest{}
		if err := c.Bind(activityReq); err != nil {
			return err
		}
		consumerId := headerutil.GetConsumerIDFromHeader(c)
		erract := activityRepo.SetActivity(activityReq, consumerId)

		if erract != nil {
			return c.JSON(http.StatusInternalServerError, errorutil.NewCustomError(http.StatusInternalServerError,
				errorutil.ErrSomethingWentWrong))
		}

		return c.JSON(http.StatusCreated, dtos.SuccessResponse{
			Success: true,
		})

	}
}

func UpdateActivity() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id, _ := strconv.Atoi(c.Param("id"))
		activityReq := &dtos.ActivityRequest{}
		if err := c.Bind(activityReq); err != nil {
			return err
		}
		consumerId := headerutil.GetConsumerIDFromHeader(c)
		erract := activityRepo.UpdateActivity(activityReq, id, consumerId)

		if erract != nil {
			return c.JSON(http.StatusInternalServerError, errorutil.NewCustomError(http.StatusInternalServerError,
				errorutil.ErrSomethingWentWrong))
		}

		return c.JSON(http.StatusOK, dtos.SuccessResponse{
			Success: true,
		})
	}
}


