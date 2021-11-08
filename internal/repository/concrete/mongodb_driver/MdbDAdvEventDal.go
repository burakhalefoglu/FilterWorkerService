package mongodb_driver

import (
	"FilterWorkerService/internal/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MdbDAdvEventDal struct {
	Client *mongo.Client
}

func (m *MdbDAdvEventDal) Add(data *model.AdvEventRespondModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("MLDatabase").Collection("AdvEventModel")
	var _, err = collection.InsertOne(ctx, bson.D{
		{"ClientId", data.ClientId},
		{"ProjectId", data.ProjectId},
		{"CustomerId", data.CustomerId},
		{"LevelIndex", data.LevelIndex},
		{"TotalAdvDay", data.TotalAdvDay},
		{"TotalAdvCount", data.TotalAdvCount},
		{"TotalVideoAdvCount", data.TotalVideoAdvCount},
		{"TotalInterstitialAdvCount", data.TotalInterstitialAdvCount},
		{"LevelBasedAverageInterstitialAdvCount", data.LevelBasedAverageInterstitialAdvCount},
		{"LevelBasedAverageVideoAdvCount", data.LevelBasedAverageVideoAdvCount},
		{"AverageAdvDailyVideoAdvClickCount", data.AverageAdvDailyVideoClickCount},
		{"FirstAdvYearOfDay", data.FirstAdvYearOfDay},
		{"FirstAdvClickHour", data.FirstAdvClickHour},
		{"FirstVideoClickYearOfDay", data.FirstVideoClickYearOfDay},
		{"FirstVideoClickHour", data.FirstVideoClickHour},
		{"LastAdvYearOfDay", data.LastAdvYearOfDay},
		{"LastVideoClickYearOfDay", data.LastVideoClickYearOfDay},
		{"LastVideoClickHour", data.LastAdvClickHour},
		{"FirstDayVideoClickCount", data.FirstDayVideoClickCount},
		{"LastDayVideoClickCount", data.LastDayVideoClickCount},
		{"LastMinusFirstDayVideoClickCount", data.LastMinusFirstDayVideoClickCount},
		{"LastDayVideoClickCountMinusAverageDailyVideoAdvClickCount", data.LastDayVideoClickCountMinusAverageDailyVideoAdvClickCount},
		{"SundayVideoAdvClickCount", data.SundayVideoAdvClickCount},
		{"MondayVideoAdvClickCount", data.MondayVideoAdvClickCount},
		{"TuesdayVideoAdvClickCount", data.TuesdayVideoAdvClickCount},
		{"WednesdayVideoAdvClickCount", data.WednesdayVideoAdvClickCount},
		{"ThursdayVideoAdvClickCount", data.ThursdayVideoAdvClickCount},
		{"FridayVideoAdvClickCount", data.FridayVideoAdvClickCount},
		{"SaturdayVideoAdvClickCount", data.SaturdayVideoAdvClickCount},
		{"AmVideoAdvClickCount", data.AmVideoAdvClickCount},
		{"PmVideoAdvClickCount", data.PmVideoAdvClickCount},
		{"VideoAdvClick0To5HourCount", data.VideoAdvClick0To5HourCount},
		{"VideoAdvClick6To11HourCount", data.VideoAdvClick6To11HourCount},
		{"VideoAdvClick12To17HourCount", data.VideoAdvClick12To17HourCount},
		{"VideoAdvClick18To23HourCount", data.VideoAdvClick18To23HourCount},
	})
	if err != nil {
		return err
	}
	return nil
}

func (m *MdbDBuyingEventDal) GetAdvEventByCustomerId(CustomerId string) (*model.AdvEventRespondModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("MLDatabase").Collection("AdvEventModel")
	var result = collection.FindOne(ctx, bson.D{{
		"CustomerId", CustomerId,
	}})

	var model = model.AdvEventRespondModel{}
	if result.Err() != nil {
		return &model, result.Err()
	}
	var err = result.Decode(&model)
	if err != nil {
		return &model, err
	}
	return &model, nil
}

func (m *MdbDBuyingEventDal) UpdateAdvEventByCustomerId(CustomerId string, data *model.AdvEventRespondModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	update := bson.D{{"$set", bson.D{
		{"ClientId", data.ClientId},
		{"ProjectId", data.ProjectId},
		{"CustomerId", data.CustomerId},
		{"LevelIndex", data.LevelIndex},
		{"TotalAdvDay", data.TotalAdvDay},
		{"TotalAdvCount", data.TotalAdvCount},
		{"TotalVideoAdvCount", data.TotalVideoAdvCount},
		{"TotalInterstitialAdvCount", data.TotalInterstitialAdvCount},
		{"LevelBasedAverageInterstitialAdvCount", data.LevelBasedAverageInterstitialAdvCount},
		{"LevelBasedAverageVideoAdvCount", data.LevelBasedAverageVideoAdvCount},
		{"AverageAdvDailyVideoAdvClickCount", data.AverageAdvDailyVideoClickCount},
		//{"FirstAdvYearOfDay",data.FirstAdvYearOfDay},
		//{"FirstAdvClickHour",data.FirstAdvClickHour},
		//{"FirstVideoClickYearOfDay",data.FirstVideoClickYearOfDay},
		//{"FirstVideoClickHour",data.FirstVideoClickHour},
		{"LastAdvYearOfDay", data.LastAdvYearOfDay},
		{"LastVideoClickYearOfDay", data.LastVideoClickYearOfDay},
		{"LastVideoClickHour", data.LastAdvClickHour},
		{"FirstDayVideoClickCount", data.FirstDayVideoClickCount},
		{"LastDayVideoClickCount", data.LastDayVideoClickCount},
		{"LastMinusFirstDayVideoClickCount", data.LastMinusFirstDayVideoClickCount},
		{"LastDayVideoClickCountMinusAverageDailyVideoAdvClickCount", data.LastDayVideoClickCountMinusAverageDailyVideoAdvClickCount},
		{"SundayVideoAdvClickCount", data.SundayVideoAdvClickCount},
		{"MondayVideoAdvClickCount", data.MondayVideoAdvClickCount},
		{"TuesdayVideoAdvClickCount", data.TuesdayVideoAdvClickCount},
		{"WednesdayVideoAdvClickCount", data.WednesdayVideoAdvClickCount},
		{"ThursdayVideoAdvClickCount", data.ThursdayVideoAdvClickCount},
		{"FridayVideoAdvClickCount", data.FridayVideoAdvClickCount},
		{"SaturdayVideoAdvClickCount", data.SaturdayVideoAdvClickCount},
		{"AmVideoAdvClickCount", data.AmVideoAdvClickCount},
		{"PmVideoAdvClickCount", data.PmVideoAdvClickCount},
		{"VideoAdvClick0To5HourCount", data.VideoAdvClick0To5HourCount},
		{"VideoAdvClick6To11HourCount", data.VideoAdvClick6To11HourCount},
		{"VideoAdvClick12To17HourCount", data.VideoAdvClick12To17HourCount},
		{"VideoAdvClick18To23HourCount", data.VideoAdvClick18To23HourCount},
	}}}

	collection := m.Client.Database("MLDatabase").Collection("BuyingEventModel")
	updateResult := collection.FindOneAndUpdate(ctx, bson.D{{
		"CustomerId", CustomerId,
	}}, update)
	if updateResult.Err() != nil {
		return updateResult.Err()
	}
	return nil
}
