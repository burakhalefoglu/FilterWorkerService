package mongodb_driver

import (
	"FilterWorkerService/internal/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MdbDScreenSwipeDal struct {
	Client *mongo.Client
}

func (m *MdbDScreenSwipeDal) Add(data *model.ScreenSwipeRespondModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := m.Client.Database("MLDatabase").Collection("ScreenSwipeModel")
	var _, err = collection.InsertOne(ctx, bson.D{
			{"ClientId", data.ClientId},
		{"ProjectId", data.ProjectId},
		{"CustomerId", data.CustomerId},
		{"LevelIndex", data.LevelIndex},
		{"TotalSwipeSessionCount", data.TotalSwipeSessionCount},
		{"FirstSwipeYearOfDay", data.FirstSwipeYearOfDay},
		{"FirstSwipeYear", data.FirstSwipeYear},
		{"FirstSwipeHour", data.FirstSwipeHour},
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
		{"PenultimateSwipeDirection", data.PenultimateSwipeDirection},
		{"PenultimateSwipeStartXCor", data.PenultimateSwipeStartXCor},
		{"PenultimateSwipeStartYCor", data.PenultimateSwipeStartYCor},
		{"PenultimateSwipeFinishXCor", data.PenultimateSwipeFinishXCor},
		{"PenultimateSwipeFinishYCor", data.PenultimateSwipeFinishYCor},
		{"LastSwipeDirection", data.LastSwipeDirection},
		{"LastSwipeStartXCor", data.LastSwipeStartXCor},
		{"LastSwipeStartYCor", data.LastSwipeStartYCor},
		{"LastSwipeFinishXCor", data.LastSwipeFinishXCor},
		{"LastSwipeFinishYCor", data.LastSwipeFinishYCor},
		{"FirstDayTotalSwipeUpCount", data.FirstDayTotalSwipeUpCount},
		{"FirstDayTotalSwipeDownCount", data.FirstDayTotalSwipeDownCount},
		{"FirstDayTotalSwipeRightCount", data.FirstDayTotalSwipeRightCount},
		{"FirstDayTotalSwipeLeftCount", data.FirstDayTotalSwipeLeftCount},
		{"FirstDaySwipeTotalStartXCor", data.FirstDaySwipeTotalStartXCor},
		{"FirstDaySwipeTotalStartYCor", data.FirstDaySwipeTotalStartYCor},
		{"FirstDaySwipeTotalFinishXCor", data.FirstDaySwipeTotalFinishXCor},
		{"FirstDaySwipeTotalFinishYCor", data.FirstDaySwipeTotalFinishYCor},
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

func (m *MdbDScreenSwipeDal) GetScreenSwipeByCustomerId(CustomerId string) (*model.ScreenSwipeRespondModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := m.Client.Database("MLDatabase").Collection("ScreenSwipeModel")
	var result = collection.FindOne(ctx, bson.D{primitive.E{
		Key: "CustomerId", Value: CustomerId,
	}})
	var model = model.ScreenSwipeRespondModel{}
	if result.Err() != nil {
		return &model, result.Err()
	}
	var err = result.Decode(&model)
	if err != nil {
		return &model, err
	}
	return &model, nil
}

func (m *MdbDScreenSwipeDal) UpdateScreenSwipeByCustomerId(CustomerId string, data *model.ScreenSwipeRespondModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	update := bson.D{{"$set", bson.D{
		{"ClientId", data.ClientId},
		{"ProjectId", data.ProjectId},
		{"CustomerId", data.CustomerId},
		{"LevelIndex", data.LevelIndex},
		{"TotalSwipeSessionCount", data.TotalSwipeSessionCount},
		{"FirstSwipeYearOfDay", data.FirstSwipeYearOfDay},
		{"FirstSwipeYear", data.FirstSwipeYear},
		{"FirstSwipeHour", data.FirstSwipeHour},
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
		{"PenultimateSwipeDirection", data.PenultimateSwipeDirection},
		{"PenultimateSwipeStartXCor", data.PenultimateSwipeStartXCor},
		{"PenultimateSwipeStartYCor", data.PenultimateSwipeStartYCor},
		{"PenultimateSwipeFinishXCor", data.PenultimateSwipeFinishXCor},
		{"PenultimateSwipeFinishYCor", data.PenultimateSwipeFinishYCor},
		{"LastSwipeDirection", data.LastSwipeDirection},
		{"LastSwipeStartXCor", data.LastSwipeStartXCor},
		{"LastSwipeStartYCor", data.LastSwipeStartYCor},
		{"LastSwipeFinishXCor", data.LastSwipeFinishXCor},
		{"LastSwipeFinishYCor", data.LastSwipeFinishYCor},
		{"FirstDayTotalSwipeUpCount", data.FirstDayTotalSwipeUpCount},
		{"FirstDayTotalSwipeDownCount", data.FirstDayTotalSwipeDownCount},
		{"FirstDayTotalSwipeRightCount", data.FirstDayTotalSwipeRightCount},
		{"FirstDayTotalSwipeLeftCount", data.FirstDayTotalSwipeLeftCount},
		{"FirstDaySwipeTotalStartXCor", data.FirstDaySwipeTotalStartXCor},
		{"FirstDaySwipeTotalStartYCor", data.FirstDaySwipeTotalStartYCor},
		{"FirstDaySwipeTotalFinishXCor", data.FirstDaySwipeTotalFinishXCor},
		{"FirstDaySwipeTotalFinishYCor", data.FirstDaySwipeTotalFinishYCor},
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
	updateResult := collection.FindOneAndUpdate(ctx, bson.D{primitive.E{
		Key: "CustomerId", Value: CustomerId,
	}}, update)
	if updateResult.Err() != nil {
		return updateResult.Err()
	}
	return nil
}
