package mongodb_driver

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/pkg/database/mongodb"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mdbDLevelBaseSessionDal struct {
	Client *mongo.Client
}

func MdbDLevelBaseSessionDalConstructor() *mdbDLevelBaseSessionDal {
	return &mdbDLevelBaseSessionDal{Client: mongodb.GetMongodbClient()}
}

func (m *mdbDLevelBaseSessionDal) Add(data *model.LevelBaseSessionRespondModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("MLDatabase").Collection("levelBaseSessions")
	var _, err = collection.InsertOne(ctx, bson.D{
		{"ClientId", data.ClientId},
		{"ProjectId", data.ProjectId},
		{"CustomerId", data.CustomerId},

		{"TotalLevelBaseSessionMinute", data.TotalLevelBaseSessionMinute},
		{"TotalLevelBaseSessionCount", data.TotalLevelBaseSessionCount},

		{"FirstLevelSessionYearOfDay", data.FirstLevelSessionYearOfDay},
		{"FirstLevelSessionYear", data.FirstLevelSessionYear},
		{"FirstLevelSessionWeekDay", data.FirstLevelSessionWeekDay},
		{"FirstLevelSessionHour", data.FirstLevelSessionHour},
		{"FirstLevelSessionMinute", data.FirstLevelSessionMinute},

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

		{"FirstQuarterHourTotalLevelBaseSessionCount", data.FirstQuarterHourTotalLevelBaseSessionCount},
		{"FirstHalfHourTotalLEvelBaseSessionCount", data.FirstHalfHourTotalLevelBaseSessionCount},
		{"FirstHourTotalLevelBaseSessionCount", data.FirstHourTotalLevelBaseSessionCount},
		{"FirstTwoHourTotalLevelBaseSessionCount", data.FirstTwoHourTotalLevelBaseSessionCount},
		{"FirstThreeHourTotalLevelBaseSessionCount", data.FirstThreeHourTotalLevelBaseSessionCount},
		{"FirstSixHourTotalLevelBaseSessionCount", data.FirstSixHourTotalLevelBaseSessionCount},
		{"FirstTwelveHourTotalLevelBaseSessionCount", data.FirstTwelveHourTotalLevelBaseSessionCount},
		{"FirstDayTotalLevelBaseSessionCount", data.FirstDayTotalLevelBaseSessionCount},

		{"PenultimateLevelSessionLevelIndex", data.PenultimateLevelSessionLevelIndex},
		{"PenultimateLevelSessionLevelDuration", data.PenultimateLevelSessionLevelDuration},
		{"LastLevelSessionLevelIndex", data.LastLevelSessionLevelIndex},
		{"LastLevelSessionLevelDuration", data.LastLevelSessionLevelDuration},

		{"LastLevelSessionYearOfDay", data.LastLevelSessionYearOfDay},
		{"LastLevelSessionYear", data.LastLevelSessionYear},
		{"LastLevelSessionWeekDay", data.LastLevelSessionWeekDay},
		{"LastLevelSessionHour", data.LastLevelSessionHour},
		{"LastLevelSessionMinute", data.LastLevelSessionMinute},
	})
	if err != nil {
		return err
	}
	return nil
}

func (m *mdbDLevelBaseSessionDal) GetLevelBaseSessionById(ClientId string) (*model.LevelBaseSessionRespondModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("MLDatabase").Collection("LevelBaseSession")
	var result = collection.FindOne(ctx, bson.D{{
		"ClientId", ClientId,
	}})

	var model = model.LevelBaseSessionRespondModel{}
	if result.Err() != nil && result.Err().Error() == "mongo: no documents in result" {
		return &model, errors.New("null data error")
	}
	if result.Err() != nil && result.Err().Error() != "mongo: no documents in result" {
		return &model, result.Err()
	}
	var err = result.Decode(&model)
	if err != nil {
		return &model, err
	}
	return &model, nil
}

func (m *mdbDLevelBaseSessionDal) UpdateLevelBaseSessionById(ClientId string, data *model.LevelBaseSessionRespondModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	update := bson.D{{"$set", bson.D{
		{"ClientId", data.ClientId},
		{"ProjectId", data.ProjectId},
		{"CustomerId", data.CustomerId},

		{"TotalLevelBaseSessionMinute", data.TotalLevelBaseSessionMinute},
		{"TotalLevelBaseSessionCount", data.TotalLevelBaseSessionCount},

		{"FirstLevelSessionYearOfDay", data.FirstLevelSessionYearOfDay},
		{"FirstLevelSessionYear", data.FirstLevelSessionYear},
		{"FirstLevelSessionWeekDay", data.FirstLevelSessionWeekDay},
		{"FirstLevelSessionHour", data.FirstLevelSessionHour},
		{"FirstLevelSessionMinute", data.FirstLevelSessionMinute},

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

		{"FirstQuarterHourTotalLevelBaseSessionCount", data.FirstQuarterHourTotalLevelBaseSessionCount},
		{"FirstHalfHourTotalLEvelBaseSessionCount", data.FirstHalfHourTotalLevelBaseSessionCount},
		{"FirstHourTotalLevelBaseSessionCount", data.FirstHourTotalLevelBaseSessionCount},
		{"FirstTwoHourTotalLevelBaseSessionCount", data.FirstTwoHourTotalLevelBaseSessionCount},
		{"FirstThreeHourTotalLevelBaseSessionCount", data.FirstThreeHourTotalLevelBaseSessionCount},
		{"FirstSixHourTotalLevelBaseSessionCount", data.FirstSixHourTotalLevelBaseSessionCount},
		{"FirstTwelveHourTotalLevelBaseSessionCount", data.FirstTwelveHourTotalLevelBaseSessionCount},
		{"FirstDayTotalLevelBaseSessionCount", data.FirstDayTotalLevelBaseSessionCount},

		{"PenultimateLevelSessionLevelIndex", data.PenultimateLevelSessionLevelIndex},
		{"PenultimateLevelSessionLevelDuration", data.PenultimateLevelSessionLevelDuration},
		{"LastLevelSessionLevelIndex", data.LastLevelSessionLevelIndex},
		{"LastLevelSessionLevelDuration", data.LastLevelSessionLevelDuration},

		{"LastLevelSessionYearOfDay", data.LastLevelSessionYearOfDay},
		{"LastLevelSessionYear", data.LastLevelSessionYear},
		{"LastLevelSessionWeekDay", data.LastLevelSessionWeekDay},
		{"LastLevelSessionHour", data.LastLevelSessionHour},
		{"LastLevelSessionMinute", data.LastLevelSessionMinute},
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
