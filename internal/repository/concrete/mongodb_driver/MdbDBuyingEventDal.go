package mongodb_driver

import (
	"FilterWorkerService/internal/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MdbDBuyingEventDal struct {
	Client *mongo.Client
}

func (m *MdbDBuyingEventDal) Add(data *model.BuyingEventRespondModel) error{
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := m.Client.Database("MLDatabase").Collection("BuyingEventModel")
	var _, err := collection.InsertOne(ctx, bson.D{
		{"ClientId",data.ClientId},
		{"ProjectId", data.ProjectId},
		{"CustomerId",data.CustomerId},
		{"LevelIndex", data.LevelIndex},
		{"TotalBuyingCount", data.TotalBuyingCount},
		{"TotalBuyingDay", data.TotalBuyingDay},
		{"TotalBuyingSession", data.TotalBuyingSession},
		{"TotalSession", data.TotalSession},
		{"TotalDay", data.TotalDay},
		{"FirstBuyingMonth", data.FirstBuyingMonth},
		{"FirstBuyingWeek", data.FirstBuyingWeek},
		{"FirstBuyingDay", data.FirstBuyingDay},
		{"FirstBuyingHour", data.FirstBuyingHour},
		{"LastBuyingMonth", data.LastBuyingMonth},
		{"LastBuyingWeek", data.LastBuyingWeek},
		{"LastBuyingDay", data.LastBuyingDay},
		{"LastBuyingHour", data.LastBuyingHour},
		{"FirstDayBuyingCount", data.FirstDayBuyingCount},
		{"PenultimateDayBuyingCount", data.PenultimateDayBuyingCount},
		{"LastDayBuyingCount", data.LastDayBuyingCount},
		{"LastMinusPenultimateDayBuyingCount", data.LastMinusPenultimateDayBuyingCount},
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
		{"DailyAverageBuyingCount", data.DailyAverageBuyingCount},
		{"BuyingDayAverageBuyingCount", data.BuyingDayAverageBuyingCount},
		{"LevelBasedAverageBuyingCount", data.LevelBasedAverageBuyingCount},
		{"SessionBasedAverageBuyingCount", data.SessionBasedAverageBuyingCount},
		{"BuyingSessionBasedAverageBuyingCount", data.BuyingSessionBasedAverageBuyingCount},
		{"FirstBuyingDayMinusFirstSessionDay", data.FirstBuyingDayMinusFirstSessionDay},
		{"FirstBuyingMonthMinusFirstSessionMonth", data.FirstBuyingMonthMinusFirstSessionMonth},
		{"TotalDifferenceBetweenFirstBuyingDayAndFirstSessionDay", data.TotalDifferenceBetweenFirstBuyingDayAndFirstSessionDay},
		{"IsDeadAndBuyingItemCount", data.IsDeadAndBuyingItemCount},

	})
	if err != nil {
			return err
		}
		return nil
}

func (m *MdbDBuyingEventDal) GetByCustomerId(CustomerId string)(*model.BuyingEventRespondModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("MLDatabase").Collection("BuyingEventModel")
	var result = collection.FindOne(ctx, bson.D{{
		"CustomerId",CustomerId,
	}})

	var model = model.BuyingEventRespondModel{}
	if result.Err() != nil {
		return &model, result.Err()
	}
	var err = result.Decode(&model)
	if err != nil{
		return &model, err
	}
	return &model, nil
}

func (m *MdbDBuyingEventDal) UpdateByCustomerId(CustomerId string, data *model.BuyingEventRespondModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.D{{"$set", bson.D{
		{"ClientId",data.ClientId},
		{"ProjectId", data.ProjectId},
		{"CustomerId",data.CustomerId},
		{"LevelIndex", data.LevelIndex},
		{"TotalBuyingCount", data.TotalBuyingCount},
		{"TotalBuyingDay", data.TotalBuyingDay},
		{"TotalBuyingSession", data.TotalBuyingSession},
		{"TotalSession", data.TotalSession},
		{"TotalDay", data.TotalDay},
		// {"FirstBuyingMonth", data.FirstBuyingMonth},
		// {"FirstBuyingWeek", data.FirstBuyingWeek},
		// {"FirstBuyingDay", data.FirstBuyingDay},
		// {"FirstBuyingHour", data.FirstBuyingHour},
		{"LastBuyingMonth", data.LastBuyingMonth},
		{"LastBuyingWeek", data.LastBuyingWeek},
		{"LastBuyingDay", data.LastBuyingDay},
		{"LastBuyingHour", data.LastBuyingHour},
		//{"FirstDayBuyingCount", data.FirstDayBuyingCount},
		{"PenultimateDayBuyingCount", data.PenultimateDayBuyingCount},
		{"LastDayBuyingCount", data.LastDayBuyingCount},
		{"LastMinusPenultimateDayBuyingCount", data.LastMinusPenultimateDayBuyingCount},
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
		{"DailyAverageBuyingCount", data.DailyAverageBuyingCount},
		{"BuyingDayAverageBuyingCount", data.BuyingDayAverageBuyingCount},
		{"LevelBasedAverageBuyingCount", data.LevelBasedAverageBuyingCount},
		{"SessionBasedAverageBuyingCount", data.SessionBasedAverageBuyingCount},
		{"BuyingSessionBasedAverageBuyingCount", data.BuyingSessionBasedAverageBuyingCount},
		{"FirstBuyingDayMinusFirstSessionDay", data.FirstBuyingDayMinusFirstSessionDay},
		{"FirstBuyingMonthMinusFirstSessionMonth", data.FirstBuyingMonthMinusFirstSessionMonth},
		{"TotalDifferenceBetweenFirstBuyingDayAndFirstSessionDay", data.TotalDifferenceBetweenFirstBuyingDayAndFirstSessionDay},
		{"IsDeadAndBuyingItemCount", data.IsDeadAndBuyingItemCount},
	}}}

	collection := m.Client.Database("MLDatabase").Collection("BuyingEventModel")
	updateResult := collection.FindOneAndUpdate(ctx, bson.D{{
		"CustomerId",CustomerId,
	}}, update)
	if updateResult.Err() != nil{
		return updateResult.Err()
	}
	return nil
}