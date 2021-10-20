package locationManager

import (
	"FilterWorkerService/internal/jsonParser"
	baseManager "FilterWorkerService/internal/manager/base"
	"log"
)

type Location struct{
	Continent string 
}

var loc = Location{}

func AddOrUpdateContinent( topic string, message []byte) {

	jsonParser.DecodeJson(message,loc)

	_, err := baseManager.ManageCache(topic,loc.Continent)

	if(err != nil) {
		log.Fatal("err: ", err)
	}

	//Todo: Bu id'yi kullanarak kullanının bilgilerini güncelle!

}