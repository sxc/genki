package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"size:255"`
	Content   string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func NotesAll() *[]Note {
	var notes []Note
	DB.Where("deleted_at IS NULL").Order("created_at DESC").Find(&notes)
	return &notes
}
