package service

import (
	"errors"
	"shadowDemo/model"
	"shadowDemo/model/do"
	"shadowDemo/zframework/datasource"
	shadowsecurity "shadowDemo/zframework/security"
	"strings"

	"github.com/sirupsen/logrus"

)

type RoleService struct {
}

var roleTreeService *RoleService

func NewRoleService() *RoleService {
	if roleTreeService == nil {
		l.Lock()
		if roleTreeService == nil {
			roleTreeService = &RoleService{}
		}
		l.Unlock()
	}
	return roleTreeService
}

func (service RoleService) GetRoleByRoleCodes(site string, roleCode string) ([]do.UpmsRole, error) {
	return model.GetModel().UpmsRoleDao.GetRolesByCodes(strings.Split(roleCode, ","))
}

func (service RoleService) GetRoleByID(site string, id int64) (do.UpmsRole, error) {
	role := do.UpmsRole{
		ID: id,
	}
	err := model.GetModel().UpmsRoleDao.Get(&role)
	return role, err
}

//SearchRole 查询所有角色
func (service RoleService) SearchRole(site string) (result []*do.UpmsRole, err error) {
	return model.GetModel().UpmsRoleDao.Find(&do.UpmsRole{})
}

//CreateRole 创建角色
func (service RoleService) CreateRole(site string, m *do.UpmsRole) (err error) {
	urdao := model.GetModel().UpmsRoleDao
	var role do.UpmsRole
	role.Code = "ROLE_" + m.Code
	//增加校验 名字和code都不能一样
	if ok := urdao.Existed(&role); ok {
		return RoleExistError{
			error: errors.New("role is exist"),
		}
	}
	role.Name = m.Name
	role.MaxAmount = m.MaxAmount
	//赋予角色初始化请求权限
	if ok := shadowsecurity.GetCasbinEnforcer().AddGroupingPolicy(role.Code, "ROLE_BASIC", site); !ok {
		Log.WithFields(logrus.Fields{
			"rolename": m.Name,
		}).Warn("failed to assign basic role to role in db ")
	}
	// menus := shadowsecurity.GetCasbinEnforcer().GetFilteredNamedPolicy("p", 0, "ROLE_BASIC", "sitea")

	// for _, v := range menus {
	// 	if ok := shadowsecurity.GetCasbinEnforcer().AddPolicy(role.Code, site, v[2], v[3], "allow"); !ok {
	// 		err = errors.New("failed to assign role to role in db ")
	// 		Log.WithField("name", m.Name).Error(err)
	// 		return err
	// 	}
	// }
	// err = shadowsecurity.GetCasbinEnforcer().LoadPolicy()
	// if err != nil{
	// 	Log.Error(err)
	// 	return
	// }
	return urdao.Create(&role)
}

//DeleteRole 删除角色
func (service RoleService) DeleteRole(site string, m *model.Role) (err error) {
	//coredb := datasource.ShardingDatasourceInstance().SDatasource("core")
	db := datasource.ShardingDatasourceInstance().SDatasource(site)
	urdao := dao.NewRoleDao(db)
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
		err = RoleAccountExistError{
			error: errors.New("is exist account no can delete "),
		}
		Log.Error(err)
		return
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

	err = urdao.Delete(m)
	if err != nil {
		return err
	}

	return nil
}

//UpdateRole 修改角色
func (service RoleService) UpdateRole(site string, m *model.Role) (err error) {
	db := datasource.ShardingDatasourceInstance().SDatasource(site)
	urdao := dao.NewRoleDao(db)
	return urdao.Updates(m.ID, map[string]interface{}{"name": m.Name, "max_amount": m.MaxAmount})
}

func (service RoleService) Pmenu(mdao *dao.MenuDao, menu model.Menu, menuMap map[int64]model.Menu) (err error) {
	if _, ok := menuMap[menu.PNodeID]; !ok && menu.PNodeID != 1 {
		pmenu := &model.Menu{NodeID: menu.PNodeID}
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
func (service RoleService) SetPermission(site string, m *model.Role, menu []model.Menu) (err error) {
	if menu == nil && len(menu) == 0 {
		err = errors.New("menu is nil ")
		Log.Error(err)
		return err
	}
	//查询赋权菜单的父级菜单
	db := datasource.ShardingDatasourceInstance().SDatasource("core")
	mdao := dao.NewMenuDao(db)
	menuMap := make(map[int64]model.Menu)
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
	// if err := shadowsecurity.GetCasbinEnforcer().GetAdapter().RemovePolicy("p", "p", []string{m.Code, site}); err != nil {
	// 	Log.WithField("name", m.Name).Error(err)
	// 	// return err
	// }
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
	// err = shadowsecurity.GetCasbinEnforcer().LoadPolicy()
	// if err != nil{
	// 	Log.Error(err)
	// 	return
	// }
	return nil
}

//SelectPermission 查询某个角色的权限
func (service RoleService) SelectPermission(site string, role *model.Role) (result []model.Menu, err error) {
	coredb := datasource.ShardingDatasourceInstance().SDatasource("core")
	muDao := dao.NewMenuDao(coredb)
	allMenus, err := muDao.Find(&model.Menu{})
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
