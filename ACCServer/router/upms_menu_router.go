package router

import (
	"net/http"
	"shadowDemo/ACCServer/router/vo"
	"shadowDemo/middleware"
	"shadowDemo/model/do"
	"shadowDemo/service"

	"github.com/gin-gonic/gin"
)

func upmsMenuRouter(r *gin.RouterGroup) {
	// swagger:route POST /menu menu createMenu
	//
	// 创建菜单;
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       200: genericSuccess
	//       default: genericError
	r.POST("/menu", createMenu)

	// swagger:route GET /menu menu getMenuByID
	//
	// 菜单查询通过id;
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       200: getMenuByIDResult
	//       default: genericError
	r.GET("/menu", getMenuByID)

	// swagger:route GET /menus menu searchMenus
	//
	// 查询所有菜单;
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       200: searchMenuResponse
	//       default: genericError
	r.GET("/menus", searchMenus)

	// swagger:route PUT /menu menu UpdateMenu
	//
	// 菜单修改;
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       200: genericSuccess
	//       default: genericError
	r.PUT("/menu", UpdateMenu)

	// swagger:route DELETE /menu menu deleteMenu
	//
	// 删除菜单。删除父节点，则她的子节点一起删除：
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       200: genericSuccess
	//       default: genericError
	r.DELETE("/menu", deleteMenu)
}

func getMenuByID(c *gin.Context) {
	menuService := service.NewUpmsMenuService()
	id, err := bindID(c)
	if err != nil {
		newClientError(c, err)
		return
	}
	profile := c.MustGet(PROFILE).(middleware.Profile)
	Menu, err := menuService.GetUpmsMenuByID(profile.CurrentSite, id)
	if err != nil {
		newServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, vo.GetMenuByIDResponse{
		Result: Menu,
	})

}

func searchMenus(c *gin.Context) {
	MenuService := service.NewUpmsMenuService()
	profile := c.MustGet(PROFILE).(middleware.Profile)
	searchResult, err := MenuService.SearchUpmsMenu(profile.CurrentSite)
	if err != nil {
		newServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, vo.SearchMenuBody{
		Result: searchResult,
	})
}

func deleteMenu(c *gin.Context) {
	MenuService := service.NewUpmsMenuService()
	request := &vo.DeleteMenuRequest{}
	if err := c.Bind(&request.Body); err != nil {
		newClientError(c, err)
		return
	}
	profile := c.MustGet(PROFILE).(middleware.Profile)

	if err := MenuService.DeleteMenu(profile.CurrentSite, request.Body.IDset); err != nil {
		newServerError(c, err)
		return
	}
	newSuccess(c)
}

func UpdateMenu(c *gin.Context) {
	MenuService := service.NewUpmsMenuService()
	request := &vo.UpdateMenuRequest{}
	id, err := bindID(c)
	if err != nil {
		newClientError(c, err)
		return
	}
	if err := c.Bind(&request.Body); err != nil {
		newClientError(c, err)
		return
	}
	profile := c.MustGet(PROFILE).(middleware.Profile)
	Menu, err := MenuService.GetUpmsMenuByID(profile.CurrentSite, id)
	if err != nil {
		newServerError(c, err)
		return
	}
	Menu.Name = request.Body.Name
	Menu.URL = request.Body.URL
	Menu.Method = request.Body.Method
	Menu.PNodeID = request.Body.PNodeID
	Menu.NodeID = request.Body.NodeID
	Menu.Sequence = request.Body.Sequence
	Menu.NodeType = request.Body.NodeType
	Menu.Level = request.Body.Level
	Menu.Path = request.Body.Path

	if err := MenuService.UpdateMenu(profile.CurrentSite, &Menu); err != nil {
		newServerError(c, err)
		return
	}

	newSuccess(c)
}

func createMenu(c *gin.Context) {
	MenuService := service.NewUpmsMenuService()
	request := &vo.CreateMenuRequest{}
	profile := c.MustGet(PROFILE).(middleware.Profile)

	if err := c.Bind(&request.Body); err != nil {
		newClientError(c, err)
		return
	}

	if err := MenuService.CreateUpmsMenu(profile.CurrentSite, &do.UpmsMenu{
		Name:     request.Body.Name,
		URL:      request.Body.URL,
		Method:   request.Body.Method,
		PNodeID:  request.Body.PNodeID,
		NodeID:   request.Body.NodeID,
		Sequence: request.Body.Sequence,
		NodeType: request.Body.NodeType,
		Level:    request.Body.Level,
		Path:     request.Body.Path,
	}); err != nil {
		newServerError(c, err)
		return
	}

	newSuccess(c)
}
