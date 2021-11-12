package mongodb_driver

import (
	"FilterWorkerService/internal/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MdbdDGameSessionEveryLoginDal struct {
	Client *mongo.Client
}

func (m *MdbdDGameSessionEveryLoginDal) Add(data *model.GameSessionEveryLoginRespondModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := m.Client.Database("MLDatabase").Collection("GameSessionEveryLoginModel")
	var _, err = collection.InsertOne(ctx, bson.D{
		{"ClientId", data.ClientId},
		{"ProjectId", data.ProjectId},
		{"CustomerId", data.CustomerId},
		{"FirstSessionYearOfDay", data.FirstSessionYearOfDay},
		{"FirstSessionYear", data.FirstSessionYear},
		{"FirstSessionWeekDay", data.FirstSessionWeekDay},
		{"FirstSessionHour", data.FirstSessionHour},		
		{"FirstSessionDuration", data.FirstSessionDuration},
		{"FirstSessionMinute", data.FirstSessionMinute},
		{"SecondSessionHour", data.SecondSessionHour},
		{"SecondSessionDuration", data.SecondSessionDuration},
		{"SecondSessionMinute", data.SecondSessionMinute},
		{"ThirdSessionHour", data.ThirdSessionHour},
		{"ThirdSessionDuration", data.ThirdSessionDuration},
		{"ThirdSessinMinute", data.ThirdSessinMinute},
		{"PenultimateSessionHour", data.PenultimateSessionHour},
		{"PenultimateSessionDuration", data.PenultimateSessionDuration},
		{"PenultimateSessionMinute", data.PenultimateSessionMinute},
		{"LastSessionYearOfDay", data.LastSessionYearOfDay},
		{"LastSessionYear", data.LastSessionYear},
		{"LastSessionHour", data.LastSessionHour},
		{"LastSessionDuration", data.LastSessionDuration},
		{"LastSessionMinute", data.LastSessionMinute},
		{"LastDurationMinusPenultimateDuration", data.LastDurationMinusPenultimateDuration},
		{"TotalSessionDay", data.TotalSessionDay},
		{"TotalSessionDuration", data.TotalSessionDuration},
		{"TotalSessionCount", data.TotalSessionCount},
		{"FirstDayTotalSessionCount", data.FirstDayTotalSessionCount},
		{"FirstDayTotalSessionDuration", data.FirstDayTotalSessionDuration},
		{"PenultimateDayTotalSessionDuration", data.PenultimateDayTotalSessionDuration},
		{"PenultimateDayTotalSessionCount", data.PenultimateDayTotalSessionCount},
		{"LastDayTotalSessionCount", data.LastDayTotalSessionCount},
		{"LastDayTotalSessionDuration", data.LastDayTotalSessionDuration},
		{"MinSessionDuration", data.MinSessionDuration},
		{"MaxSessionDuration", data.MaxSessionDuration},
		{"DailyAvegareSessionCount", data.DailyAvegareSessionCount},
		{"DailyAverageSessionDuration", data.DailyAverageSessionDuration},
		{"SessionBasedAvegareSessionDuration", data.SessionBasedAvegareSessionDuration},
		{"DailyAvegareSessionCountMinusFirstDaySessionCount", data.DailyAvegareSessionCountMinusFirstDaySessionCount},
		{"DailyAvegareSessionDurationMinusFirstDaySessionDuration", data.DailyAvegareSessionDurationMinusFirstDaySessionDuration},
		{"SessionBasedAvegareSessionDurationMinusFirstSessionDuration", data.SessionBasedAvegareSessionDurationMinusFirstSessionDuration},
		{"SessionBasedAvegareSessionDurationMinusLastSessionDuration", data.SessionBasedAvegareSessionDurationMinusLastSessionDuration},
		{"SundaySessionCount", data.SundaySessionCount},
		{"MondaySessionCount", data.MondaySessionCount},
		{"TuesdaySessionCount", data.TuesdaySessionCount},
		{"WednesdaySessionCount", data.WednesdaySessionCount},
		{"ThursdaySessionCount", data.ThursdaySessionCount},
		{"FridaySessionCount", data.FridaySessionCount},
		{"SaturdaySessionCount", data.SaturdaySessionCount},
		{"AmSessionCount", data.AmSessionCount},
		{"PmSessionCount", data.PmSessionCount},
		{"Session0To5HourCount", data.Session0To5HourCount},
		{"Session6To11HourCount", data.Session6To11HourCount},
		{"Session12To17HourCount", data.Session12To17HourCount},
		{"Session18To23HourCount", data.Session18To23HourCount},
	})
	if err != nil {
		return err
	}
	return nil
}

func (m *MdbdDGameSessionEveryLoginDal) GetGameSessionEveryLoginByCustomerId(CustomerId string) (*model.GameSessionEveryLoginRespondModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := m.Client.Database("MLDatabase").Collection("GameSessionEveryLoginModel")
	var result = collection.FindOne(ctx, bson.D{{
		"CustomerId", CustomerId,
	}})
	var model = model.GameSessionEveryLoginRespondModel{}
	if result.Err() != nil {
		return &model, result.Err()
	}
	var err = result.Decode(&model)
	if err != nil {
		return &model, err
	}
	return &model, nil
}

