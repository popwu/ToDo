package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name   string `json:"name"`
	UserID uint   `json:"user_id"`
	Items  []Item `json:"items"`
}
