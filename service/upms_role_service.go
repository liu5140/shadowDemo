package service

import (
	"errors"
	"shadowDemo/model"
	"shadowDemo/model/dao"
	"shadowDemo/model/do"
	"shadowDemo/zframework/bizerr"
	shadowsecurity "shadowDemo/zframework/security"
	"strings"

	"github.com/sirupsen/logrus"
)

type UpmsRoleService struct {
}

var upmsRoleService *UpmsRoleService

func NewUpmsRoleService() *UpmsRoleService {
	if upmsRoleService == nil {
		l.Lock()
		if upmsRoleService == nil {
			upmsRoleService = &UpmsRoleService{}
		}
		l.Unlock()
	}
	return upmsRoleService
}

func (service *UpmsRoleService) GetRoleByRoleCodes(site string, roleCode string) ([]do.UpmsRole, error) {
	return model.GetModel().UpmsRoleDao.GetRolesByCodes(strings.Split(roleCode, ","))
}

func (service *UpmsRoleService) GetRoleByID(site string, id int64) (do.UpmsRole, error) {
	role := do.UpmsRole{
		ID: id,
	}
	err := model.GetModel().UpmsRoleDao.Get(&role)
	return role, err
}

//SearchRole 查询所有角色
func (service *UpmsRoleService) SearchRole(site string) (result []*do.UpmsRole, err error) {
	return model.GetModel().UpmsRoleDao.Find(&do.UpmsRole{})
}

//CreateRole 创建角色
func (service *UpmsRoleService) CreateRole(site string, m *do.UpmsRole) (err error) {
	urdao := model.GetModel().UpmsRoleDao
	var role do.UpmsRole
	role.Code = "ROLE_" + m.Code
	//增加校验 名字和code都不能一样
	if ok := urdao.Found(&role); ok {
		return bizerr.GenErr("key_role_exist")
	}
	role.Name = m.Name
	//赋予角色初始化请求权限
	if ok := shadowsecurity.GetCasbinEnforcer().AddGroupingPolicy(role.Code, "ROLE_BASIC", site); !ok {
		Log.WithFields(logrus.Fields{
			"rolename": m.Name,
		}).Warn("failed to assign basic role to role in db ")
	}
	return urdao.Create(&role)
}

//DeleteRole 删除角色
func (service *UpmsRoleService) DeleteRole(site string, m *do.UpmsRole) (err error) {
	//coredb := datasource.ShardingDatasourceInstance().SDatasource("core")
	//增加判断，如果改角色下面存在有效账号则不能进行删除
	// adminDao := dao.NewAdminDao(coredb)
	// adminList, err := adminDao.Find(&model.Admin{RoleID: m.ID, State: model.Normal})
	// if err != nil {
	// 	return err
	// }
	// if len(adminList) > 0 {
	// 	return RoleAccountExistError{
	// 		error: errors.New("is exist account no can delete "),
	// 	}
	// }

	adminsList := shadowsecurity.GetCasbinEnforcer().GetFilteredGroupingPolicy(1, m.Code, site)
	if len(adminsList) > 0 {
		return bizerr.GenErr("key_role_account_exist")
	}

	if err = shadowsecurity.GetCasbinEnforcer().GetAdapter().RemoveFilteredPolicy("p", "p", 0, m.Code, site); err != nil {
		Log.WithField("name", m.Name).Error(err)
		// return err
	}

	//删除玩家的权限,没建立玩家则是ok =false
	if ok := shadowsecurity.GetCasbinEnforcer().RemoveGroupingPolicy(m.Code, "ROLE_BASIC", site); !ok {
		err = errors.New("delete role in db is error")
		Log.WithField("name", m.Name).Error(err)
		//	return err
	}

	err = model.GetModel().UpmsRoleDao.Delete(m)
	if err != nil {
		return err
	}

	return nil
}

//UpdateRole 修改角色
func (service *UpmsRoleService) UpdateRole(site string, m *do.UpmsRole) (err error) {
	return model.GetModel().UpmsRoleDao.Updates(m.ID, map[string]interface{}{"name": m.Name})
}

func (service *UpmsRoleService) Pmenu(mdao *dao.UpmsMenuDao, menu do.UpmsMenu , menuMap map[int64]do.UpmsMenu) (err error) {
	if _, ok := menuMap[menu.PNodeID]; !ok && menu.PNodeID != 1 {
		pmenu := &do.UpmsMenu{NodeID: menu.PNodeID}
		err := mdao.FindOne(pmenu)
		if err != nil {
			Log.Error(err)
			return err
		}
		menuMap[menu.PNodeID] = *pmenu
		err = service.Pmenu(mdao, *pmenu, menuMap)
		if err != nil {
			Log.Error(err)
			return err
		}
	}
	//本身菜单
	menuMap[menu.NodeID] = menu
	return nil
}

//SetPermission 权限的赋予
func (service *UpmsRoleService) SetPermission(site string, m *do.UpmsRole, menu []do.UpmsMenu) (err error) {
	if menu == nil && len(menu) == 0 {
		err = errors.New("menu is nil ")
		Log.Error(err)
		return err
	}
	//查询赋权菜单的父级菜单
	mdao := model.GetModel().UpmsMenuDao
	menuMap := make(map[int64]do.UpmsMenu)
	for _, v := range menu {
		err = service.Pmenu(mdao, v, menuMap)
		if err != nil {
			Log.Error(err)
			return err
		}
	}
	//每次给权限都先删除，然后重新添加
	if err = shadowsecurity.GetCasbinEnforcer().GetAdapter().RemoveFilteredPolicy("p", "p", 0, m.Code, site); err != nil {
		Log.WithField("name", m.Name).Error(err)
		// return err
	}
	if err := shadowsecurity.GetCasbinEnforcer().GetAdapter().RemovePolicy("g", "g", []string{m.Code, "ROLE_BASIC", site}); err != nil {
		Log.WithField("name", m.Name).Error(err)
		// return err
	}
	//删除内存中的基础权限
	shadowsecurity.GetCasbinEnforcer().GetModel().ClearPolicy()

	for _, v := range menuMap {
		if err := shadowsecurity.GetCasbinEnforcer().GetAdapter().AddPolicy("p", "p", []string{m.Code, site, v.URL, v.Method, "allow"}); err != nil {
			Log.WithFields(logrus.Fields{
				"v.URL":    v.URL,
				"v.Method": v.Method,
				"rolename": m.Name,
			}).Warn("failed to assign menu role to role in db ")
			return err
		}
	}

	//赋予角色初始化请求权限
	if ok := shadowsecurity.GetCasbinEnforcer().AddGroupingPolicy(m.Code, "ROLE_BASIC", site); !ok {
		Log.WithFields(logrus.Fields{
			"rolename": m.Name,
		}).Warn("failed to assign basic role to role in db ")
	}
	return nil
}

//SelectPermission 查询某个角色的权限
func (service *UpmsRoleService) SelectPermission(site string, role *do.UpmsRole) (result []*do.UpmsMenu, err error) {
	muDao := model.GetModel().UpmsMenuDao
	allMenus, err := muDao.Find(&do.UpmsMenu{})
	if err != nil {
		Log.Error(err)
		return nil, err
	}

	for i, menu := range allMenus {
		ok := shadowsecurity.GetCasbinEnforcer().HasPolicy(role.Code, site, menu.URL, menu.Method, "allow")
		if ok {
			allMenus[i].Selected = true
		}
	}
	return allMenus, nil
}
