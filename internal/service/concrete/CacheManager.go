package concrete

import (
	"FilterWorkerService/internal/IoC"
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/repository/abstract"
	Cache "FilterWorkerService/pkg/Cache"
	Ijsonparser "FilterWorkerService/pkg/jsonParser"
	"FilterWorkerService/pkg/logger"
	"strconv"
)

type CacheManager struct {
	Cache *Cache.ICache
	ITypeStandardizationDal *abstract.ITypeStandardizationDal
	IJsonParser *Ijsonparser.IJsonParser
	ILog *logger.ILog
}

func CacheManagerConstructor() *CacheManager {
	return &CacheManager{Cache: &IoC.RedisCache,
		ITypeStandardizationDal: &IoC.TypeStandardizationDal,
		IJsonParser: &IoC.JsonParser,
		ILog: &IoC.Logger}
}

func (c *CacheManager) ManageCache(tableName string, key string) (v int64, s bool, m string){

	// cache sorgula,
	value, err := (*c.Cache).Get(key)
	if err != nil && err.Error() == "null data error"{

		//Bilgi yok ise veri tabanına sor
		m, getErr := (*c.ITypeStandardizationDal).GetByKey(tableName, key)
		if getErr != nil{

			(*c.ILog).SendErrorLog("CacheManager",
				"ManageCache",
				"ITypeStandardizationDal_GetByKey", getErr.Error())
			return  int64(0), false, getErr.Error()
		}

		//bu bilgi var ise veriyi dön ve cache'i güncelle,
		if m!= nil{
			_, err := (*c.Cache).Set(m.Key, m.Value, 10 )
			if err != nil {
				(*c.ILog).SendErrorLog("CacheManager",
					"ManageCache",
					"Cache_Set", err.Error())
			}
			return m.Value, true, ""
		}
		//bilgi yok ise yenisini yarat ve cache'i güncelle
		var max, maxErr = (*c.ITypeStandardizationDal).GetMaxByValue(tableName)
		if maxErr != nil {
			(*c.ILog).SendErrorLog("CacheManager",
				"ManageCache",
				"ITypeStandardizationDal_GetMaxByValue", maxErr.Error())
			return int64(0), false, maxErr.Error()
		}

		if err := (*c.ITypeStandardizationDal).Add(tableName, &model.TypeStandardizationModel{
			Key:   key,
			Value: max + 1,
		}) ; err != nil {
			(*c.ILog).SendErrorLog("CacheManager",
				"ManageCache",
				"ITypeStandardizationDal_Add", err.Error())
			return int64(0), false, err.Error()
		}
		return max + 1, true, ""
	}

	if err != nil {
		(*c.ILog).SendErrorLog("CacheManager",
			"ManageCache",
			err.Error())
		return 0, false, err.Error()
	}
	//Bu bilgi var ise dön,
	i, logErr := strconv.Atoi(value)
	if logErr != nil {
		(*c.ILog).SendErrorLog("CacheManager",
			"ManageCache",
			"strconv", err.Error())

		return int64(0), false, logErr.Error()
	}
	return int64(i), true, ""
}

