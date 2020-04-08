package controller

import (
	"strconv"
	"easyweb/service"
	"easyweb/lib/htt"
	"easyweb/lib/util"
)

func Login(params map[string]interface{}, ctx htt.HttpCtx)int{
	token := util.Sign(strconv.Itoa(int(params["userid"].(float64))) + strconv.Itoa(int(params["t"].(float64))))
	service.Login(params)
	ctx.Send(map[string]string{"token":token})
	return 0
}

func Profile(params map[string]interface{}, ctx htt.HttpCtx)int{
	profile := service.UserInfo(ctx.UserId,ctx.GameId)
	profile["gameid"] = ctx.GameId
	ctx.Send(profile)
	return 0
}

func GetUserInfos(params map[string]interface{}, ctx htt.HttpCtx)int{
	profileList := service.GetUserInfos(ctx.GameId,params["userids"].(string))
	ctx.Send(profileList)
	return 0
}

func SetUserInfo(params map[string]interface{}, ctx htt.HttpCtx)int{
	service.SetUserInfo(ctx.UserId,params)
	ctx.Send(map[string]string{
		"msg":"synced",
	})
	return 0
}
