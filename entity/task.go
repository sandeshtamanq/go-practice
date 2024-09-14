package entity

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"type:varchar(100);not null;unique_index" json:"title"`
	Description string         `gorom:"type:varchar(250); not null" json:"description"`
	Done        bool           `gorm:"default:false" json:"done"`
	UserID      uint           `json:"user_id"`
	User        User           `json:"user"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`
}
