package cassandra_dal

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/pkg/database/cassandra"
	"fmt"

	//logger "github.com/appneuroncompany/light-logger"
	//"github.com/appneuroncompany/light-logger/clogger"
	"github.com/gocql/gocql"
)

type cassLevelBaseSessionDal struct {
	Client *gocql.Session
	Table  string
}

func NewCassLevelBaseSessionDal(Table string) *cassLevelBaseSessionDal {
	return &cassLevelBaseSessionDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassLevelBaseSessionDal) Add(data *model.LevelBaseSessionResponseModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO MLDatabase.%s(id, client_id, project_id, customer_id, total_level_base_session_minute ,total_level_base_session_count ,first_level_session_level_index ,first_level_session_duration ,first_level_session_year_of_day ,first_level_session_year ,first_level_session_week_day ,first_level_session_hour ,first_level_session_minute ,second_level_session_level_index ,second_level_session_duration ,third_level_session_level_index ,third_level_session_duration ,four_level_session_level_index ,four_level_session_duration ,five_level_session_level_index ,five_level_session_duration ,six_level_session_level_index ,six_level_session_duration ,seven_level_session_level_index ,seven_level_session_duration ,first_five_minutes_total_level_base_session_count ,first_ten_minutes_total_level_base_session_count ,first_quarter_hour_total_level_base_session_count ,first_half_hour_total_level_base_session_count ,first_hour_total_level_base_session_count ,first_two_hour_total_level_base_session_count ,first_three_hour_total_level_base_session_count , first_six_hour_total_level_base_session_count ,first_twelve_hour_total_level_base_session_count ,first_day_total_level_base_session_count ,penultimate_level_session_level_index ,penultimate_level_session_level_duration ,last_level_session_level_index ,last_level_session_level_duration ,last_level_session_year_of_day ,last_level_session_year ,last_level_session_week_day ,last_level_session_hour ,last_level_session_minute , status) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", m.Table),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId,
		data.TotalLevelBaseSessionMinute,
		data.TotalLevelBaseSessionCount,
		data.FirstLevelSessionLevelIndex,
		data.FirstLevelSessionDuration,
		data.FirstLevelSessionYearOfDay,
		data.FirstLevelSessionYear,
		data.FirstLevelSessionWeekDay,
		data.FirstLevelSessionHour,
		data.FirstLevelSessionMinute,
		data.SecondLevelSessionLevelIndex,
		data.SecondLevelSessionDuration,
		data.ThirdLevelSessionLevelIndex,
		data.ThirdLevelSessionDuration,
		data.FourLevelSessionLevelIndex,
		data.FourLevelSessionDuration,
		data.FiveLevelSessionLevelIndex,
		data.FiveLevelSessionDuration,
		data.SixLevelSessionLevelIndex,
		data.SixLevelSessionDuration,
		data.SevenLevelSessionLevelIndex,
		data.SevenLevelSessionDuration,
		data.FirstFiveMinutesTotalLevelBaseSessionCount,
		data.FirstTenMinutesTotalLevelBaseSessionCount,
		data.FirstQuarterHourTotalLevelBaseSessionCount,
		data.FirstHalfHourTotalLevelBaseSessionCount,
		data.FirstHourTotalLevelBaseSessionCount,
		data.FirstTwoHourTotalLevelBaseSessionCount,
		data.FirstThreeHourTotalLevelBaseSessionCount,
		data.FirstSixHourTotalLevelBaseSessionCount,
		data.FirstTwelveHourTotalLevelBaseSessionCount,
		data.FirstDayTotalLevelBaseSessionCount,
		data.PenultimateLevelSessionLevelIndex,
		data.PenultimateLevelSessionLevelDuration,
		data.LastLevelSessionLevelIndex,
		data.LastLevelSessionLevelDuration,
		data.LastLevelSessionYearOfDay,
		data.LastLevelSessionYear,
		data.LastLevelSessionWeekDay,
		data.LastLevelSessionHour,
		data.LastLevelSessionMinute,
		data.Status).Exec(); err != nil {

		return err
	}

	return nil
}

