package repos

import (
	"github.com/imran103019/audit_logs/conn"
	dtos "github.com/imran103019/audit_logs/dtos"
	"github.com/imran103019/audit_logs/utils/paginationutil"
	"github.com/imran103019/audit_logs/utils/filterutil"
	"github.com/imran103019/audit_logs/models"
	log "github.com/sirupsen/logrus"
)

type ActivityRepo struct {
	db *conn.DB
}

func NewActivityRepo(db *conn.DB) Activity {
	return &ActivityRepo{db}
}

func (a *ActivityRepo) GetActivities(f filterutil.ActivityFilterElements, pgntr paginationutil.Paginator ) ([]models.Activities, error) {
	data := []models.Activities{}
	query := a.db.Model(&models.Activities{}).Preload("Changes")
	query = f.GetFilteredActivityQuery(query).Order("activities.id DESC")
	
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
		log.Error("Failed to get activities: ",query.Error.Error())
		return nil, query.Error
	}
	return data, nil
}


func (a *ActivityRepo) SetActivity(activity *dtos.ActivityRequest, consumerId uint64) error {

	changeLogs := []models.Changes{}

	for _, change := range activity.Changes {
		changeLogs = append(changeLogs, models.Changes{Field: *change.Field, NewValue: *change.NewValue, OldValue: *change.OldValue})
	}
	activitymodel :=  models.Activities{
		ConsumerID:   consumerId,
	    Source:       activity.Source, 
	    Type:         activity.Type, 
	    EntityName:   activity.EntityName,
	    EntityID:     activity.EntityID,
	    Data:         activity.Data, 
	    Description:  activity.Description, 
	    ActionBy:     activity.ActionBy,
	    Changes:      changeLogs,
    }

	if err := a.db.Create(&activitymodel).Error; err != nil {
		log.Error("Failed to save activity: ",err.Error())
		return err
	}
	return nil
}


func (a *ActivityRepo) UpdateActivity(activity *dtos.ActivityRequest, id int, consumerId uint64) error {
	err := a.db.
		Table("activities AS a").
		Where("a.id = ?", id).
		Where("a.consumer_id = ?", consumerId).
		Updates(map[string]interface{}{
			"description": activity.Description,
			"data":        activity.Data,
			"type":        activity.Type,
			"entity_name": activity.EntityName,
			"entity_id":   activity.EntityID,
		}).Error
	if err != nil {
		log.Error("Failed to update activity: ",err.Error())
		return err
	}
	return nil
}


