package controller

import (
	"easyweb/service"
	"easyweb/lib/htt"
)

func SaveScore(params map[string]interface{}, ctx htt.HttpCtx)int{
	service.SaveScore(ctx.UserId,ctx.GameId,int(params["score"].(float64)),params["tag"].(string))
	ctx.Send(map[string]string{
		"msg":"synced",
	})
	return 0
}

func LoadScore(params map[string]interface{}, ctx htt.HttpCtx)int{
	score := service.LoadScore(ctx.UserId,ctx.GameId,params["tag"].(string))
	ctx.Send(score)
	return 0
}

func AddExp(params map[string]interface{}, ctx htt.HttpCtx)int{
	service.AddExp(ctx.UserId,ctx.GameId,int(params["exp"].(float64)))
	ctx.Send(map[string]string{
		"msg":"synced",
	})
	return 0
}

func UpdateInfo(params map[string]interface{}, ctx htt.HttpCtx)int{
	service.UpdateInfo(ctx.UserId,ctx.GameId,params)
	ctx.Send(map[string]string{
		"msg":"synced",
	})
	return 0
}

func SignIn(params map[string]interface{}, ctx htt.HttpCtx)int{
	if service.SignIn(ctx.UserId,ctx.GameId,params["tag"].(string)) {
		ctx.Send(map[string]string{
			"msg":"synced",
		})
	}else{
		ctx.Send(map[string]string{
			"msg":"today is already signed",
		})
	}

	return 0
}

func GetSignList(params map[string]interface{}, ctx htt.HttpCtx)int{
	signList := service.GetSignList(ctx.UserId,ctx.GameId)
	ctx.Send(signList)

	return 0
}