package repos

import (
	"github.com/imran103019/audit_logs/conn"
	dtos "github.com/imran103019/audit_logs/api/dtos"
	"github.com/imran103019/audit_logs/utils/paginationutil"
	"github.com/imran103019/audit_logs/utils/filterutil"
	"github.com/imran103019/audit_logs/api/models"
	"time"
)

type ActivityRepo struct {
	db *conn.DB
}

func NewActivityRepo(db *conn.DB) Activity {
	return &ActivityRepo{db}
}

func (a *ActivityRepo) GetActivities(f filterutil.ActivityFilterElements, pgntr paginationutil.Paginator ) ([]dtos.ActivityResponse, error) {
	data := []dtos.ActivityResponse{}
	query := a.db.Table("activities AS a")
	query = f.GetFilteredActivityQuery(query).Where("deleted_at is null").Order("a.id DESC")

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

	return data, nil
}


func (a *ActivityRepo) SetActivity(activity *dtos.ActivityRequest) error {

	activitymodel :=  models.Activities{
	    Source:       activity.Source, 
	    Type:         activity.Type, 
	    EntityType:   activity.EntityType,
	    EntityId:     activity.EntityId,
	    Field:        activity.Field,
	    OldValue:     activity.OldValue,
	    NewValue:     activity.NewValue,
	    Data:         activity.Data, 
	    Description:  activity.Description, 
	    ActionBy:     activity.ActionBy,
    }

	if err := a.db.Create(&activitymodel).Error; err != nil {
		return err
	}
	return nil
}


func (a *ActivityRepo) UpdateActivity(activity *dtos.ActivityRequest, id int) error {
	err := a.db.
		Table("activities AS a").
		Where("a.id = ?", id).
		Updates(map[string]interface{}{
			"description": activity.Description,
			"data":        activity.Data,
			"old_value":   activity.OldValue,
			"new_value":   activity.NewValue,
			"type":        activity.Type,
			"entity_type": activity.EntityType,
			"entity_id":   activity.EntityId,
			"field":       activity.Field,
		}).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *ActivityRepo) DeleteActivity(id int) error {
	err := a.db.
		Table("activities AS a").
		Where("a.id = ?", id).
		Updates(map[string]interface{}{
			"deleted_at": time.Now().UTC(),
		}).Error
	if err != nil {
		return err
	}
	return nil
}

