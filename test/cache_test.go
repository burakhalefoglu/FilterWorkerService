package test_go

import (
	cache "FilterWorkerService/pkg"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CacheMap map[string]interface{}

var fakemap = CacheMap{"veri1": 1, "veri2": 2}

func (c CacheMap) Get (key string)(interface{}, error){
		for k, v := range c {
		if(k==key) {
			return v, nil
		}
	}
	return nil, errors.New("Veri bulunamadı!")
}

func (c CacheMap) Add(key string, value int)(error){
	c[key] = value
	return nil
}

func (c CacheMap) Delete(key string)(error){
    delete(c,key)
	return nil
}

func Test_GetFromCache(t *testing.T) {


	v1, _ := cache.GetFromCache("veri3",fakemap)
	assert.Nil(t, v1)

	
	v2, _ := cache.GetFromCache("veri2",fakemap)
	assert.NotNil(t, v2)

	v3, _ := cache.GetFromCache("veri2", fakemap)
	assert.NotNil(t, v2)

	assert.Equal(t, 2, v3, "they should be equal")
}

func Test_AppendCache(t *testing.T) {

 err := cache.AppendToCache("veri3", 8, fakemap)

 if(err != nil) {

 }

assert.NotNil(t, fakemap["veri3"])
assert.Equal(t, 8, fakemap["veri3"], "they should be equal")

}

func Test_DeleteCache(t *testing.T) {

 err := cache.DeleteToCache("veri3",fakemap)

 if(err != nil) {

 }
 assert.Equal(t, 0, fakemap["veri3"], "they should be equal")
}