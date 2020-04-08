package model

import (
	"strconv"
	"time"
	"easyweb/lib/redis"
)

func gameAccountKey(userid int, gameid int) string{
    return "user:gameaccount:" + strconv.Itoa(gameid) + ":" + strconv.Itoa(gameid)
}

func gameSignKey(userid int, gameid int) string{
	return "game:sign:" + strconv.Itoa(gameid) + ":" + strconv.Itoa(gameid)
}

func GetGameUserInfo(userid int, gameid int)map[string]interface{}{
	gameAccontKey := gameAccountKey(userid, gameid)
	profile,_ := redis.Get().Do("HGETALL", gameAccontKey)
	return profile.(map[string]interface{})
}

func SaveScore(userid int, gameid int, score int, tag string){
	gameAccontKey := gameAccountKey(userid, gameid)
	redis.Get().Do("HSET", gameAccontKey, "score:" + tag, score)
}

func LoadScore(userid int, gameid int, tag string)int{
	gameAccontKey := gameAccountKey(userid, gameid)
	score,_ := redis.Get().Do("HGET", gameAccontKey, "score:" + tag)
	return score.(int)
}

func IncExp(userid int, gameid int, exp int){
	gameAccontKey := gameAccountKey(userid, gameid)
	redis.Get().Do("HINCRBY", gameAccontKey, "exp", exp)
}

func SaveLevel(userid int, gameid int, level int){
	gameAccontKey := gameAccountKey(userid, gameid)
	redis.Get().Do("HSET", gameAccontKey, "level", level)
}

func LoadExp(userid int, gameid int)int{
	gameAccontKey := gameAccountKey(userid, gameid)
	score,_ := redis.Get().Do("HGET", gameAccontKey, "exp")
	return score.(int)
}

func UpdateInfo(userid int, gameid int, infos map[string]interface{}){
	gameAccontKey := gameAccountKey(userid, gameid)
	redis.Get().Do("HMSET", gameAccontKey, infos)
}

func SignIn(userid int, gameid int, tag string) bool{
	now := time.Now()
	timeStr := now.Format("2006-01-02 15:04:05")
    signKey := gameSignKey(userid, gameid)
	exist,_ := redis.Get().Do("EXISTS", signKey, timeStr)
	if exist != nil{
		return false
	}
	
	redis.Get().Do("ZADD", signKey, now.Unix(), timeStr, "EX", "7")
	return true
}

func GetSignList(userid int, gameid int)[]map[string]interface{}{
	signList,_ := redis.Get().Do("ZREVERANGE", gameSignKey(userid, gameid), 0, 7, "WITHSCORES")
	return signList.([]map[string]interface{})
}