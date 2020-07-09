package dto

import "shadowDemo/model/do"

type MenuNode struct {
	ID       int64
	Name     string
	NodeID   int64 //当前节点ID
	Menu     do.UpmsMenu
	Children []*MenuNode
}

func NewMenuNode(id int64, name string, node int64, ntype do.UpmsMenu) *MenuNode {
	return &MenuNode{
		ID:     id,
		Name:   name,
		NodeID: node,
		Menu:   ntype,
	}
}
