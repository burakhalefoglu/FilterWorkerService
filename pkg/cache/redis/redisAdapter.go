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

var CacheForRedis = RedisCache{
	client: rdb,
}

func (r RedisCache) Get (key string)(map[string]string, error){

	result := r.client.HGetAll(key)
	if result.Err() != nil {
		return nil, result.Err()
	}
	return r.client.HGetAll(key).Val(), nil

}

func (r RedisCache) Add(key string, value map[string]interface{})(interface{}, error){
	result := r.client.HMSet(key, value)
	if result.Err() != nil{
		return nil, result.Err()
	}
	return result.Val(), nil
}


func (r RedisCache) Delete(key string)(interface{}, error){
	result := r.client.HDel(key)
	if result.Err() !=nil{
		return nil, result.Err()
	}
	return result.Val(), nil
}


