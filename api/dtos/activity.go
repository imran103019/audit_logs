package dtos

//Responses
type ActivityResponse struct {
	Id            uint32    `json:"id,omitempty"`
	Source        string    `json:"source"`
	Type          string    `json:"type"`
	EntityId      string    `json:"entity_id"`
	EntityType    string    `json:"entity_type"`
	Field         string    `json:"field"`
	OldValue      string    `json:"old_value"`
	NewValue      string    `json:"new_value"`
	Data          string    `json:"data"`
	Description   string    `json:"description"`
	ActionBy      string    `json:"action_by"`
	CreatedAt     string    `json:"created_at"`

}


//Requests
type ActivityRequest struct {
	Source        string    `json:"source"      form:"source"`
	Type          string    `json:"type"        form:"type"`
	EntityId      string    `json:"entity_id"   form:"event_id"`
	EntityType    string    `json:"entity_type" form:"event_type"`
	Field         string    `json:"field"       form:"field"`
	OldValue      string    `json:"old_value"   form:"old_value"`
	NewValue      string    `json:"new_value"   form:"new_value"`
	Data          string    `json:"data"        form:"data"`
	Description   string    `json:"description" form:"description"`
	ActionBy      string    `json:"action_by"   form:"action_by"`
}