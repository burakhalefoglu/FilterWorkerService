package locationManager

import (
	"FilterWorkerService/internal/dataAccess/mongodb"
	baseManager "FilterWorkerService/internal/manager/base"
	"go.mongodb.org/mongo-driver/bson"
)

type Location struct{
	ProjectId string
	ClientId string
	Continent string
}

type Client struct{
	ProjectId string
	ClientId string
	Continent int
}

func AddOrUpdateParameter( topic string, collection string, key string, loc *Location) error {

	value, err := baseManager.ManageCache(collection,loc.Continent)

	if err != nil {
		return err
	}
	resultId, err := mongodb.GetCollectionCount(topic, bson.D{
		{"ClientId",loc.ClientId},
		{ "ProjectId", loc.ClientId},})

	if err != nil {
		return err
	}
	if resultId>0 {
		_, err := mongodb.UpdateCollection(topic, bson.D{
			{"ClientId",loc.ClientId},
		{ "ProjectId", loc.ClientId},},
		bson.D{{"Continent",value}})
		if err != nil{
			return err
		}
		return nil
	}

	_, err = mongodb.AddCollection(topic, Client{ProjectId: loc.ProjectId,
		ClientId:  loc.ClientId,
		Continent: value})
	if err != nil {
		return err
	}
	return  nil
}