package cassandra_dal

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/pkg/database/cassandra"
	"fmt"

	"github.com/gocql/gocql"
)

type cassGameSessionDal struct {
	Client *gocql.Session
	Table  string
}

func NewCassGameSessionDal(Table string) *cassGameSessionDal {
	return &cassGameSessionDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassGameSessionDal) Add(data *model.GameSessionResponseModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO MLDatabase.%s(id, client_id, project_id, customer_id,first_session_year_of_day ,first_session_year ,first_session_week_day ,first_session_hour ,first_session_duration ,first_session_minute ,second_session_hour ,second_session_duration ,second_session_minute ,third_session_hour ,third_session_duration ,third_session_minute ,fourth_session_hour ,fourth_session_duration ,fourth_session_minute ,fifth_session_hour ,fifth_session_duration ,fifth_session_minute ,sixth_session_hour ,sixth_session_duration , sixth_session_minute ,   seventh_session_hour ,   seventh_session_duration ,seventh_session_minute ,penultimate_session_hour ,penultimate_session_duration ,  penultimate_session_minute ,last_session_year_of_day ,last_session_year ,last_session_hour ,last_session_duration ,last_session_minute ,last_duration_minus_penultimate_duration ,first_five_minutes_total_session_count ,  first_five_minutes_total_session_duration ,first_ten_minutes_total_session_count ,   first_ten_minutes_total_session_duration ,first_quarter_hour_total_session_count ,  first_quarter_hour_total_session_duration ,first_half_hour_total_session_count ,first_half_hour_total_session_duration ,first_hour_total_session_count ,first_hour_total_session_duration ,first_two_hour_total_session_count ,first_two_hour_total_session_duration ,first_three_hour_total_session_count ,first_three_hour_total_session_duration ,first_six_hour_total_session_count ,first_six_hour_total_session_duration ,first_twelve_hour_total_session_count ,first_twelve_hour_total_session_duration ,total_session_day ,total_session_hour ,total_session_minute ,total_session_duration ,total_session_count ,first_day_total_session_count ,first_day_total_session_duration ,second_day_total_session_count ,second_day_total_session_duration ,third_day_total_session_count ,third_day_total_session_duration ,fourth_day_total_session_count ,fourth_day_total_session_duration ,fifth_day_total_session_count ,fifth_day_total_session_duration ,sixth_day_total_session_count ,sixth_day_total_session_duration ,seventh_day_total_session_dount ,seventh_day_total_session_duration ,min_session_duration ,max_session_duration ,daily_avegare_session_count ,daily_average_session_duration ,session_based_avegare_session_duration ,daily_avegare_session_count_minus_first_day_session_count ,daily_avegare_session_duration_minus_first_day_session_duration ,   session_based_avegare_session_duration_minus_first_session_duration ,session_based_avegare_session_duration_minus_last_session_duration ,sunday_session_count ,monday_session_count ,tuesday_session_count ,wednesday_session_count ,thursday_session_count ,friday_session_count ,saturday_session_count ,am_session_count ,pm_session_count ,session_0_to_5_hour_count ,session_6_to_11_hour_count ,session_12_to_17_hour_count ,session_18_to_23_hour_count , status) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", m.Table),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId, 
		data.FirstSessionYearOfDay,
		data.FirstSessionYear,
		data.FirstSessionWeekDay,
		data.FirstSessionHour,
		data.FirstSessionDuration,
		data.FirstSessionMinute,
		data.SecondSessionHour,
		data.SecondSessionDuration,
		data.SecondSessionMinute,
		data.ThirdSessionHour,
		data.ThirdSessionDuration,
		data.ThirdSessionMinute,
		data.FourthSessionHour,
		data.FourthSessionDuration,
		data.FourthSessionMinute,
		data.FifthSessionHour,
		data.FifthSessionDuration,
		data.FifthSessionMinute,
		data.SixthSessionHour,
		data.SixthSessionDuration,
		data.SixthSessionMinute,
		data.SeventhSessionHour,
		data.SeventhSessionDuration,
		data.SeventhSessionMinute,
		data.PenultimateSessionHour,
		data.PenultimateSessionDuration,
		data.PenultimateSessionMinute,
		data.LastSessionYearOfDay,
		data.LastSessionYear,
		data.LastSessionHour,
		data.LastSessionDuration,
		data.LastSessionMinute,
		data.LastDurationMinusPenultimateDuration,
		data.FirstFiveMinutesTotalSessionCount,
		data.FirstFiveMinutesTotalSessionDuration,
		data.FirstTenMinutesTotalSessionCount,
		data.FirstTenMinutesTotalSessionDuration,
		data.FirstQuarterHourTotalSessionCount,
		data.FirstQuarterHourTotalSessionDuration,
		data.FirstHalfHourTotalSessionCount,
		data.FirstHalfHourTotalSessionDuration,
		data.FirstHourTotalSessionCount,
		data.FirstHourTotalSessionDuration,
		data.FirstTwoHourTotalSessionCount,
		data.FirstTwoHourTotalSessionDuration,
		data.FirstThreeHourTotalSessionCount,
		data.FirstThreeHourTotalSessionDuration,
		data.FirstSixHourTotalSessionCount,
		data.FirstSixHourTotalSessionDuration,
		data.FirstTwelveHourTotalSessionCount,
		data.FirstTwelveHourTotalSessionDuration,
		data.TotalSessionDay,
		data.TotalSessionHour,
		data.TotalSessionMinute,
		data.TotalSessionDuration,
		data.TotalSessionCount,
		data.FirstDayTotalSessionCount,
		data.FirstDayTotalSessionDuration,
		data.SecondDayTotalSessionCount,
		data.SecondDayTotalSessionDuration,
		data.ThirdDayTotalSessionCount,
		data.ThirdDayTotalSessionDuration,
		data.FourthDayTotalSessionCount,
		data.FourthDayTotalSessionDuration,
		data.FifthDayTotalSessionCount,
		data.FifthDayTotalSessionDuration,
		data.SixthDayTotalSessionCount,
		data.SixthDayTotalSessionDuration,
		data.SeventhDayTotalSessionCount,
		data.SeventhDayTotalSessionDuration,
		data.MinSessionDuration,
		data.MaxSessionDuration,
		data.DailyAvegareSessionCount,
		data.DailyAverageSessionDuration,
		data.SessionBasedAvegareSessionDuration,
		data.DailyAvegareSessionCountMinusFirstDaySessionCount,
		data.DailyAvegareSessionDurationMinusFirstDaySessionDuration,
		data.SessionBasedAvegareSessionDurationMinusFirstSessionDuration,
		data.SessionBasedAvegareSessionDurationMinusLastSessionDuration,
		data.SundaySessionCount,
		data.MondaySessionCount,
		data.TuesdaySessionCount,
		data.WednesdaySessionCount,
		data.ThursdaySessionCount,
		data.FridaySessionCount,
		data.SaturdaySessionCount,
		data.AmSessionCount,
		data.PmSessionCount,
		data.Session0To5HourCount,
		data.Session6To11HourCount,
		data.Session12To17HourCount,
		data.Session18To23HourCount, data.Status).Exec(); err != nil {
	
		return err
	}
	
	return nil
}

