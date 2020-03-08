package models
import (
  "github.com/jinzhu/gorm"
)

type Consumers struct {
	gorm.Model `json:"-"`
	AppName        string    `gorm:"size:255;default:null" json:"app_name"`
	Token          string    `gorm:"size:255;default:null" json:"type"`
}
