package note

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	NID  string `gorm:"index"`
	Data []byte // Encoded Payload

	Title   []byte `gorm:"-"` // Decoded Title
	Content []byte `gorm:"-"` // Decoded Content
}

type DisplayNote struct {
	NID     string `json:"nid"`
	Lock    bool   `json:"lock"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
