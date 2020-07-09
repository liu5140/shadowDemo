package model

import (
	"shadowDemo/model/dao"
	"shadowDemo/zframework/datasource"
)

type Model struct {
	APIAccessReqLogDao *dao.APIAccessReqLogDao
	APIAccessResLogDao *dao.APIAccessResLogDao
	IPWhiteListDao     *dao.IPWhiteListDao
	MerchantDao        *dao.MerchantDao
	OrderDao           *dao.OrderDao
	PlayerDao          *dao.PlayerDao
	ProgConfigDao      *dao.ProgConfigDao
	UserDao            *dao.UserDao
}

var model *Model = nil

func ModelInit() {
	model = &Model{}

	db := datasource.DataSourceInstance().Master()

	model.APIAccessReqLogDao = dao.NewAPIAccessReqLogDao(db)
	model.APIAccessResLogDao = dao.NewAPIAccessResLogDao(db)
	model.IPWhiteListDao = dao.NewIPWhiteListDao(db)
	model.MerchantDao = dao.NewMerchantDao(db)
	model.OrderDao = dao.NewOrderDao(db)
	model.PlayerDao = dao.NewPlayerDao(db)
	model.ProgConfigDao = dao.NewProgConfigDao(db)
	model.UserDao = dao.NewUserDao(db)

}
