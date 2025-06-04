package table

import (
	"github.com/Sn0wo2/QuickNote/note"
)

func Init() error {
	return note.InitNoteSchema()
}