func (m *cassGameSessionDal) GetById(ClientId int64, ProjectId int64) (*model.GameSessionResponseModel, error) {
	data := &model.GameSessionResponseModel{}
	if err := m.Client.Query(fmt.Sprintf("SELECT id , client_id , project_id , customer_id , first_session_year_of_day ,first_session_year ,first_session_week_day ,first_session_hour ,first_session_duration ,first_session_minute ,second_session_hour ,second_session_duration ,second_session_minute ,third_session_hour ,third_session_duration ,third_session_minute ,fourth_session_hour ,fourth_session_duration ,fourth_session_minute ,fifth_session_hour ,fifth_session_duration ,fifth_session_minute ,sixth_session_hour ,sixth_session_duration , sixth_session_minute ,   seventh_session_hour ,   seventh_session_duration ,seventh_session_minute ,penultimate_session_hour ,penultimate_session_duration ,  penultimate_session_minute ,last_session_year_of_day ,last_session_year ,last_session_hour ,last_session_duration ,last_session_minute ,last_duration_minus_penultimate_duration ,first_five_minutes_total_session_count ,  first_five_minutes_total_session_duration ,first_ten_minutes_total_session_count ,   first_ten_minutes_total_session_duration ,first_quarter_hour_total_session_count ,  first_quarter_hour_total_session_duration ,first_half_hour_total_session_count ,first_half_hour_total_session_duration ,first_hour_total_session_count ,first_hour_total_session_duration ,first_two_hour_total_session_count ,first_two_hour_total_session_duration ,first_three_hour_total_session_count ,first_three_hour_total_session_duration ,first_six_hour_total_session_count ,first_six_hour_total_session_duration ,first_twelve_hour_total_session_count ,first_twelve_hour_total_session_duration ,total_session_day ,total_session_hour ,total_session_minute ,total_session_duration ,total_session_count ,first_day_total_session_count ,first_day_total_session_duration ,second_day_total_session_count ,second_day_total_session_duration ,third_day_total_session_count ,third_day_total_session_duration ,fourth_day_total_session_count ,fourth_day_total_session_duration ,fifth_day_total_session_count ,fifth_day_total_session_duration ,sixth_day_total_session_count ,sixth_day_total_session_duration ,seventh_day_total_session_dount ,seventh_day_total_session_duration ,min_session_duration ,max_session_duration ,daily_avegare_session_count ,daily_average_session_duration ,session_based_avegare_session_duration ,daily_avegare_session_count_minus_first_day_session_count ,daily_avegare_session_duration_minus_first_day_session_duration ,session_based_avegare_session_duration_minus_first_session_duration ,session_based_avegare_session_duration_minus_last_session_duration ,sunday_session_count ,monday_session_count ,tuesday_session_count ,wednesday_session_count ,thursday_session_count ,friday_session_count ,saturday_session_count ,am_session_count ,pm_session_count ,session_0_to_5_hour_count ,session_6_to_11_hour_count ,session_12_to_17_hour_count ,session_18_to_23_hour_count , status FROM MLDatabase.%s WHERE client_id = ? AND project_id = ? LIMIT 1", m.Table),
		ClientId, ProjectId).Scan(&data.Id, &data.ClientId, &data.ProjectId, &data.CustomerId,
			&data.FirstSessionYearOfDay,
			&data.FirstSessionYear,
			&data.FirstSessionWeekDay,
			&data.FirstSessionHour,
			&data.FirstSessionDuration,
			&data.FirstSessionMinute,
			&data.SecondSessionHour,
			&data.SecondSessionDuration,
			&data.SecondSessionMinute,
			&data.ThirdSessionHour,
			&data.ThirdSessionDuration,
			&data.ThirdSessionMinute,
			&data.FourthSessionHour,
			&data.FourthSessionDuration,
			&data.FourthSessionMinute,
			&data.FifthSessionHour,
			&data.FifthSessionDuration,
			&data.FifthSessionMinute,
			&data.SixthSessionHour,
			&data.SixthSessionDuration,
			&data.SixthSessionMinute,
			&data.SeventhSessionHour,
			&data.SeventhSessionDuration,
			&data.SeventhSessionMinute,
			&data.PenultimateSessionHour,
			&data.PenultimateSessionDuration,
			&data.PenultimateSessionMinute,
			&data.LastSessionYearOfDay,
			&data.LastSessionYear,
			&data.LastSessionHour,
			&data.LastSessionDuration,
			&data.LastSessionMinute,
			&data.LastDurationMinusPenultimateDuration,
			&data.FirstFiveMinutesTotalSessionCount,
			&data.FirstFiveMinutesTotalSessionDuration,
			&data.FirstTenMinutesTotalSessionCount,
			&data.FirstTenMinutesTotalSessionDuration,
			&data.FirstQuarterHourTotalSessionCount,
			&data.FirstQuarterHourTotalSessionDuration,
			&data.FirstHalfHourTotalSessionCount,
			&data.FirstHalfHourTotalSessionDuration,
			&data.FirstHourTotalSessionCount,
			&data.FirstHourTotalSessionDuration,
			&data.FirstTwoHourTotalSessionCount,
			&data.FirstTwoHourTotalSessionDuration,
			&data.FirstThreeHourTotalSessionCount,
			&data.FirstThreeHourTotalSessionDuration,
			&data.FirstSixHourTotalSessionCount,
			&data.FirstSixHourTotalSessionDuration,
			&data.FirstTwelveHourTotalSessionCount,
			&data.FirstTwelveHourTotalSessionDuration,
			&data.TotalSessionDay,
			&data.TotalSessionHour,
			&data.TotalSessionMinute,
			&data.TotalSessionDuration,
			&data.TotalSessionCount,
			&data.FirstDayTotalSessionCount,
			&data.FirstDayTotalSessionDuration,
			&data.SecondDayTotalSessionCount,
			&data.SecondDayTotalSessionDuration,
			&data.ThirdDayTotalSessionCount,
			&data.ThirdDayTotalSessionDuration,
			&data.FourthDayTotalSessionCount,
			&data.FourthDayTotalSessionDuration,
			&data.FifthDayTotalSessionCount,
			&data.FifthDayTotalSessionDuration,
			&data.SixthDayTotalSessionCount,
			&data.SixthDayTotalSessionDuration,
			&data.SeventhDayTotalSessionCount,
			&data.SeventhDayTotalSessionDuration,
			&data.MinSessionDuration,
			&data.MaxSessionDuration,
			&data.DailyAvegareSessionCount,
			&data.DailyAverageSessionDuration,
			&data.SessionBasedAvegareSessionDuration,
			&data.DailyAvegareSessionCountMinusFirstDaySessionCount,
			&data.DailyAvegareSessionDurationMinusFirstDaySessionDuration,
			&data.SessionBasedAvegareSessionDurationMinusFirstSessionDuration,
			&data.SessionBasedAvegareSessionDurationMinusLastSessionDuration,
			&data.SundaySessionCount,
			&data.MondaySessionCount,
			&data.TuesdaySessionCount,
			&data.WednesdaySessionCount,
			&data.ThursdaySessionCount,
			&data.FridaySessionCount,
			&data.SaturdaySessionCount,
			&data.AmSessionCount,
			&data.PmSessionCount,
			&data.Session0To5HourCount,
			&data.Session6To11HourCount,
			&data.Session12To17HourCount,
			&data.Session18To23HourCount, &data.Status); err != nil {
	
		return nil, err
	}

	return data, nil
}

