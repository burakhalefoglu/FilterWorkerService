package concrete

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/repository/abstract"
	"FilterWorkerService/pkg/Cache"
	Ijsonparser "FilterWorkerService/pkg/jsonParser"
	"log"
	"strconv"
)

type CacheManager struct {
	Cache Cache.ICache
	ITypeStandardizationDal abstract.ITypeStandardizationDal
	IJsonParser Ijsonparser.IJsonParser
}


func (c *CacheManager) ManageCache(tableName string, key string) (v int64, s bool, m string){

	// cache sorgula,
	value, err := c.Cache.Get(key)
	if err != nil{
		//Bilgi yok ise veri tabanına sor
		m, getErr := c.ITypeStandardizationDal.GetByKey(tableName, key)
		if getErr != nil{
			log.Fatal("err:", getErr)
		}

		//bu bilgi var ise veriyi dön ve cache'i güncelle,
		if m!= nil{
			_, err := c.Cache.Set(m.Key, m.Value, 10 )
			if err != nil {
				log.Print(err)
			}
			return m.Value, true, ""
		}
		//bilgi yok ise yenisini yarat ve cache'i güncelle
		var max, maxErr = c.ITypeStandardizationDal.GetMaxByValue(tableName)
		if maxErr != nil{
			log.Fatal(maxErr)
		}

		err := c.ITypeStandardizationDal.Add(tableName, &model.TypeStandardizationModel{
			Key:   key,
			Value: max + 1,
		})
		if err != nil {
			log.Fatal(err)
		}

		return max + 1, true, ""
	}
	//Bu bilgi var ise dön,
	i, _ := strconv.Atoi(value)
	return int64(i), true, ""

}

