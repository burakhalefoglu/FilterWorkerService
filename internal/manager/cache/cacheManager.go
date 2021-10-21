package cacheManager

import (
	redisAdapter "FilterWorkerService/pkg/cache/redis"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	cache "FilterWorkerService/pkg/cache"
	"strconv"
)

func ManageCache(NameOfMap string, elementKey string) (int, error) {
	ctx := mgm.Ctx()
	type DictStruct struct {
		mgm.DefaultModel `bson:",inline"`
		Key   string `json:"Field Str"`
		Value int `json:"Field Int"`
	}

	//* map[string]map[string]string
	//* {NameOfMap: {elementKey:elementvalue}, {elementKey: elementValue}}
	//* {"Continent": {"Türkiye":"1"}, {"Yunanistan":"2"}}
	//* {"org": {"türknet":"1"}, {"turkcell":"2"}}

	//* Cache'ten Id yi control et!
	v, err := cache.GetFromCache(NameOfMap, redisAdapter.CacheForRedis)
	if err != nil {
		return 0, err
	}

	for key, value := range v {
		//* Cache'te Id var ise bu veriyi kullan!
		if key == elementKey {
			log.Printf(key, value)
			i, err := strconv.Atoi(value)
			if err != nil {
				return 0, err
			}
			return i, nil
		}
	}

	//* Cache'ten Id yok ise bu veriyi mongodb den getir. Yoksa ekle!
	log.Println("finding document")
	coll := mgm.Coll(&DictStruct{})
	result := coll.FindOne(ctx, bson.M{"key" : elementKey})

	if result.Err() != nil {
		return 0, result.Err()
	}
 	var resultModel = &DictStruct{};
	result.Decode(resultModel)
	log.Println(resultModel)

	//if elementArray[0].Value == 0{
	//	log.Println("adding all document")
	//	elementArray.key = elementKey
	//	elementArray.value = 1
	//	_, err = mongodb.AddCollection(NameOfMap,
	//		bson.D{
	//		{"key",elementKey},
	//		{"value", 1}})
	//	if err != nil {
	//		return 0, err
	//	}
	//}

	////* bu topiği cache'e ekle!
	//_, err = cache.DeleteToCache(NameOfMap, redisAdapter.CacheForRedis)
	//if err != nil {
	//	return 0, err
	//}

	return len(v) + 1, nil
}
