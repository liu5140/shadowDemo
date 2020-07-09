package model

import "shadowDemo/shadow-framework/logger"

var Log *logger.Logger

func init() {
	Log = logger.InitLog()
}

// swagger:model
type (
	Sex          int
	UserType     int
	YesOrNo      int
)

type IEnum interface {
	Val() int
}

func init() {
	Log = logger.InitLog()
}


const (
	_   YesOrNo = iota
	Yes         //	1
	No          //	2
)

const (
	_    Sex = iota
	Man      //	1
	Male     //	2
)

const (
	IDSpaceUser          = "user"
	IDSpaceLabel         = "label"
	IDSpaceRecommendCode = "recommendcode"
	IDSpaceOrder         = "order"
	IDSpaceSiteBill      = "sitebill"
)

const (
	_                  UserType = iota
	UserTypePlayer              //玩家	1
	UserTypeMockPlayer          //虚拟玩家	2
)


