package table

import (
	"github.com/Sn0wo2/QuickNote/Backend/internal/note"
)

func Init() error {
	return note.InitNoteSchema()
}
