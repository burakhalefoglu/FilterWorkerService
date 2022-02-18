package cassandra_dal

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/pkg/database/cassandra"
	"fmt"

	logger "github.com/appneuroncompany/light-logger"
	"github.com/appneuroncompany/light-logger/clogger"
	"github.com/gocql/gocql"
)

type cassAdvEventDal struct {
	Client *gocql.Session
	Table  string
}

func NewCassAdvEventDal(Table string) *cassAdvEventDal {
	return &cassAdvEventDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassAdvEventDal) Add(data *model.AdvEventResponseModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO %s (id, client_id, project_id, customer_id, level_index, total_adv_day, total_adv_count, total_adv_hour, total_adv_minute, level_based_average_adv_count, average_adv_daily_click_count, first_adv_year_of_day, first_adv_year, first_week_day, first_adv_click_hour, first_adv_click_minute, first_adv_type, second_adv_year_of_day, second_adv_hour, second_adv_minute, second_adv_type, third_adv_year_of_day, third_adv_hour, third_adv_minute, third_adv_type, fourth_adv_year_of_day, fourth_adv_hour, fourth_adv_minute, fourth_adv_type, fifth_adv_year_of_day, fifth_adv_hour, fifth_adv_minute, fifth_adv_type, sixth_adv_year_of_day, sixth_adv_hour, sixth_adv_minute, sixth_adv_type, seventh_adv_year_of_day, seventh_adv_hour, seventh_adv_minute, seventh_adv_type, penultimate_adv_year_of_day, penultimate_adv_hour, penultimate_adv_minute, penultimate_adv_type, last_adv_year_of_day, last_adv_year, last_adv_click_hour, last_adv_click_minute, last_adv_type, first_five_minutes_adv_click_count, first_ten_minutes_adv_click_count, first_quarter_hour_adv_click_count, first_half_hour_adv_click_count, first_hour_adv_click_count, first_two_hour_adv_click_count, firs_three_hour_adv_click_count, first_six_hour_adv_click_count, first_twelve_hour_adv_click_count, first_day_adv_click_count, second_day_adv_click_count, third_day_adv_click_count, fourth_day_adv_click_count, fifth_day_adv_click_count, sixth_day_adv_click_count, seventh_day_adv_click_count, penultimate_day_adv_click_count, last_day_adv_click_count, last_minus_first_day_adv_click_count, last_minus_penultimate_day_adv_click_count, last_day_adv_click_count_minus_average_daily_adv_click_count, sunday_adv_click_count, monday_adv_click_count, tuesday_adv_click_count, wednesday_adv_click_count, thursday_adv_click_count, friday_adv_click_count, saturday_adv_click_count, am_adv_click_count, pm_adv_click_count, adv_click_0_to_5_hour_count, adv_click_6_to_11_hour_count, adv_click_12_to_17_hour_count, adv_click_18_to_23_hour_count, status) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", m.Table),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId, data.LevelIndex, data.TotalAdvDay,
		data.TotalAdvCount, data.TotalAdvHour, data.TotalAdvMinute, data.LevelBasedAverageAdvCount,
		data.AverageAdvDailyClickCount, data.FirstAdvYearOfDay, data.FirstAdvYear,
		data.FirstWeekDay, data.FirstAdvClickHour, data.FirstAdvClickMinute, data.FirstAdvType,
		data.SecondAdvYearOfDay, data.SecondAdvHour, data.SecondAdvMinute, data.SecondAdvType,
		data.ThirdAdvYearOfDay, data.ThirdAdvHour, data.ThirdAdvMinute, data.ThirdAdvType,
		data.FourthAdvYearOfDay, data.FourthAdvHour, data.FourthAdvMinute, data.FourthAdvType,
		data.FifthAdvYearOfDay, data.FifthAdvHour, data.FifthAdvMinute, data.FifthAdvType,
		data.SixthAdvYearOfDay, data.SixthAdvHour, data.SixthAdvMinute, data.SixthAdvType,
		data.SeventhAdvYearOfDay, data.SeventhAdvHour, data.SeventhAdvMinute, data.SeventhAdvType,
		data.PenultimateAdvYearOfDay, data.PenultimateAdvHour, data.PenultimateAdvMinute, data.PenultimateAdvType,
		data.LastAdvYearOfDay, data.LastAdvYear, data.LastAdvClickHour, data.LastAdvClickMinute, data.LastAdvType,
		data.FirstFiveMinutesAdvClickCount, data.FirstTenMinutesAdvClickCount,
		data.FirstQuarterHourAdvClickCount, data.FirstHalfHourAdvClickCount,
		data.FirstHourAdvClickCount, data.FirstTwoHourAdvClickCount, data.FirstThreeHourAdvClickCount,
		data.FirstSixHourAdvClickCount, data.FirstTwelveHourAdvClickCount,
		data.FirstDayAdvClickCount, data.SecondDayAdvClickCount, data.ThirdDayAdvClickCount,
		data.FourthDayAdvClickCount, data.FifthDayAdvClickCount, data.SixthDayAdvClickCount,
		data.SeventhDayAdvClickCount, data.PenultimateDayAdvClickCount, data.LastDayAdvClickCount,
		data.LastMinusFirstDayAdvClickCount, data.LastMinusPenultimateDayAdvClickCount,
		data.LastDayAdvClickCountMinusAverageDailyAdvClickCount,
		data.SundayAdvClickCount, data.MondayAdvClickCount, data.TuesdayAdvClickCount, data.WednesdayAdvClickCount,
		data.ThursdayAdvClickCount, data.FridayAdvClickCount, data.SaturdayAdvClickCount, data.AmAdvClickCount,
		data.PmAdvClickCount, data.AdvClick0To5HourCount, data.AdvClick6To11HourCount,
		data.AdvClick12To17HourCount, data.AdvClick18To23HourCount, data.Status).Exec(); err != nil {
		// clogger.Error(&logger.Messages{
		// 	"Insert adv_event_data err: ": err.Error(),
		// })
		return err
	}
	// clogger.Info(&logger.Messages{
	// 	"Insert adv_event_data  : ": "SUCCESS",
	// })
	return nil
}


