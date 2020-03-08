package dtos

//Requests
type ActivityRequest struct {
	Source        string          `json:"source"      form:"source"`
	Type          string          `json:"type"        form:"type"`
	EntityID      string          `json:"entity_id"   form:"event_id"`
	EntityName    string          `json:"entity_name" form:"event_name"`
	Data          string          `json:"data"        form:"data"`
	Description   string          `json:"description" form:"description"`
	ActionBy      string          `json:"action_by"   form:"action_by"`
	Changes      []ChangeRequest `json:"changes" form:"changes`
}


type ChangeRequest struct {
	Field         *string    `gorm:"type:text;default:null" json:"field, omitempty"`
	NewValue      *string    `gorm:"type:text;default:null" json:"new_value, omitempty"`
	OldValue      *string    `gorm:"type:text;default:null" json:"old_value, omitempty"`
}


type ChangeResponse struct {
	Field         string    `gorm:"type:text;default:null" json:"field"`
	NewValue      string    `gorm:"type:text;default:null" json:"new_value"`
	OldValue      string    `gorm:"type:text;default:null" json:"old_value"`
}
