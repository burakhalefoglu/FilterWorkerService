package mongodb_driver

import (
	"FilterWorkerService/internal/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MdbDScreenClickDal struct {
	Client *mongo.Client
}

func (m *MdbDScreenClickDal) Add(data *model.ScreenClickRespondModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := m.Client.Database("MLDatabase").Collection("ScreenClickModel")
	var _, err = collection.InsertOne(ctx, bson.D{
		{"ClientId", data.ClientId},
		{"ProjectId", data.ProjectId},
		{"CustomerId", data.CustomerId},
		{"LevelIndex", data.LevelIndex},
		{"FirstClickSessionYearOfDay", data.FirstClickSessionYearOfDay},
		{"FirstClickSessionYear", data.FirstClickSessionYear},
		{"FirstClickSessionHour", data.FirstClickSessionHour},
		{"FirstClickSessionMinute", data.FirstClickSessionMinute},
		{"FirstTouchCount", data.FirstTouchCount},
		{"SecondClickSessionHour", data.SecondClickSessionHour},
		{"SecondClickSessionMinute", data.SecondClickSessionMinute},
		{"SecondTouchCount", data.SecondTouchCount},
		{"ThirdClickSessionHour", data.ThirdClickSessionHour},
		{"ThirdClickSessionMinute", data.ThirdClickSessionMinute},
		{"ThirdTouchCount", data.ThirdTouchCount},
		{"PenultimateClickSessionHour", data.PenultimateClickSessionHour},
		{"PenultimateClickSessionMinute", data.PenultimateClickSessionMinute},
		{"PenultimateTouchCount", data.PenultimateTouchCount},
		{"LastClickSessionYearOfDay", data.LastClickSessionYearOfDay},
		{"LastClickSessionYear", data.LastClickSessionYear},
		{"LastClickSessionHour", data.LastClickSessionHour},
		{"LastClickSessionMinute", data.LastClickSessionMinute},
		{"LastTouchCount", data.LastTouchCount},		
		{"FirstStartXCor", data.FirstStartXCor},
		{"FirstStartYCor", data.FirstStartYCor},
		{"FirstFinishXCor", data.FirstFinishXCor},
		{"FirstFinishYCor", data.FirstFinishYCor},		
		{"SecondStartXCor", data.SecondStartXCor},
		{"SecondStartYCor", data.SecondStartYCor},
		{"SecondFinishXCor", data.SecondFinishXCor},
		{"SecondFinishYCor", data.SecondFinishYCor},
		{"ThirdStartXCor", data.ThirdStartXCor},
		{"ThirdStartYCor", data.ThirdStartYCor},
		{"ThirdFinishXCor", data.ThirdFinishXCor},
		{"ThirdFinishYCor", data.ThirdFinishYCor},	
		{"PenultimateStartXCor", data.PenultimateStartXCor},
		{"PenultimateStartYCor", data.PenultimateStartYCor},
		{"PenultimateFinishXCor", data.PenultimateFinishXCor},
		{"PenultimateFinishYCor", data.PenultimateFinishYCor},
		{"LastStartXCor", data.LastStartXCor},
		{"LastStartYCor", data.LastStartYCor},
		{"LastFinishXCor", data.LastFinishXCor},
		{"LastFinishYCor", data.LastFinishYCor},	
		{"FirstMinusLastTouchCount", data.FirstMinusLastTouchCount},
		{"FirstFingerId", data.FirstFingerId},
		{"PenultimateFingerId", data.PenultimateFingerId},
		{"LastFingerId", data.LastFingerId},
		{"FirstDayClickCount", data.FirstDayClickCount},
		{"PenultimateDayClickCount", data.PenultimateDayClickCount},
		{"LastDayClickCount", data.LastDayClickCount},
		{"TotalClickDay", data.TotalClickDay},
		{"TotalClickCount", data.TotalClickCount},
		{"TotalClickSessionCount", data.TotalClickSessionCount},
		{"TotalStartXCor", data.TotalStartXCor},
		{"TotalStartYCor", data.TotalStartYCor},
		{"TotalFinishXCor", data.TotalFinishXCor},
		{"TotalFinishYCor", data.TotalFinishYCor},
		{"SessionBasedAvegareStartXCor", data.SessionBasedAvegareStartXCor},
		{"SessionBasedAvegareStartYCor", data.SessionBasedAvegareStartYCor},
		{"SessionBasedAvegareFinishXcor", data.SessionBasedAvegareFinishXCor},
		{"SessionBasedAvegareFinishYCor", data.SessionBasedAvegareFinishYCor},
		{"SessionBasedAvegareClickCount", data.SessionBasedAvegareClickCount},
		{"DailyAvegareClickCount", data.DailyAvegareClickCount},
		{"LastTouchCountMinusSessionBasedAvegareClickCount", data.LastTouchCountMinusSessionBasedAvegareClickCount},
	})
	if err != nil {
		return err
	}
	return nil
}

func (m *MdbDScreenClickDal) GetScreenClickById(ClientId string) (*model.ScreenClickRespondModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := m.Client.Database("MLDatabase").Collection("ScreenClickModel")
	var result = collection.FindOne(ctx, bson.D{{
		"ClientId", ClientId,
	}})
	var model = model.ScreenClickRespondModel{}
	if result.Err() != nil {
		return &model, result.Err()
	}
	var err = result.Decode(&model)
	if err != nil {
		return &model, err
	}
	return &model, nil
}

