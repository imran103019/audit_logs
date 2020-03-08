package repos

import (
	"github.com/imran103019/audit_logs/conn"
	dtos "github.com/imran103019/audit_logs/dtos"
	"github.com/imran103019/audit_logs/models"
	"github.com/imran103019/audit_logs/utils/paginationutil"
	"github.com/imran103019/audit_logs/utils/filterutil"
	"time"
	"fmt"
	"github.com/imran103019/audit_logs/utils/cacheutil"
	"github.com/spf13/viper"
	log "github.com/sirupsen/logrus"
)

type ConsumerRepo struct {
	db *conn.DB
}

func NewConsumerRepo(db *conn.DB) Consumer {
	return &ConsumerRepo{db}
}


func (a *ConsumerRepo) GetConsumers(f filterutil.ConsumerFilterElements, pgntr paginationutil.Paginator ) ([]dtos.ConsumerResponse, error) {
	data := []dtos.ConsumerResponse{}
	query := a.db.Table("consumers AS c")
	query = f.GetFilteredConsumerQuery(query).Where("deleted_at is null").Order("c.id DESC")

	if err := query.Limit(pgntr.PageLimit()).Offset(pgntr.PageOffset()).Scan(&data).Error; err != nil {
	    return nil, err
	}

	var count uint
	query.Count(&count)

	if pgntr != nil {
		pgntr.SetTotalPage(count)
		query.
			Offset(pgntr.PageOffset()).
			Limit(pgntr.PageLimit()).
			Find(&data)
	} else {
		query.Find(&data)
	}

	if query.Error != nil {
		return nil, query.Error
	}

	for _, consumer := range data {
		key := fmt.Sprintf("consumer_%s",consumer.Token)
		if err := cacheutil.SetCache(key, string(consumer.Id),
			viper.GetDuration("GLOBAL_CACHE_DURATION")*time.Second); err != nil {
			log.Error("Failed to set cache for consumers:", err.Error())
		}
	}

	return data, nil
}



func (a *ConsumerRepo) SetConsumer(consumer *dtos.ConsumerRequest) error {

	consumerModel :=  models.Consumers{
	    AppName:      consumer.AppName,
    }

	if err := a.db.Create(&consumerModel).Error; err != nil {
		log.Error("Failed to set consumer: ", err.Error())
		return err
	}
	return nil
}

func (a *ConsumerRepo) SetAllConsumerAppTokenCache() error {

	consumers := []dtos.ConsumerResponse{}
	query := a.db.Table("consumers AS c").Where("c.deleted_at is null")
	if err := query.Scan(&consumers).Error; err != nil {
	    return err
	}

	for _, consumer := range consumers {
		key := fmt.Sprintf("consumer_%s",consumer.Token)
		if newerr := cacheutil.SetCache(key, string(consumer.Id),
			viper.GetDuration("GLOBAL_CACHE_DURATION")*time.Second); newerr != nil {
			log.Error("Failed to set cache for consumers: ", newerr.Error())
		}
	}

	return nil
}



func (a *ConsumerRepo) UpdateConsumer(consumer *dtos.ConsumerRequest, token string) error {
	err := a.db.
		Table("consumers AS a").
		Where("a.token = ?", token).
		Updates(map[string]interface{}{
			"app_name":    consumer.AppName,
		}).Error
	if err != nil {
		log.Error("Failed to update consumer: ", err.Error())
		return err
	}
	return nil
}

func (a *ConsumerRepo) DeleteConsumer(token string) error {
	err := a.db.
		Table("consumers AS a").
		Where("a.token = ?", token).
		Updates(map[string]interface{}{
			"deleted_at": time.Now().UTC(),
		}).Error
	if err != nil {
		log.Error("Failed to delete consumer: ", err.Error())
		return err
	}
	return nil
}

func (a *ConsumerRepo) GetConsumerByToken(token string) (dtos.ConsumerResponse,error) {
	data := dtos.ConsumerResponse{}
	query := a.db.Table("consumers AS c").Where("token = ?", token).Where("c.deleted_at is null")
	if err := query.Scan(&data).Error; err != nil {
		log.Error("Failed to get consumer by token: ", err.Error())
	    return dtos.ConsumerResponse{},err
	}

	return data, nil
}



