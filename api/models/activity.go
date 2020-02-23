package models
import (
  "github.com/jinzhu/gorm"
)

type Activities struct {
	gorm.Model
	ID            int       `gorm:"primary_key;auto_increment" json:"id"`
	Source        string    `gorm:"size:255;default:null" json:"source"`
	Type          string    `gorm:"size:255;default:null" json:"type"`
	EntityId      string    `gorm:"size:255;default:null" json:"entity_id"`
	EntityType    string    `gorm:"size:255;default:null" json:"entity_type"`
	Field         string    `gorm:"type:text;default:null" json:"field"`
	OldValue      string    `gorm:"size:255;default:null" json:"old_value"`
	NewValue      string    `gorm:"size:255;default:null" json:"new_value"`
	Data          string    `gorm:"size:255;default:null" json:"data"`
	Description   string    `gorm:"type:text;default:null" json:"description"`
	ActionBy      string    `gorm:"size:255;default:null" json:"action_by"`
}
