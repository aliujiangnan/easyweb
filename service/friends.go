package service

import (
	"easyweb/model"
)

func SaveFriends(userid int, userids string){
	var useridArr []string
	model.SaveUserFriends(userid, useridArr)
}

func LoadFriends(userid int)[]string{
	return model.LoadFriends(userid)
}