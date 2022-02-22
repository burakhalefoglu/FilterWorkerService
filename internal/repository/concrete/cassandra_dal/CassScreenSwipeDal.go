package cassandra_dal

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/pkg/database/cassandra"
	"fmt"

	//logger "github.com/appneuroncompany/light-logger"
	//"github.com/appneuroncompany/light-logger/clogger"
	"github.com/gocql/gocql"
)

type cassScreenSwipeDal struct {
	Client *gocql.Session
	Table  string
}

func NewCassScreenSwipeDal(Table string) *cassScreenSwipeDal {
	return &cassScreenSwipeDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassScreenSwipeDal) Add(data *model.ScreenSwipeResponseModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO MLDatabase.%s(id, client_id, project_id, customer_id, level_index ,total_swipe_session_count ,total_swipe_hour ,first_swipe_year_of_day ,first_swipe_year ,first_swipe_hour ,first_swipe_week_day ,first_swipe_minute ,fist_swipe_direction ,first_swipe_start_x_cor ,first_swipe_start_y_cor ,first_swipe_finish_x_cor ,first_swipe_finish_y_cor ,second_swipe_direction ,second_swipe_start_x_cor ,second_swipe_start_y_cor ,second_swipe_finish_x_cor ,second_swipe_finish_y_cor ,third_swipe_direction ,third_swipe_start_x_cor ,third_swipe_start_y_cor ,third_swipe_finish_x_cor ,third_swipe_finish_y_cor ,fourth_swipe_direction ,fourth_swipe_start_x_cor ,fourth_swipe_start_y_cor ,fourth_swipe_finish_x_cor ,fourth_swipe_finish_y_cor ,fifth_swipe_direction ,fifth_swipe_start_x_cor ,fifth_swipe_start_y_cor ,fifth_swipe_finish_x_cor ,fifth_swipe_finish_y_cor ,sixth_swipe_direction ,  sixth_swipe_start_x_cor ,  sixth_swipe_start_y_cor ,  sixth_swipe_finish_x_cor , sixth_swipe_finish_y_cor , seventh_swipe_direction ,seventh_swipe_start_x_cor ,seventh_swipe_start_y_cor ,seventh_swipe_finish_x_cor ,seventh_swipe_finish_y_cor ,penultimate_swipe_direction ,penultimate_swipe_start_x_cor ,Penultimate_swipe_start_y_cor ,penultimate_swipe_finish_x_cor ,penultimate_swipe_finish_y_cor ,penultimate_swipe_year_of_day ,penultimate_swipe_year ,penultimate_swipe_hour ,penultimate_swipe_week_day ,penultimate_swipe_minute ,last_swipe_direction ,last_swipe_start_x_cor ,last_swipe_start_y_cor ,last_swipe_finish_x_cor ,last_swipe_finish_y_cor ,last_swipe_year_of_day ,last_swipe_year ,last_swipe_hour ,last_swipe_week_day ,last_swipe_minute ,first_day_total_swipe_up_count ,first_day_total_swipe_down_count ,  first_day_total_swipe_right_count , first_day_total_swipe_left_count ,   first_day_swipe_total_start_x_cor ,  first_day_swipe_total_start_y_cor ,  first_day_swipe_total_finish_x_cor , first_day_swipe_total_finish_y_cor , second_day_total_swipe_up_count ,   second_day_total_swipe_down_count , second_day_total_swipe_right_count ,second_day_total_swipe_left_count , second_day_swipe_total_start_x_cor ,  second_day_swipe_total_start_y_cor , second_day_swipe_total_finish_x_cor ,second_day_swipe_total_finish_y_cor ,third_day_total_swipe_up_count ,    third_day_total_swipe_down_count ,  third_day_total_swipe_right_count , third_day_total_swipe_left_count ,   third_day_Swipetotal_start_x_cor ,  third_day_Swipetotal_start_y_cor ,  third_day_Swipetotal_finish_x_cor ,  third_day_Swipetotal_finish_y_cor , fourth_day_total_swipe_up_count ,   fourth_day_total_swipe_down_count , fourth_day_total_swipe_right_count ,fourth_day_total_swipe_left_count , fourth_day_swipe_total_start_x_cor , fourth_day_swipe_total_start_y_cor , fourth_day_swipe_total_finish_x_cor ,fourth_day_swipe_total_finish_y_cor ,fifth_day_total_swipe_up_count ,fifth_day_total_swipe_down_count ,  fifth_day_total_swipe_right_count , fifth_day_total_swipe_left_count ,  fifth_day_swipe_total_start_x_cor ,  fifth_day_swipe_total_start_y_cor ,  fifth_day_swipe_total_finish_x_cor , fifth_day_swipe_total_finish_y_cor , sixth_day_total_swipe_up_count ,    sixth_day_total_swipe_down_count ,  sixth_day_total_swipe_right_count , sixth_day_total_swipe_left_count ,  sixth_day_swipe_total_start_x_cor ,  sixth_day_swipe_total_start_y_cor ,  sixth_day_swipe_total_finish_x_cor , sixth_day_swipe_total_finish_y_cor , seventh_day_total_swipe_up_count ,  seventh_day_total_swipe_down_count ,seventh_day_total_swipe_right_count ,seventh_day_total_swipe_left_count ,seventh_day_swipe_total_start_x_cor ,seventh_day_swipe_total_start_y_cor ,seventh_day_swipe_total_finish_x_cor ,seventh_day_swipe_total_finish_y_cor ,total_swipe_up_count ,total_swipe_down_count ,total_swipe_right_count ,total_swipe_left_count ,total_swipe_start_x_cor ,total_swipe_start_y_cor ,total_swipe_finish_x_cor ,total_swipe_finish_y_cor ,  status) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", m.Table),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId,
		data.LevelIndex,
		data.TotalSwipeSessionCount,
		data.TotalSwipeHour,
		data.FirstSwipeYearOfDay,
		data.FirstSwipeYear,
		data.FirstSwipeHour,
		data.FirstSwipeWeekDay,
		data.FirstSwipeMinute,
		data.FistSwipeDirection,
		data.FirstSwipeStartXCor,
		data.FirstSwipeStartYCor,
		data.FirstSwipeFinishXCor,
		data.FirstSwipeFinishYCor,
		data.SecondSwipeDirection,
		data.SecondSwipeStartXCor,
		data.SecondSwipeStartYCor,
		data.SecondSwipeFinishXCor,
		data.SecondSwipeFinishYCor,
		data.ThirdSwipeDirection,
		data.ThirdSwipeStartXCor,
		data.ThirdSwipeStartYCor,
		data.ThirdSwipeFinishXCor,
		data.ThirdSwipeFinishYCor,
		data.FourthSwipeDirection,
		data.FourthSwipeStartXCor,
		data.FourthSwipeStartYCor,
		data.FourthSwipeFinishXCor,
		data.FourthSwipeFinishYCor,
		data.FifthSwipeDirection,
		data.FifthSwipeStartXCor,
		data.FifthSwipeStartYCor,
		data.FifthSwipeFinishXCor,
		data.FifthSwipeFinishYCor,
		data.SixthSwipeDirection,
		data.SixthSwipeStartXCor,
		data.SixthSwipeStartYCor,
		data.SixthSwipeFinishXCor,
		data.SixthSwipeFinishYCor,
		data.SeventhSwipeDirection,
		data.SeventhSwipeStartXCor,
		data.SeventhSwipeStartYCor,
		data.SeventhSwipeFinishXCor,
		data.SeventhSwipeFinishYCor,
		data.PenultimateSwipeDirection,
		data.PenultimateSwipeStartXCor,
		data.PenultimateSwipeStartYCor,
		data.PenultimateSwipeFinishXCor,
		data.PenultimateSwipeFinishYCor,
		data.PenultimateSwipeYearOfDay,
		data.PenultimateSwipeYear,
		data.PenultimateSwipeHour,
		data.PenultimateSwipeWeekDay,
		data.PenultimateSwipeMinute,
		data.LastSwipeDirection,
		data.LastSwipeStartXCor,
		data.LastSwipeStartYCor,
		data.LastSwipeFinishXCor,
		data.LastSwipeFinishYCor,
		data.LastSwipeYearOfDay,
		data.LastSwipeYear,
		data.LastSwipeHour,
		data.LastSwipeWeekDay,
		data.LastSwipeMinute,
		data.FirstDayTotalSwipeUpCount,
		data.FirstDayTotalSwipeDownCount,
		data.FirstDayTotalSwipeRightCount,
		data.FirstDayTotalSwipeLeftCount,
		data.FirstDaySwipeTotalStartXCor,
		data.FirstDaySwipeTotalStartYCor,
		data.FirstDaySwipeTotalFinishXCor,
		data.FirstDaySwipeTotalFinishYCor,
		data.SecondDayTotalSwipeUpCount,
		data.SecondDayTotalSwipeDownCount,
		data.SecondDayTotalSwipeRightCount,
		data.SecondDayTotalSwipeLeftCount,
		data.SecondDaySwipeTotalStartXCor,
		data.SecondDaySwipeTotalStartYCor,
		data.SecondDaySwipeTotalFinishXCor,
		data.SecondDaySwipeTotalFinishYCor,
		data.ThirdDayTotalSwipeUpCount,
		data.ThirdDayTotalSwipeDownCount,
		data.ThirdDayTotalSwipeRightCount,
		data.ThirdDayTotalSwipeLeftCount,
		data.ThirdDaySwipeTotalStartXCor,
		data.ThirdDaySwipeTotalStartYCor,
		data.ThirdDaySwipeTotalFinishXCor,
		data.ThirdDaySwipeTotalFinishYCor,
		data.FourthDayTotalSwipeUpCount,
		data.FourthDayTotalSwipeDownCount,
		data.FourthDayTotalSwipeRightCount,
		data.FourthDayTotalSwipeLeftCount,
		data.FourthDaySwipeTotalStartXCor,
		data.FourthDaySwipeTotalStartYCor,
		data.FourthDaySwipeTotalFinishXCor,
		data.FourthDaySwipeTotalFinishYCor,
		data.FifthDayTotalSwipeUpCount,
		data.FifthDayTotalSwipeDownCount,
		data.FifthDayTotalSwipeRightCount,
		data.FifthDayTotalSwipeLeftCount,
		data.FifthDaySwipeTotalStartXCor,
		data.FifthDaySwipeTotalStartYCor,
		data.FifthDaySwipeTotalFinishXCor,
		data.FifthDaySwipeTotalFinishYCor,
		data.SixthDayTotalSwipeUpCount,
		data.SixthDayTotalSwipeDownCount,
		data.SixthDayTotalSwipeRightCount,
		data.SixthDayTotalSwipeLeftCount,
		data.SixthDaySwipeTotalStartXCor,
		data.SixthDaySwipeTotalStartYCor,
		data.SixthDaySwipeTotalFinishXCor,
		data.SixthDaySwipeTotalFinishYCor,
		data.SeventhDayTotalSwipeUpCount,
		data.SeventhDayTotalSwipeDownCount,
		data.SeventhDayTotalSwipeRightCount,
		data.SeventhDayTotalSwipeLeftCount,
		data.SeventhDaySwipeTotalStartXCor,
		data.SeventhDaySwipeTotalStartYCor,
		data.SeventhDaySwipeTotalFinishXCor,
		data.SeventhDaySwipeTotalFinishYCor,
		data.TotalSwipeUpCount,
		data.TotalSwipeDownCount,
		data.TotalSwipeRightCount,
		data.TotalSwipeLeftCount,
		data.TotalSwipeStartXCor,
		data.TotalSwipeStartYCor,
		data.TotalSwipeFinishXCor,
		data.TotalSwipeFinishYCor,
		data.Status).Exec(); err != nil {
		
		return err
	}
	
	return nil
}

