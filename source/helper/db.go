package helper

import (
	"gorm.io/gorm"
)

func StartTransaction(inputtx *gorm.DB) (tx *gorm.DB) {
	tx = inputtx.Begin()
	if tx.Error != nil {
		PanicIfError(tx.Error)
	}
	return tx
}
func Commit(tx *gorm.DB) {
	errorCommit := tx.Commit().Error
	PanicIfError(errorCommit)
}
func Rollback(tx *gorm.DB) {
	errorCommit := tx.Rollback().Error
	PanicIfError(errorCommit)
}
func RollbackIfError(tx *gorm.DB) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback().Error
		if errorRollback != nil {
			PanicIfError(errorRollback)
		} else {
			PanicIfError(err.(error))
		}
	}
}
