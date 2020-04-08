package validator

import (
	"easyweb/lib/util"
	"easyweb/lib/erro"
	"easyweb/lib/htt"
)

func AddExp(p map[string]interface{}, ctx htt.HttpCtx)int{
	return util.Validate(p,ctx,map[string]erro.Error{
		"exp":erro.NewError(11100, "params exp is not correct"),
		"t":erro.NewError(11101, "params t is not correct"),
	})
}

func SaveScore(p map[string]interface{}, ctx htt.HttpCtx)int{
	return util.Validate(p,ctx,map[string]erro.Error{
		"score":erro.NewError(11200, "params score is not correct"),
		"tag":erro.NewError(11201, "params tag is not correct"),
	})
}

func GetUserInfos(p map[string]interface{}, ctx htt.HttpCtx)int{
	return util.Validate(p,ctx,map[string]erro.Error{
		"userids":erro.NewError(11300, "params userids is not correct"),
	})
}

func SaveFriends(p map[string]interface{}, ctx htt.HttpCtx)int{
	return util.Validate(p,ctx,map[string]erro.Error{
		"userids":erro.NewError(11400, "params userids is not correct"),
	})
}

func Login(p map[string]interface{}, ctx htt.HttpCtx)int{
	return util.Validate(p,ctx,map[string]erro.Error{
		"userid":erro.NewError(11600, "params userid is not correct"),
		"gameid":erro.NewError(11601, "params gameid is not correct"),
		"username":erro.NewError(11602, "params username is not correct"),
		"nickname":erro.NewError(11603, "params nickname is not correct"),
		"gender":erro.NewError(11604, "params gender is not correct"),
		"avatar":erro.NewError(11605, "params avatar is not correct"),
		"t":erro.NewError(11606, "params t is not correct"),
	})
}