package mongodb_driver

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/pkg/database/mongodb"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mdbDHardwareInformationDal struct {
	Client *mongo.Client
}

func MdbDHardwareInformationDalConstructor() *mdbDHardwareInformationDal {
	return &mdbDHardwareInformationDal{Client: mongodb.GetMongodbClient()}
}

func (m *mdbDHardwareInformationDal) Add(data *model.HardwareInformationResponseModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := m.Client.Database("MLDatabase").Collection("HardwareInformationModel")
	var _, err = collection.InsertOne(ctx, bson.D{
		{"ClientId", data.ClientId},
		{"ProjectId", data.ProjectId},
		{"CustomerId", data.CustomerId},
		{"DeviceType", data.DeviceType},
		{"GraphicsDeviceType", data.GraphicsDeviceType},
		{"GraphicsMemorySize", data.GraphicsMemorySize},
		{"OperatingSystem", data.OperatingSystem},
		{"ProcessorCount", data.ProcessorCount},
		{"ProcessorType", data.ProcessorType},
		{"SystemMemorySize", data.SystemMemorySize},
	})
	if err != nil {
		return err
	}
	return nil
}
