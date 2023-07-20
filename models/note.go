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

func NoteCreate(name string, content string) *Note {
	entry := Note{
		Name:    name,
		Content: content,
	}
	DB.Create(&entry)
	return &entry
}

func NotesFind(id uint64) *Note {
	var note Note
	// DB.First(&note, id)
	DB.Where("id = ?", id).First(&note)
	return &note
}
