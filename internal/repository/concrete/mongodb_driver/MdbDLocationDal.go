package mongodb_driver

import (
	"FilterWorkerService/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MdbLocationDal struct {
	Client *mongo.Client
}


func (m *MdbLocationDal) Add(data *model.LocationResponseModel) error{

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := m.Client.Database("MLDatabase").Collection("LocationModel")
	var _, err = collection.InsertOne(ctx, bson.D{
		{"ClientId", data.ClientId},
		{"ProjectId",data.ProjectId},
		{"CustomerId",data.CustomerId},
		{"City", data.City},
		{"Country", data.Country},
		{"Org", data.Org},
		{"Region", data.Region},
		{"Continent", data.Continent},

	})
	if err != nil {
		return err
	}
	return nil

}