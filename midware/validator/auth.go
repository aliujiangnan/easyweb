package validator

import (
	"strconv"
	"strings"
	"easyweb/lib/htt"
	"easyweb/lib/erro"
	"easyweb/lib/util"
)

func verify(token string, userid int, time int) bool {
	return token == util.Sign(strconv.Itoa(userid) + strconv.Itoa(time))
}

func Auth(p map[string]interface{}, ctx htt.HttpCtx)int{
	authString := ctx.Writer.Header().Get("authorization")

	if authString == ""{
		ctx.Send(erro.NewError(10002, "token error").ToMap())
		return -1
	}

	authArray := strings.Split(authString, " ")
	authType := authArray[0]
	if authType != "Bearer"{
		ctx.Send(erro.NewError(10002, "token error").ToMap())
		return -1
	}
	authToken := authArray[1]
	userId := int(p["userid"].(float64))
	if verify(authToken, userId, int(p["t"].(float64))){
		ctx.Send(erro.NewError(10003, "token error").ToMap())
		return -1
	}
	ctx.UserId = userId
	return 0
}