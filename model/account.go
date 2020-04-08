package model

import (
	"strconv"
	"easyweb/lib/redis"
)

func userAccountKey(userid int) string{
	return "u:a:" + strconv.Itoa(userid)
}

func SetUserProfile(userinfos map[string]interface{}){
	accountUserKey := userAccountKey(userinfos["userid"].(int))
	redis.Get().Do("HMSET", accountUserKey, userinfos)
}

func GetUserProfile(userid int)map[string]interface{}{
	accountUserKey := userAccountKey(userid)
	profile,_ := redis.Get().Do("HGETALL", accountUserKey)
	return profile.(map[string]interface{})
}

func SetUserInfo(userid int, userinfo map[string]interface{}){
	accountUserKey := userAccountKey(userid)
	redis.Get().Do("HMSET", accountUserKey, userinfo)
}
