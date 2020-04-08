package controller

import (
	"easyweb/service"
	"easyweb/lib/htt"
)

func SaveFriends(params map[string]interface{}, ctx htt.HttpCtx)int{
	service.SaveFriends(ctx.UserId, params["userids"].(string))
	ctx.Send(map[string]string{
		"msg":"synced",
	})
	return 0
}

func LoadFriends(params map[string]interface{}, ctx htt.HttpCtx)int{
	profileList := service.LoadFriends(ctx.UserId)
	ctx.Send(profileList)
	return 0
}