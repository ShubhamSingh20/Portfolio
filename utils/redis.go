package utils

import (
	"github.com/gomodule/redigo/redis"
)


//Cache pointer to redis.Conn
var Cache redis.Conn


//CloseRedisConnection close the connection to redis
func CloseRedisConnection() {
	Cache.Close()
}

//SetupRedisCache everything releated to Redis conf goes here
func SetupRedisCache() {
	conn, err := redis.DialURL("redis://localhost")
	
	if err != nil {
		panic(err)
	}
	
	Cache = conn

}