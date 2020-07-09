package service

import (
	"errors"
	"fmt"
	"shadowDemo/model"
	"shadowDemo/model/do"
	"shadowDemo/service/dto"
	"shadowDemo/zframework/bizerr"
	modelc "shadowDemo/zframework/model"
	"sync"

	"github.com/sirupsen/logrus"

	"shadowDemo/zframework/security"
)

type UpmsAdminService struct {
	PwdEncode security.IPasswordEncoder
	MenuList  []*do.UpmsMenu
}

var adminService *UpmsAdminService

// 安全密码错误次数记录
var wrongSecurePwdCountHolder sync.Map

func NewUpmsAdminService() *UpmsAdminService {
	if adminService == nil {
		l.Lock()
		if adminService == nil {
			menu, err := model.GetModel().UpmsMenuDao.Find(&do.UpmsMenu{})
			if err != nil {
				Log.Error(err)
			}
			adminService = &UpmsAdminService{
				PwdEncode: security.PasswordEncoderInstance(security.PASSWORD_ENCODER),
				MenuList:  menu,
			}
		}
		l.Unlock()
	}
	return adminService
}

func UpmsAdminUserDetailService() interface{} {
	return NewUpmsAdminService()
}

func (adminService *UpmsAdminService) LoadUserByUsername(username string) security.IUserDetails {
	admin := &do.UpmsAdmin{
		Account: username,
	}
	if err := model.GetModel().UpmsAdminDao.FindOne(admin); err != nil {
		Log.Errorln(err)
	}
	return admin
}

func (adminService *UpmsAdminService) LoadBaccUserByUsername(site string, username string) (admin do.UpmsAdmin, err error) {
	if username == "" {
		return do.UpmsAdmin{}, errors.New("no login name")
	}
	adminDao := model.GetModel().UpmsAdminDao
	result := do.UpmsAdmin{
		Account: username,
	}
	Log.Infoln("GetUpmsAdminByLoginName", username)
	err = adminDao.FindOne(&result)
	if err != nil {
		Log.Error(err)
		return do.UpmsAdmin{}, err
	}
	return result, nil
}

func (adminManager *UpmsAdminService) GetUpmsAdminByLoginName(login string) (admin do.UpmsAdmin, err error) {
	if login == "" {
		return do.UpmsAdmin{}, errors.New("no login name")
	}
	adminDao := model.GetModel().UpmsAdminDao
	result := do.UpmsAdmin{
		Account: login,
	}
	Log.Infoln("GetUpmsAdminByLoginName", login)
	err = adminDao.FindOne(&result)
	if err != nil {
		Log.Error(err)
		return do.UpmsAdmin{}, err
	}
	return result, nil
}

func (adminManager *UpmsAdminService) GetAuthority(site string, id int64) (menu []do.UpmsMenu, err error) {
	var menulist []do.UpmsMenu
	for _, menu := range adminManager.MenuList {
		ok := security.GetCasbinEnforcer().Enforce(fmt.Sprintf("%d", id), site, menu.URL, menu.Method)
		if ok {
			menu.Selected = true
			menulist = append(menulist, *menu)
		}
	}
	return menulist, nil
}

func (adminService UpmsAdminService) GetUpmsAdminByID(adminID int64) (admin do.UpmsAdmin, err error) {
	admin = do.UpmsAdmin{
		ID: adminID,
	}
	err = model.GetModel().UpmsAdminDao.Get(&admin)
	return admin, err
}

func (adminService UpmsAdminService) VerifySecurePwd(adminID int64, securePwd string) (err error) {
	adminDao := model.GetModel().UpmsAdminDao
	admin := do.UpmsAdmin{
		ID: adminID,
	}
	err = adminDao.Get(&admin)
	if err != nil {
		return err
	}
	passwordEncoder := adminService.PwdEncode
	//默认最多可重试6次
	defaultTryTimes := 6
	wrongTimes := 0
	if !passwordEncoder.Matches(securePwd, admin.SecurePassword) {
		//记录输入错误次数
		val, ok := wrongSecurePwdCountHolder.Load(adminID)
		if ok {
			wrongTimes = val.(int) + 1
		} else {
			wrongTimes = 1
		}
		//如果错误次数达到最大，则冻结帐号
		if wrongTimes == defaultTryTimes {
			model.GetModel().UpmsAdminDao.Update(adminID, "state", modelc.Frozen)
			wrongSecurePwdCountHolder.Delete(adminID)
			return security.AccountLockedError{}
		}
		wrongSecurePwdCountHolder.Store(adminID, wrongTimes)
		return bizerr.GenErr("key_invalid_seruce_pwd_error", defaultTryTimes-wrongTimes)
	}
	wrongSecurePwdCountHolder.Delete(adminID)
	return nil
}

func (adminService UpmsAdminService) UpdateUpmsAdminPwd(userid int64, usertype do.UserType, password dto.UpdatePwd) (err error) {
	Log.WithFields(logrus.Fields{
		"userid":   userid,
		"usertype": usertype,
		"password": password,
	}).Debug("UpdateUpmsAdminPwd")
	if password.NewPwd == password.OldPwd {
		return bizerr.GenErr("key_alert_new_pwd_same_as_old_error")
	}

	if password.NewPwd != password.EnsurePwd {
		return bizerr.GenErr("key_alert_fielderror_field_pwdmatch")
	}
	adminDao := model.GetModel().UpmsAdminDao
	admin := do.UpmsAdmin{
		ID:       userid,
		UserType: usertype,
	}
	err = adminDao.Get(&admin)
	if err != nil {
		return
	}
	passwordEncoder := adminService.PwdEncode
	var oldPwd string
	if password.UpdatePwdType == do.UpdatePwdTypeLogin {
		oldPwd = admin.LoginPassword
	} else if password.UpdatePwdType == do.UpdatePwdTypeSecure {
		oldPwd = admin.SecurePassword
	}
	if password.UpdatePwdType == do.UpdatePwdTypeLogin && oldPwd == "" {
		return bizerr.GenErr("key_alert_nologin_password_error")
	} else if password.UpdatePwdType == do.UpdatePwdTypeSecure && oldPwd == "" {
		Log.Error(err)
		return bizerr.GenErr("key_alert_nosecure_password_error")
	}
	if !passwordEncoder.Matches(password.OldPwd, oldPwd) {
		Log.Error(err)
		return bizerr.GenErr("key_current_password_verify_error")
	}
	var param string
	if password.UpdatePwdType == do.UpdatePwdTypeLogin {
		param = "login_password"
	} else if password.UpdatePwdType == do.UpdatePwdTypeSecure {
		param = "secure_password"
	}
	err = adminDao.Update(userid, param, passwordEncoder.Encode(password.NewPwd))
	if err != nil {
		return
	}
	return
}
