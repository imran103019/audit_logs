package filterutil
import (
	"github.com/jinzhu/gorm"
)


type ActivityFilterElements struct {
	Source        string    `json:"source"      query:"source"`
	Type          string    `json:"type"        query:"type"`
	EntityId      string    `json:"entity_id"   query:"event_id"`
	EntityType    string    `json:"entity_type" query:"event_type"`
	ActionBy      string    `json:"action_by"   query:"action_by"`
	ToDate        string    `json:"to_date"     query:"to_date"`
	FromDate      string    `json:"from_date"   query:"from_date"`
}


// GetFilteredCategoryQuery adds filter to query and return it
func (f *ActivityFilterElements) GetFilteredActivityQuery(query *gorm.DB) *gorm.DB {
	if(f.Source != "") {
	    query = query.Where("a.source=?", f.Source)
	}
	if(f.FromDate != "") {
		query = query.Where("a.created_at>=?", f.FromDate)
	}
	if(f.ToDate != "") {
		query = query.Where("a.created_at<=?", f.ToDate)
	}
	if(f.ActionBy != "") {
	    query = query.Where("a.action_by=?", f.ActionBy)
	}
	if(f.EntityType != "") {
	    query = query.Where("a.entity_type=?", f.EntityType)
	}
	if(f.EntityId != "") {
	    query = query.Where("a.entity_id=?", f.EntityId)
	}
	if(f.Type != "") {
	    query = query.Where("a.type=?", f.Type)
	}
	return query
}