func (m *cassLevelBaseSessionDal) GetById(ClientId int64, ProjectId int64) (*model.LevelBaseSessionResponseModel, error) {
	data := &model.LevelBaseSessionResponseModel{}
	if err := m.Client.Query(fmt.Sprintf("SELECT id , client_id , project_id , customer_id , total_level_base_session_minute ,total_level_base_session_count ,first_level_session_level_index ,first_level_session_duration ,first_level_session_year_of_day ,first_level_session_year ,first_level_session_week_day ,first_level_session_hour ,first_level_session_minute ,second_level_session_level_index ,second_level_session_duration ,third_level_session_level_index ,third_level_session_duration ,four_level_session_level_index ,four_level_session_duration ,five_level_session_level_index ,five_level_session_duration ,six_level_session_level_index ,six_level_session_duration ,seven_level_session_level_index ,seven_level_session_duration ,  first_five_minutes_total_level_base_session_count ,first_ten_minutes_total_level_base_session_count ,first_quarter_hour_total_level_base_session_count ,first_half_hour_total_level_base_session_count ,   first_hour_total_level_base_session_count ,first_two_hour_total_level_base_session_count ,   first_three_hour_total_level_base_session_count , first_six_hour_total_level_base_session_count ,   first_twelve_hour_total_level_base_session_count ,first_day_total_level_base_session_count ,penultimate_level_session_level_index ,penultimate_level_session_level_duration ,last_level_session_level_index ,last_level_session_level_duration ,last_level_session_year_of_day ,last_level_session_year ,last_level_session_week_day ,last_level_session_hour ,last_level_session_minute ,status FROM MLDatabase.%s WHERE client_id = ? AND project_id = ? LIMIT 1", m.Table),
		ClientId, ProjectId).Scan(&data.Id, &data.ClientId, &data.ProjectId, &data.CustomerId,
		&data.TotalLevelBaseSessionMinute,
		&data.TotalLevelBaseSessionCount,
		&data.FirstLevelSessionLevelIndex,
		&data.FirstLevelSessionDuration,
		&data.FirstLevelSessionYearOfDay,
		&data.FirstLevelSessionYear,
		&data.FirstLevelSessionWeekDay,
		&data.FirstLevelSessionHour,
		&data.FirstLevelSessionMinute,
		&data.SecondLevelSessionLevelIndex,
		&data.SecondLevelSessionDuration,
		&data.ThirdLevelSessionLevelIndex,
		&data.ThirdLevelSessionDuration,
		&data.FourLevelSessionLevelIndex,
		&data.FourLevelSessionDuration,
		&data.FiveLevelSessionLevelIndex,
		&data.FiveLevelSessionDuration,
		&data.SixLevelSessionLevelIndex,
		&data.SixLevelSessionDuration,
		&data.SevenLevelSessionLevelIndex,
		&data.SevenLevelSessionDuration,
		&data.FirstFiveMinutesTotalLevelBaseSessionCount,
		&data.FirstTenMinutesTotalLevelBaseSessionCount,
		&data.FirstQuarterHourTotalLevelBaseSessionCount,
		&data.FirstHalfHourTotalLevelBaseSessionCount,
		&data.FirstHourTotalLevelBaseSessionCount,
		&data.FirstTwoHourTotalLevelBaseSessionCount,
		&data.FirstThreeHourTotalLevelBaseSessionCount,
		&data.FirstSixHourTotalLevelBaseSessionCount,
		&data.FirstTwelveHourTotalLevelBaseSessionCount,
		&data.FirstDayTotalLevelBaseSessionCount,
		&data.PenultimateLevelSessionLevelIndex,
		&data.PenultimateLevelSessionLevelDuration,
		&data.LastLevelSessionLevelIndex,
		&data.LastLevelSessionLevelDuration,
		&data.LastLevelSessionYearOfDay,
		&data.LastLevelSessionYear,
		&data.LastLevelSessionWeekDay,
		&data.LastLevelSessionHour,
		&data.LastLevelSessionMinute, &data.Status); err != nil {

		return nil, err
	}

	return data, nil
}

