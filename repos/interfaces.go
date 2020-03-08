package repos

import (
	"github.com/imran103019/audit_logs/dtos"
	"github.com/imran103019/audit_logs/utils/filterutil"
	"github.com/imran103019/audit_logs/utils/paginationutil"
	"github.com/imran103019/audit_logs/models"
)


type Activity interface {
	GetActivities(filterutil.ActivityFilterElements, paginationutil.Paginator) ([]models.Activities, error)
	SetActivity(* dtos.ActivityRequest, uint64) error
	UpdateActivity(* dtos.ActivityRequest, int, uint64) error
}


type Consumer interface {
	GetConsumers(filterutil.ConsumerFilterElements, paginationutil.Paginator) ([]dtos.ConsumerResponse, error)
	SetConsumer(* dtos.ConsumerRequest) error
	UpdateConsumer(* dtos.ConsumerRequest, string) error
	DeleteConsumer(string) error
	SetAllConsumerAppTokenCache() error
	GetConsumerByToken(string) (dtos.ConsumerResponse,error)
}
