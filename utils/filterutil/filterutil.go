package filterutil
import (
	"github.com/jinzhu/gorm"
)


type ActivityFilterElements struct {
	Source        string    `json:"source"      query:"source"`
	Type          string    `json:"type"        query:"type"`
	EntityID      string    `json:"entity_id"   query:"event_id"`
	EntityName    string    `json:"event_name"  query:"event_name"`
	ActionBy      string    `json:"action_by"   query:"action_by"`
	ToDate        string    `json:"to_date"     query:"to_date"`
	FromDate      string    `json:"from_date"   query:"from_date"`
	ConsumerID    uint64    `json:"consumer_id" query:"consumer_id"`
}

type ConsumerFilterElements struct {
	AppName       string    `json:"app_name"    query:"app_name"`
	ToDate        string    `json:"to_date"     query:"to_date"`
	FromDate      string    `json:"from_date"   query:"from_date"`
}



func (f *ActivityFilterElements) GetFilteredActivityQuery(query *gorm.DB) *gorm.DB {
	query = query.Where("activities.deleted_at is null")

	if(f.ConsumerID > 0) {
	    query = query.Where("activities.consumer_id=?", f.ConsumerID)
	}

	if(f.Source != "") {
	    query = query.Where("activities.source=?", f.Source)
	}
	if(f.FromDate != "") {
		query = query.Where("activities.created_at>=?", f.FromDate)
	}
	if(f.ToDate != "") {
		query = query.Where("activities.created_at<=?", f.ToDate)
	}
	if(f.ActionBy != "") {
	    query = query.Where("activities.action_by=?", f.ActionBy)
	}
	if(f.EntityName != "") {
	    query = query.Where("activities.entity_name=?", f.EntityName)
	}
	if(f.EntityID != "") {
	    query = query.Where("activities.entity_id=?", f.EntityID)
	}
	if(f.Type != "") {
	    query = query.Where("activities.type=?", f.Type)
	}
	return query
}


func (f *ConsumerFilterElements) GetFilteredConsumerQuery(query *gorm.DB) *gorm.DB {
	if(f.AppName != "") {
	    query = query.Where("c.app_name=?", f.AppName)
	}
	if(f.FromDate != "") {
		query = query.Where("c.created_at>=?", f.FromDate)
	}
	if(f.ToDate != "") {
		query = query.Where("c.created_at<=?", f.ToDate)
	}
	return query
}
