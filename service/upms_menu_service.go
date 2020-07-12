package service

import (
	"github.com/jinzhu/gorm"
	"shadowDemo/model"
	"shadowDemo/model/dao"
	"shadowDemo/model/do"
	"shadowDemo/service/dto"
	"shadowDemo/zframework/datasource"
	"strconv"
)

type UpmsMenuService struct {
}

var menuService *UpmsMenuService

var tree *dto.MenuNode

func NewUpmsMenuService() *UpmsMenuService {
	if menuService == nil {
		l.Lock()
		if menuService == nil {
			menuService = &UpmsMenuService{}
		}
		l.Unlock()
	}
	return menuService
}


func (service *UpmsMenuService) GetUpmsMenuByID(coresite string, id int64) (do.UpmsMenu, error) {
	menu := do.UpmsMenu{
		ID: id,
	}
	err := model.GetModel().UpmsMenuDao.Get(&menu)
	return menu, err
}

func (service *UpmsMenuService) SearchUpmsMenu(coresite string) (result dto.MenuNode, err error) {
	db := datasource.DataSourceInstance().Master()
	node := dto.MenuNode{
		NodeID: 1,
	}
	if err := service.FindTree(db, &node); err != nil {
		Log.Errorln(err)
		return dto.MenuNode{}, err
	}
	return node, nil
}

//CreateMenu 创建一个菜单  传入父节点，并且传入一个顺序   顺序主要用户节点的显示 还有类型
func (service *UpmsMenuService) CreateUpmsMenu(coresite string, m *do.UpmsMenu) (err error) {
	urdao := model.GetModel().UpmsMenuDao
	cid := GenAdminID() //生成 一个唯一id
	if m.PNodeID > 0 && m.NodeType != "1" {
		pNode := do.UpmsMenu{
			NodeID: m.PNodeID,
		}
		err = urdao.FindOne(&pNode)
		if err != nil {
			Log.Error(err)
			return err
		}
		return urdao.Create(&do.UpmsMenu{
			Name:     m.Name,
			URL:      m.URL,
			Sequence: m.Sequence,
			PNodeID:  m.PNodeID,
			NodeID:   cid,
			NodeType: m.NodeType,
			Level:    pNode.Level + 1,
			Path:     pNode.Path + strconv.FormatInt(cid, 10) + ":",
		})
	} else {
		return urdao.Create(&do.UpmsMenu{
			Name:     m.Name,
			URL:      m.URL,
			Sequence: m.Sequence,
			PNodeID:  m.PNodeID,
			NodeID:   cid,
			NodeType: m.NodeType,
			Level:    1,
			Path:     strconv.FormatInt(cid, 10) + ":",
		})
	}

}

//DeleteMenu 删除选择菜单 ,则把他的子菜单 和本身全部删掉
func (service *UpmsMenuService) DeleteMenu(coresite string, pid []int64) (err error) {
	db := datasource.ShardingDatasourceInstance().SDatasource(coresite)
	tx := db.Begin()
	defer closeTx(tx, &err)
	udao := dao.NewUpmsMenuDao(tx)
	for _, v := range pid {
		pNode := do.UpmsMenu{
			NodeID: v,
		}

		err = udao.FindOne(&pNode)
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				//如果没找到则继续循环，
				continue
			}
			Log.Error(err)
			return err
		}
		var result []do.UpmsMenu
		result, _, err = udao.SearchUpmsMenus(&dao.UpmsMenuSearchCondition{Path: pNode.Path},nil)
		if err != nil {
			Log.Error(err)
			return err
		}
		menuid := make([]int64, len(result))
		for i, m := range result {
			menuid[i] = m.ID
		}

		err = udao.BatchDelete(menuid)
		if err != nil {
			Log.Error(err)
			return err
		}

	}
	return nil
}

//UpdateMenu 将已存在的子节点移动另一个父节点下  pid  是需要移动的， cid 是移动到的那个点
func (service *UpmsMenuService) UpdateMenu(coresite string, m *do.UpmsMenu) (err error) {
	db := datasource.ShardingDatasourceInstance().SDatasource(coresite)
	tx := db.Begin()
	defer closeTx(tx, &err)
	udao := dao.NewUpmsMenuDao(tx)

	toNodeModel := do.UpmsMenu{
		NodeID: m.NodeID,
	}

	err = udao.FindOne(&toNodeModel)
	if err != nil {
		Log.Error(err)
		return err
	}
	//父节点相同，则不修改path信息 。只是修改名称，和对应url
	if m.PNodeID == toNodeModel.PNodeID {
		err = udao.Save(m)
		if err != nil {
			Log.Error(err)
			return err
		}
	} else {
		fromNodePath := m.Path
		toNodePath := toNodeModel.Path

		err = udao.Save(m)

		if err != nil {
			Log.Error(err)
			return err
		}

		formNodeIDString := strconv.FormatInt(m.NodeID, 10)

		toPath := toNodePath + formNodeIDString + ":"
		Log.Debug("toPath", toPath)
		//2、更新所有pid path 开头的数据 为
		err = udao.UpdatePath(fromNodePath, toPath)
		if err != nil {
			Log.Error(err)
			return err
		}
	}

	return nil
}

//查找一个节点的直接子节点
func (service *UpmsMenuService) FindChildren(tx *gorm.DB, nid int64) (children []*dto.MenuNode, err error) {
	udao := dao.NewUpmsMenuDao(tx)
	node := do.UpmsMenu{
		PNodeID: nid,
	}

	result, err := udao.Find(&node)
	if err != nil {
		return nil, err
	}

	for _, e := range result {
		children = append(children, dto.NewMenuNode(e.ID, e.Name, e.NodeID, *e))
	}
	return
}

//查询下级，例如 代理查询下面代理有多少
func (service *UpmsMenuService) FindUser(tx *gorm.DB, node do.UpmsMenu) (children []dto.MenuNode, err error) {
	udao := dao.NewUpmsMenuDao(tx)

	result, err := udao.Find(&node)
	if err != nil {
		return nil, err
	}

	for _, e := range result {
		children = append(children, *dto.NewMenuNode(e.ID, e.Name, e.NodeID, *e))
	}
	return
}

//递归实现(返回树状结果得数据)
func (service *UpmsMenuService) FindTree(tx *gorm.DB, node *dto.MenuNode) (err error) {
	children, _ := service.FindChildren(tx, node.NodeID)
	if len(children) > 0 {
		node.Children = children
		for _, child := range children {
			err := service.FindTree(tx, child)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (service *UpmsMenuService) GetNode(tx *gorm.DB, id int64) (node *dto.MenuNode, err error) {
	udao := dao.NewUpmsMenuDao(tx)
	ur := do.UpmsMenu{
		NodeID: id,
	}
	err = udao.FindOne(&ur)
	if err != nil {
		return node, err
	}
	node = dto.NewMenuNode(ur.ID, ur.Name, ur.NodeID, ur)
	return node, err
}
