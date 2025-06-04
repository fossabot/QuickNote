package note

import (
	"gorm.io/gorm"
)

type Payload struct {
	Title   []byte `gorm:"-"`
	Content []byte `gorm:"-"`
}

type Note struct {
	gorm.Model
	NID  string `gorm:"uniqueIndex"`
	Lock bool
	Data []byte
	Key  []byte `gorm:"-"`

	Payload
}

type DisplayNote struct {
	NID     string `json:"nid"`
	Lock    bool   `json:"lock"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
