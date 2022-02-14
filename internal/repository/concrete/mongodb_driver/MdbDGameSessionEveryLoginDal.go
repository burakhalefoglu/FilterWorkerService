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

type mdbdDGameSessionEveryLoginDal struct {
	Client *mongo.Client
}

func MdbDGameSessionEveryLoginDalConstructor() *mdbdDGameSessionEveryLoginDal {
	return &mdbdDGameSessionEveryLoginDal{Client: mongodb.GetMongodbClient()}
}

func (m *mdbdDGameSessionEveryLoginDal) Add(data *model.GameSessionEveryLoginRespondModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := m.Client.Database("MLDatabase").Collection("gameSessionEveryLoginModels")
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
		{"ThirdSessinMinute", data.ThirdSessionMinute},
		{"FourthSessionHour", data.FourthSessionHour},
		{"FourthSessionDuration", data.FourthSessionDuration},
		{"FourthSessinMinute", data.FourthSessionMinute},
		{"FifthSessionHour", data.FifthSessionHour},
		{"FifthSessionDuration", data.FifthSessionDuration},
		{"FifthSessinMinute", data.FifthSessionMinute},
		{"PenultimateSessionHour", data.PenultimateSessionHour},
		{"PenultimateSessionDuration", data.PenultimateSessionDuration},
		{"PenultimateSessionMinute", data.PenultimateSessionMinute},
		{"LastSessionYearOfDay", data.LastSessionYearOfDay},
		{"LastSessionYear", data.LastSessionYear},
		{"LastSessionHour", data.LastSessionHour},
		{"LastSessionDuration", data.LastSessionDuration},
		{"LastSessionMinute", data.LastSessionMinute},
		{"LastDurationMinusPenultimateDuration", data.LastDurationMinusPenultimateDuration},
		{"FirstHalfHourTotalSessionCount", data.FirstHalfHourTotalSessionCount},
		{"FirstHalfHourTotalSessionDuration", data.FirstHalfHourTotalSessionDuration},
		{"FirstHourTotalSessionCount", data.FirstHourTotalSessionCount},
		{"FirstHourTotalSessionDuration", data.FirstHourTotalSessionDuration},
		{"FirstTwoHourTotalSessionCount", data.FirstTwoHourTotalSessionCount},
		{"FirstTwoHourTotalSessionDuration", data.FirstTwoHourTotalSessionDuration},
		{"FirstThreeHourTotalSessionCount", data.FirstThreeHourTotalSessionCount},
		{"FirstThreeHourTotalSessionDuration", data.FirstThreeHourTotalSessionDuration},
		{"FirstSixHourTotalSessionCount", data.FirstSixHourTotalSessionCount},
		{"FirstSixHourTotalSessionDuration", data.FirstSixHourTotalSessionDuration},
		{"FirstTwelveHourTotalSessionCount", data.FirstTwelveHourTotalSessionCount},
		{"FirstTwelveHourTotalSessionDuration", data.FirstTwelveHourTotalSessionDuration},
		{"TotalSessionDay", data.TotalSessionDay},
		{"TotalSessionHour", data.TotalSessionHour},
		{"TotalSessionMinute", data.TotalSessionMinute},
		{"TotalSessionDuration", data.TotalSessionDuration},
		{"TotalSessionCount", data.TotalSessionCount},
		{"FirstDayTotalSessionCount", data.FirstDayTotalSessionCount},
		{"FirstDayTotalSessionDuration", data.FirstDayTotalSessionDuration},
		{"SecondDayTotalSessionCount", data.SecondDayTotalSessionCount},
		{"SecondDayTotalSessionDuration", data.SecondDayTotalSessionDuration},
		{"ThirdDayTotalSessionCount", data.ThirdDayTotalSessionCount},
		{"ThirdDayTotalSessionDuration", data.ThirdDayTotalSessionDuration},
		{"FourthDayTotalSessionCount", data.FourthDayTotalSessionCount},
		{"FourthDayTotalSessionDuration", data.FourthDayTotalSessionDuration},
		{"FifthDayTotalSessionCount", data.FifthDayTotalSessionCount},
		{"FifthDayTotalSessionDuration", data.FifthDayTotalSessionDuration},
		{"SixthDayTotalSessionCount", data.SixthDayTotalSessionCount},
		{"SixthDayTotalSessionDuration", data.SixthDayTotalSessionDuration},
		{"SeventhDayTotalSessionCount", data.SeventhDayTotalSessionCount},
		{"SeventhDayTotalSessionDuration", data.SeventhDayTotalSessionDuration},
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

func (m *mdbdDGameSessionEveryLoginDal) GetGameSessionEveryLoginById(ClientId string) (*model.GameSessionEveryLoginRespondModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := m.Client.Database("MLDatabase").Collection("GameSessionEveryLoginModel")
	var result = collection.FindOne(ctx, bson.D{{
		"ClientId", ClientId,
	}})
	var model = model.GameSessionEveryLoginRespondModel{}
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

func (m *mdbdDGameSessionEveryLoginDal) UpdateGameSessionEveryLoginById(ClientId string, data *model.GameSessionEveryLoginRespondModel) error {
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
		{"ThirdSessinMinute", data.ThirdSessionMinute},
		{"FourthSessionHour", data.FourthSessionHour},
		{"FourthSessionDuration", data.FourthSessionDuration},
		{"FourthSessinMinute", data.FourthSessionMinute},
		{"FifthSessionHour", data.FifthSessionHour},
		{"FifthSessionDuration", data.FifthSessionDuration},
		{"FifthSessinMinute", data.FifthSessionMinute},
		{"PenultimateSessionHour", data.PenultimateSessionHour},
		{"PenultimateSessionDuration", data.PenultimateSessionDuration},
		{"PenultimateSessionMinute", data.PenultimateSessionMinute},
		{"LastSessionYearOfDay", data.LastSessionYearOfDay},
		{"LastSessionYear", data.LastSessionYear},
		{"LastSessionHour", data.LastSessionHour},
		{"LastSessionDuration", data.LastSessionDuration},
		{"LastSessionMinute", data.LastSessionMinute},
		{"LastDurationMinusPenultimateDuration", data.LastDurationMinusPenultimateDuration},
		{"FirstHalfHourTotalSessionCount", data.FirstHalfHourTotalSessionCount},
		{"FirstHalfHourTotalSessionDuration", data.FirstHalfHourTotalSessionDuration},
		{"FirstHourTotalSessionCount", data.FirstHourTotalSessionCount},
		{"FirstHourTotalSessionDuration", data.FirstHourTotalSessionDuration},
		{"FirstTwoHourTotalSessionCount", data.FirstTwoHourTotalSessionCount},
		{"FirstTwoHourTotalSessionDuration", data.FirstTwoHourTotalSessionDuration},
		{"FirstThreeHourTotalSessionCount", data.FirstThreeHourTotalSessionCount},
		{"FirstThreeHourTotalSessionDuration", data.FirstThreeHourTotalSessionDuration},
		{"FirstSixHourTotalSessionCount", data.FirstSixHourTotalSessionCount},
		{"FirstSixHourTotalSessionDuration", data.FirstSixHourTotalSessionDuration},
		{"FirstTwelveHourTotalSessionCount", data.FirstTwelveHourTotalSessionCount},
		{"FirstTwelveHourTotalSessionDuration", data.FirstTwelveHourTotalSessionDuration},
		{"TotalSessionDay", data.TotalSessionDay},
		{"TotalSessionHour", data.TotalSessionHour},
		{"TotalSessionMinute", data.TotalSessionMinute},
		{"TotalSessionDuration", data.TotalSessionDuration},
		{"TotalSessionCount", data.TotalSessionCount},
		{"FirstDayTotalSessionCount", data.FirstDayTotalSessionCount},
		{"FirstDayTotalSessionDuration", data.FirstDayTotalSessionDuration},
		{"SecondDayTotalSessionCount", data.SecondDayTotalSessionCount},
		{"SecondDayTotalSessionDuration", data.SecondDayTotalSessionDuration},
		{"ThirdDayTotalSessionCount", data.ThirdDayTotalSessionCount},
		{"ThirdDayTotalSessionDuration", data.ThirdDayTotalSessionDuration},
		{"FourthDayTotalSessionCount", data.FourthDayTotalSessionCount},
		{"FourthDayTotalSessionDuration", data.FourthDayTotalSessionDuration},
		{"FifthDayTotalSessionCount", data.FifthDayTotalSessionCount},
		{"FifthDayTotalSessionDuration", data.FifthDayTotalSessionDuration},
		{"SixthDayTotalSessionCount", data.SixthDayTotalSessionCount},
		{"SixthDayTotalSessionDuration", data.SixthDayTotalSessionDuration},
		{"SeventhDayTotalSessionCount", data.SeventhDayTotalSessionCount},
		{"SeventhDayTotalSessionDuration", data.SeventhDayTotalSessionDuration},
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
		"ClientId", ClientId,
	}}, update)
	if updateResult.Err() != nil {
		return updateResult.Err()
	}
	return nil
}