func (m *MdbDScreenClickDal) UpdateScreenClickById(ClientId string, data *model.ScreenClickRespondModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	update := bson.D{{"$set", bson.D{
		{"ClientId", data.ClientId},
		{"ProjectId", data.ProjectId},
		{"CustomerId", data.CustomerId},
		{"LevelIndex", data.LevelIndex},
		{"FirstClickSessionYearOfDay", data.FirstClickSessionYearOfDay},
		{"FirstClickSessionYear", data.FirstClickSessionYear},
		{"FirstClickSessionHour", data.FirstClickSessionHour},
		{"FirstClickSessionMinute", data.FirstClickSessionMinute},
		{"FirstTouchCount", data.FirstTouchCount},
		{"SecondClickSessionHour", data.SecondClickSessionHour},
		{"SecondClickSessionMinute", data.SecondClickSessionMinute},
		{"SecondTouchCount", data.SecondTouchCount},
		{"ThirdClickSessionHour", data.ThirdClickSessionHour},
		{"ThirdClickSessionMinute", data.ThirdClickSessionMinute},
		{"ThirdTouchCount", data.ThirdTouchCount},
		{"PenultimateClickSessionHour", data.PenultimateClickSessionHour},
		{"PenultimateClickSessionMinute", data.PenultimateClickSessionMinute},
		{"PenultimateTouchCount", data.PenultimateTouchCount},
		{"LastClickSessionYearOfDay", data.LastClickSessionYearOfDay},
		{"LastClickSessionYear", data.LastClickSessionYear},
		{"LastClickSessionHour", data.LastClickSessionHour},
		{"LastClickSessionMinute", data.LastClickSessionMinute},
		{"LastTouchCount", data.LastTouchCount},		
		{"FirstStartXCor", data.FirstStartXCor},
		{"FirstStartYCor", data.FirstStartYCor},
		{"FirstFinishXCor", data.FirstFinishXCor},
		{"FirstFinishYCor", data.FirstFinishYCor},		
		{"SecondStartXCor", data.SecondStartXCor},
		{"SecondStartYCor", data.SecondStartYCor},
		{"SecondFinishXCor", data.SecondFinishXCor},
		{"SecondFinishYCor", data.SecondFinishYCor},
		{"ThirdStartXCor", data.ThirdStartXCor},
		{"ThirdStartYCor", data.ThirdStartYCor},
		{"ThirdFinishXCor", data.ThirdFinishXCor},
		{"ThirdFinishYCor", data.ThirdFinishYCor},	
		{"PenultimateStartXCor", data.PenultimateStartXCor},
		{"PenultimateStartYCor", data.PenultimateStartYCor},
		{"PenultimateFinishXCor", data.PenultimateFinishXCor},
		{"PenultimateFinishYCor", data.PenultimateFinishYCor},
		{"LastStartXCor", data.LastStartXCor},
		{"LastStartYCor", data.LastStartYCor},
		{"LastFinishXCor", data.LastFinishXCor},
		{"LastFinishYCor", data.LastFinishYCor},	
		{"FirstMinusLastTouchCount", data.FirstMinusLastTouchCount},
		{"FirstFingerId", data.FirstFingerId},
		{"PenultimateFingerId", data.PenultimateFingerId},
		{"LastFingerId", data.LastFingerId},
		{"FirstDayClickCount", data.FirstDayClickCount},
		{"PenultimateDayClickCount", data.PenultimateDayClickCount},
		{"LastDayClickCount", data.LastDayClickCount},
		{"TotalClickDay", data.TotalClickDay},
		{"TotalClickCount", data.TotalClickCount},
		{"TotalClickSessionCount", data.TotalClickSessionCount},
		{"TotalStartXCor", data.TotalStartXCor},
		{"TotalStartYCor", data.TotalStartYCor},
		{"TotalFinishXCor", data.TotalFinishXCor},
		{"TotalFinishYCor", data.TotalFinishYCor},
		{"SessionBasedAvegareStartXCor", data.SessionBasedAvegareStartXCor},
		{"SessionBasedAvegareStartYCor", data.SessionBasedAvegareStartYCor},
		{"SessionBasedAvegareFinishXcor", data.SessionBasedAvegareFinishXCor},
		{"SessionBasedAvegareFinishYCor", data.SessionBasedAvegareFinishYCor},
		{"SessionBasedAvegareClickCount", data.SessionBasedAvegareClickCount},
		{"DailyAvegareClickCount", data.DailyAvegareClickCount},
		{"LastTouchCountMinusSessionBasedAvegareClickCount", data.LastTouchCountMinusSessionBasedAvegareClickCount},
	}}}
	collection := m.Client.Database("MLDatabase").Collection("ScreenClickModel")
	updateResult := collection.FindOneAndUpdate(ctx, bson.D{{
		"ClientId", ClientId,
	}}, update)
	if updateResult.Err() != nil {
		return updateResult.Err()
	}
	return nil
}
