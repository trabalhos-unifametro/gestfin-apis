package models

import (
	"gorm.io/gorm"
	"time"
)

type Expenses struct {
	gorm.Model
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `gorm:"type:varchar(200)" rql:"filter" json:"name"`
	Value     float64        `json:"value"`
	MonthYear string         `gorm:"type:varchar(5)" json:"month_year"`
	UserID    int            `rql:"filter" json:"user_id"`
	User      User           `gorm:"foreignKey:UserID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (Expenses) TableName() string {
	return "gestfin.expenses"
}