func (m *cassScreenSwipeDal) GetById(ClientId int64, ProjectId int64) (*model.ScreenSwipeResponseModel, error) {
	data := &model.ScreenSwipeResponseModel{}
	if err := m.Client.Query(fmt.Sprintf("SELECT * FROM MLDatabase.%s WHERE client_id = ? AND project_id = ? LIMIT 1", m.Table),
		ClientId, ProjectId).Scan(&data.Id, &data.ClientId, &data.ProjectId, &data.CustomerId,
		&data.LevelIndex,
		&data.TotalSwipeSessionCount,
		&data.TotalSwipeHour,
		&data.FirstSwipeYearOfDay,
		&data.FirstSwipeYear,
		&data.FirstSwipeHour,
		&data.FirstSwipeWeekDay,
		&data.FirstSwipeMinute,
		&data.FistSwipeDirection,
		&data.FirstSwipeStartXCor,
		&data.FirstSwipeStartYCor,
		&data.FirstSwipeFinishXCor,
		&data.FirstSwipeFinishYCor,
		&data.SecondSwipeDirection,
		&data.SecondSwipeStartXCor,
		&data.SecondSwipeStartYCor,
		&data.SecondSwipeFinishXCor,
		&data.SecondSwipeFinishYCor,
		&data.ThirdSwipeDirection,
		&data.ThirdSwipeStartXCor,
		&data.ThirdSwipeStartYCor,
		&data.ThirdSwipeFinishXCor,
		&data.ThirdSwipeFinishYCor,
		&data.FourthSwipeDirection,
		&data.FourthSwipeStartXCor,
		&data.FourthSwipeStartYCor,
		&data.FourthSwipeFinishXCor,
		&data.FourthSwipeFinishYCor,
		&data.FifthSwipeDirection,
		&data.FifthSwipeStartXCor,
		&data.FifthSwipeStartYCor,
		&data.FifthSwipeFinishXCor,
		&data.FifthSwipeFinishYCor,
		&data.SixthSwipeDirection,
		&data.SixthSwipeStartXCor,
		&data.SixthSwipeStartYCor,
		&data.SixthSwipeFinishXCor,
		&data.SixthSwipeFinishYCor,
		&data.SeventhSwipeDirection,
		&data.SeventhSwipeStartXCor,
		&data.SeventhSwipeStartYCor,
		&data.SeventhSwipeFinishXCor,
		&data.SeventhSwipeFinishYCor,
		&data.PenultimateSwipeDirection,
		&data.PenultimateSwipeStartXCor,
		&data.PenultimateSwipeStartYCor,
		&data.PenultimateSwipeFinishXCor,
		&data.PenultimateSwipeFinishYCor,
		&data.PenultimateSwipeYearOfDay,
		&data.PenultimateSwipeYear,
		&data.PenultimateSwipeHour,
		&data.PenultimateSwipeWeekDay,
		&data.PenultimateSwipeMinute,
		&data.LastSwipeDirection,
		&data.LastSwipeStartXCor,
		&data.LastSwipeStartYCor,
		&data.LastSwipeFinishXCor,
		&data.LastSwipeFinishYCor,
		&data.LastSwipeYearOfDay,
		&data.LastSwipeYear,
		&data.LastSwipeHour,
		&data.LastSwipeWeekDay,
		&data.LastSwipeMinute,
		&data.FirstDayTotalSwipeUpCount,
		&data.FirstDayTotalSwipeDownCount,
		&data.FirstDayTotalSwipeRightCount,
		&data.FirstDayTotalSwipeLeftCount,
		&data.FirstDaySwipeTotalStartXCor,
		&data.FirstDaySwipeTotalStartYCor,
		&data.FirstDaySwipeTotalFinishXCor,
		&data.FirstDaySwipeTotalFinishYCor,
		&data.SecondDayTotalSwipeUpCount,
		&data.SecondDayTotalSwipeDownCount,
		&data.SecondDayTotalSwipeRightCount,
		&data.SecondDayTotalSwipeLeftCount,
		&data.SecondDaySwipeTotalStartXCor,
		&data.SecondDaySwipeTotalStartYCor,
		&data.SecondDaySwipeTotalFinishXCor,
		&data.SecondDaySwipeTotalFinishYCor,
		&data.ThirdDayTotalSwipeUpCount,
		&data.ThirdDayTotalSwipeDownCount,
		&data.ThirdDayTotalSwipeRightCount,
		&data.ThirdDayTotalSwipeLeftCount,
		&data.ThirdDaySwipeTotalStartXCor,
		&data.ThirdDaySwipeTotalStartYCor,
		&data.ThirdDaySwipeTotalFinishXCor,
		&data.ThirdDaySwipeTotalFinishYCor,
		&data.FourthDayTotalSwipeUpCount,
		&data.FourthDayTotalSwipeDownCount,
		&data.FourthDayTotalSwipeRightCount,
		&data.FourthDayTotalSwipeLeftCount,
		&data.FourthDaySwipeTotalStartXCor,
		&data.FourthDaySwipeTotalStartYCor,
		&data.FourthDaySwipeTotalFinishXCor,
		&data.FourthDaySwipeTotalFinishYCor,
		&data.FifthDayTotalSwipeUpCount,
		&data.FifthDayTotalSwipeDownCount,
		&data.FifthDayTotalSwipeRightCount,
		&data.FifthDayTotalSwipeLeftCount,
		&data.FifthDaySwipeTotalStartXCor,
		&data.FifthDaySwipeTotalStartYCor,
		&data.FifthDaySwipeTotalFinishXCor,
		&data.FifthDaySwipeTotalFinishYCor,
		&data.SixthDayTotalSwipeUpCount,
		&data.SixthDayTotalSwipeDownCount,
		&data.SixthDayTotalSwipeRightCount,
		&data.SixthDayTotalSwipeLeftCount,
		&data.SixthDaySwipeTotalStartXCor,
		&data.SixthDaySwipeTotalStartYCor,
		&data.SixthDaySwipeTotalFinishXCor,
		&data.SixthDaySwipeTotalFinishYCor,
		&data.SeventhDayTotalSwipeUpCount,
		&data.SeventhDayTotalSwipeDownCount,
		&data.SeventhDayTotalSwipeRightCount,
		&data.SeventhDayTotalSwipeLeftCount,
		&data.SeventhDaySwipeTotalStartXCor,
		&data.SeventhDaySwipeTotalStartYCor,
		&data.SeventhDaySwipeTotalFinishXCor,
		&data.SeventhDaySwipeTotalFinishYCor,
		&data.TotalSwipeUpCount,
		&data.TotalSwipeDownCount,
		&data.TotalSwipeRightCount,
		&data.TotalSwipeLeftCount,
		&data.TotalSwipeStartXCor,
		&data.TotalSwipeStartYCor,
		&data.TotalSwipeFinishXCor,
		&data.TotalSwipeFinishYCor,
		&data.Status); err != nil {
		
		return nil, err
	}
	
	return data, nil
}

