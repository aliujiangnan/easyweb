package main
 
import (
	"net/http"

	"easyweb/controller"
	"easyweb/lib/router"
	"easyweb/lib/redis"
	"easyweb/midware/validator"
)
 
func main() {
	c := redis.Connect("127.0.0.1:6379")
	defer c.Close()

	router.Get("/login", validator.Login, controller.Login)

	router.Get("/profile", validator.Auth, controller.Profile)

	router.Post("/setuserinfo", validator.Auth, controller.SetUserInfo, controller.SignIn)

	router.Post("/sign/game", validator.Auth, controller.SignIn)

	router.Get("/sign/list", validator.Auth, controller.GetSignList)

	router.Post("/userinfo/game", validator.Auth, controller.UpdateInfo)

	router.Post("/friends/sync", validator.SaveFriends, validator.Auth, controller.SaveFriends)

	router.Get("/friends/list", validator.Auth, controller.LoadFriends)

	router.Get("/Getuserinfos", validator.GetUserInfos, validator.Auth, controller.GetUserInfos)

	router.Post("/score/report", validator.SaveScore, validator.Auth, controller.SaveScore)

	router.Post("/exp/add", validator.AddExp, validator.Auth, controller.AddExp)

	router.Get("/ranks/global", validator.Auth, controller.GlobalRanks)

	router.Get("/ranks/friends", validator.Auth, controller.FriendsRanks)

	http.ListenAndServe(":443", nil)
}

