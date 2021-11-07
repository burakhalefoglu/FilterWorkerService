package mongodb_driver

import (
	"FilterWorkerService/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MdbDAdvEventDal struct {
	Client *mongo.Client
}

func (m *MdbDAdvEventDal) Add(data *model.AdvEventRespondModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := m.Client.Database("MLDatabase").Collection("AdvEventModel")
	var _, err := collection.InsertOne(ctx, bson.D{
		{"ClientId",data.ClientId},
		{"ProjectId", data.ProjectId},
		{"CustomerId",data.CustomerId},
		{"LevelIndex", data.LevelIndex},
		{"VideoAdvCount", data.VideoAdvCount},
		{"InterstitialAdvCount",data.InterstitialAdvCount},
		{"VideoClickMonth",data.VideoClickMonth},
		{"VideoClickWeek",data.VideoClickWeek},
		{"VideoClickDay",data.VideoClickDay},
		{"VideoClickHour",data.VideoClickHour},
	})
	if err != nil {
			return err
		}
		return nil
}
