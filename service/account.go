package service

import (
	"easyweb/model"
)

func Login(info map[string]interface{}){
	saveInfo := map[string]interface{}{
		"userid" : int(info["userid"].(float64)),
		"username" : info["username"].(string),
		"nickname" : info["nickname"].(string),
		"gender" : int(info["gender"].(float64)),
		"avatar" : info["avatar"].(string),
		"t" : int(info["t"].(float64)),
	}
	model.SetUserProfile(saveInfo)
}

func UserInfo(userid int, gameid int)map[string]interface{}{
	profile := model.GetUserProfile(userid)
	userInfo := model.GetGameUserInfo(userid, gameid)

	for k,v := range profile {
		userInfo[k] = v
	}

	return userInfo
}

func GetUserInfos(gameid int, userids string)[]map[string]interface{}{
	var useridArr []string
	data := make([]map[string]interface{}, 0)
	for userid := range useridArr{
		data = append(data, UserInfo(userid, gameid))
	}
	return data
}

func SetUserInfo(userid int, params map[string]interface{}){
	model.SetUserInfo(userid, params)
}