func (m *cassScreenSwipeDal) UpdateById(ClientId int64, ProjectId int64, data *model.ScreenSwipeResponseModel) error {
	if err := m.Client.Query(fmt.Sprintf("UPDATE MLDatabase.%s SET id=?, customer_id=?, level_index=? ,total_swipe_session_count=? ,total_swipe_hour=? ,first_swipe_year_of_day=? ,first_swipe_year=? ,first_swipe_hour=? ,first_swipe_week_day=? ,first_swipe_minute=? ,fist_swipe_direction=? ,first_swipe_start_x_cor=? ,first_swipe_start_y_cor=? ,first_swipe_finish_x_cor=? ,first_swipe_finish_y_cor=? ,second_swipe_direction=? ,second_swipe_start_x_cor=? ,second_swipe_start_y_cor=? ,second_swipe_finish_x_cor=? ,second_swipe_finish_y_cor=? ,third_swipe_direction=? ,third_swipe_start_x_cor=? ,third_swipe_start_y_cor=? ,third_swipe_finish_x_cor=? ,third_swipe_finish_y_cor=? ,fourth_swipe_direction=? ,fourth_swipe_start_x_cor=? ,fourth_swipe_start_y_cor=? ,fourth_swipe_finish_x_cor=? ,fourth_swipe_finish_y_cor=? ,fifth_swipe_direction=? ,fifth_swipe_start_x_cor=? ,fifth_swipe_start_y_cor=? ,fifth_swipe_finish_x_cor=? ,fifth_swipe_finish_y_cor=? ,sixth_swipe_direction=? ,  sixth_swipe_start_x_cor=? ,  sixth_swipe_start_y_cor=? ,  sixth_swipe_finish_x_cor=? , sixth_swipe_finish_y_cor=? , seventh_swipe_direction=? ,seventh_swipe_start_x_cor=? ,seventh_swipe_start_y_cor=? ,seventh_swipe_finish_x_cor=? ,seventh_swipe_finish_y_cor=? ,penultimate_swipe_direction=? ,penultimate_swipe_start_x_cor=? ,Penultimate_swipe_start_y_cor=? ,penultimate_swipe_finish_x_cor=? ,penultimate_swipe_finish_y_cor=? ,penultimate_swipe_year_of_day=? ,penultimate_swipe_year=? ,penultimate_swipe_hour=? ,penultimate_swipe_week_day=? ,penultimate_swipe_minute=? ,last_swipe_direction=? ,last_swipe_start_x_cor=? ,last_swipe_start_y_cor=? ,last_swipe_finish_x_cor=? ,last_swipe_finish_y_cor=? ,last_swipe_year_of_day=? ,last_swipe_year=? ,last_swipe_hour=? ,last_swipe_week_day=? ,last_swipe_minute=? ,first_day_total_swipe_up_count=? ,first_day_total_swipe_down_count=? ,  first_day_total_swipe_right_count=? , first_day_total_swipe_left_count=? ,   first_day_swipe_total_start_x_cor=? ,  first_day_swipe_total_start_y_cor=? ,  first_day_swipe_total_finish_x_cor=? , first_day_swipe_total_finish_y_cor=? , second_day_total_swipe_up_count=? ,   second_day_total_swipe_down_count=? , second_day_total_swipe_right_count=? ,second_day_total_swipe_left_count=? , second_day_swipe_total_start_x_cor=? ,  second_day_swipe_total_start_y_cor=? , second_day_swipe_total_finish_x_cor=? ,second_day_swipe_total_finish_y_cor=? ,third_day_total_swipe_up_count=? ,    third_day_total_swipe_down_count=? ,  third_day_total_swipe_right_count=? , third_day_total_swipe_left_count=? ,   third_day_Swipetotal_start_x_cor=? ,  third_day_Swipetotal_start_y_cor=? ,  third_day_Swipetotal_finish_x_cor=? ,  third_day_Swipetotal_finish_y_cor=? , fourth_day_total_swipe_up_count=? ,   fourth_day_total_swipe_down_count=? , fourth_day_total_swipe_right_count=? ,fourth_day_total_swipe_left_count=? , fourth_day_swipe_total_start_x_cor=? , fourth_day_swipe_total_start_y_cor=? , fourth_day_swipe_total_finish_x_cor=? ,fourth_day_swipe_total_finish_y_cor=? ,fifth_day_total_swipe_up_count=? ,fifth_day_total_swipe_down_count=? ,  fifth_day_total_swipe_right_count=? , fifth_day_total_swipe_left_count=? ,  fifth_day_swipe_total_start_x_cor=? ,  fifth_day_swipe_total_start_y_cor=? ,  fifth_day_swipe_total_finish_x_cor=? , fifth_day_swipe_total_finish_y_cor=? , sixth_day_total_swipe_up_count=? ,    sixth_day_total_swipe_down_count=? ,  sixth_day_total_swipe_right_count=? , sixth_day_total_swipe_left_count=? ,  sixth_day_swipe_total_start_x_cor=? ,  sixth_day_swipe_total_start_y_cor=? ,  sixth_day_swipe_total_finish_x_cor=? , sixth_day_swipe_total_finish_y_cor=? , seventh_day_total_swipe_up_count=? ,  seventh_day_total_swipe_down_count=? ,seventh_day_total_swipe_right_count=? ,seventh_day_total_swipe_left_count=? ,seventh_day_swipe_total_start_x_cor=? ,seventh_day_swipe_total_start_y_cor=? ,seventh_day_swipe_total_finish_x_cor=? ,seventh_day_swipe_total_finish_y_cor=? ,total_swipe_up_count=? ,total_swipe_down_count=? ,total_swipe_right_count=? ,total_swipe_left_count=? ,total_swipe_start_x_cor=? ,total_swipe_start_y_cor=? ,total_swipe_finish_x_cor=? ,total_swipe_finish_y_cor=? , status=? WHERE client_id = ? AND project_id = ?", m.Table),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId,
		data.LevelIndex,
		data.TotalSwipeSessionCount,
		data.TotalSwipeHour,
		data.FirstSwipeYearOfDay,
		data.FirstSwipeYear,
		data.FirstSwipeHour,
		data.FirstSwipeWeekDay,
		data.FirstSwipeMinute,
		data.FistSwipeDirection,
		data.FirstSwipeStartXCor,
		data.FirstSwipeStartYCor,
		data.FirstSwipeFinishXCor,
		data.FirstSwipeFinishYCor,
		data.SecondSwipeDirection,
		data.SecondSwipeStartXCor,
		data.SecondSwipeStartYCor,
		data.SecondSwipeFinishXCor,
		data.SecondSwipeFinishYCor,
		data.ThirdSwipeDirection,
		data.ThirdSwipeStartXCor,
		data.ThirdSwipeStartYCor,
		data.ThirdSwipeFinishXCor,
		data.ThirdSwipeFinishYCor,
		data.FourthSwipeDirection,
		data.FourthSwipeStartXCor,
		data.FourthSwipeStartYCor,
		data.FourthSwipeFinishXCor,
		data.FourthSwipeFinishYCor,
		data.FifthSwipeDirection,
		data.FifthSwipeStartXCor,
		data.FifthSwipeStartYCor,
		data.FifthSwipeFinishXCor,
		data.FifthSwipeFinishYCor,
		data.SixthSwipeDirection,
		data.SixthSwipeStartXCor,
		data.SixthSwipeStartYCor,
		data.SixthSwipeFinishXCor,
		data.SixthSwipeFinishYCor,
		data.SeventhSwipeDirection,
		data.SeventhSwipeStartXCor,
		data.SeventhSwipeStartYCor,
		data.SeventhSwipeFinishXCor,
		data.SeventhSwipeFinishYCor,
		data.PenultimateSwipeDirection,
		data.PenultimateSwipeStartXCor,
		data.PenultimateSwipeStartYCor,
		data.PenultimateSwipeFinishXCor,
		data.PenultimateSwipeFinishYCor,
		data.PenultimateSwipeYearOfDay,
		data.PenultimateSwipeYear,
		data.PenultimateSwipeHour,
		data.PenultimateSwipeWeekDay,
		data.PenultimateSwipeMinute,
		data.LastSwipeDirection,
		data.LastSwipeStartXCor,
		data.LastSwipeStartYCor,
		data.LastSwipeFinishXCor,
		data.LastSwipeFinishYCor,
		data.LastSwipeYearOfDay,
		data.LastSwipeYear,
		data.LastSwipeHour,
		data.LastSwipeWeekDay,
		data.LastSwipeMinute,
		data.FirstDayTotalSwipeUpCount,
		data.FirstDayTotalSwipeDownCount,
		data.FirstDayTotalSwipeRightCount,
		data.FirstDayTotalSwipeLeftCount,
		data.FirstDaySwipeTotalStartXCor,
		data.FirstDaySwipeTotalStartYCor,
		data.FirstDaySwipeTotalFinishXCor,
		data.FirstDaySwipeTotalFinishYCor,
		data.SecondDayTotalSwipeUpCount,
		data.SecondDayTotalSwipeDownCount,
		data.SecondDayTotalSwipeRightCount,
		data.SecondDayTotalSwipeLeftCount,
		data.SecondDaySwipeTotalStartXCor,
		data.SecondDaySwipeTotalStartYCor,
		data.SecondDaySwipeTotalFinishXCor,
		data.SecondDaySwipeTotalFinishYCor,
		data.ThirdDayTotalSwipeUpCount,
		data.ThirdDayTotalSwipeDownCount,
		data.ThirdDayTotalSwipeRightCount,
		data.ThirdDayTotalSwipeLeftCount,
		data.ThirdDaySwipeTotalStartXCor,
		data.ThirdDaySwipeTotalStartYCor,
		data.ThirdDaySwipeTotalFinishXCor,
		data.ThirdDaySwipeTotalFinishYCor,
		data.FourthDayTotalSwipeUpCount,
		data.FourthDayTotalSwipeDownCount,
		data.FourthDayTotalSwipeRightCount,
		data.FourthDayTotalSwipeLeftCount,
		data.FourthDaySwipeTotalStartXCor,
		data.FourthDaySwipeTotalStartYCor,
		data.FourthDaySwipeTotalFinishXCor,
		data.FourthDaySwipeTotalFinishYCor,
		data.FifthDayTotalSwipeUpCount,
		data.FifthDayTotalSwipeDownCount,
		data.FifthDayTotalSwipeRightCount,
		data.FifthDayTotalSwipeLeftCount,
		data.FifthDaySwipeTotalStartXCor,
		data.FifthDaySwipeTotalStartYCor,
		data.FifthDaySwipeTotalFinishXCor,
		data.FifthDaySwipeTotalFinishYCor,
		data.SixthDayTotalSwipeUpCount,
		data.SixthDayTotalSwipeDownCount,
		data.SixthDayTotalSwipeRightCount,
		data.SixthDayTotalSwipeLeftCount,
		data.SixthDaySwipeTotalStartXCor,
		data.SixthDaySwipeTotalStartYCor,
		data.SixthDaySwipeTotalFinishXCor,
		data.SixthDaySwipeTotalFinishYCor,
		data.SeventhDayTotalSwipeUpCount,
		data.SeventhDayTotalSwipeDownCount,
		data.SeventhDayTotalSwipeRightCount,
		data.SeventhDayTotalSwipeLeftCount,
		data.SeventhDaySwipeTotalStartXCor,
		data.SeventhDaySwipeTotalStartYCor,
		data.SeventhDaySwipeTotalFinishXCor,
		data.SeventhDaySwipeTotalFinishYCor,
		data.TotalSwipeUpCount,
		data.TotalSwipeDownCount,
		data.TotalSwipeRightCount,
		data.TotalSwipeLeftCount,
		data.TotalSwipeStartXCor,
		data.TotalSwipeStartYCor,
		data.TotalSwipeFinishXCor,
		data.TotalSwipeFinishYCor, data.Status, ClientId, ProjectId).Exec(); err != nil {

		return err
	}

	return nil
}
