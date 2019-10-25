package models

import "github.com/jinzhu/gorm"

type VoteRecord struct {
	gorm.Model
	UserId      uint  `json:"user_id" gorm:"type:int(10);not null"`
	CandidateId uint `json:"candidate_id" gorm:"type:int(10);not null"`
	SeasonId    uint `json:"season_id" gorm:"type:int(10);not null"`
}
