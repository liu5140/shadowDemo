package dao

import "github.com/jinzhu/gorm"

// model recover tx
func RecoverTransaction(tx *gorm.DB) {
	err := recover()
	if err != nil {
		Log.Errorf("panic captured, will rollback tx, err = %v", err)
		tx.Rollback()
	}
}
