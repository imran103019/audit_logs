package models

import (
  "github.com/jinzhu/gorm"
)


type Changes struct {
	gorm.Model `json:"-"`
	ActivityID    *uint64   `gorm:"size:11;default:null"  json:"-"`
	Field         string    `gorm:"type:text;default:null" json:"field,omitempty"`
	NewValue      string    `gorm:"type:text;default:null" json:"new_value,omitempty"`
	OldValue      string    `gorm:"type:text;default:null" json:"old_value,omitempty"`
}
