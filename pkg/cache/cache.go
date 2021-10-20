package cache

import (
	"errors"
)


type CacheMap map[string]map[string]interface{}


type cacheInterface interface{
	Get(key string)(interface{}, error)
	Add(key string, value interface{})(interface{}, error)
	Delete(key string)(interface{}, error)
}


func GetFromCache(key string, c cacheInterface) (interface{}, error) {
	v, err := c.Get(key)
	if(err != nil) {
		return nil, errors.New("veri bulunamadı!")
	}
	return v, nil
}

func AppendToCache(key string, v map[string]interface{}, c cacheInterface) (interface{}, error) {

	result, err := c.Add(key, v)
	if(err != nil) {
		return nil, errors.New("veri eklenemedi!")
	}
	return result, nil

}


func DeleteToCache(key string, c cacheInterface) (interface{}, error) {

	result, err := c.Delete(key)
	if(err != nil) {
		return nil, errors.New("veri silinemedi!")
	}
	return result, nil
}