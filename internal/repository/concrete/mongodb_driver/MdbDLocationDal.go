package mongodb_driver

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/pkg/database/mongodb"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mdbLocationDal struct {
	Client *mongo.Client
}

func MdbDLocationDalConstructor() *mdbLocationDal {
	return &mdbLocationDal{Client: mongodb.GetMongodbClient()}
}

func (m *mdbLocationDal) Add(data *model.LocationResponseModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := m.Client.Database("MLDatabase").Collection("locationModels")
	var _, err = collection.InsertOne(ctx, bson.D{
		{"ClientId", data.ClientId},
		{"ProjectId", data.ProjectId},
		{"CustomerId", data.CustomerId},
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
