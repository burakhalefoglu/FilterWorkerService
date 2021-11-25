package mongodb_driver

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/pkg/database/mongodb"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mdbDAdvEventDal struct {
	Client *mongo.Client
}

func MdbDAdvEventDalConstructor() *mdbDAdvEventDal {
	return &mdbDAdvEventDal{Client: mongodb.GetMongodbClient()}
}

func (m *mdbDAdvEventDal) Add(data *model.AdvEventRespondModel) error {
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
		{"TotalAdvHour", data.TotalAdvHour},
		{"TotalAdvMinute", data.TotalAdvMinute},
		{"LevelBasedAverageAdvCount", data.LevelBasedAverageAdvCount},
		{"AverageAdvDailyClickCount", data.AverageAdvDailyClickCount},
		{"FirstAdvYearOfDay", data.FirstAdvYearOfDay},
		{"FirstAdvYear", data.FirstAdvYear},
		{"FirstWeekDay",data.FirstWeekDay},
		{"FirstAdvClickHour", data.FirstAdvClickHour},
		{"FirstADvClickMinute", data.FirstADvClickMinute},
		{"FirstAdvType", data.FirstAdvType},
		{"SecondAdvYearOfDay", data.SecondAdvYearOfDay},
		{"SecondAdvHour", data.SecondAdvHour},
		{"SecondAdvMinute", data.SecondAdvMinute},
		{"SecondAdvType", data.SecondAdvType},
		{"ThirdAdvYearOfDay", data.ThirdAdvYearOfDay},
		{"ThirdAdvHour", data.ThirdAdvHour},
		{"ThirdAdvMinute", data.ThirdAdvMinute},
		{"ThirdAdvType", data.ThirdAdvType},
		{"FourthAdvYearOfDay", data.FourthAdvYearOfDay},
		{"FourthAdvHour", data.FourthAdvHour},
		{"FourthAdvMinute", data.FourthAdvMinute},
		{"FourthAdvType", data.FourthAdvType},
		{"FifthAdvYearOfDay", data.FifthAdvYearOfDay},
		{"FifthAdvHour", data.FifthAdvHour},
		{"FifthAdvMinute", data.FifthAdvMinute},
		{"FifthAdvType", data.FifthAdvType},
		{"PenultimateAdvYearOfDay", data.PenultimateAdvYearOfDay},
		{"PenultimateAdvHour", data.PenultimateAdvHour},
		{"PenultimateAdvMinute", data.PenultimateAdvMinute},
		{"PenultimateAdvType", data.PenultimateAdvType},
		{"LastAdvYearOfDay", data.LastAdvYearOfDay},
		{"LastAdvYear", data.LastAdvYear},
		{"LastVideoClickHour", data.LastAdvClickHour},
		{"LastAdvClickMinute", data.LastAdvClickMinute},
		{"LastAdvType", data.LastAdvType},
		{"FirstHalfHourAdvClickCount", data.FirstHalfHourAdvClickCount},
		{"FirstHourAdvClickCount", data.FirstHourAdvClickCount},
		{"FirstTwoHourAdvClickCount", data.FirstTwoHourAdvClickCount},
		{"FirstThreeHourAdvClickCount", data.FirstThreeHourAdvClickCount},
		{"FirstSixHourAdvClickCount", data.FirstSixHourAdvClickCount},
		{"FirstTwelveHourAdvClickCount", data.FirstTwelveHourAdvClickCount},
		{"FirstDayAdvClickCount", data.FirstDayAdvClickCount},
		{"SecondDayAdvClickCount", data.SecondDayAdvClickCount},
		{"ThirdDayAdvClickCount", data.ThirdDayAdvClickCount},
		{"FourthDayAdvClickCount", data.FourthDayAdvClickCount},
		{"FifthDayAdvClickCount", data.FifthDayAdvClickCount},
		{"SixthDayAdvClickCount", data.SixthDayAdvClickCount},
		{"SeventhDayAdvClickCount", data.SeventhDayAdvClickCount},
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

func (m *mdbDAdvEventDal) GetAdvEventById(ClientId string) (*model.AdvEventRespondModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.Client.Database("MLDatabase").Collection("AdvEventModel")
	var result = collection.FindOne(ctx, bson.D{{
		"ClientId", ClientId,
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

func (m *mdbDAdvEventDal) UpdateAdvEventById(ClientId string, data *model.AdvEventRespondModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	update := bson.D{{"$set", bson.D{
		{"ClientId", data.ClientId},
		{"ProjectId", data.ProjectId},
		{"CustomerId", data.CustomerId},
		{"LevelIndex", data.LevelIndex},
		{"TotalAdvDay", data.TotalAdvDay},
		{"TotalAdvCount", data.TotalAdvCount},
		{"TotalAdvHour", data.TotalAdvHour},
		{"TotalAdvMinute", data.TotalAdvMinute},
		{"LevelBasedAverageAdvCount", data.LevelBasedAverageAdvCount},
		{"AverageAdvDailyClickCount", data.AverageAdvDailyClickCount},
		{"FirstAdvYearOfDay", data.FirstAdvYearOfDay},
		{"FirstAdvYear", data.FirstAdvYear},
		{"FirstWeekDay",data.FirstWeekDay},
		{"FirstAdvClickHour", data.FirstAdvClickHour},
		{"FirstADvClickMinute", data.FirstADvClickMinute},
		{"FirstAdvType", data.FirstAdvType},
		{"SecondAdvYearOfDay", data.SecondAdvYearOfDay},
		{"SecondAdvHour", data.SecondAdvHour},
		{"SecondAdvMinute", data.SecondAdvMinute},
		{"SecondAdvType", data.SecondAdvType},
		{"ThirdAdvYearOfDay", data.ThirdAdvYearOfDay},
		{"ThirdAdvHour", data.ThirdAdvHour},
		{"ThirdAdvMinute", data.ThirdAdvMinute},
		{"ThirdAdvType", data.ThirdAdvType},
		{"FourthAdvYearOfDay", data.FourthAdvYearOfDay},
		{"FourthAdvHour", data.FourthAdvHour},
		{"FourthAdvMinute", data.FourthAdvMinute},
		{"FourthAdvType", data.FourthAdvType},
		{"FifthAdvYearOfDay", data.FifthAdvYearOfDay},
		{"FifthAdvHour", data.FifthAdvHour},
		{"FifthAdvMinute", data.FifthAdvMinute},
		{"FifthAdvType", data.FifthAdvType},
		{"PenultimateAdvYearOfDay", data.PenultimateAdvYearOfDay},
		{"PenultimateAdvHour", data.PenultimateAdvHour},
		{"PenultimateAdvMinute", data.PenultimateAdvMinute},
		{"PenultimateAdvType", data.PenultimateAdvType},
		{"LastAdvYearOfDay", data.LastAdvYearOfDay},
		{"LastAdvYear", data.LastAdvYear},
		{"LastVideoClickHour", data.LastAdvClickHour},
		{"LastAdvClickMinute", data.LastAdvClickMinute},
		{"LastAdvType", data.LastAdvType},
		{"FirstHalfHourAdvClickCount", data.FirstHalfHourAdvClickCount},
		{"FirstHourAdvClickCount", data.FirstHourAdvClickCount},
		{"FirstTwoHourAdvClickCount", data.FirstTwoHourAdvClickCount},
		{"FirstThreeHourAdvClickCount", data.FirstThreeHourAdvClickCount},
		{"FirstSixHourAdvClickCount", data.FirstSixHourAdvClickCount},
		{"FirstTwelveHourAdvClickCount", data.FirstTwelveHourAdvClickCount},
		{"FirstDayAdvClickCount", data.FirstDayAdvClickCount},
		{"SecondDayAdvClickCount", data.SecondDayAdvClickCount},
		{"ThirdDayAdvClickCount", data.ThirdDayAdvClickCount},
		{"FourthDayAdvClickCount", data.FourthDayAdvClickCount},
		{"FifthDayAdvClickCount", data.FifthDayAdvClickCount},
		{"SixthDayAdvClickCount", data.SixthDayAdvClickCount},
		{"SeventhDayAdvClickCount", data.SeventhDayAdvClickCount},
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
		"ClientId", ClientId,
	}}, update)
	if updateResult.Err() != nil {
		return updateResult.Err()
	}
	return nil
}
