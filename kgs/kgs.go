// Go package for key generation system
package kgs

import (
	"github.com/go-redis/redis"
)

var client *redis.Client
var key int

func Init() {
	client = redis.NewClient(&redis.Options{
        Addr : "localhost:6379",
    })
}

// Initial implementation would give the next incremented number
func GetRandomKey() int {
	ret := key
	key++
	return ret
}

func returnShortUrl(url string) string {
	tempKey := GetRandomKey()
	client.Set(string(tempKey), url, 0)
	return string(tempKey)
}