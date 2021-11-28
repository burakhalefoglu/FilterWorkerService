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

type mdbDScreenSwipeDal struct {
	Client *mongo.Client
}

func MdbDScreenSwipeDalConstructor() *mdbDScreenSwipeDal {
	return &mdbDScreenSwipeDal{Client: mongodb.GetMongodbClient()}
}


func (m *mdbDScreenSwipeDal) Add(data *model.ScreenSwipeRespondModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := m.Client.Database("MLDatabase").Collection("ScreenSwipeModel")
	var _, err = collection.InsertOne(ctx, bson.D{
		{"ClientId", data.ClientId},
		{"ProjectId", data.ProjectId},
		{"CustomerId", data.CustomerId},
		{"LevelIndex", data.LevelIndex},
		{"TotalSwipeSessionCount", data.TotalSwipeSessionCount},
		{"TotalSwipeHour", data.TotalSwipeHour},
		{"FirstSwipeYearOfDay", data.FirstSwipeYearOfDay},
		{"FirstSwipeYear", data.FirstSwipeYear},
		{"FirstSwipeHour", data.FirstSwipeHour},
		{"FirstSwipeWeekDay", data.FirstSwipeWeekDay},
		{"FirstSwipeMinute", data.FirstSwipeMinute},
		{"FistSwipeDirection", data.FistSwipeDirection},
		{"FirstSwipeStartXCor", data.FirstSwipeStartXCor},
		{"FirstSwipeStartYCor", data.FirstSwipeStartYCor},
		{"FirstSwipeFinishXCor", data.FirstSwipeFinishXCor},
		{"FirstSwipeFinishYCor", data.FirstSwipeFinishYCor},
		{"SecondSwipeDirection", data.SecondSwipeDirection},
		{"SecondSwipeStartXCor", data.SecondSwipeStartXCor},
		{"SecondSwipeStartYCor", data.SecondSwipeStartYCor},
		{"SecondSwipeFinishXCor", data.SecondSwipeFinishXCor},
		{"SecondSwipeFinishYCor", data.SecondSwipeFinishYCor},
		{"ThirdSwipeDirection", data.ThirdSwipeDirection},
		{"ThirdSwipeStartXCor", data.ThirdSwipeStartXCor},
		{"ThirdSwipeStartYCor", data.ThirdSwipeStartYCor},
		{"ThirdSwipeFinishXCor", data.ThirdSwipeFinishXCor},
		{"ThirdSwipeFinishYCor", data.ThirdSwipeFinishYCor},
		{"FourthSwipeDirection", data.FourthSwipeDirection},
		{"FourthSwipeStartXCor", data.FourthSwipeStartXCor},
		{"FourthSwipeStartYCor", data.FourthSwipeStartYCor},
		{"FourthSwipeFinishXCor", data.FourthSwipeFinishXCor},
		{"FourthSwipeFinishYCor", data.FourthSwipeFinishYCor},
		{"FifthSwipeDirection", data.FifthSwipeDirection},
		{"FifthSwipeStartXCor", data.FifthSwipeStartXCor},
		{"FifthSwipeStartYCor", data.FifthSwipeStartYCor},
		{"FifthSwipeFinishXCor", data.FifthSwipeFinishXCor},
		{"FifthSwipeFinishYCor", data.FifthSwipeFinishYCor},
		{"PenultimateSwipeDirection", data.PenultimateSwipeDirection},
		{"PenultimateSwipeStartXCor", data.PenultimateSwipeStartXCor},
		{"PenultimateSwipeStartYCor", data.PenultimateSwipeStartYCor},
		{"PenultimateSwipeFinishXCor", data.PenultimateSwipeFinishXCor},
		{"PenultimateSwipeFinishYCor", data.PenultimateSwipeFinishYCor},
		{"PenultimateSwipeYearOfDay", data.PenultimateSwipeYearOfDay},
		{"PenultimateSwipeYear", data.PenultimateSwipeYear},
		{"PenultimateSwipeHour", data.PenultimateSwipeHour},
		{"PenultimateSwipeWeekDay", data.PenultimateSwipeWeekDay},
		{"PenultimateSwipeMinute", data.PenultimateSwipeMinute},
		{"LastSwipeDirection", data.LastSwipeDirection},
		{"LastSwipeStartXCor", data.LastSwipeStartXCor},
		{"LastSwipeStartYCor", data.LastSwipeStartYCor},
		{"LastSwipeFinishXCor", data.LastSwipeFinishXCor},
		{"LastSwipeFinishYCor", data.LastSwipeFinishYCor},
		{"LastSwipeYearOfDay", data.LastSwipeYearOfDay},
		{"LastSwipeYear", data.LastSwipeYear},
		{"LastSwipeHour", data.LastSwipeHour},
		{"LastSwipeWeekDay", data.LastSwipeWeekDay},
		{"LastSwipeMinute", data.LastSwipeMinute},

		{"FirstDayTotalSwipeUpCount", data.FirstDayTotalSwipeUpCount},
		{"FirstDayTotalSwipeDownCount", data.FirstDayTotalSwipeDownCount},
		{"FirstDayTotalSwipeRightCount", data.FirstDayTotalSwipeRightCount},
		{"FirstDayTotalSwipeLeftCount", data.FirstDayTotalSwipeLeftCount},
		{"FirstDaySwipeTotalStartXCor", data.FirstDaySwipeTotalStartXCor},
		{"FirstDaySwipeTotalStartYCor", data.FirstDaySwipeTotalStartYCor},
		{"FirstDaySwipeTotalFinishXCor", data.FirstDaySwipeTotalFinishXCor},
		{"FirstDaySwipeTotalFinishYCor", data.FirstDaySwipeTotalFinishYCor},

		{"SecondDayTotalSwipeUpCount", data.SecondDayTotalSwipeUpCount},
		{"SecondDayTotalSwipeDownCount", data.SecondDayTotalSwipeDownCount},
		{"SecondDayTotalSwipeRightCount", data.SecondDayTotalSwipeRightCount},
		{"SecondDayTotalSwipeLeftCount", data.SecondDayTotalSwipeLeftCount},
		{"SecondDaySwipeTotalStartXCor", data.SecondDaySwipeTotalStartXCor},
		{"SecondDaySwipeTotalStartYCor", data.SecondDaySwipeTotalStartYCor},
		{"SecondDaySwipeTotalFinishXCor", data.SecondDaySwipeTotalFinishXCor},
		{"SecondDaySwipeTotalFinishYCor", data.SecondDaySwipeTotalFinishYCor},

		{"ThirdDayTotalSwipeUpCount", data.ThirdDayTotalSwipeUpCount},
		{"ThirdDayTotalSwipeDownCount", data.ThirdDayTotalSwipeDownCount},
		{"ThirdDayTotalSwipeRightCount", data.ThirdDayTotalSwipeRightCount},
		{"ThirdDayTotalSwipeLeftCount", data.ThirdDayTotalSwipeLeftCount},
		{"ThirdDaySwipeTotalStartXCor", data.ThirdDaySwipeTotalStartXCor},
		{"ThirdDaySwipeTotalStartYCor", data.ThirdDaySwipeTotalStartYCor},
		{"ThirdDaySwipeTotalFinishXCor", data.ThirdDaySwipeTotalFinishXCor},
		{"ThirdDaySwipeTotalFinishYCor", data.ThirdDaySwipeTotalFinishYCor},

		{"FourthDayTotalSwipeUpCount", data.FourthDayTotalSwipeUpCount},
		{"FourthDayTotalSwipeDownCount", data.FourthDayTotalSwipeDownCount},
		{"FourthDayTotalSwipeRightCount", data.FourthDayTotalSwipeRightCount},
		{"FourthDayTotalSwipeLeftCount", data.FourthDayTotalSwipeLeftCount},
		{"FourthDaySwipeTotalStartXCor", data.FourthDaySwipeTotalStartXCor},
		{"FourthDaySwipeTotalStartYCor", data.FourthDaySwipeTotalStartYCor},
		{"FourthDaySwipeTotalFinishXCor", data.FourthDaySwipeTotalFinishXCor},
		{"FourthDaySwipeTotalFinishYCor", data.FourthDaySwipeTotalFinishYCor},

		{"FifthDayTotalSwipeUpCount", data.FifthDayTotalSwipeUpCount},
		{"FifthDayTotalSwipeDownCount", data.FifthDayTotalSwipeDownCount},
		{"FifthDayTotalSwipeRightCount", data.FifthDayTotalSwipeRightCount},
		{"FifthDayTotalSwipeLeftCount", data.FifthDayTotalSwipeLeftCount},
		{"FifthDaySwipeTotalStartXCor", data.FifthDaySwipeTotalStartXCor},
		{"FifthDaySwipeTotalStartYCor", data.FifthDaySwipeTotalStartYCor},
		{"FifthDaySwipeTotalFinishXCor", data.FifthDaySwipeTotalFinishXCor},
		{"FifthDaySwipeTotalFinishYCor", data.FifthDaySwipeTotalFinishYCor},

		{"SixthDayTotalSwipeUpCount", data.SixthDayTotalSwipeUpCount},
		{"SixthDayTotalSwipeDownCount", data.SixthDayTotalSwipeDownCount},
		{"SixthDayTotalSwipeRightCount", data.SixthDayTotalSwipeRightCount},
		{"SixthDayTotalSwipeLeftCount", data.SixthDayTotalSwipeLeftCount},
		{"SixthDaySwipeTotalStartXCor", data.SixthDaySwipeTotalStartXCor},
		{"SixthDaySwipeTotalStartYCor", data.SixthDaySwipeTotalStartYCor},
		{"SixthDaySwipeTotalFinishXCor", data.SixthDaySwipeTotalFinishXCor},
		{"SixthDaySwipeTotalFinishYCor", data.SixthDaySwipeTotalFinishYCor},

		{"SeventhDayTotalSwipeUpCount", data.SeventhDayTotalSwipeUpCount},
		{"SeventhDayTotalSwipeDownCount", data.SeventhDayTotalSwipeDownCount},
		{"SeventhDayTotalSwipeRightCount", data.SeventhDayTotalSwipeRightCount},
		{"SeventhDayTotalSwipeLeftCount", data.SeventhDayTotalSwipeLeftCount},
		{"SeventhDaySwipeTotalStartXCor", data.SeventhDaySwipeTotalStartXCor},
		{"SeventhDaySwipeTotalStartYCor", data.SeventhDaySwipeTotalStartYCor},
		{"SeventhDaySwipeTotalFinishXCor", data.SeventhDaySwipeTotalFinishXCor},
		{"SeventhDaySwipeTotalFinishYCor", data.SeventhDaySwipeTotalFinishYCor},

		{"TotalSwipeUpCount", data.TotalSwipeUpCount},
		{"TotalSwipeDownCount", data.TotalSwipeDownCount},
		{"TotalSwipeRightCount", data.TotalSwipeRightCount},
		{"TotalSwipeLeftCount", data.TotalSwipeLeftCount},
		{"TotalSwipeStartXCor", data.TotalSwipeStartXCor},
		{"TotalSwipeStartYCor", data.TotalSwipeStartYCor},
		{"TotalSwipeFinishXCor", data.TotalSwipeFinishXCor},
		{"TotalSwipeFinishYCor", data.TotalSwipeFinishYCor},
	})
	if err != nil {
		return err
	}
	return nil
}

