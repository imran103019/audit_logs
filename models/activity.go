package models

import (
  "github.com/jinzhu/gorm"
  //"errors"
  //log "github.com/sirupsen/logrus"
)


type Activities struct {
	gorm.Model `json:"-"`
	ConsumerID    uint64    `gorm:"size:11;default:null"  json:"-"`
	Source        string    `gorm:"size:255;default:null" json:"source"`
	Type          string    `gorm:"size:255;default:null" json:"type"`
	EntityID      string    `gorm:"size:255;default:null" json:"entity_id"`
	EntityName    string    `gorm:"size:255;default:null" json:"entity_name"`
	Data          string    `gorm:"size:255;default:null" json:"data"`
	Description   string    `gorm:"type:text;default:null" json:"description"`
	ActionBy      string    `gorm:"size:255;default:null" json:"action_by"`
	Changes      []Changes  `gorm:"foreignkey:ActivityID" json:"changes"`
}
