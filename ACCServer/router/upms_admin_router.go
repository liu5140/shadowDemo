package router

// func upmsAdminRouter(r *gin.RouterGroup) {
// 	// swagger:route POST /upmsAdmin upmsAdmin createUpmsAdmin
// 	//
// 	// 创建;
// 	//
// 	//     Consumes:
// 	//     - application/json
// 	//
// 	//     Produces:
// 	//     - application/json
// 	//
// 	//     Schemes: http
// 	//
// 	//     Security:
// 	//       api_key:
// 	//       oauth: read, write
// 	//
// 	//     Responses:
// 	//       200: genericSuccess
// 	//       default: genericError
// 	r.POST("/upmsAdmin", createUpmsAdmin)

// 	// swagger:route POST /upmsAdmins upmsAdmin searchUpmsAdmin
// 	//
// 	// 查询列表;
// 	//
// 	//     Consumes:
// 	//     - application/json
// 	//
// 	//     Produces:
// 	//     - application/json
// 	//
// 	//     Schemes: http
// 	//
// 	//     Security:
// 	//       api_key:
// 	//       oauth: read, write
// 	//
// 	//     Responses:
// 	//       200: searchUpmsAdminResponse
// 	//       default: genericError
// 	r.POST("/upmsAdmins", SearchUpmsAdmin)

// 	// swagger:route GET /upmsAdmin upmsAdmin getUpmsAdmin
// 	//
// 	// 详情;
// 	//
// 	//     Consumes:
// 	//     - application/json
// 	//
// 	//     Produces:
// 	//     - application/json
// 	//
// 	//     Schemes: http
// 	//
// 	//     Security:
// 	//       api_key:
// 	//       oauth: read, write
// 	//
// 	//     Responses:
// 	//       200: getUpmsAdminResponse
// 	//       default: genericError
// 	r.GET("/upmsAdmin", GetUpmsAdmin)

// 	// swagger:route PUT /upmsAdmin upmsAdmin updateUpmsAdmin
// 	//
// 	// 修改;
// 	//
// 	//     Consumes:
// 	//     - application/json
// 	//
// 	//     Produces:
// 	//     - application/json
// 	//
// 	//     Schemes: http
// 	//
// 	//     Security:
// 	//       api_key:
// 	//       oauth: read, write
// 	//
// 	//     Responses:
// 	//       200: genericSuccess
// 	//       default: genericError
// 	r.PUT("/upmsAdmin", UpdateUpmsAdmin)

// 	// swagger:route DELETE /upmsAdmin upmsAdmin deleteUpmsAdmin
// 	//
// 	// 删除;
// 	//
// 	//     Consumes:
// 	//     - application/json
// 	//
// 	//     Produces:
// 	//     - application/json
// 	//
// 	//     Schemes: http
// 	//
// 	//     Security:
// 	//       api_key:
// 	//       oauth: read, write
// 	//
// 	//     Responses:
// 	//       200: genericSuccess
// 	//       default: genericError
// 	r.DELETE("/upmsAdmin", DeleteUpmsAdmin)

// }

// func createUpmsAdmin(c *gin.Context) {
// 	upmsAdminService := service.NewUpmsAdminService()
// 	request := &vo.CreateUpmsAdminRequest{}
// 	err := c.Bind(&request.Body)
// 	if err != nil {
// 		newClientError(c, err)
// 		return
// 	}
// 	profile := c.MustGet(PROFILE).(middleware.Profile)
// 	upmsAdmin := &do.UpmsAdmin{
// 		CreatedBy:  profile.Account,
// 	}

// 	err = upmsAdminService.CreateUpmsAdmin(upmsAdmin)
// 	if err != nil {
// 		newServerError(c, err)
// 		return
// 	}

// 	newSuccess(c)
// }

// func SearchUpmsAdmin(c *gin.Context) {
// 	upmsAdminService := service.NewUpmsAdminService()
// 	request := &vo.SearchUpmsAdminRequest{}
// 	err := c.Bind(&request.Body)
// 	if err != nil {
// 		newClientError(c, err)
// 		return
// 	}

// 	condition := &dao.UpmsAdminSearchCondition{
// 		IDS:             request.Body.IDS,
// 		CreateStartTime: request.Body.CreateStartTime,
// 		CreateEndTime:   request.Body.CreateEndTime,
// 	}

// 	result, count, err := upmsAdminService.SearchUpmsAdminPaging(condition, request.Body.PageNum, request.Body.PageSize)
// 	if err != nil {
// 		newServerError(c, err)
// 		return
// 	}

// 	c.JSON(http.StatusOK, vo.SearchUpmsAdminResponseBody{
// 		Result: result,
// 		Count:  count,
// 	})
// }

// func GetUpmsAdmin(c *gin.Context) {
// 	id, err := bindID(c)
// 	if err != nil {
// 		newClientError(c, err)
// 		return
// 	}
// 	upmsAdminService := service.NewUpmsAdminService()
// 	upmsAdmin, err := upmsAdminService.GetUpmsAdminByID(id)
// 	if err != nil {
// 		newServerError(c, err)
// 		return
// 	}

// 	c.JSON(http.StatusOK, vo.GetUpmsAdminResponseBody{
// 		Result: *upmsAdmin,
// 	})
// }

// func UpdateUpmsAdmin(c *gin.Context) {
// 	upmsAdminService := service.NewUpmsAdminService()
// 	profile := c.MustGet(PROFILE).(middleware.Profile)
// 	request := &vo.UpdateUpmsAdminRequest{}
// 	id, err := bindID(c)
// 	if err != nil {
// 		newClientError(c, err)
// 		return
// 	}
// 	if err := c.Bind(&request.Body); err != nil {
// 		newClientError(c, err)
// 		return
// 	}
// 	attrs := map[string]interface{}{}
// 	//修改人
// 	attrs["created_by"] = profile.Account

// 	err = upmsAdminService.UpdateUpmsAdmin(id, attrs)
// 	if err != nil {
// 		newServerError(c, err)
// 		return
// 	}
// 	newSuccess(c)
// }

// func DeleteUpmsAdmin(c *gin.Context) {
// 	upmsAdminService := service.NewUpmsAdminService()
// 	id, err := bindID(c)
// 	if err != nil {
// 		newClientError(c, err)
// 		return
// 	}
// 	err = upmsAdminService.DeleteUpmsAdminByID(id)
// 	if err != nil {
// 		newServerError(c, err)
// 		return
// 	}
// 	newSuccess(c)
// }
