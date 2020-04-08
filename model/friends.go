package model

import (
	"strconv"
	"easyweb/lib/redis"
)

func userFriendsKey(userid int) string{
    return "user:friends:"+strconv.Itoa(userid)
}

func SaveUserFriends(userid int, userids []string){
    friendsKey := userFriendsKey(userid)
	redis.Get().Do("SADD", friendsKey, userids)
}

func LoadFriends(userid int)[]string{
	friendsKey := userFriendsKey(userid)
	friends,_ := redis.Get().Do("SMEMBERS", friendsKey)
	return friends.([]string)
}