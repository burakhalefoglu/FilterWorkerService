package locationManager

import (
	cacheManager "FilterWorkerService/internal/manager/cache"
	"FilterWorkerService/pkg/jsonParser"
	"log"
)

type Location struct {
	ProjectId string
	ClientId  string
	Continent string
	Country   string
	City      string
	Region    string
	Org       string
}

var loc = Location{}

// type Client struct {
// 	ProjectId string
// 	ClientId  string
// 	Continent int
// 	Country   int
// 	City      int
// 	Region    int
// 	Org       int
// }

func AddOrUpdateContient(topic string, data []byte) error {

	jsonParser.DecodeJson(data, &loc)
	err := addOrUpdateParameter(topic, "Continent", loc.ProjectId, loc.ClientId, loc.Continent)
	if err != nil {
		return err
	}
	return nil
}

func AddOrUpdateCountry(topic string, data []byte) error {

	jsonParser.DecodeJson(data, &loc)
	err := addOrUpdateParameter(topic, "Country", loc.ProjectId, loc.ClientId, loc.Country)
	if err != nil {
		return err
	}
	return nil
}

func AddOrUpdateCity(topic string, data []byte) error {

	jsonParser.DecodeJson(data, &loc)
	err := addOrUpdateParameter(topic, "City", loc.ProjectId, loc.ClientId, loc.City)
	if err != nil {
		return err
	}
	return nil
}

func AddOrUpdateRegion(topic string, data []byte) error {

	jsonParser.DecodeJson(data, &loc)
	err := addOrUpdateParameter(topic, "Region", loc.ProjectId, loc.ClientId, loc.Region)
	if err != nil {
		return err
	}
	return nil
}

func AddOrUpdateOrg(topic string, data []byte) error {

	jsonParser.DecodeJson(data, &loc)
	err := addOrUpdateParameter(topic, "Org", loc.ProjectId, loc.ClientId, loc.Org)
	if err != nil {
		return err
	}
	return nil
}

func addOrUpdateParameter(topic string, collection string, clientId string, projectId string, data string) error {

	log.Printf(clientId, projectId, data)
	_, err := cacheManager.ManageCache(collection, data)

	if err != nil {
		return err
	}
	//resultId, err := mongodb.GetCollectionCount(topic, bson.D{
	//	{"ClientId", clientId},
	//	{"ProjectId", projectId}})
	//
	//if err != nil {
	//	return err
	//}
	//if resultId > 0 {
	//	_, err := mongodb.UpdateCollection(topic, bson.D{
	//		{"ClientId", clientId},
	//		{"ProjectId", projectId}},
	//		bson.D{{collection, value}})
	//	if err != nil {
	//		return err
	//	}
	//	return nil
	//}
	//
	//// Client{ProjectId: projectId,
	//// 		ClientId:  clientId,
	//// 		Continent: value}
	//
	//_, err = mongodb.AddCollection(topic, bson.D{
	//	{"ClientId", clientId},
	//	{"ProjectId", projectId},
	//	{collection, value}})
	//if err != nil {
	//	return err
	//}
	return nil
}
