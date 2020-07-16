package service

import (
	"errors"
	"fmt"
	"shadowDemo/model"
	"shadowDemo/model/do"
	"shadowDemo/service/dto"
	"shadowDemo/zframework/bizerr"
	modelc "shadowDemo/zframework/model"
	shadowsecurity "shadowDemo/zframework/security"
	"strings"
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
		return bizerr.GenErr("key_alert_nosecure_password_error")
	}

	if !passwordEncoder.Matches(password.OldPwd, oldPwd) {
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

func (service UpmsAdminService) CreateSubAccount(subAccount *do.UpmsAdmin, roleids []int64, appid string) (err error) {
	adminDao := model.GetModel().UpmsAdminDao
	rdao := model.GetModel().UpmsRoleDao
	subAccount.Account = strings.TrimSpace(subAccount.Account)
	if ok := adminDao.Found(&do.UpmsAdmin{
		Account: subAccount.Account,
	}); ok {
		Log.Error("account is exist")
		return bizerr.GenErr("key_account_exist")
	}

	passwordEncoder := shadowsecurity.PasswordEncoderInstance(shadowsecurity.PASSWORD_ENCODER)
	loginPassword := passwordEncoder.Encode(strings.TrimSpace(subAccount.LoginPassword))
	securePassword := passwordEncoder.Encode(strings.TrimSpace(subAccount.SecurePassword))
	subAccount.LoginPassword = loginPassword
	subAccount.SecurePassword = securePassword
	subAccount.ID = GenAdminID()
	subAccount.State = modelc.Normal
	subAccount.UserType = do.UserTypeAdmin

	roles, err := rdao.BatchGet(roleids)
	if err != nil || len(roles) == 0 {
		return bizerr.GenErr("key_roles_not_exist")
	}
	var roleid, rolecode, roleName string

	//为该子账号指定角色
	for _, role := range roles {
		roleid = roleid + fmt.Sprint(role.ID) + ";"
		rolecode = rolecode + role.Code + ";"
		roleName = roleName + role.Name + ";"
		if ok := shadowsecurity.GetCasbinEnforcer().AddGroupingPolicy(fmt.Sprintf("%d", subAccount.ID), role.Code, appid); !ok {
			err = errors.New("failed to assign role to admin in db")
			Log.WithField("Account", subAccount.Account).Error(err)
			return err
		}
	}
	subAccount.RoleCode = rolecode
	subAccount.RoleID = roleid
	subAccount.RoleName = roleName
	//创建admin
	err = adminDao.Create(subAccount)
	if err != nil {
		Log.Error(err)
		return err
	}

	return nil
}

func (service UpmsAdminService) UpdateRole(id int64, roleids []int64, appid string) (err error) {
	sdao := model.GetModel().UpmsAdminDao
	rdao := model.GetModel().UpmsRoleDao
	old := &do.UpmsAdmin{ID: id}
	err = sdao.FindOne(old)
	if err != nil {
		Log.Error(err)
		return err
	}
	//删除修改前的角色
	oldRoles := strings.Split(old.RoleCode, ";")
	for _, code := range oldRoles {
		//删除已有角色
		if code != "" {
			if ok := shadowsecurity.GetCasbinEnforcer().RemoveGroupingPolicy(fmt.Sprintf("%d", id), code, appid); !ok {
				err = errors.New("failed to delete role to admin in db")
				Log.WithField("Account", old.Account).Error(err)
				return err
			}
		}
	}

	roles, err := rdao.BatchGet(roleids)
	if err != nil || len(roles) == 0 {
		return bizerr.GenErr("key_roles_not_exist")
	}
	var roleid, rolecode, roleName string

	//为该子账号指定角色
	for _, role := range roles {
		roleid = roleid + fmt.Sprint(role.ID) + ";"
		rolecode = rolecode + role.Code + ";"
		roleName = roleName + role.Name + ";"
		if ok := shadowsecurity.GetCasbinEnforcer().AddGroupingPolicy(fmt.Sprintf("%d", id), role.Code, appid); !ok {
			err = errors.New("failed to assign role to admin in db")
			Log.WithField("Account", old.Account).Error(err)
			return err
		}
	}

	err = sdao.Updates(id, map[string]interface{}{"role_id": roleid, "role_code": rolecode, "role_name": roleName})
	if err != nil {
		Log.Error(err)
		return err
	}
	return nil
}
