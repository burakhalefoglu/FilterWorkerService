package mongodb_driver

import (
	"FilterWorkerService/internal/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MdbDBuyingEventDal struct {
	Client *mongo.Client
}

func (m *MdbDBuyingEventDal) Add(data *model.BuyingEventRespondModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := m.Client.Database("MLDatabase").Collection("BuyingEventModel")
	var _, err = collection.InsertOne(ctx, bson.D{
		{"ClientId", data.ClientId},
		{"ProjectId", data.ProjectId},
		{"CustomerId", data.CustomerId},
		{"LevelIndex", data.LevelIndex},
		{"TotalBuyingCount", data.TotalBuyingCount},
		{"TotalBuyingDay", data.TotalBuyingDay},
		{"FirstBuyingYearOfDay", data.FirstBuyingYearOfDay},
		{"FirstBuyingYear", data.FirstBuyingYear},
		{"FirstBuyingHour", data.FirstBuyingHour},
		{"FirstBuyingMinute", data.FirstBuyingMinute},
		{"SecondBuyingYearOfDay", data.SecondBuyingYearOfDay},
		{"SecondBuyingHour", data.SecondBuyingHour},
		{"ThirdBuyingYearOfDay", data.ThirdBuyingYearOfDay},
		{"ThirdBuyingHour", data.ThirdBuyingHour},
		{"PenultimateBuyingYearOfDay", data.PenultimateBuyingYearOfDay},
		{"PenultimateBuyingHour", data.PenultimateBuyingHour},
		{"LastBuyingYearOfDay", data.LastBuyingYearOfDay},
		{"LastBuyingYear", data.LastBuyingYear},
		{"LastBuyingHour", data.LastBuyingHour},
		{"LastBuyingMinute", data.LastBuyingMinute},
		{"FirstDayBuyingCount", data.FirstDayBuyingCount},
		{"PenultimateDayBuyingCount", data.PenultimateDayBuyingCount},
		{"LastDayBuyingCount", data.LastDayBuyingCount},
		{"LastMinusFirstDayBuyingCount", data.LastMinusFirstDayBuyingCount},
		{"SundayBuyingCount", data.SundayBuyingCount},
		{"MondayBuyingCount", data.MondayBuyingCount},
		{"TuesdayBuyingCount", data.TuesdayBuyingCount},
		{"WednesdayBuyingCount", data.WednesdayBuyingCount},
		{"ThursdayBuyingCount", data.ThursdayBuyingCount},
		{"FridayBuyingCount", data.FridayBuyingCount},
		{"SaturdayBuyingCount", data.SaturdayBuyingCount},
		{"AmBuyingCount", data.AmBuyingCount},
		{"PmBuyingCount", data.PmBuyingCount},
		{"Buying0To5HourCount", data.Buying0To5HourCount},
		{"Buying6To11HourCount", data.Buying6To11HourCount},
		{"Buying12To17HourCount", data.Buying12To17HourCount},
		{"Buying18To23HourCount", data.Buying18To23HourCount},
		{"BuyingDayAverageBuyingCount", data.BuyingDayAverageBuyingCount},
		{"LevelBasedAverageBuyingCount", data.LevelBasedAverageBuyingCount},
	})
	if err != nil {
		return err
	}
	return nil
}

func (m *MdbDBuyingEventDal) GetBuyingEventById(ClientId string) (*model.BuyingEventRespondModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//"BuyingEventModel"
	collection := m.Client.Database("MLDatabase").Collection("BuyingEventModel")
	var result = collection.FindOne(ctx, bson.D{{
		"ClientId", ClientId,
	}})

	var model = model.BuyingEventRespondModel{}
	if result.Err() != nil {
		return &model, result.Err()
	}
	var err = result.Decode(&model)
	if err != nil {
		return &model, err
	}
	return &model, nil
}

func (m *MdbDBuyingEventDal) UpdateBuyingEventById(ClientId string, data *model.BuyingEventRespondModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.D{{"$set", bson.D{
		{"ClientId", data.ClientId},
		{"ProjectId", data.ProjectId},
		{"CustomerId", data.CustomerId},
		{"LevelIndex", data.LevelIndex},
		{"TotalBuyingCount", data.TotalBuyingCount},
		{"TotalBuyingDay", data.TotalBuyingDay},
		{"FirstBuyingYearOfDay", data.FirstBuyingYearOfDay},
		{"FirstBuyingYear", data.FirstBuyingYear},
		{"FirstBuyingHour", data.FirstBuyingHour},
		{"FirstBuyingMinute", data.FirstBuyingMinute},
		{"SecondBuyingYearOfDay", data.SecondBuyingYearOfDay},
		{"SecondBuyingHour", data.SecondBuyingHour},
		{"ThirdBuyingYearOfDay", data.ThirdBuyingYearOfDay},
		{"ThirdBuyingHour", data.ThirdBuyingHour},
		{"PenultimateBuyingYearOfDay", data.PenultimateBuyingYearOfDay},
		{"PenultimateBuyingHour", data.PenultimateBuyingHour},
		{"LastBuyingYearOfDay", data.LastBuyingYearOfDay},
		{"LastBuyingYear", data.LastBuyingYear},
		{"LastBuyingHour", data.LastBuyingHour},
		{"LastBuyingMinute", data.LastBuyingMinute},
		{"FirstDayBuyingCount", data.FirstDayBuyingCount},
		{"PenultimateDayBuyingCount", data.PenultimateDayBuyingCount},
		{"LastDayBuyingCount", data.LastDayBuyingCount},
		{"LastMinusFirstDayBuyingCount", data.LastMinusFirstDayBuyingCount},
		{"SundayBuyingCount", data.SundayBuyingCount},
		{"MondayBuyingCount", data.MondayBuyingCount},
		{"TuesdayBuyingCount", data.TuesdayBuyingCount},
		{"WednesdayBuyingCount", data.WednesdayBuyingCount},
		{"ThursdayBuyingCount", data.ThursdayBuyingCount},
		{"FridayBuyingCount", data.FridayBuyingCount},
		{"SaturdayBuyingCount", data.SaturdayBuyingCount},
		{"AmBuyingCount", data.AmBuyingCount},
		{"PmBuyingCount", data.PmBuyingCount},
		{"Buying0To5HourCount", data.Buying0To5HourCount},
		{"Buying6To11HourCount", data.Buying6To11HourCount},
		{"Buying12To17HourCount", data.Buying12To17HourCount},
		{"Buying18To23HourCount", data.Buying18To23HourCount},
		{"BuyingDayAverageBuyingCount", data.BuyingDayAverageBuyingCount},
		{"LevelBasedAverageBuyingCount", data.LevelBasedAverageBuyingCount},
	}}}
	collection := m.Client.Database("MLDatabase").Collection("BuyingEventModel")
	updateResult := collection.FindOneAndUpdate(ctx, bson.D{{
		"ClientId", ClientId,
	}}, update)
	if updateResult.Err() != nil {
		return updateResult.Err()
	}
	return nil
}
