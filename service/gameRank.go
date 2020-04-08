package service

import (
	"strconv"
	"easyweb/model"
)

var periodType = "weekly"

func RankItWeekly(params map[string]interface{}){
	userId := params["userId"].(int)
	gameId := params["gameId"].(int)
	tag := params["tag"].(string)
	value := params["value"].(int)
	friendsList := LoadFriends(userId)
	friendsList = append(friendsList, strconv.Itoa(userId))
	model.GameAddRankGlobal(gameId, periodType, tag, userId, value);

	for friendid := range friendsList{
		model.GameAddRank(gameId, periodType, tag, strconv.Itoa(friendid), userId, value)
	}
}

func GlobalRanks(userid int, gameid int, tag string)[]map[string]interface{}{
	rankList := model.GetGlobal(userid, periodType, tag)
	return rankList
}

func FriendsRanks(userid int, gameid int, tag string)[]map[string]interface{}{
	rankList := model.GetFriends(userid, periodType, userid, tag, 1)
	return rankList
}
