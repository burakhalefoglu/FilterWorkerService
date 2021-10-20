package redisAdapter

import (
	"github.com/go-redis/redis"
)

type RedisCache struct {
	client *redis.Client
}

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func (r RedisCache) Get (key string)(interface{}, error){

	return r.client.HGetAll(key), nil

}

func (r RedisCache) Add(key string, value map[string]interface{})(interface{}, error){
	result := r.client.HMSet(key, value)
	return result, nil
}


func (r RedisCache) Delete(key string)(interface{}, error){
	result := r.client.HDel(key)
	return result, nil
}


