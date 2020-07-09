package vo

// swagger:parameters player createPlayer
type CreateMerchantRequest struct {
	//in:body
	Body struct {
		// 帐号
		MerchantNo string `binding:"required"`
		// 登录密码
		LoginPassword string
	}
}