func (m *cassAdvEventDal) GetById(ClientId int64, ProjectId int64) (*model.AdvEventResponseModel, error) {
	data := &model.AdvEventResponseModel{}
	if err := m.Client.Query(fmt.Sprintf("SELECT * FROM %s WHERE client_id = ? AND project_id = ? LIMIT 1", m.Table),
	ClientId, ProjectId).Scan(&data.Id, &data.ClientId, &data.ProjectId , &data.CustomerId, &data.LevelIndex, &data.TotalAdvDay, 
			&data.TotalAdvCount, &data.TotalAdvHour, &data.TotalAdvMinute, &data.LevelBasedAverageAdvCount, 
			&data.AverageAdvDailyClickCount, &data.FirstAdvYearOfDay, &data.FirstAdvYear, &data.FirstWeekDay, 
			&data.FirstAdvClickHour, &data.FirstAdvClickMinute, &data.FirstAdvType, &data.SecondAdvYearOfDay, 
			&data.SecondAdvHour, &data.SecondAdvMinute, &data.SecondAdvType, &data.ThirdAdvYearOfDay, &data.ThirdAdvHour, 
			&data.ThirdAdvMinute, &data.ThirdAdvType, &data.FourthAdvYearOfDay, &data.FourthAdvHour, 
			&data.FourthAdvMinute, &data.FourthAdvType, &data.FourthAdvType, &data.FifthAdvYearOfDay, &data.FifthAdvHour, 
			&data.FifthAdvMinute, &data.FifthAdvType, &data.SixthAdvYearOfDay, &data.SixthAdvHour, &data.SixthAdvMinute, 
			&data.SixthAdvType, &data.SeventhAdvYearOfDay, &data.SeventhAdvHour, &data.SeventhAdvMinute, 
			&data.SeventhAdvType, &data.PenultimateAdvYearOfDay, &data.PenultimateAdvHour, &data.PenultimateAdvMinute, 
			&data.PenultimateAdvType, &data.LastAdvYearOfDay, &data.LastAdvYear, &data.LastAdvClickHour, 
			&data.LastAdvClickMinute, &data.LastAdvType, &data.FirstFiveMinutesAdvClickCount, &data.FirstTenMinutesAdvClickCount, 
			&data.FirstQuarterHourAdvClickCount, &data.FirstHalfHourAdvClickCount, &data.FirstHourAdvClickCount, 
			&data.FirstTwoHourAdvClickCount, &data.FirstThreeHourAdvClickCount, &data.FirstSixHourAdvClickCount, 
			&data.FirstTwelveHourAdvClickCount, &data.FirstDayAdvClickCount, &data.SecondDayAdvClickCount, &data.ThirdDayAdvClickCount, 
			&data.FourthDayAdvClickCount, &data.FifthDayAdvClickCount, &data.SixthDayAdvClickCount, &data.SeventhDayAdvClickCount, 
			&data.PenultimateDayAdvClickCount, &data.LastDayAdvClickCount, &data.LastMinusFirstDayAdvClickCount, 
			&data.LastMinusPenultimateDayAdvClickCount, &data.LastDayAdvClickCountMinusAverageDailyAdvClickCount, 
			&data.SundayAdvClickCount, &data.MondayAdvClickCount, &data.TuesdayAdvClickCount, &data.WednesdayAdvClickCount, 
			&data.ThursdayAdvClickCount, &data.FridayAdvClickCount, &data.SaturdayAdvClickCount, &data.AmAdvClickCount, 
			&data.PmAdvClickCount, &data.AdvClick0To5HourCount, &data.AdvClick6To11HourCount, &data.AdvClick12To17HourCount, 
			&data.AdvClick18To23HourCount, &data.Status); err != nil {
			clogger.Error(&logger.Messages{
				"Get adv_event_data err: ": err.Error(),
			})
		return nil, err
	}
	clogger.Info(&logger.Messages{
		"Get adv_event_data  : ": "SUCCESS",
	})
	return data, nil
}


             

