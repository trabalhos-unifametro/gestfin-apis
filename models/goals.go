package models

import (
	"gorm.io/gorm"
	"time"
)

type Goals struct {
	gorm.Model
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Value     float64        `json:"value"`
	UserID    int            `rql:"filter" json:"user_id"`
	User      User           `gorm:"foreignKey:UserID"`
	Deadline  int            `json:"deadline"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (Goals) TableName() string {
	return "gestfin.goals"
}
