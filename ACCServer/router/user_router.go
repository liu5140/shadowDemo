package router

import (
	"net/http"
	"shadowDemo/ACCServer/router/vo"
	"shadowDemo/middleware"
	"shadowDemo/model/dao"
	"shadowDemo/model/do"
	"shadowDemo/service"

	"github.com/gin-gonic/gin"
)

func userRouter(r *gin.RouterGroup) {
	// swagger:route POST /user user createUser
	//
	// 创建;
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
	r.POST("/user", createUser)

	// swagger:route POST /users user searchUser
	//
	// 查询列表;
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
	//       200: searchUserResponse
	//       default: genericError
	r.POST("/users", SearchUser)

	// swagger:route GET /user user getUser
	//
	// 详情;
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
	//       200: getUserResponse
	//       default: genericError
	r.GET("/user", GetUser)

	// swagger:route PUT /user user updateUser
	//
	// 修改;
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
	r.PUT("/user", UpdateUser)

	// swagger:route DELETE /user user deleteUser
	//
	// 删除;
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
	r.DELETE("/user", DeleteUser)

}

func createUser(c *gin.Context) {
	userService := service.NewUserService()
	request := &vo.CreateUserRequest{}
	err := c.Bind(&request.Body)
	if err != nil {
		newClientError(c, err)
		return
	}
	profile := c.MustGet(PROFILE).(middleware.Profile)
	user := &do.User{
		CreatedBy:  profile.Account,
	}

	err = userService.CreateUser(user)
	if err != nil {
		newServerError(c, err)
		return
	}

	newSuccess(c)
}

func SearchUser(c *gin.Context) {
	userService := service.NewUserService()
	request := &vo.SearchUserRequest{}
	err := c.Bind(&request.Body)
	if err != nil {
		newClientError(c, err)
		return
	}

	condition := &dao.UserSearchCondition{
		IDS:             request.Body.IDS,
		CreateStartTime: request.Body.CreateStartTime,
		CreateEndTime:   request.Body.CreateEndTime,
	}

	result, count, err := userService.SearchUserPaging(condition, request.Body.PageNum, request.Body.PageSize)
	if err != nil {
		newServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, vo.SearchUserResponseBody{
		Result: result,
		Count:  count,
	})
}

func GetUser(c *gin.Context) {
	id, err := bindID(c)
	if err != nil {
		newClientError(c, err)
		return
	}
	userService := service.NewUserService()
	user, err := userService.GetUserByID(id)
	if err != nil {
		newServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, vo.GetUserResponseBody{
		Result: *user,
	})
}

func UpdateUser(c *gin.Context) {
	userService := service.NewUserService()
	profile := c.MustGet(PROFILE).(middleware.Profile)
	request := &vo.UpdateUserRequest{}
	id, err := bindID(c)
	if err != nil {
		newClientError(c, err)
		return
	}
	if err := c.Bind(&request.Body); err != nil {
		newClientError(c, err)
		return
	}
	attrs := map[string]interface{}{}
	//修改人
	attrs["created_by"] = profile.Account

	err = userService.UpdateUser(id, attrs)
	if err != nil {
		newServerError(c, err)
		return
	}
	newSuccess(c)
}

func DeleteUser(c *gin.Context) {
	userService := service.NewUserService()
	id, err := bindID(c)
	if err != nil {
		newClientError(c, err)
		return
	}
	err = userService.DeleteUserByID(id)
	if err != nil {
		newServerError(c, err)
		return
	}
	newSuccess(c)
}
