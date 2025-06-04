package note

import (
	"errors"

	"github.com/Sn0wo2/QuickNote/database/orm"
	"gorm.io/gorm"
)

func InitNoteSchema() error {
	return orm.Instance.Get().AutoMigrate(&Note{})
}

func GetNote(nid string) (Note, error) {
	var note Note
	err := orm.Instance.Get().First(&note, "n_id = ?", nid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Note{}, nil
	}
	return note, err
}

func SetNote(note Note) error {
	db := orm.Instance.Get()

	return db.Transaction(func(tx *gorm.DB) error {
		var existing Note
		err := tx.First(&existing, "n_id = ?", note.NID).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return tx.Create(&note).Error
		} else if err != nil {
			return err
		}
		return tx.Model(&existing).Updates(note).Error
	})
}

func DeleteNote(nid string) error {
	db := orm.Instance.Get()

	return db.Transaction(func(tx *gorm.DB) error {
		var note Note
		err := tx.First(&note, "n_id = ?", nid).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		} else if err != nil {
			return err
		}
		return tx.Unscoped().Delete(&note).Error
	})
}
