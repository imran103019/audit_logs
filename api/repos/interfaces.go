package repos

import (
	"github.com/imran103019/audit_logs/api/dtos"
	"github.com/imran103019/audit_logs/utils/filterutil"
	"github.com/imran103019/audit_logs/utils/paginationutil"
)


type Activity interface {
	GetActivities(filterutil.ActivityFilterElements, paginationutil.Paginator) ([]dtos.ActivityResponse, error)
	SetActivity(* dtos.ActivityRequest) error
	UpdateActivity(* dtos.ActivityRequest, int) error
	DeleteActivity(int) error
}
