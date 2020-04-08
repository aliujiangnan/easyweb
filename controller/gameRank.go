package controller

import (
	"easyweb/service"
	"easyweb/lib/htt"
)

func GlobalRanks(params map[string]interface{}, ctx htt.HttpCtx)int{
	rankList := service.GlobalRanks(ctx.UserId,ctx.GameId,params["tag"].(string))
	ctx.Send(rankList)
	return 0
}

func FriendsRanks(params map[string]interface{}, ctx htt.HttpCtx)int{
	rankList := service.FriendsRanks(ctx.UserId,ctx.GameId,params["tag"].(string))
	ctx.Send(rankList)
	return 0
}
