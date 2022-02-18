package cassandra_dal

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/pkg/database/cassandra"
	"fmt"

	//logger "github.com/appneuroncompany/light-logger"
	//"github.com/appneuroncompany/light-logger/clogger"
	"github.com/gocql/gocql"
)

type cassScreenClickDal struct {
	Client *gocql.Session
	Table  string
}

func NewCassScreenClickDal(Table string) *cassScreenClickDal {
	return &cassScreenClickDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassScreenClickDal) Add(data *model.ScreenClickResponseModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO %s (id, client_id, project_id, customer_id, level_index ,first_click_session_year_of_day ,first_click_session_year ,first_click_session_hour ,first_click_session_minute ,first_touch_count ,second_click_session_hour ,second_click_session_minute ,second_touch_count ,third_click_session_hour ,third_click_session_minute ,third_touch_count ,fourth_click_session_hour ,fourth_click_session_minute ,fourth_touch_count ,fifth_click_session_hour ,fifth_click_session_minute ,fifth_touch_count ,sixth_click_session_hour ,sixth_click_session_minute , sixth_touch_count ,seventh_click_session_hour , seventh_click_session_minute ,seventh_touch_count ,penultimate_click_session_hour ,penultimate_click_session_minute ,penultimate_touch_count ,last_click_session_year_of_day ,last_click_session_year ,last_click_session_hour ,last_click_session_minute ,last_touch_count ,first_start_x_cor ,first_start_y_cor ,first_finish_x_cor ,first_finish_y_cor ,second_start_x_cor ,second_start_y_cor ,second_finish_x_cor ,second_finish_y_cor ,third_start_x_cor ,third_start_y_cor ,third_finish_x_cor ,third_finish_y_cor ,fourth_start_x_cor ,fourth_start_y_cor ,fourth_finish_x_cor ,fourth_finish_y_cor ,fifth_start_x_cor ,fifth_start_y_cor ,fifth_finish_x_cor ,fifth_finish_y_cor ,sixth_start_x_cor ,  sixth_start_y_cor ,   sixth_finish_x_cor , sixth_finish_y_cor , seventh_start_x_cor ,seventh_start_y_cor ,seventh_finish_x_cor ,seventh_finish_y_cor ,penultimate_start_x_cor ,penultimate_start_y_cor ,penultimate_finish_x_cor ,penultimate_finish_y_cor ,last_start_x_cor ,last_start_y_cor ,last_finish_x_cor ,last_finish_y_cor ,first_half_hour_touch_count ,first_hour_touch_count ,first_two_hour_touch_count ,first_three_hour_touch_count ,first_six_hour_touch_count ,first_twelve_hour_touch_count ,first_minus_last_touch_count ,first_finger_id ,penultimate_finger_id ,last_finger_id ,first_day_click_count ,second_day_click_count ,third_day_click_count ,fourth_day_click_count ,fifth_day_click_count ,sixth_day_click_count ,seventh_day_click_count ,total_click_day ,total_click_count ,total_click_session_count ,total_click_hour ,total_click_minute ,total_start_x_cor ,total_start_y_cor ,total_finish_x_cor ,total_finish_y_cor ,session_based_avegare_start_x_cor ,session_based_avegare_start_y_cor ,session_based_avegare_finish_x_cor ,session_based_avegare_finish_y_cor ,session_based_avegare_click_count ,daily_avegare_click_count ,last_touch_count_minus_session_based_avegare_click_count ,  status) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", m.Table),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId,
		data.LevelIndex,
		data.FirstClickSessionYearOfDay,
		data.FirstClickSessionYear,
		data.FirstClickSessionHour,
		data.FirstClickSessionMinute,
		data.FirstTouchCount,
		data.SecondClickSessionHour,
		data.SecondClickSessionMinute,
		data.SecondTouchCount,
		data.ThirdClickSessionHour,
		data.ThirdClickSessionMinute,
		data.ThirdTouchCount,
		data.FourthClickSessionHour,
		data.FourthClickSessionMinute,
		data.FourthTouchCount,
		data.FifthClickSessionHour,
		data.FifthClickSessionMinute,
		data.FifthTouchCount,
		data.SixthClickSessionHour,
		data.SixthClickSessionMinute,
		data.SixthTouchCount,
		data.SeventhClickSessionHour,
		data.SeventhClickSessionMinute,
		data.SeventhTouchCount,
		data.PenultimateClickSessionHour,
		data.PenultimateClickSessionMinute,
		data.PenultimateTouchCount,
		data.LastClickSessionYearOfDay,
		data.LastClickSessionYear,
		data.LastClickSessionHour,
		data.LastClickSessionMinute,
		data.LastTouchCount,
		data.FirstStartXCor,
		data.FirstStartYCor,
		data.FirstFinishXCor,
		data.FirstFinishYCor,
		data.SecondStartXCor,
		data.SecondStartYCor,
		data.SecondFinishXCor,
		data.SecondFinishYCor,
		data.ThirdStartXCor,
		data.ThirdStartYCor,
		data.ThirdFinishXCor,
		data.ThirdFinishYCor,
		data.FourthStartXCor,
		data.FourthStartYCor,
		data.FourthFinishXCor,
		data.FourthFinishYCor,
		data.FifthStartXCor,
		data.FifthStartYCor,
		data.FifthFinishXCor,
		data.FifthFinishYCor,
		data.SixthStartXCor,
		data.SixthStartYCor,
		data.SixthFinishXCor,
		data.SixthFinishYCor,
		data.SeventhStartXCor,
		data.SeventhStartYCor,
		data.SeventhFinishXCor,
		data.SeventhFinishYCor,
		data.PenultimateStartXCor,
		data.PenultimateStartYCor,
		data.PenultimateFinishXCor,
		data.PenultimateFinishYCor,
		data.LastStartXCor,
		data.LastStartYCor,
		data.LastFinishXCor,
		data.LastFinishYCor,
		data.FirstFiveMinutesTouchCount,
		data.FirstTenMinutesTouchCount,
		data.FirstQuarterHourTouchCount,
		data.FirstHalfHourTouchCount,
		data.FirstHourTouchCount,
		data.FirstTwoHourTouchCount,
		data.FirstThreeHourTouchCount,
		data.FirstSixHourTouchCount,
		data.FirstTwelveHourTouchCount,
		data.FirstMinusLastTouchCount,
		data.FirstFingerId,
		data.PenultimateFingerId,
		data.LastFingerId,
		data.FirstDayClickCount,
		data.SecondDayClickCount,
		data.ThirdDayClickCount,
		data.FourthDayClickCount,
		data.FifthDayClickCount,
		data.SixthDayClickCount,
		data.SeventhDayClickCount,
		data.TotalClickDay,
		data.TotalClickCount,
		data.TotalClickSessionCount,
		data.TotalClickHour,
		data.TotalClickMinute,
		data.TotalStartXCor,
		data.TotalStartYCor,
		data.TotalFinishXCor,
		data.TotalFinishYCor,
		data.SessionBasedAvegareStartXCor,
		data.SessionBasedAvegareStartYCor,
		data.SessionBasedAvegareFinishXCor,
		data.SessionBasedAvegareFinishYCor,
		data.SessionBasedAvegareClickCount,
		data.DailyAvegareClickCount,
		data.LastTouchCountMinusSessionBasedAvegareClickCount,
		data.Status).Exec(); err != nil {
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

func (m *cassScreenClickDal) GetById(ClientId int64, ProjectId int64) (*model.ScreenClickResponseModel, error) {
	data := &model.ScreenClickResponseModel{}
	if err := m.Client.Query(fmt.Sprintf("SELECT * FROM %s WHERE client_id = ? AND project_id = ? LIMIT 1", m.Table),
		ClientId, ProjectId).Scan(&data.Id, &data.ClientId, &data.ProjectId, &data.CustomerId,
		&data.LevelIndex,
		&data.FirstClickSessionYearOfDay,
		&data.FirstClickSessionYear,
		&data.FirstClickSessionHour,
		&data.FirstClickSessionMinute,
		&data.FirstTouchCount,
		&data.SecondClickSessionHour,
		&data.SecondClickSessionMinute,
		&data.SecondTouchCount,
		&data.ThirdClickSessionHour,
		&data.ThirdClickSessionMinute,
		&data.ThirdTouchCount,
		&data.FourthClickSessionHour,
		&data.FourthClickSessionMinute,
		&data.FourthTouchCount,
		&data.FifthClickSessionHour,
		&data.FifthClickSessionMinute,
		&data.FifthTouchCount,
		&data.SixthClickSessionHour,
		&data.SixthClickSessionMinute,
		&data.SixthTouchCount,
		&data.SeventhClickSessionHour,
		&data.SeventhClickSessionMinute,
		&data.SeventhTouchCount,
		&data.PenultimateClickSessionHour,
		&data.PenultimateClickSessionMinute,
		&data.PenultimateTouchCount,
		&data.LastClickSessionYearOfDay,
		&data.LastClickSessionYear,
		&data.LastClickSessionHour,
		&data.LastClickSessionMinute,
		&data.LastTouchCount,
		&data.FirstStartXCor,
		&data.FirstStartYCor,
		&data.FirstFinishXCor,
		&data.FirstFinishYCor,
		&data.SecondStartXCor,
		&data.SecondStartYCor,
		&data.SecondFinishXCor,
		&data.SecondFinishYCor,
		&data.ThirdStartXCor,
		&data.ThirdStartYCor,
		&data.ThirdFinishXCor,
		&data.ThirdFinishYCor,
		&data.FourthStartXCor,
		&data.FourthStartYCor,
		&data.FourthFinishXCor,
		&data.FourthFinishYCor,
		&data.FifthStartXCor,
		&data.FifthStartYCor,
		&data.FifthFinishXCor,
		&data.FifthFinishYCor,
		&data.SixthStartXCor,
		&data.SixthStartYCor,
		&data.SixthFinishXCor,
		&data.SixthFinishYCor,
		&data.SeventhStartXCor,
		&data.SeventhStartYCor,
		&data.SeventhFinishXCor,
		&data.SeventhFinishYCor,
		&data.PenultimateStartXCor,
		&data.PenultimateStartYCor,
		&data.PenultimateFinishXCor,
		&data.PenultimateFinishYCor,
		&data.LastStartXCor,
		&data.LastStartYCor,
		&data.LastFinishXCor,
		&data.LastFinishYCor,
		&data.FirstFiveMinutesTouchCount,
		&data.FirstTenMinutesTouchCount,
		&data.FirstQuarterHourTouchCount,
		&data.FirstHalfHourTouchCount,
		&data.FirstHourTouchCount,
		&data.FirstTwoHourTouchCount,
		&data.FirstThreeHourTouchCount,
		&data.FirstSixHourTouchCount,
		&data.FirstTwelveHourTouchCount,
		&data.FirstMinusLastTouchCount,
		&data.FirstFingerId,
		&data.PenultimateFingerId,
		&data.LastFingerId,
		&data.FirstDayClickCount,
		&data.SecondDayClickCount,
		&data.ThirdDayClickCount,
		&data.FourthDayClickCount,
		&data.FifthDayClickCount,
		&data.SixthDayClickCount,
		&data.SeventhDayClickCount,
		&data.TotalClickDay,
		&data.TotalClickCount,
		&data.TotalClickSessionCount,
		&data.TotalClickHour,
		&data.TotalClickMinute,
		&data.TotalStartXCor,
		&data.TotalStartYCor,
		&data.TotalFinishXCor,
		&data.TotalFinishYCor,
		&data.SessionBasedAvegareStartXCor,
		&data.SessionBasedAvegareStartYCor,
		&data.SessionBasedAvegareFinishXCor,
		&data.SessionBasedAvegareFinishYCor,
		&data.SessionBasedAvegareClickCount,
		&data.DailyAvegareClickCount,
		&data.LastTouchCountMinusSessionBasedAvegareClickCount, &data.Status); err != nil {
		// clogger.Error(&logger.Messages{
		// 	"Get adv_event_data err: ": err.Error(),
		// })
		return nil, err
	}
	// clogger.Info(&logger.Messages{
	// 	"Get adv_event_data  : ": "SUCCESS",
	// })
	return data, nil
}

func (m *cassScreenClickDal) UpdateById(ClientId int64, ProjectId int64, data *model.ScreenClickResponseModel) error {
	if err := m.Client.Query(fmt.Sprintf("UPDATE %s SET id=?, customer_id=?, level_index=? ,first_click_session_year_of_day=? ,first_click_session_year=? ,first_click_session_hour=? ,first_click_session_minute=? ,first_touch_count=? ,second_click_session_hour=? ,second_click_session_minute=? ,second_touch_count=? ,third_click_session_hour=? ,third_click_session_minute=? ,third_touch_count=? ,fourth_click_session_hour=? ,fourth_click_session_minute=? ,fourth_touch_count=? ,fifth_click_session_hour=? ,fifth_click_session_minute=? ,fifth_touch_count=? ,sixth_click_session_hour=? ,sixth_click_session_minute=? , sixth_touch_count=? ,seventh_click_session_hour=? , seventh_click_session_minute=? ,seventh_touch_count=? ,penultimate_click_session_hour=? ,penultimate_click_session_minute=? ,penultimate_touch_count=? ,last_click_session_year_of_day=? ,last_click_session_year=? ,last_click_session_hour=? ,last_click_session_minute=? ,last_touch_count=? ,first_start_x_cor=? ,first_start_y_cor=? ,first_finish_x_cor=? ,first_finish_y_cor=? ,second_start_x_cor=? ,second_start_y_cor=? ,second_finish_x_cor=? ,second_finish_y_cor=? ,third_start_x_cor=? ,third_start_y_cor=? ,third_finish_x_cor=? ,third_finish_y_cor=? ,fourth_start_x_cor=? ,fourth_start_y_cor=? ,fourth_finish_x_cor=? ,fourth_finish_y_cor=? ,fifth_start_x_cor=? ,fifth_start_y_cor=? ,fifth_finish_x_cor=? ,fifth_finish_y_cor=? ,sixth_start_x_cor=? ,  sixth_start_y_cor=? ,   sixth_finish_x_cor=? , sixth_finish_y_cor=? , seventh_start_x_cor=? ,seventh_start_y_cor=? ,seventh_finish_x_cor=? ,seventh_finish_y_cor=? ,penultimate_start_x_cor=? ,penultimate_start_y_cor=? ,penultimate_finish_x_cor=? ,penultimate_finish_y_cor=? ,last_start_x_cor=? ,last_start_y_cor=? ,last_finish_x_cor=? ,last_finish_y_cor=? ,first_half_hour_touch_count=? ,first_hour_touch_count=? ,first_two_hour_touch_count=? ,first_three_hour_touch_count=? ,first_six_hour_touch_count=? ,first_twelve_hour_touch_count=? ,first_minus_last_touch_count=? ,first_finger_id=? ,penultimate_finger_id=? ,last_finger_id=? ,first_day_click_count=? ,second_day_click_count=? ,third_day_click_count=? ,fourth_day_click_count=? ,fifth_day_click_count=? ,sixth_day_click_count=? ,seventh_day_click_count=? ,total_click_day=? ,total_click_count=? ,total_click_session_count=? ,total_click_hour=? ,total_click_minute=? ,total_start_x_cor=? ,total_start_y_cor=? ,total_finish_x_cor=? ,total_finish_y_cor=? ,session_based_avegare_start_x_cor=? ,session_based_avegare_start_y_cor=? ,session_based_avegare_finish_x_cor=? ,session_based_avegare_finish_y_cor=? ,session_based_avegare_click_count=? ,daily_avegare_click_count=? ,last_touch_count_minus_session_based_avegare_click_count=? , status=? WHERE client_id = %d AND project_id = %d", m.Table, ClientId, ProjectId),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId,
		data.LevelIndex,
		data.FirstClickSessionYearOfDay,
		data.FirstClickSessionYear,
		data.FirstClickSessionHour,
		data.FirstClickSessionMinute,
		data.FirstTouchCount,
		data.SecondClickSessionHour,
		data.SecondClickSessionMinute,
		data.SecondTouchCount,
		data.ThirdClickSessionHour,
		data.ThirdClickSessionMinute,
		data.ThirdTouchCount,
		data.FourthClickSessionHour,
		data.FourthClickSessionMinute,
		data.FourthTouchCount,
		data.FifthClickSessionHour,
		data.FifthClickSessionMinute,
		data.FifthTouchCount,
		data.SixthClickSessionHour,
		data.SixthClickSessionMinute,
		data.SixthTouchCount,
		data.SeventhClickSessionHour,
		data.SeventhClickSessionMinute,
		data.SeventhTouchCount,
		data.PenultimateClickSessionHour,
		data.PenultimateClickSessionMinute,
		data.PenultimateTouchCount,
		data.LastClickSessionYearOfDay,
		data.LastClickSessionYear,
		data.LastClickSessionHour,
		data.LastClickSessionMinute,
		data.LastTouchCount,
		data.FirstStartXCor,
		data.FirstStartYCor,
		data.FirstFinishXCor,
		data.FirstFinishYCor,
		data.SecondStartXCor,
		data.SecondStartYCor,
		data.SecondFinishXCor,
		data.SecondFinishYCor,
		data.ThirdStartXCor,
		data.ThirdStartYCor,
		data.ThirdFinishXCor,
		data.ThirdFinishYCor,
		data.FourthStartXCor,
		data.FourthStartYCor,
		data.FourthFinishXCor,
		data.FourthFinishYCor,
		data.FifthStartXCor,
		data.FifthStartYCor,
		data.FifthFinishXCor,
		data.FifthFinishYCor,
		data.SixthStartXCor,
		data.SixthStartYCor,
		data.SixthFinishXCor,
		data.SixthFinishYCor,
		data.SeventhStartXCor,
		data.SeventhStartYCor,
		data.SeventhFinishXCor,
		data.SeventhFinishYCor,
		data.PenultimateStartXCor,
		data.PenultimateStartYCor,
		data.PenultimateFinishXCor,
		data.PenultimateFinishYCor,
		data.LastStartXCor,
		data.LastStartYCor,
		data.LastFinishXCor,
		data.LastFinishYCor,
		data.FirstFiveMinutesTouchCount,
		data.FirstTenMinutesTouchCount,
		data.FirstQuarterHourTouchCount,
		data.FirstHalfHourTouchCount,
		data.FirstHourTouchCount,
		data.FirstTwoHourTouchCount,
		data.FirstThreeHourTouchCount,
		data.FirstSixHourTouchCount,
		data.FirstTwelveHourTouchCount,
		data.FirstMinusLastTouchCount,
		data.FirstFingerId,
		data.PenultimateFingerId,
		data.LastFingerId,
		data.FirstDayClickCount,
		data.SecondDayClickCount,
		data.ThirdDayClickCount,
		data.FourthDayClickCount,
		data.FifthDayClickCount,
		data.SixthDayClickCount,
		data.SeventhDayClickCount,
		data.TotalClickDay,
		data.TotalClickCount,
		data.TotalClickSessionCount,
		data.TotalClickHour,
		data.TotalClickMinute,
		data.TotalStartXCor,
		data.TotalStartYCor,
		data.TotalFinishXCor,
		data.TotalFinishYCor,
		data.SessionBasedAvegareStartXCor,
		data.SessionBasedAvegareStartYCor,
		data.SessionBasedAvegareFinishXCor,
		data.SessionBasedAvegareFinishYCor,
		data.SessionBasedAvegareClickCount,
		data.DailyAvegareClickCount,
		data.LastTouchCountMinusSessionBasedAvegareClickCount, data.Status).Exec(); err != nil {

		return err
	}

	return nil
}
