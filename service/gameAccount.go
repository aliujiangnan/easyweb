package service

import (
	"easyweb/model"
)

func getLevelByExp(gameid int, exp int) int{
	return 0
}

func SaveScore(userid int, gameid int, score int, tag string){
	pscore := LoadScore(userid, gameid, tag)
	if score > pscore{
		RankItWeekly(map[string]interface{}{
			"gameId": gameid,
            "periodType": "weekly",
            "userId": userid,
            "tag": "score:" + tag,
            "value": score,
		})
	}
	model.SaveScore(userid, gameid, score, tag)
}

func LoadScore(userid int, gameid int, tag string)int{
	return model.LoadScore(userid, gameid, tag)
}

func AddExp(userid int, gameid int, exp int){
	curExp := model.LoadExp(userid, gameid)
	if exp > 0{
		curLevel := getLevelByExp(gameid, curExp + exp)
        model.SaveLevel(userid, gameid, curLevel)
        model.IncExp(userid, gameid, exp)
	}
}

func UpdateInfo(userid int, gameid int, params map[string]interface{}){
	model.UpdateInfo(userid, gameid, params)
}

func SignIn(userid int, gameid int, tag string)bool{
	return model.SignIn(userid, gameid, tag)
}

func GetSignList(userid int, gameid int)[]map[string]interface{}{
	return model.GetSignList(userid, gameid)
}