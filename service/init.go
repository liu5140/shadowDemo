package service

import (
	"sync"

	"shadowDemo/shadow-framework/logger"
	"github.com/jinzhu/gorm"

	"github.com/nicksnyder/go-i18n/i18n"
)

var (
	Log *logger.Logger
	l   sync.Mutex
)

func init() {
	Log = logger.InitLog()
}

func closeTx(tx *gorm.DB, err *error) {
	r := recover()
	if r != nil {
		tx.Rollback()
		Log.Error(r)
		return
	}

	if *err != nil {
		tx.Rollback()
		Log.Errorf("%+v", *err)
		return
	}
	tx.Commit()
}

func CloseTx(tx *gorm.DB, err *error) {
	closeTx(tx, err)
}

type TanslateModel struct {
	Number int
	Value  string
}

func TranslateTypeToName(T i18n.TranslateFunc) (result map[string]string) {
	tanslateRecordTypeMap := make(map[string]string)
	tanslateRecordTypeMap["fundSourceType"] = T("key_t_fundSourceType")
	tanslateRecordTypeMap["bannerType"] = T("key_t_bannerType")
	return tanslateRecordTypeMap
}

func TranslateModelToName(T i18n.TranslateFunc) (result map[string]map[string]string) {
	modelMap := make(map[string]map[string]string)
	modelMap["bannerType"] = map[string]string{
		// fmt.Sprint(model.Sy): T("key_t_sy"),
		// fmt.Sprint(model.Sc): T("key_t_sc"),
		// fmt.Sprint(model.Zs): T("key_t_zs"),
	}
	return modelMap
}
