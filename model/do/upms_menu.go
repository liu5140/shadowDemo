package do

import (
	"time"
)

//菜单
// sagger:model
type UpmsMenu struct {
	//ID
	ID int64 `gorm:"primary_key"`
	//创建时间
	CreatedAt *time.Time
	//修改时间
	UpdatedAt *time.Time
	//创建人
	CreatedBy string
	//名称
	Name string
	//具体Url
	URL string
	//请求method
	Method string
	//父节点ID
	PNodeID int64
	//当前节点ID
	NodeID int64
	//显示顺序
	Sequence int
	//1 文件夹  2、url 页面跳转 3、 具体按钮的url
	NodeType string
	//层级
	Level int
	//当前节点到根结点的线路
	Path string
	//前端对应码
	Code string
	//是否存在该权限
	Selected bool

}
