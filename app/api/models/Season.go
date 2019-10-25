package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Season struct {
	gorm.Model
	Month uint `json:"month" gorm:"type:int(10);unique_index;not null"`
	StartTime  time.Time `json:"start_time" gorm:"not null"`
	EndTime   time.Time `json:"end_time" gorm:"not null"`
}
