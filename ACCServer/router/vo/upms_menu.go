package vo

import (
	"shadowDemo/model/do"
	"shadowDemo/service/dto"
)

// swagger:parameters menu searchMenu
type MenuID struct {
	//in:query
	ID string `json:"ID"`
}

// swagger:parameters menu createMenu
type CreateMenuRequest struct {
	//in:body
	Body do.UpmsMenu
}

// swagger:parameters menu UpdateMenu
type UpdateMenuRequest struct {
	// in:query
	ID int64
	// in:body
	Body do.UpmsMenu
}

// swagger:parameters menu deleteMenu
type DeleteMenuRequest struct {
	// in:body
	Body struct {
		IDset []int64
	}
}

//层级信息
// swagger:response searchMenuResponse
type SearchMenuResponse struct {
	// in: body
	Result SearchMenuBody
}


//swagger:model
type SearchMenuBody struct {
	Result dto.MenuNode
}

// swagger:parameters menu getMenuByID
type GetMenuByIDRequst struct {
	//in: query
	ID int64 `binding:"required"`
}

//层级信息
//swagger:response getMenuByIDResult
type GetMenuByIDResponse struct {
	//in: body
	Result do.UpmsMenu
}