func (m *cassGameSessionDal) UpdateById(ClientId int64, ProjectId int64, data *model.GameSessionResponseModel) error {
	if err := m.Client.Query(fmt.Sprintf("UPDATE MLDatabase.%s SET id=?, customer_id=?, first_session_year_of_day=? ,first_session_year=? ,first_session_week_day=? ,first_session_hour=? ,first_session_duration=? ,first_session_minute=? ,second_session_hour=? ,second_session_duration=? ,second_session_minute=? ,third_session_hour=? ,third_session_duration=? ,third_session_minute=? ,fourth_session_hour=? ,fourth_session_duration=? ,fourth_session_minute=? ,fifth_session_hour=? ,fifth_session_duration=? ,fifth_session_minute=? ,sixth_session_hour=? ,sixth_session_duration=? , sixth_session_minute=? ,   seventh_session_hour=? ,   seventh_session_duration=? ,seventh_session_minute=? ,penultimate_session_hour=? ,penultimate_session_duration=? ,  penultimate_session_minute=? ,last_session_year_of_day=? ,last_session_year=? ,last_session_hour=? ,last_session_duration=? ,last_session_minute=? ,last_duration_minus_penultimate_duration=? ,first_five_minutes_total_session_count=? ,  first_five_minutes_total_session_duration=? ,first_ten_minutes_total_session_count=? ,   first_ten_minutes_total_session_duration=? ,first_quarter_hour_total_session_count=? ,  first_quarter_hour_total_session_duration=? ,first_half_hour_total_session_count=? ,first_half_hour_total_session_duration=?,first_hour_total_session_count=? ,first_hour_total_session_duration=? ,first_two_hour_total_session_count=? ,first_two_hour_total_session_duration=? ,first_three_hour_total_session_count=? ,first_three_hour_total_session_duration=? ,first_six_hour_total_session_count=? ,first_six_hour_total_session_duration=? ,first_twelve_hour_total_session_count=? ,first_twelve_hour_total_session_duration=? ,total_session_day=? ,total_session_hour=? ,total_session_minute=? ,total_session_duration=? ,total_session_count=? ,first_day_total_session_count=? ,first_day_total_session_duration=? ,second_day_total_session_count=? ,second_day_total_session_duration=? ,third_day_total_session_count=? ,third_day_total_session_duration=? ,fourth_day_total_session_count=? ,fourth_day_total_session_duration=? ,fifth_day_total_session_count=? ,fifth_day_total_session_duration=? ,sixth_day_total_session_count=? ,sixth_day_total_session_duration=? ,seventh_day_total_session_dount=? ,seventh_day_total_session_duration=? ,min_session_duration=? ,max_session_duration=? ,daily_avegare_session_count=? ,daily_average_session_duration=? ,session_based_avegare_session_duration=? ,daily_avegare_session_count_minus_first_day_session_count=? ,daily_avegare_session_duration_minus_first_day_session_duration=? ,session_based_avegare_session_duration_minus_first_session_duration=? ,session_based_avegare_session_duration_minus_last_session_duration=? ,sunday_session_count=? ,monday_session_count=? ,tuesday_session_count=? ,wednesday_session_count=? ,thursday_session_count=? ,friday_session_count=? ,saturday_session_count=? ,am_session_count=? ,pm_session_count=? ,session_0_to_5_hour_count=? ,session_6_to_11_hour_count=? ,session_12_to_17_hour_count=? ,session_18_to_23_hour_count=? , status=? WHERE client_id = ? AND project_id = ?", m.Table),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId,
		data.FirstSessionYearOfDay,
		data.FirstSessionYear,
		data.FirstSessionWeekDay,
		data.FirstSessionHour,
		data.FirstSessionDuration,
		data.FirstSessionMinute,
		data.SecondSessionHour,
		data.SecondSessionDuration,
		data.SecondSessionMinute,
		data.ThirdSessionHour,
		data.ThirdSessionDuration,
		data.ThirdSessionMinute,
		data.FourthSessionHour,
		data.FourthSessionDuration,
		data.FourthSessionMinute,
		data.FifthSessionHour,
		data.FifthSessionDuration,
		data.FifthSessionMinute,
		data.SixthSessionHour,
		data.SixthSessionDuration,
		data.SixthSessionMinute,
		data.SeventhSessionHour,
		data.SeventhSessionDuration,
		data.SeventhSessionMinute,
		data.PenultimateSessionHour,
		data.PenultimateSessionDuration,
		data.PenultimateSessionMinute,
		data.LastSessionYearOfDay,
		data.LastSessionYear,
		data.LastSessionHour,
		data.LastSessionDuration,
		data.LastSessionMinute,
		data.LastDurationMinusPenultimateDuration,
		data.FirstFiveMinutesTotalSessionCount,
		data.FirstFiveMinutesTotalSessionDuration,
		data.FirstTenMinutesTotalSessionCount,
		data.FirstTenMinutesTotalSessionDuration,
		data.FirstQuarterHourTotalSessionCount,
		data.FirstQuarterHourTotalSessionDuration,
		data.FirstHalfHourTotalSessionCount,
		data.FirstHalfHourTotalSessionDuration,
		data.FirstHourTotalSessionCount,
		data.FirstHourTotalSessionDuration,
		data.FirstTwoHourTotalSessionCount,
		data.FirstTwoHourTotalSessionDuration,
		data.FirstThreeHourTotalSessionCount,
		data.FirstThreeHourTotalSessionDuration,
		data.FirstSixHourTotalSessionCount,
		data.FirstSixHourTotalSessionDuration,
		data.FirstTwelveHourTotalSessionCount,
		data.FirstTwelveHourTotalSessionDuration,
		data.TotalSessionDay,
		data.TotalSessionHour,
		data.TotalSessionMinute,
		data.TotalSessionDuration,
		data.TotalSessionCount,
		data.FirstDayTotalSessionCount,
		data.FirstDayTotalSessionDuration,
		data.SecondDayTotalSessionCount,
		data.SecondDayTotalSessionDuration,
		data.ThirdDayTotalSessionCount,
		data.ThirdDayTotalSessionDuration,
		data.FourthDayTotalSessionCount,
		data.FourthDayTotalSessionDuration,
		data.FifthDayTotalSessionCount,
		data.FifthDayTotalSessionDuration,
		data.SixthDayTotalSessionCount,
		data.SixthDayTotalSessionDuration,
		data.SeventhDayTotalSessionCount,
		data.SeventhDayTotalSessionDuration,
		data.MinSessionDuration,
		data.MaxSessionDuration,
		data.DailyAvegareSessionCount,
		data.DailyAverageSessionDuration,
		data.SessionBasedAvegareSessionDuration,
		data.DailyAvegareSessionCountMinusFirstDaySessionCount,
		data.DailyAvegareSessionDurationMinusFirstDaySessionDuration,
		data.SessionBasedAvegareSessionDurationMinusFirstSessionDuration,
		data.SessionBasedAvegareSessionDurationMinusLastSessionDuration,
		data.SundaySessionCount,
		data.MondaySessionCount,
		data.TuesdaySessionCount,
		data.WednesdaySessionCount,
		data.ThursdaySessionCount,
		data.FridaySessionCount,
		data.SaturdaySessionCount,
		data.AmSessionCount,
		data.PmSessionCount,
		data.Session0To5HourCount,
		data.Session6To11HourCount,
		data.Session12To17HourCount,
		data.Session18To23HourCount, data.Status , ClientId, ProjectId).Exec(); err != nil {

		return err
	}

	return nil
}
