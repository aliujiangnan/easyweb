package redis

import (
	"log"
	"github.com/garyburd/redigo/redis"
)

var redisConn redis.Conn
var redisErr error
func Connect(host string) redis.Conn{
	redisConn, redisErr = redis.Dial("tcp", "127.0.0.1:6379")
	if redisErr != nil {
		log.Println("Connect to redis error", redisErr)
		return redisConn
	}

	return nil
}

func Get() redis.Conn{
	return redisConn
}