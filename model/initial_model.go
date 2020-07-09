

package model

import "shadowDemo/model/do"

func GetModel() *Model {
	return model
}

var initialModels []interface{} = []interface{}{
	//new(do.Player),
	new(do.APIAccessReqLog),
new(do.APIAccessResLog),
new(do.IPWhiteList),
new(do.Merchant),
new(do.Order),
new(do.Player),
new(do.ProgConfig),
new(do.User),

}

func GetInitialModels() []interface{} {
	return initialModels
}

