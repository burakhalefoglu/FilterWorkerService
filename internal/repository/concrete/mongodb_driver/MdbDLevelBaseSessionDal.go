package mongodb_driver

import (
	"FilterWorkerService/internal/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MdbDLevelBaseSessionDal struct {
	Client *mongo.Client
}

func (m *MdbDLevelBaseSessionDal) Add(data *model.LevelBaseSessionRespondModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("MLDatabase").Collection("LevelBaseSession")
	var _, err = collection.InsertOne(ctx, bson.D{
		{"ClientId", data.ClientId},
		{"ProjectId", data.ProjectId},
		{"CustomerId", data.CustomerId},
		{"FirstLevelSessionLevelIndex", data.FirstLevelSessionLevelIndex},
		{"FirstLevelSessionDuration", data.FirstLevelSessionDuration},
		{"SecondLevelSessionLevelIndex", data.SecondLevelSessionLevelIndex},
		{"SecondLevelSessionDuration", data.SecondLevelSessionDuration},
		{"ThirdLevelSessionLevelIndex", data.ThirdLevelSessionLevelIndex},
		{"ThirdLevelSessionDuration", data.ThirdLevelSessionDuration},
		{"FourLevelSessionLevelIndex", data.FourLevelSessionLevelIndex},
		{"FourLevelSessionDuration", data.FourLevelSessionDuration},
		{"FiveLevelSessionLevelIndex", data.FiveLevelSessionLevelIndex},
		{"FiveLevelSessionDuration", data.FiveLevelSessionDuration},
		{"SixLevelSessionLevelIndex", data.SixLevelSessionLevelIndex},
		{"SixLevelSessionDuration", data.SixLevelSessionDuration},
		{"SevenLevelSessionLevelIndex", data.SevenLevelSessionLevelIndex},
		{"SevenLevelSessionDuration", data.SevenLevelSessionDuration},
		{"PenultimateLevelSessionLevelIndex", data.PenultimateLevelSessionLevelIndex},
		{"PenultimateLevelSessionLevelDuration", data.PenultimateLevelSessionLevelDuration},
		{"LastLevelSessionLevelIndex", data.LastLevelSessionLevelIndex},
		{"LastLevelSessionLevelDuration", data.LastLevelSessionLevelDuration},
	})
	if err != nil {
		return err
	}
	return nil
}

func (m *MdbDBuyingEventDal) GetLevelBaseSessionById(ClientId string) (*model.LevelBaseSessionRespondModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("MLDatabase").Collection("LevelBaseSession")
	var result = collection.FindOne(ctx, bson.D{{
		"ClientId", ClientId,
	}})

	var model = model.LevelBaseSessionRespondModel{}
	if result.Err() != nil {
		return &model, result.Err()
	}
	var err = result.Decode(&model)
	if err != nil {
		return &model, err
	}
	return &model, nil
}

func (m *MdbDBuyingEventDal) UpdateLevelBaseSessionById(ClientId string, data *model.LevelBaseSessionRespondModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	update := bson.D{{"$set", bson.D{
		{"ClientId", data.ClientId},
		{"ProjectId", data.ProjectId},
		{"CustomerId", data.CustomerId},
		{"FirstLevelSessionLevelIndex", data.FirstLevelSessionLevelIndex},
		{"FirstLevelSessionDuration", data.FirstLevelSessionDuration},
		{"SecondLevelSessionLevelIndex", data.SecondLevelSessionLevelIndex},
		{"SecondLevelSessionDuration", data.SecondLevelSessionDuration},
		{"ThirdLevelSessionLevelIndex", data.ThirdLevelSessionLevelIndex},
		{"ThirdLevelSessionDuration", data.ThirdLevelSessionDuration},
		{"FourLevelSessionLevelIndex", data.FourLevelSessionLevelIndex},
		{"FourLevelSessionDuration", data.FourLevelSessionDuration},
		{"FiveLevelSessionLevelIndex", data.FiveLevelSessionLevelIndex},
		{"FiveLevelSessionDuration", data.FiveLevelSessionDuration},
		{"SixLevelSessionLevelIndex", data.SixLevelSessionLevelIndex},
		{"SixLevelSessionDuration", data.SixLevelSessionDuration},
		{"SevenLevelSessionLevelIndex", data.SevenLevelSessionLevelIndex},
		{"SevenLevelSessionDuration", data.SevenLevelSessionDuration},
		{"PenultimateLevelSessionLevelIndex", data.PenultimateLevelSessionLevelIndex},
		{"PenultimateLevelSessionLevelDuration", data.PenultimateLevelSessionLevelDuration},
		{"LastLevelSessionLevelIndex", data.LastLevelSessionLevelIndex},
		{"LastLevelSessionLevelDuration", data.LastLevelSessionLevelDuration},
	}}}

	collection := m.Client.Database("MLDatabase").Collection("LevelBaseSession")
	updateResult := collection.FindOneAndUpdate(ctx, bson.D{{
		"ClientId", ClientId,
	}}, update)
	if updateResult.Err() != nil {
		return updateResult.Err()
	}
	return nil
}
