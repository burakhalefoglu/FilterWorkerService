package baseManager

import (
	redisAdapter "FilterWorkerService/pkg/cache/redis"
	"go.mongodb.org/mongo-driver/bson"
)
import cache "FilterWorkerService/pkg/cache"
import "strconv"
import mongodb "FilterWorkerService/internal/dataAccess/mongodb"
func ManageCache(collection string, k string) (int, error) {

	//Todo: Cache'ten Id yi control et!
	v, err := cache.GetFromCache(collection, redisAdapter.CacheForRedis)
	if err != nil{
		return 0, err
	}
	for key, value := range v {
		//Todo: Cache'ten Id var ise bu veriyi kullan!
		if(key == k){
			i, err := strconv.Atoi(value)
			if err !=nil{
				return 0, err
			}
			return i, nil
		}
	}
	 // key Name: k
	// value: len(v) + 1
	//Todo: Cache'ten Id yok ise bu veriyi mongodbye gönder!
	_, err = mongodb.AddCollection(collection, bson.D{
		{k, len(v) + 1},
	})
	if err != nil{
		return 0, err
	}

	//Todo: bu topiği cache'ten sil!
	_, err = cache.DeleteToCache(collection, redisAdapter.CacheForRedis)
	if err != nil{
		return 0, err
	}

	return len(v) + 1, nil
}