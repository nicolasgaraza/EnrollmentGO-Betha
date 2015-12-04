package redisSession

import (
	"gopkg.in/redis.v3"
)


func New() *redis.Client {
	client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
	
	return client
}