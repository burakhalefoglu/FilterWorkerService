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

type mdbDBuyingEventDal struct {
	Client *mongo.Client
}

func MdbDBuyingEventDalConstructor() *mdbDBuyingEventDal {
	return &mdbDBuyingEventDal{Client: mongodb.GetMongodbClient()}
}

func (m *mdbDBuyingEventDal) Add(data *model.BuyingEventResponseModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := m.Client.Database("MLDatabase").Collection("buyingEventModels")
	var _, err = collection.InsertOne(ctx, bson.D{
		{"ClientId", data.ClientId},
		{"ProjectId", data.ProjectId},
		{"CustomerId", data.CustomerId},
		{"LevelIndex", data.LevelIndex},
		{"TotalBuyingCount", data.TotalBuyingCount},
		{"TotalBuyingDay", data.TotalBuyingDay},
		{"TotalBuyingHour", data.TotalBuyingHour},
		{"FirstBuyingYearOfDay", data.FirstBuyingYearOfDay},
		{"FirstBuyingYear", data.FirstBuyingYear},
		{"FirstBuyingHour", data.FirstBuyingHour},
		{"FirstBuyingMinute", data.FirstBuyingMinute},
		{"FirstBuyingProductType", data.FirstBuyingProductType},
		{"SecondBuyingYearOfDay", data.SecondBuyingYearOfDay},
		{"SecondBuyingHour", data.SecondBuyingHour},
		{"SecondBuyingMinute", data.SecondBuyingMinute},
		{"SecondBuyingProductType", data.SecondBuyingProductType},
		{"ThirdBuyingYearOfDay", data.ThirdBuyingYearOfDay},
		{"ThirdBuyingHour", data.ThirdBuyingHour},
		{"ThirdBuyingMinute", data.ThirdBuyingMinute},
		{"ThirdBuyingProductType", data.ThirdBuyingProductType},

		{"FourthBuyingYearOfDay", data.FourthBuyingYearOfDay},
		{"FourthBuyingHour", data.FourthBuyingHour},
		{"FourthBuyingMinute", data.FourthBuyingMinute},
		{"FourthBuyingProductType", data.FourthBuyingProductType},

		{"FifthBuyingYearOfDay", data.FifthBuyingYearOfDay},
		{"FifthBuyingHour", data.FifthBuyingHour},
		{"FifthBuyingMinute", data.FifthBuyingMinute},
		{"FifthBuyingProductType", data.FifthBuyingProductType},

		{"PenultimateBuyingYearOfDay", data.PenultimateBuyingYearOfDay},
		{"PenultimateBuyingHour", data.PenultimateBuyingHour},
		{"PenultimateBuyingMinute", data.PenultimateBuyingMinute},
		{"PenultimateBuyingProductType", data.PenultimateBuyingProductType},
		{"LastBuyingYearOfDay", data.LastBuyingYearOfDay},
		{"LastBuyingYear", data.LastBuyingYear},
		{"LastBuyingHour", data.LastBuyingHour},
		{"LastBuyingMinute", data.LastBuyingMinute},
		{"LastBuyingProductType", data.LastBuyingProductType},

		{"FirstDayBuyingCount", data.FirstDayBuyingCount},
		{"SecondDayBuyingCount", data.SecondDayBuyingCount},
		{"ThirdDayBuyingCount", data.ThirdDayBuyingCount},
		{"FourthDayBuyingCount", data.FourthDayBuyingCount},
		{"FifthDayBuyingCount", data.FifthDayBuyingCount},
		{"SixthDayBuyingCount", data.SixthDayBuyingCount},
		{"SeventhDayBuyingCount", data.SeventhDayBuyingCount},

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

func (m *mdbDBuyingEventDal) GetBuyingEventById(ClientId string) (*model.BuyingEventResponseModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := m.Client.Database("MLDatabase").Collection("BuyingEventModel")
	var result = collection.FindOne(ctx, bson.D{{
		"ClientId", ClientId,
	}})

	var model = model.BuyingEventResponseModel{}
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

func (m *mdbDBuyingEventDal) UpdateBuyingEventById(ClientId string, data *model.BuyingEventResponseModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.D{{"$set", bson.D{
		{"ClientId", data.ClientId},
		{"ProjectId", data.ProjectId},
		{"CustomerId", data.CustomerId},
		{"LevelIndex", data.LevelIndex},
		{"TotalBuyingCount", data.TotalBuyingCount},
		{"TotalBuyingDay", data.TotalBuyingDay},
		{"TotalBuyingHour", data.TotalBuyingHour},
		{"FirstBuyingYearOfDay", data.FirstBuyingYearOfDay},
		{"FirstBuyingYear", data.FirstBuyingYear},
		{"FirstBuyingHour", data.FirstBuyingHour},
		{"FirstBuyingMinute", data.FirstBuyingMinute},
		{"FirstBuyingProductType", data.FirstBuyingProductType},
		{"SecondBuyingYearOfDay", data.SecondBuyingYearOfDay},
		{"SecondBuyingHour", data.SecondBuyingHour},
		{"SecondBuyingMinute", data.SecondBuyingMinute},
		{"SecondBuyingProductType", data.SecondBuyingProductType},
		{"ThirdBuyingYearOfDay", data.ThirdBuyingYearOfDay},
		{"ThirdBuyingHour", data.ThirdBuyingHour},
		{"ThirdBuyingMinute", data.ThirdBuyingMinute},
		{"ThirdBuyingProductType", data.ThirdBuyingProductType},

		{"FourthBuyingYearOfDay", data.FourthBuyingYearOfDay},
		{"FourthBuyingHour", data.FourthBuyingHour},
		{"FourthBuyingMinute", data.FourthBuyingMinute},
		{"FourthBuyingProductType", data.FourthBuyingProductType},

		{"FifthBuyingYearOfDay", data.FifthBuyingYearOfDay},
		{"FifthBuyingHour", data.FifthBuyingHour},
		{"FifthBuyingMinute", data.FifthBuyingMinute},
		{"FifthBuyingProductType", data.FifthBuyingProductType},

		{"PenultimateBuyingYearOfDay", data.PenultimateBuyingYearOfDay},
		{"PenultimateBuyingHour", data.PenultimateBuyingHour},
		{"PenultimateBuyingMinute", data.PenultimateBuyingMinute},
		{"PenultimateBuyingProductType", data.PenultimateBuyingProductType},
		{"LastBuyingYearOfDay", data.LastBuyingYearOfDay},
		{"LastBuyingYear", data.LastBuyingYear},
		{"LastBuyingHour", data.LastBuyingHour},
		{"LastBuyingMinute", data.LastBuyingMinute},
		{"LastBuyingProductType", data.LastBuyingProductType},

		{"FirstDayBuyingCount", data.FirstDayBuyingCount},
		{"SecondDayBuyingCount", data.SecondDayBuyingCount},
		{"ThirdDayBuyingCount", data.ThirdDayBuyingCount},
		{"FourthDayBuyingCount", data.FourthDayBuyingCount},
		{"FifthDayBuyingCount", data.FifthDayBuyingCount},
		{"SixthDayBuyingCount", data.SixthDayBuyingCount},
		{"SeventhDayBuyingCount", data.SeventhDayBuyingCount},

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