func (m *cassLevelBaseSessionDal) UpdateById(ClientId int64, ProjectId int64, data *model.LevelBaseSessionResponseModel) error {
	if err := m.Client.Query(fmt.Sprintf("UPDATE MLDatabase.%s SET id=?, customer_id=?, total_level_base_session_minute=? ,total_level_base_session_count=? ,first_level_session_level_index=? ,first_level_session_duration=? ,first_level_session_year_of_day=? ,first_level_session_year=? ,first_level_session_week_day=? ,first_level_session_hour=? ,first_level_session_minute=? ,second_level_session_level_index=? ,second_level_session_duration=? ,third_level_session_level_index=? ,third_level_session_duration=? ,four_level_session_level_index=? ,four_level_session_duration=? ,five_level_session_level_index=? ,five_level_session_duration=? ,six_level_session_level_index=? ,six_level_session_duration=? ,seven_level_session_level_index=? ,seven_level_session_duration=? ,first_five_minutes_total_level_base_session_count=? ,first_ten_minutes_total_level_base_session_count=? ,first_quarter_hour_total_level_base_session_count=? ,first_half_hour_total_level_base_session_count=? ,first_hour_total_level_base_session_count=? ,first_two_hour_total_level_base_session_count=? ,first_three_hour_total_level_base_session_count=? , first_six_hour_total_level_base_session_count=? ,first_twelve_hour_total_level_base_session_count=? ,first_day_total_level_base_session_count=? ,penultimate_level_session_level_index=? ,penultimate_level_session_level_duration=? ,last_level_session_level_index=? ,last_level_session_level_duration=? ,last_level_session_year_of_day=? ,last_level_session_year=? ,last_level_session_week_day=? ,last_level_session_hour=? ,last_level_session_minute=? , status=? WHERE client_id = ? AND project_id = ?", m.Table),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId,
		data.TotalLevelBaseSessionMinute,
		data.TotalLevelBaseSessionCount,
		data.FirstLevelSessionLevelIndex,
		data.FirstLevelSessionDuration,
		data.FirstLevelSessionYearOfDay,
		data.FirstLevelSessionYear,
		data.FirstLevelSessionWeekDay,
		data.FirstLevelSessionHour,
		data.FirstLevelSessionMinute,
		data.SecondLevelSessionLevelIndex,
		data.SecondLevelSessionDuration,
		data.ThirdLevelSessionLevelIndex,
		data.ThirdLevelSessionDuration,
		data.FourLevelSessionLevelIndex,
		data.FourLevelSessionDuration,
		data.FiveLevelSessionLevelIndex,
		data.FiveLevelSessionDuration,
		data.SixLevelSessionLevelIndex,
		data.SixLevelSessionDuration,
		data.SevenLevelSessionLevelIndex,
		data.SevenLevelSessionDuration,
		data.FirstFiveMinutesTotalLevelBaseSessionCount,
		data.FirstTenMinutesTotalLevelBaseSessionCount,
		data.FirstQuarterHourTotalLevelBaseSessionCount,
		data.FirstHalfHourTotalLevelBaseSessionCount,
		data.FirstHourTotalLevelBaseSessionCount,
		data.FirstTwoHourTotalLevelBaseSessionCount,
		data.FirstThreeHourTotalLevelBaseSessionCount,
		data.FirstSixHourTotalLevelBaseSessionCount,
		data.FirstTwelveHourTotalLevelBaseSessionCount,
		data.FirstDayTotalLevelBaseSessionCount,
		data.PenultimateLevelSessionLevelIndex,
		data.PenultimateLevelSessionLevelDuration,
		data.LastLevelSessionLevelIndex,
		data.LastLevelSessionLevelDuration,
		data.LastLevelSessionYearOfDay,
		data.LastLevelSessionYear,
		data.LastLevelSessionWeekDay,
		data.LastLevelSessionHour,
		data.LastLevelSessionMinute, data.Status, ClientId, ProjectId).Exec(); err != nil {

		return err
	}

	return nil
}
