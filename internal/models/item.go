package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ProjectID   uint   `json:"project_id"`
	ProjectName string `json:"project_name"`
	UserID      uint   `json:"user_id"`
	Name        string `json:"name"`
	// ParentID    *uint   `json:"parent_id"`
	ParentTime float64 `json:"parent_time"`
	Status     string  `json:"status"`
	// SubItems   []Item  `json:"sub_items" gorm:"foreignkey:ParentID"`
}
