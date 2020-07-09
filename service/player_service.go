package service

import (
	"errors"
	"shadowDemo/model"
	"shadowDemo/model/dao"
	"shadowDemo/model/do"
	"shadowDemo/zframework/datasource"
	modelc "shadowDemo/zframework/model"
	"time"

	shadowsecurity "shadowDemo/zframework/security"
)

type PlayerService struct {
	PwdEncode shadowsecurity.IPasswordEncoder
}

var playerService *PlayerService

func NewPlayerService() *PlayerService {
	if playerService == nil {
		l.Lock()
		if playerService == nil {
			playerService = &PlayerService{
				PwdEncode: shadowsecurity.PasswordEncoderInstance(shadowsecurity.PASSWORD_ENCODER),
			}
		}
		l.Unlock()
	}
	return playerService
}

func PlayerUserDetailService() interface{} {
	return NewPlayerService()
}

func (service *PlayerService) GetPlayerByID(id int64) (do.Player, error) {

	return do.Player{ID: 1}, nil
}
func (service *PlayerService) GetPlayerByLoginName(loginName string) (player do.Player, err error) {
	if loginName == "" {
		return do.Player{}, errors.New("no login name")
	}
	//db := datasource.DatasourceManagerInstance(datasource.DATASOURCE_MANAGER).Datasource()
	//playerDao := dao.NewPlayerDao(db)
	mPlayer := do.Player{
		Account: loginName,
	}
	// err = playerDao.FindPlayerByLoginName(&mPlayer)
	// if err != nil {
	// 	Log.Error(err)
	// 	err = AccountNotExistError{
	// 		error: errors.New("account not exist"),
	// 	}
	// 	return do.Player{}, err
	// }
	return mPlayer, nil
}

func (service *PlayerService) LoadUserByUsername(userName string) shadowsecurity.IUserDetails {
	mPlayer := do.Player{
		Account: userName,
	}
	if mPlayer.ID == 0 {
		return nil
	}

	return &mPlayer
}

// // CreatePlayer 创建玩家
func (service *PlayerService) CreatePlayer(agentID int64, player *do.Player) (err error) {
	db := datasource.DatasourceManagerInstance(datasource.DATASOURCE_MANAGER).Datasource()
	tx := db.Begin()
	defer closeTx(tx, &err)
	return nil
}

func (service *PlayerService) UpdateLastLoginTime(id int64, ip string, ipAddr string, lastLogUrl string, device modelc.Device) (err error) {
	m := model.GetModel()
	now := time.Now()
	attrs := map[string]interface{}{}
	attrs["last_login_at"] = &now
	attrs["last_login_ip"] = ip
	attrs["last_login_ip_addr"] = ipAddr
	attrs["last_log_url"] = lastLogUrl
	attrs["last_dev_info"] = device
	attrs["online_flag"] = true
	err = m.PlayerDao.Updates(id, attrs)
	if err != nil {
		Log.Error(err)
		return
	}
	return err
}

func (service *PlayerService) SearchPlayerPaging(condition *dao.PlayerSearchCondition, pageNum int, pageSize int) (request []do.Player, count int, err error) {
	rowbound := modelc.NewRowBound(pageNum, pageSize)
	return service.searchPlayer(condition, &rowbound)
}

func (service *PlayerService) SearchPlayerWithOutPaging(condition *dao.PlayerSearchCondition) (request []do.Player, count int, err error) {
	return service.searchPlayer(condition, nil)
}

func (service *PlayerService) searchPlayer(condition *dao.PlayerSearchCondition, rowbound *modelc.RowBound) (request []do.Player, count int, err error) {
	result, count, err := model.GetModel().PlayerDao.SearchPlayers(condition, rowbound)
	if err != nil {
		Log.Error(err)
		return nil, 0, err
	}
	return result, count, err
}
