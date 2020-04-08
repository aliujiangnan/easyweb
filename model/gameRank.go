package model

import (
	"strconv"
	"easyweb/lib/redis"
)

var g = "global"

func gameRankKey(gameid int, period string, tag string, forWhat string) string{
	return "game:rank:"+forWhat+":"+period+":"+strconv.Itoa(gameid)+":"+tag
}

func GameAddRank(gameid int, period string, tag string, forWhat string, userid int, value int) {
	rankKey := gameRankKey(gameid, period, tag, forWhat)	
	redis.Get().Do("ZADD", rankKey, value, userid, "EX", "7")
}
func GameAddRankGlobal(gameid int, period string, tag string, userid int, value int) {
	redis.Get().Do("ZADD", gameRankKey(gameid, period, tag, g), value, userid)
}

func getRankByuserid(gameid int, period string, tag string, userid int) map[string]interface{}{
	globalRank,_ := redis.Get().Do("ZREVRANK", gameRankKey(gameid, period, tag, g), userid)
	globalScore,_ := redis.Get().Do("ZSCORE", gameRankKey(gameid, period, tag, g), userid)
	friendsRank,_ := redis.Get().Do("ZREVRANK", gameRankKey(gameid, period, tag, strconv.Itoa(userid)), userid)
	friendsScore,_ := redis.Get().Do("ZSCORE", gameRankKey(gameid, period, tag, strconv.Itoa(userid)), userid)

	return map[string]interface{}{
	  "g": map[string]interface{}{
		"score": globalScore,
		"rank": globalRank.(int) + 1,
	  },
	  "f": map[string]interface{}{
		"score": friendsScore,
		"rank": friendsRank.(int) + 1,
	  },
	}
  }

func GetGlobal(gameid int, period string, tag string) []map[string]interface{}{
	rankInfos := make([]map[string]interface{}, 0)
	ret,_ := redis.Get().Do("ZREVRANGE", gameRankKey(gameid, period, tag, g), 0, 99, "WITHSCORES")
	top100 := ret.([][]int)
	for i := 0; i < len(top100); i+=1{
		rankinfo := top100[i]
		rankInfos = append(rankInfos, map[string]interface{}{
			"userid": rankinfo[0],
			"score": rankinfo[1],
			"rank": i + 1,
		})
	}

	return rankInfos
}

func GetFriends(gameid int, period string, userid int, tag string, page int) []map[string]interface{}{
	presize := 100
   	start := presize * (page - 1)
   	end := (start + presize) - 1
	rankInfos := make([]map[string]interface{}, 0)
	ret,_ := redis.Get().Do("ZREVRANGE", gameRankKey(gameid, period, tag, strconv.Itoa(userid)), start, end, "WITHSCORES")
	top100 := ret.([][]int)
	for i := 0; i < len(top100); i+=1{
		rankinfo := top100[i]
		rankInfos = append(rankInfos, map[string]interface{}{
			"userid": rankinfo[0],
			"score": rankinfo[1],
			"rank": i + 1,
		})
	}

	return rankInfos
}