func (m *MdbdDGameSessionEveryLoginDal) UpdateGameSessionEveryLoginByCustomerId(CustomerId string, data *model.GameSessionEveryLoginRespondModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	update := bson.D{{"$set", bson.D{
		{"ClientId", data.ClientId},
		{"ProjectId", data.ProjectId},
		{"CustomerId", data.CustomerId},
		{"FirstSessionYearOfDay", data.FirstSessionYearOfDay},
		{"FirstSessionYear", data.FirstSessionYear},
		{"FirstSessionWeekDay", data.FirstSessionWeekDay},
		{"FirstSessionHour", data.FirstSessionHour},		
		{"FirstSessionDuration", data.FirstSessionDuration},
		{"FirstSessionMinute", data.FirstSessionMinute},
		{"SecondSessionHour", data.SecondSessionHour},
		{"SecondSessionDuration", data.SecondSessionDuration},
		{"SecondSessionMinute", data.SecondSessionMinute},
		{"ThirdSessionHour", data.ThirdSessionHour},
		{"ThirdSessionDuration", data.ThirdSessionDuration},
		{"ThirdSessinMinute", data.ThirdSessinMinute},
		{"PenultimateSessionHour", data.PenultimateSessionHour},
		{"PenultimateSessionDuration", data.PenultimateSessionDuration},
		{"PenultimateSessionMinute", data.PenultimateSessionMinute},
		{"LastSessionYearOfDay", data.LastSessionYearOfDay},
		{"LastSessionYear", data.LastSessionYear},
		{"LastSessionHour", data.LastSessionHour},
		{"LastSessionDuration", data.LastSessionDuration},
		{"LastSessionMinute", data.LastSessionMinute},
		{"LastDurationMinusPenultimateDuration", data.LastDurationMinusPenultimateDuration},
		{"TotalSessionDay", data.TotalSessionDay},
		{"TotalSessionDuration", data.TotalSessionDuration},
		{"TotalSessionCount", data.TotalSessionCount},
		{"FirstDayTotalSessionCount", data.FirstDayTotalSessionCount},
		{"FirstDayTotalSessionDuration", data.FirstDayTotalSessionDuration},
		{"PenultimateDayTotalSessionDuration", data.PenultimateDayTotalSessionDuration},
		{"PenultimateDayTotalSessionCount", data.PenultimateDayTotalSessionCount},
		{"LastDayTotalSessionCount", data.LastDayTotalSessionCount},
		{"LastDayTotalSessionDuration", data.LastDayTotalSessionDuration},
		{"MinSessionDuration", data.MinSessionDuration},
		{"MaxSessionDuration", data.MaxSessionDuration},
		{"DailyAvegareSessionCount", data.DailyAvegareSessionCount},
		{"DailyAverageSessionDuration", data.DailyAverageSessionDuration},
		{"SessionBasedAvegareSessionDuration", data.SessionBasedAvegareSessionDuration},
		{"DailyAvegareSessionCountMinusFirstDaySessionCount", data.DailyAvegareSessionCountMinusFirstDaySessionCount},
		{"DailyAvegareSessionDurationMinusFirstDaySessionDuration", data.DailyAvegareSessionDurationMinusFirstDaySessionDuration},
		{"SessionBasedAvegareSessionDurationMinusFirstSessionDuration", data.SessionBasedAvegareSessionDurationMinusFirstSessionDuration},
		{"SessionBasedAvegareSessionDurationMinusLastSessionDuration", data.SessionBasedAvegareSessionDurationMinusLastSessionDuration},
		{"SundaySessionCount", data.SundaySessionCount},
		{"MondaySessionCount", data.MondaySessionCount},
		{"TuesdaySessionCount", data.TuesdaySessionCount},
		{"WednesdaySessionCount", data.WednesdaySessionCount},
		{"ThursdaySessionCount", data.ThursdaySessionCount},
		{"FridaySessionCount", data.FridaySessionCount},
		{"SaturdaySessionCount", data.SaturdaySessionCount},
		{"AmSessionCount", data.AmSessionCount},
		{"PmSessionCount", data.PmSessionCount},
		{"Session0To5HourCount", data.Session0To5HourCount},
		{"Session6To11HourCount", data.Session6To11HourCount},
		{"Session12To17HourCount", data.Session12To17HourCount},
		{"Session18To23HourCount", data.Session18To23HourCount},
	}}}
	collection := m.Client.Database("MLDatabase").Collection("GameSessionEveryLoginModel")
	updateResult := collection.FindOneAndUpdate(ctx, bson.D{{
		"CustomerId", CustomerId,
	}}, update)
	if updateResult.Err() != nil {
		return updateResult.Err()
	}
	return nil
}