func (m *mdbDScreenSwipeDal) GetScreenSwipeById(ClientId string) (*model.ScreenSwipeRespondModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := m.Client.Database("MLDatabase").Collection("ScreenSwipeModel")
	var result = collection.FindOne(ctx, bson.D{{
		"ClientId", ClientId,
	}})
	var model = model.ScreenSwipeRespondModel{}
	if result.Err() != nil && result.Err().Error() == "mongo: no documents in result"{
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

func (m *mdbDScreenSwipeDal) UpdateScreenSwipeById(ClientId string, data *model.ScreenSwipeRespondModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	update := bson.D{{"$set", bson.D{
		{"ClientId", data.ClientId},
		{"ProjectId", data.ProjectId},
		{"CustomerId", data.CustomerId},
		{"LevelIndex", data.LevelIndex},
		{"TotalSwipeSessionCount", data.TotalSwipeSessionCount},
		{"TotalSwipeHour", data.TotalSwipeHour},
		{"FirstSwipeYearOfDay", data.FirstSwipeYearOfDay},
		{"FirstSwipeYear", data.FirstSwipeYear},
		{"FirstSwipeHour", data.FirstSwipeHour},
		{"FirstSwipeWeekDay", data.FirstSwipeWeekDay},
		{"FirstSwipeMinute", data.FirstSwipeMinute},
		{"FistSwipeDirection", data.FistSwipeDirection},
		{"FirstSwipeStartXCor", data.FirstSwipeStartXCor},
		{"FirstSwipeStartYCor", data.FirstSwipeStartYCor},
		{"FirstSwipeFinishXCor", data.FirstSwipeFinishXCor},
		{"FirstSwipeFinishYCor", data.FirstSwipeFinishYCor},
		{"SecondSwipeDirection", data.SecondSwipeDirection},
		{"SecondSwipeStartXCor", data.SecondSwipeStartXCor},
		{"SecondSwipeStartYCor", data.SecondSwipeStartYCor},
		{"SecondSwipeFinishXCor", data.SecondSwipeFinishXCor},
		{"SecondSwipeFinishYCor", data.SecondSwipeFinishYCor},
		{"ThirdSwipeDirection", data.ThirdSwipeDirection},
		{"ThirdSwipeStartXCor", data.ThirdSwipeStartXCor},
		{"ThirdSwipeStartYCor", data.ThirdSwipeStartYCor},
		{"ThirdSwipeFinishXCor", data.ThirdSwipeFinishXCor},
		{"ThirdSwipeFinishYCor", data.ThirdSwipeFinishYCor},
		{"FourthSwipeDirection", data.FourthSwipeDirection},
		{"FourthSwipeStartXCor", data.FourthSwipeStartXCor},
		{"FourthSwipeStartYCor", data.FourthSwipeStartYCor},
		{"FourthSwipeFinishXCor", data.FourthSwipeFinishXCor},
		{"FourthSwipeFinishYCor", data.FourthSwipeFinishYCor},
		{"FifthSwipeDirection", data.FifthSwipeDirection},
		{"FifthSwipeStartXCor", data.FifthSwipeStartXCor},
		{"FifthSwipeStartYCor", data.FifthSwipeStartYCor},
		{"FifthSwipeFinishXCor", data.FifthSwipeFinishXCor},
		{"FifthSwipeFinishYCor", data.FifthSwipeFinishYCor},
		{"PenultimateSwipeDirection", data.PenultimateSwipeDirection},
		{"PenultimateSwipeStartXCor", data.PenultimateSwipeStartXCor},
		{"PenultimateSwipeStartYCor", data.PenultimateSwipeStartYCor},
		{"PenultimateSwipeFinishXCor", data.PenultimateSwipeFinishXCor},
		{"PenultimateSwipeFinishYCor", data.PenultimateSwipeFinishYCor},
		{"PenultimateSwipeYearOfDay", data.PenultimateSwipeYearOfDay},
		{"PenultimateSwipeYear", data.PenultimateSwipeYear},
		{"PenultimateSwipeHour", data.PenultimateSwipeHour},
		{"PenultimateSwipeWeekDay", data.PenultimateSwipeWeekDay},
		{"PenultimateSwipeMinute", data.PenultimateSwipeMinute},
		{"LastSwipeDirection", data.LastSwipeDirection},
		{"LastSwipeStartXCor", data.LastSwipeStartXCor},
		{"LastSwipeStartYCor", data.LastSwipeStartYCor},
		{"LastSwipeFinishXCor", data.LastSwipeFinishXCor},
		{"LastSwipeFinishYCor", data.LastSwipeFinishYCor},
		{"LastSwipeYearOfDay", data.LastSwipeYearOfDay},
		{"LastSwipeYear", data.LastSwipeYear},
		{"LastSwipeHour", data.LastSwipeHour},
		{"LastSwipeWeekDay", data.LastSwipeWeekDay},
		{"LastSwipeMinute", data.LastSwipeMinute},

		{"FirstDayTotalSwipeUpCount", data.FirstDayTotalSwipeUpCount},
		{"FirstDayTotalSwipeDownCount", data.FirstDayTotalSwipeDownCount},
		{"FirstDayTotalSwipeRightCount", data.FirstDayTotalSwipeRightCount},
		{"FirstDayTotalSwipeLeftCount", data.FirstDayTotalSwipeLeftCount},
		{"FirstDaySwipeTotalStartXCor", data.FirstDaySwipeTotalStartXCor},
		{"FirstDaySwipeTotalStartYCor", data.FirstDaySwipeTotalStartYCor},
		{"FirstDaySwipeTotalFinishXCor", data.FirstDaySwipeTotalFinishXCor},
		{"FirstDaySwipeTotalFinishYCor", data.FirstDaySwipeTotalFinishYCor},

		{"SecondDayTotalSwipeUpCount", data.SecondDayTotalSwipeUpCount},
		{"SecondDayTotalSwipeDownCount", data.SecondDayTotalSwipeDownCount},
		{"SecondDayTotalSwipeRightCount", data.SecondDayTotalSwipeRightCount},
		{"SecondDayTotalSwipeLeftCount", data.SecondDayTotalSwipeLeftCount},
		{"SecondDaySwipeTotalStartXCor", data.SecondDaySwipeTotalStartXCor},
		{"SecondDaySwipeTotalStartYCor", data.SecondDaySwipeTotalStartYCor},
		{"SecondDaySwipeTotalFinishXCor", data.SecondDaySwipeTotalFinishXCor},
		{"SecondDaySwipeTotalFinishYCor", data.SecondDaySwipeTotalFinishYCor},

		{"ThirdDayTotalSwipeUpCount", data.ThirdDayTotalSwipeUpCount},
		{"ThirdDayTotalSwipeDownCount", data.ThirdDayTotalSwipeDownCount},
		{"ThirdDayTotalSwipeRightCount", data.ThirdDayTotalSwipeRightCount},
		{"ThirdDayTotalSwipeLeftCount", data.ThirdDayTotalSwipeLeftCount},
		{"ThirdDaySwipeTotalStartXCor", data.ThirdDaySwipeTotalStartXCor},
		{"ThirdDaySwipeTotalStartYCor", data.ThirdDaySwipeTotalStartYCor},
		{"ThirdDaySwipeTotalFinishXCor", data.ThirdDaySwipeTotalFinishXCor},
		{"ThirdDaySwipeTotalFinishYCor", data.ThirdDaySwipeTotalFinishYCor},

		{"FourthDayTotalSwipeUpCount", data.FourthDayTotalSwipeUpCount},
		{"FourthDayTotalSwipeDownCount", data.FourthDayTotalSwipeDownCount},
		{"FourthDayTotalSwipeRightCount", data.FourthDayTotalSwipeRightCount},
		{"FourthDayTotalSwipeLeftCount", data.FourthDayTotalSwipeLeftCount},
		{"FourthDaySwipeTotalStartXCor", data.FourthDaySwipeTotalStartXCor},
		{"FourthDaySwipeTotalStartYCor", data.FourthDaySwipeTotalStartYCor},
		{"FourthDaySwipeTotalFinishXCor", data.FourthDaySwipeTotalFinishXCor},
		{"FourthDaySwipeTotalFinishYCor", data.FourthDaySwipeTotalFinishYCor},

		{"FifthDayTotalSwipeUpCount", data.FifthDayTotalSwipeUpCount},
		{"FifthDayTotalSwipeDownCount", data.FifthDayTotalSwipeDownCount},
		{"FifthDayTotalSwipeRightCount", data.FifthDayTotalSwipeRightCount},
		{"FifthDayTotalSwipeLeftCount", data.FifthDayTotalSwipeLeftCount},
		{"FifthDaySwipeTotalStartXCor", data.FifthDaySwipeTotalStartXCor},
		{"FifthDaySwipeTotalStartYCor", data.FifthDaySwipeTotalStartYCor},
		{"FifthDaySwipeTotalFinishXCor", data.FifthDaySwipeTotalFinishXCor},
		{"FifthDaySwipeTotalFinishYCor", data.FifthDaySwipeTotalFinishYCor},

		{"SixthDayTotalSwipeUpCount", data.SixthDayTotalSwipeUpCount},
		{"SixthDayTotalSwipeDownCount", data.SixthDayTotalSwipeDownCount},
		{"SixthDayTotalSwipeRightCount", data.SixthDayTotalSwipeRightCount},
		{"SixthDayTotalSwipeLeftCount", data.SixthDayTotalSwipeLeftCount},
		{"SixthDaySwipeTotalStartXCor", data.SixthDaySwipeTotalStartXCor},
		{"SixthDaySwipeTotalStartYCor", data.SixthDaySwipeTotalStartYCor},
		{"SixthDaySwipeTotalFinishXCor", data.SixthDaySwipeTotalFinishXCor},
		{"SixthDaySwipeTotalFinishYCor", data.SixthDaySwipeTotalFinishYCor},

		{"SeventhDayTotalSwipeUpCount", data.SeventhDayTotalSwipeUpCount},
		{"SeventhDayTotalSwipeDownCount", data.SeventhDayTotalSwipeDownCount},
		{"SeventhDayTotalSwipeRightCount", data.SeventhDayTotalSwipeRightCount},
		{"SeventhDayTotalSwipeLeftCount", data.SeventhDayTotalSwipeLeftCount},
		{"SeventhDaySwipeTotalStartXCor", data.SeventhDaySwipeTotalStartXCor},
		{"SeventhDaySwipeTotalStartYCor", data.SeventhDaySwipeTotalStartYCor},
		{"SeventhDaySwipeTotalFinishXCor", data.SeventhDaySwipeTotalFinishXCor},
		{"SeventhDaySwipeTotalFinishYCor", data.SeventhDaySwipeTotalFinishYCor},

		{"TotalSwipeUpCount", data.TotalSwipeUpCount},
		{"TotalSwipeDownCount", data.TotalSwipeDownCount},
		{"TotalSwipeRightCount", data.TotalSwipeRightCount},
		{"TotalSwipeLeftCount", data.TotalSwipeLeftCount},
		{"TotalSwipeStartXCor", data.TotalSwipeStartXCor},
		{"TotalSwipeStartYCor", data.TotalSwipeStartYCor},
		{"TotalSwipeFinishXCor", data.TotalSwipeFinishXCor},
		{"TotalSwipeFinishYCor", data.TotalSwipeFinishYCor},
	}}}
	collection := m.Client.Database("MLDatabase").Collection("ScreenSwipeModel")
	updateResult := collection.FindOneAndUpdate(ctx, bson.D{{
		"ClientId", ClientId,
	}}, update)
	if updateResult.Err() != nil {
		return updateResult.Err()
	}
	return nil
}