func (m *cassAdvEventDal) UpdateById(ClientId int64, ProjectId int64, data *model.AdvEventResponseModel) error {
	if err := m.Client.Query(fmt.Sprintf("UPDATE %s SET id=?, customer_id=?, level_index=?, total_adv_day=?, total_adv_count=?, total_adv_hour=?, total_adv_minute=?, level_based_average_adv_count=?, average_adv_daily_click_count=?, first_adv_year_of_day=?, first_adv_year=?, first_week_day=?, first_adv_click_hour=?, first_adv_click_minute=?, first_adv_type=?, second_adv_year_of_day=?, second_adv_hour=?, second_adv_minute=?, second_adv_type=?, third_adv_year_of_day=?, third_adv_hour=?, third_adv_minute=?, third_adv_type=?, fourth_adv_year_of_day=?, fourth_adv_hour=?, fourth_adv_minute=?, fourth_adv_type=?, fifth_adv_year_of_day=?, fifth_adv_hour=?, fifth_adv_minute=?, fifth_adv_type=?, sixth_adv_year_of_day=?, sixth_adv_hour=?, sixth_adv_minute=?, sixth_adv_type=?, seventh_adv_year_of_day=?, seventh_adv_hour=?, seventh_adv_minute=?, seventh_adv_type=?, penultimate_adv_year_of_day=?, penultimate_adv_hour=?, penultimate_adv_minute=?, penultimate_adv_type=?, last_adv_year_of_day=?, last_adv_year=?, last_adv_click_hour=?, last_adv_click_minute=?, last_adv_type=?, first_five_minutes_adv_click_count=?, first_ten_minutes_adv_click_count=?, first_quarter_hour_adv_click_count=?, first_half_hour_adv_click_count=?, first_hour_adv_click_count=?, first_two_hour_adv_click_count=?, firs_three_hour_adv_click_count=?, first_six_hour_adv_click_count=?, first_twelve_hour_adv_click_count=?, first_day_adv_click_count=?, second_day_adv_click_count=?, third_day_adv_click_count=?, fourth_day_adv_click_count=?, fifth_day_adv_click_count=?, sixth_day_adv_click_count=?, seventh_day_adv_click_count=?, penultimate_day_adv_click_count=?, last_day_adv_click_count=?, last_minus_first_day_adv_click_count=?, last_minus_penultimate_day_adv_click_count=?, last_day_adv_click_count_minus_average_daily_adv_click_count=?, sunday_adv_click_count=?, monday_adv_click_count=?, tuesday_adv_click_count=?, wednesday_adv_click_count=?, thursday_adv_click_count=?, friday_adv_click_count=?, saturday_adv_click_count=?, am_adv_click_count=?, pm_adv_click_count=?, adv_click_0_to_5_hour_count=?, adv_click_6_to_11_hour_count=?, adv_click_12_to_17_hour_count=?, adv_click_18_to_23_hour_count=?, status=? WHERE client_id = %d AND project_id = %d", m.Table, ClientId, ProjectId),
		data.Id, data.CustomerId, data.LevelIndex, data.TotalAdvDay, data.TotalAdvCount, data.TotalAdvHour, 
		data.TotalAdvMinute, data.LevelBasedAverageAdvCount, data.AverageAdvDailyClickCount, data.FirstAdvYearOfDay, 
		data.FirstAdvYear, data.FirstWeekDay, data.FirstAdvClickHour, data.FirstAdvClickMinute, data.FirstAdvType, 
		data.SecondAdvYearOfDay, data.SecondAdvHour, data.SecondAdvMinute, data.SecondAdvType, 
		data.ThirdAdvYearOfDay, data.ThirdAdvHour, data.ThirdAdvMinute, data.ThirdAdvType, data.FourthAdvYearOfDay, 
		data.FourthAdvHour, data.FourthAdvMinute, data.FourthAdvType, data.FourthAdvType, data.FifthAdvYearOfDay, 
		data.FifthAdvHour, data.FifthAdvMinute, data.FifthAdvType, data.SixthAdvYearOfDay, data.SixthAdvHour, 
		data.SixthAdvMinute, data.SixthAdvType, data.SeventhAdvYearOfDay, data.SeventhAdvHour, 
		data.SeventhAdvMinute, data.SeventhAdvType, data.PenultimateAdvYearOfDay, data.PenultimateAdvHour, 
		data.PenultimateAdvMinute, data.PenultimateAdvType, data.LastAdvYearOfDay, data.LastAdvYear, 
		data.LastAdvClickHour, data.LastAdvClickMinute, data.LastAdvType, data.FirstFiveMinutesAdvClickCount, 
		data.FirstTenMinutesAdvClickCount, data.FirstQuarterHourAdvClickCount, data.FirstHalfHourAdvClickCount, 
		data.FirstHourAdvClickCount, data.FirstTwoHourAdvClickCount, data.FirstThreeHourAdvClickCount, 
		data.FirstSixHourAdvClickCount, data.FirstTwelveHourAdvClickCount, data.FirstDayAdvClickCount, 
		data.SecondDayAdvClickCount, data.ThirdDayAdvClickCount, data.FourthDayAdvClickCount, 
		data.FifthDayAdvClickCount, data.SixthDayAdvClickCount, data.SeventhDayAdvClickCount, 
		data.PenultimateDayAdvClickCount, data.LastDayAdvClickCount, data.LastMinusFirstDayAdvClickCount, 
		data.LastMinusPenultimateDayAdvClickCount, data.LastDayAdvClickCountMinusAverageDailyAdvClickCount, 
		data.SundayAdvClickCount, data.MondayAdvClickCount, data.TuesdayAdvClickCount, data.WednesdayAdvClickCount, 
		data.ThursdayAdvClickCount, data.FridayAdvClickCount, data.SaturdayAdvClickCount, data.AmAdvClickCount, 
		data.PmAdvClickCount, data.AdvClick0To5HourCount, data.AdvClick6To11HourCount, data.AdvClick12To17HourCount, 
		data.AdvClick18To23HourCount, data.Status).Exec(); err != nil {
		
			return err
	}

	
	return nil
}






