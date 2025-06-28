package note

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	// gorm.Model
	ID        uint           `json:"-" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	NID  string `json:"nid" gorm:"index"`
	Data []byte `json:"-"` // Encoded Payload

	Title   []byte `json:"-" gorm:"-"` // Decoded Title
	Content []byte `json:"-" gorm:"-"` // Decoded Content

	// TODO: REFACTOR THIS
	DisplayTitle   string `json:"title,omitempty"`
	DisplayContent string `json:"content,omitempty"`
}
