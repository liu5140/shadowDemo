package router

import (
	"net/http"
	"shadowDemo/ACCServer/router/vo"
	"shadowDemo/middleware"
	"shadowDemo/model/do"
	"shadowDemo/service"
	"strings"

	"github.com/gin-gonic/gin"
)

func upmsRoleRouter(r *gin.RouterGroup) {
	// swagger:route POST /role role createRole
	//
	// 创建角色;
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
	r.POST("/role", createRole)

	// swagger:route GET /roles role searchRoles
	//
	// 查询所有角色;
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
	//       200: searchRoleResponse
	//       default: genericError
	r.GET("/roles", searchRoles)

	// swagger:route GET /role/permissions role getPermissionByID
	//
	// 查询角色对应的权限;
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
	//       200: searchRoleMenuResponse
	//       default: genericError
	r.GET("/role/permissions", getPermissionByID)

	// swagger:route POST /role/permission role createPermission
	//
	// 修改角色权限，添加权限也用;
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
	r.POST("/role/permission", createPermission)

	// swagger:route PUT /role role UpdateRole
	//
	// 修改角色，只能修改name;
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
	r.PUT("/role", updateRole)

	// swagger:route DELETE /role role deleteRole
	//
	// 删除角色：
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
	r.DELETE("/role", deleteRole)
}

func searchRoles(c *gin.Context) {
	roleService := service.NewUpmsRoleService()
	profile := c.MustGet(PROFILE).(middleware.Profile)
	searchResult, err := roleService.SearchRole(profile.CurrentSite)
	if err != nil {
		newServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, vo.SearchRoleBody{
		Data: searchResult,
	})
}

func deleteRole(c *gin.Context) {
	roleService := service.NewUpmsRoleService()
	profile := c.MustGet(PROFILE).(middleware.Profile)
	id, err := bindID(c)
	if err != nil {
		newClientError(c, err)
		return
	}

	role, err := roleService.GetRoleByID(profile.CurrentSite, id)
	if err != nil {
		newServerError(c, err)
		return
	}

	if err := roleService.DeleteRole(profile.CurrentSite, &role); err != nil {
		newServerError(c, err)
		return
	}
	newSuccess(c)
}

func updateRole(c *gin.Context) {
	roleService := service.NewUpmsRoleService()
	profile := c.MustGet(PROFILE).(middleware.Profile)
	request := &vo.UpdateRoleRequest{}
	id, err := bindID(c)
	if err != nil {
		newClientError(c, err)
		return
	}
	if err := c.Bind(&request.Body); err != nil {
		newClientError(c, err)
		return
	}

	role, err := roleService.GetRoleByID(profile.CurrentSite, id)
	if err != nil {
		newServerError(c, err)
		return
	}
	if request.Body.Name != "" {
		role.Name = request.Body.Name
	}
	if err := roleService.UpdateRole(profile.CurrentSite, &role); err != nil {
		newServerError(c, err)
		return
	}

	newSuccess(c)
}

func createRole(c *gin.Context) {
	roleService := service.NewUpmsRoleService()
	request := &vo.CreateRoleRequest{}
	profile := c.MustGet(PROFILE).(middleware.Profile)
	if err := c.Bind(&request.Body); err != nil {
		newClientError(c, err)
		return
	}
	if err := roleService.CreateRole(profile.CurrentSite, &do.UpmsRole{
		Name:      strings.TrimSpace(request.Body.Name),
		Code:      request.Body.Code,
	}); err != nil {
		newServerError(c, err)
		return
	}

	newSuccess(c)
}

func getPermissionByID(c *gin.Context) {
	roleService := service.NewUpmsRoleService()
	profile := c.MustGet(PROFILE).(middleware.Profile)
	id, err := bindID(c)
	if err != nil {
		newClientError(c, err)
		return
	}
	role, err := roleService.GetRoleByID(profile.CurrentSite, id)
	if err != nil {
		newServerError(c, err)
		return
	}
	result, err := roleService.SelectPermission(profile.CurrentSite, &role)
	if err != nil {
		newServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, vo.SearchRoleMenuResponse{
		Result: result,
	})
}

func createPermission(c *gin.Context) {
	roleService := service.NewUpmsRoleService()
	request := &vo.CreateRolePermissionRequest{}
	profile := c.MustGet(PROFILE).(middleware.Profile)
	id, err := bindID(c)
	if err != nil {
		newClientError(c, err)
		return
	}
	if err := c.Bind(&request.Body); err != nil {
		newClientError(c, err)
		return
	}
	role, err := roleService.GetRoleByID(profile.CurrentSite, id)
	if err != nil {
		newServerError(c, err)
		return
	}
	if err := roleService.SetPermission(profile.CurrentSite, &role, request.Body); err != nil {
		newServerError(c, err)
		return
	}
	newSuccess(c)
}