// func (m *cassAdvEventDal) GetById(ClientId int64) (*[]model.AdvEventResponseModel, error) {
// 	var AdvEventResponseModel []model.AdvEventResponseModel
// 	 c := map[string]interface{}{}
// 	iter := m.Client.Query(fmt.Sprintf("SELECT * FROM %s WHERE client_id=?", m.Table), ClientId).Iter()
// 	for iter.MapScan(c) {
// 		AdvEventResponseModel = append(AdvEventResponseModel, model.AdvEventResponseModel{
// 			Id:                                   c["id"].(int64),
// 			ClientId:                             c["client_id"].(int64),
// 			ProjectId:                            c["project_id"].(int64),
// 			CustomerId:                           c["customer_id"].(int64),
// 			LevelIndex:                           c["level_index"].(int16),
// 			TotalAdvDay:                          c["total_adv_day"].(int32),
// 			TotalAdvCount:                        c["total_adv_count"].(int32),
// 			TotalAdvHour:                         c["total_adv_hour"].(int32),
// 			TotalAdvMinute:                       c["total_adv_minute"].(int32),
// 			LevelBasedAverageAdvCount:            c["level_based_average_adv_count"].(float32),
// 			AverageAdvDailyClickCount:            c["average_adv_daily_click_count"].(float32),
// 			FirstAdvYearOfDay:                    c["first_adv_year_of_day"].(int16),
// 			FirstAdvYear:                         c["first_adv_year"].(int16),
// 			FirstWeekDay:                         c["first_week_day"].(int16),
// 			FirstAdvClickHour:                    c["first_adv_click_hour"].(int16),
// 			FirstAdvClickMinute:                  c["first_adv_click_minute"].(int16),
// 			FirstAdvType:                         c["first_adv_type"].(byte),
// 			SecondAdvYearOfDay:                   c["second_adv_year_of_day"].(int16),
// 			SecondAdvHour:                        c["second_adv_hour"].(int16),
// 			SecondAdvMinute:                      c["second_adv_minute"].(int16),
// 			SecondAdvType:                        c["second_adv_type"].(byte),
// 			ThirdAdvYearOfDay:                    c["third_adv_year_of_day"].(int16),
// 			ThirdAdvHour:                         c["third_adv_hour"].(int16),
// 			ThirdAdvMinute:                       c["third_adv_minute"].(int16),
// 			ThirdAdvType:                         c["third_adv_type"].(byte),
// 			FourthAdvYearOfDay:                   c["fourth_adv_year_of_day"].(int16),
// 			FourthAdvHour:                        c["fourth_adv_hour"].(int16),
// 			FourthAdvMinute:                      c["fourth_adv_minute"].(int16),
// 			FourthAdvType:                        c["fourth_adv_type"].(byte),
// 			FifthAdvYearOfDay:                    c["fifth_adv_year_of_day"].(int16),
// 			FifthAdvHour:                         c["fifth_adv_hour"].(int16),
// 			FifthAdvMinute:                       c["fifth_adv_minute"].(int16),
// 			FifthAdvType:                         c["fifth_adv_type"].(byte),
// 			SixthAdvYearOfDay:                    c["sixth_adv_year_of_day"].(int16),
// 			SixthAdvHour:                         c["sixth_adv_hour"].(int16),
// 			SixthAdvMinute:                       c["sixth_adv_minute"].(int16),
// 			SixthAdvType:                         c["sixth_adv_type"].(byte),
// 			SeventhAdvYearOfDay:                  c["seventh_adv_year_of_day"].(int16),
// 			SeventhAdvHour:                       c["seventh_adv_hour"].(int16),
// 			SeventhAdvMinute:                     c["seventh_adv_minute"].(int16),
// 			SeventhAdvType:                       c["seventh_adv_type"].(byte),
// 			PenultimateAdvYearOfDay:              c["penultimate_adv_year_of_day"].(int16),
// 			PenultimateAdvHour:                   c["penultimate_adv_hour"].(int16),
// 			PenultimateAdvMinute:                 c["penultimate_adv_minute"].(int16),
// 			PenultimateAdvType:                   c["penultimate_adv_type"].(byte),
// 			LastAdvYearOfDay:                     c["last_adv_year_of_day"].(int16),
// 			LastAdvYear:                          c["last_adv_year"].(int16),
// 			LastAdvClickHour:                     c["last_adv_click_hour"].(int16),
// 			LastAdvClickMinute:                   c["last_adv_click_minute"].(int16),
// 			LastAdvType:                          c["last_adv_type"].(byte),
// 			FirstFiveMinutesAdvClickCount:        c["first_five_minutes_adv_click_count"].(int16),
// 			FirstTenMinutesAdvClickCount:         c["first_ten_minutes_adv_click_count"].(int16),
// 			FirstQuarterHourAdvClickCount:        c["first_quarter_hour_adv_click_count"].(int16),
// 			FirstHalfHourAdvClickCount:           c["first_half_hour_adv_click_count"].(int16),
// 			FirstHourAdvClickCount:               c["first_hour_adv_click_count"].(int16),
// 			FirstTwoHourAdvClickCount:            c["first_two_hour_adv_click_count"].(int16),
// 			FirstThreeHourAdvClickCount:          c["firs_three_hour_adv_click_count"].(int16),
// 			FirstSixHourAdvClickCount:            c["first_six_hour_adv_click_count"].(int16),
// 			FirstTwelveHourAdvClickCount:         c["first_twelve_hour_adv_click_count"].(int16),
// 			FirstDayAdvClickCount:                c["first_day_adv_click_count"].(int16),
// 			SecondDayAdvClickCount:               c["second_day_adv_click_count"].(int16),
// 			ThirdDayAdvClickCount:                c["third_day_adv_click_count"].(int16),
// 			FourthDayAdvClickCount:               c["fourth_day_adv_click_count"].(int16),
// 			FifthDayAdvClickCount:                c["fifth_day_adv_click_count"].(int16),
// 			SixthDayAdvClickCount:                c["sixth_day_adv_click_count"].(int16),
// 			SeventhDayAdvClickCount:              c["seventh_day_adv_click_count"].(int16),
// 			PenultimateDayAdvClickCount:          c["penultimate_day_adv_click_count"].(int16),
// 			LastDayAdvClickCount:                 c["last_day_adv_click_count"].(int16),
// 			LastMinusFirstDayAdvClickCount:       c["last_minus_first_day_adv_click_count"].(int16),
// 			LastMinusPenultimateDayAdvClickCount: c["last_minus_penultimate_day_adv_click_count"].(int16),
// 			LastDayAdvClickCountMinusAverageDailyAdvClickCount: c["last_day_adv_click_count_minus_average_daily_adv_click_count"].(float32),
// 			SundayAdvClickCount:     c["sunday_adv_click_count"].(int16),
// 			MondayAdvClickCount:     c["monday_adv_click_count"].(int16),
// 			TuesdayAdvClickCount:    c["tuesday_adv_click_count"].(int16),
// 			WednesdayAdvClickCount:  c["wednesday_adv_click_count"].(int16),
// 			ThursdayAdvClickCount:   c["thursday_adv_click_count"].(int16),
// 			FridayAdvClickCount:     c["friday_adv_click_count"].(int16),
// 			SaturdayAdvClickCount:   c["saturday_adv_click_count"].(int16),
// 			AmAdvClickCount:         c["am_adv_click_count"].(int16),
// 			PmAdvClickCount:         c["pm_adv_click_count"].(int16),
// 			AdvClick0To5HourCount:   c["adv_click_0_to_5_hour_count"].(int16),
// 			AdvClick6To11HourCount:  c["adv_click_6_to_11_hour_count"].(int16),
// 			AdvClick12To17HourCount: c["adv_click_12_to_17_hour_count"].(int16),
// 			AdvClick18To23HourCount: c["adv_click_18_to_23_hour_count"].(int16),
// 			Status:                  c["status"].(bool),
// 		})
// 		c = map[string]interface{}{}
// 	}
// 	return &AdvEventResponseModel, nil

// }
