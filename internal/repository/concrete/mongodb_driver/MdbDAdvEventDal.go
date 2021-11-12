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
		{"LevelBasedAverageAdvCount", data.LevelBasedAverageAdvCount},
		{"AverageAdvDailyClickCount", data.AverageAdvDailyClickCount},
		{"FirstAdvYearOfDay", data.FirstAdvYearOfDay},
		{"FirstAdvYear", data.FirstAdvYear},
		{"FirstAdvClickHour", data.FirstAdvClickHour},
		{"FirstADvClickMinute",data.FirstADvClickMinute},
		{"FirstAdvType", data.FirstAdvType},
		{"SecondAdvYearOfDay", data.SecondAdvYearOfDay},
		{"SecondAdvHour", data.SecondAdvHour},
		{"SecondAdvMinute", data.SecondAdvMinute},
		{"ThirdAdvYearOfDay", data.ThirdAdvYearOfDay},
		{"ThirdAdvHour", data.ThirdAdvHour},
		{"ThirdAdvMinute", data.ThirdAdvMinute},
		{"PenultimateAdvYearOfDay", data.PenultimateAdvYearOfDay},
		{"PenultimateAdvHour", data.PenultimateAdvHour},
		{"PenultimateAdvMinute", data.PenultimateAdvMinute},
		{"LastAdvYearOfDay", data.LastAdvYearOfDay},
		{"LastAdvYear", data.LastAdvYear},
		{"LastVideoClickHour", data.LastAdvClickHour},
		{"LastAdvClickMinute", data.LastAdvClickMinute},
		{"LastAdvType", data.LastAdvType},
		{"FirstDayAdvClickCount", data.FirstDayAdvClickCount},
		{"PenultimateDayAdvClickCount", data.PenultimateDayAdvClickCount},
		{"LastDayAdvClickCount", data.LastDayAdvClickCount},
		{"LastDayVideoClickCountMinusAverageDailyVideoAdvClickCount", data.LastMinusFirstDayAdvClickCount},
		{"LastMinusPenultimateDayAdvClickCount", data.LastMinusPenultimateDayAdvClickCount},
		{"LastDayAdvClickCountMinusAverageDailyAdvClickCount", data.LastDayAdvClickCountMinusAverageDailyAdvClickCount},
		{"SundayVideoAdvClickCount", data.SundayAdvClickCount},
		{"MondayVideoAdvClickCount", data.MondayAdvClickCount},
		{"TuesdayVideoAdvClickCount", data.TuesdayAdvClickCount},
		{"WednesdayVideoAdvClickCount", data.WednesdayAdvClickCount},
		{"ThursdayVideoAdvClickCount", data.ThursdayAdvClickCount},
		{"FridayVideoAdvClickCount", data.FridayAdvClickCount},
		{"SaturdayVideoAdvClickCount", data.SaturdayAdvClickCount},
		{"AmVideoAdvClickCount", data.AmAdvClickCount},
		{"PmVideoAdvClickCount", data.PmAdvClickCount},
		{"VideoAdvClick0To5HourCount", data.AdvClick0To5HourCount},
		{"VideoAdvClick6To11HourCount", data.AdvClick6To11HourCount},
		{"VideoAdvClick12To17HourCount", data.AdvClick12To17HourCount},
		{"VideoAdvClick18To23HourCount", data.AdvClick18To23HourCount},
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
		{"LevelBasedAverageAdvCount", data.LevelBasedAverageAdvCount},
		{"AverageAdvDailyClickCount", data.AverageAdvDailyClickCount},
		{"FirstAdvYearOfDay", data.FirstAdvYearOfDay},
		{"FirstAdvYear", data.FirstAdvYear},
		{"FirstAdvClickHour", data.FirstAdvClickHour},
		{"FirstADvClickMinute",data.FirstADvClickMinute},
		{"FirstAdvType", data.FirstAdvType},
		{"SecondAdvYearOfDay", data.SecondAdvYearOfDay},
		{"SecondAdvHour", data.SecondAdvHour},
		{"SecondAdvMinute", data.SecondAdvMinute},
		{"ThirdAdvYearOfDay", data.ThirdAdvYearOfDay},
		{"ThirdAdvHour", data.ThirdAdvHour},
		{"ThirdAdvMinute", data.ThirdAdvMinute},
		{"PenultimateAdvYearOfDay", data.PenultimateAdvYearOfDay},
		{"PenultimateAdvHour", data.PenultimateAdvHour},
		{"PenultimateAdvMinute", data.PenultimateAdvMinute},
		{"LastAdvYearOfDay", data.LastAdvYearOfDay},
		{"LastAdvYear", data.LastAdvYear},
		{"LastVideoClickHour", data.LastAdvClickHour},
		{"LastAdvClickMinute", data.LastAdvClickMinute},
		{"LastAdvType", data.LastAdvType},
		{"FirstDayAdvClickCount", data.FirstDayAdvClickCount},
		{"PenultimateDayAdvClickCount", data.PenultimateDayAdvClickCount},
		{"LastDayAdvClickCount", data.LastDayAdvClickCount},
		{"LastDayVideoClickCountMinusAverageDailyVideoAdvClickCount", data.LastMinusFirstDayAdvClickCount},
		{"LastMinusPenultimateDayAdvClickCount", data.LastMinusPenultimateDayAdvClickCount},
		{"LastDayAdvClickCountMinusAverageDailyAdvClickCount", data.LastDayAdvClickCountMinusAverageDailyAdvClickCount},
		{"SundayVideoAdvClickCount", data.SundayAdvClickCount},
		{"MondayVideoAdvClickCount", data.MondayAdvClickCount},
		{"TuesdayVideoAdvClickCount", data.TuesdayAdvClickCount},
		{"WednesdayVideoAdvClickCount", data.WednesdayAdvClickCount},
		{"ThursdayVideoAdvClickCount", data.ThursdayAdvClickCount},
		{"FridayVideoAdvClickCount", data.FridayAdvClickCount},
		{"SaturdayVideoAdvClickCount", data.SaturdayAdvClickCount},
		{"AmVideoAdvClickCount", data.AmAdvClickCount},
		{"PmVideoAdvClickCount", data.PmAdvClickCount},
		{"VideoAdvClick0To5HourCount", data.AdvClick0To5HourCount},
		{"VideoAdvClick6To11HourCount", data.AdvClick6To11HourCount},
		{"VideoAdvClick12To17HourCount", data.AdvClick12To17HourCount},
		{"VideoAdvClick18To23HourCount", data.AdvClick18To23HourCount},
	}}}
	collection := m.Client.Database("MLDatabase").Collection("AdvEventModel")
	updateResult := collection.FindOneAndUpdate(ctx, bson.D{{
		"CustomerId", CustomerId,
	}}, update)
	if updateResult.Err() != nil {
		return updateResult.Err()
	}
	return nil
}
