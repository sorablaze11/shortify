// Go package for key generation system
package kgs

import (
	"github.com/go-redis/redis"
	"fmt"
	"strconv"
)

var client *redis.Client
var key int

func Init() {
	key = 0
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

func ReturnShortUrl(url string) string {
	tempKey := strconv.Itoa(GetRandomKey())
	client.Set(tempKey, url, 0)
	fmt.Println(tempKey)
	return tempKey
}

func GetUrl(shortUrl string) (string, error) {
	return client.Get(shortUrl).Result()
}