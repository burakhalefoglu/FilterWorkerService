package cassandra_dal

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/pkg/database/cassandra"
	"fmt"

	"github.com/gocql/gocql"
)

type cassBuyingEventDal struct {
	Client *gocql.Session
	Table  string
}

func NewCassBuyingEventDal(Table string) *cassBuyingEventDal {
	return &cassBuyingEventDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}


func (m *cassBuyingEventDal) Add(data *model.BuyingEventResponseModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO MLDatabase.%s(id, client_id, project_id, customer_id, level_index, total_buying_count ,total_buying_day, total_buying_hour, first_buying_year_of_day, first_buying_year,  first_buying_hour,  first_buying_minute , first_buying_product_type, second_buying_year_of_day, second_buying_hour,  second_buying_minute, second_buying_product_type, third_buying_year_of_day, third_buying_hour ,third_buying_minute ,third_buying_product_type ,fourth_buying_year_of_day ,fourth_buying_hour ,fourth_buying_minute ,fourth_buying_product_type ,fifth_buying_year_of_day ,fifth_buying_hour ,fifth_buying_minute ,fifth_buying_product_type ,sixth_buying_year_of_day ,sixth_buying_hour ,sixth_buying_minute ,sixth_buying_product_type, seventh_buying_year_of_day ,seventh_buying_hour ,seventh_buying_minute ,seventh_buying_product_type ,penultimate_buying_year_of_day ,penultimate_buying_hour ,penultimate_buying_minute ,penultimate_buying_product_type ,last_buying_year_of_day ,last_buying_year ,last_buying_hour ,last_buying_minute ,last_buying_product_type ,first_day_buying_count ,second_day_buying_count ,third_day_buying_count ,fourth_day_buying_count ,fifth_day_buying_count ,sixth_day_buying_count ,seventh_day_buying_count ,sunday_buying_count ,monday_buying_count ,tuesday_buying_count ,wednesday_buying_count ,thursday_buying_count ,friday_buying_count ,saturday_buying_count ,am_buying_count ,pm_buying_count ,buying_0_to_5_hour_count ,buying_6_to_11_hour_count ,buying_12_to_17_hour_count ,buying_18_to_23_hour_count ,buying_day_average_buying_count ,level_based_average_buying_count, status) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", m.Table),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId, data.LevelIndex, data.TotalBuyingCount,
		data.TotalBuyingDay, data.TotalBuyingHour, data.FirstBuyingYearOfDay, data.FirstBuyingYear,
		data.FirstBuyingHour, data.FirstBuyingMinute, data.FirstBuyingProductType,
		data.SecondBuyingYearOfDay, data.SecondBuyingHour, data.SecondBuyingMinute, data.SecondBuyingProductType,
		data.ThirdBuyingYearOfDay, data.ThirdBuyingHour, data.ThirdBuyingMinute, data.ThirdBuyingProductType,
		data.FourthBuyingYearOfDay, data.FourthBuyingHour, data.FourthBuyingMinute, data.FourthBuyingProductType,
		data.FifthBuyingYearOfDay, data.FifthBuyingHour, data.FifthBuyingMinute, data.FifthBuyingProductType,
		data.SixthBuyingYearOfDay, data.SixthBuyingHour, data.SixthBuyingMinute, data.SixthBuyingProductType,
		data.SeventhBuyingYearOfDay, data.SeventhBuyingHour, data.SeventhBuyingMinute, data.SeventhBuyingProductType,
		data.PenultimateBuyingYearOfDay, data.PenultimateBuyingHour, data.PenultimateBuyingMinute, 
		data.PenultimateBuyingProductType, data.LastBuyingYearOfDay, data.LastBuyingYear, data.LastBuyingHour, 
		data.LastBuyingMinute, data.LastBuyingProductType, data.FirstDayBuyingCount, data.SecondDayBuyingCount, 
		data.ThirdDayBuyingCount, data.FourthDayBuyingCount, data.FifthDayBuyingCount, data.SixthDayBuyingCount,
		data.SeventhDayBuyingCount, data.SundayBuyingCount,data.MondayBuyingCount, data.TuesdayBuyingCount,
		data.WednesdayBuyingCount, data.ThursdayBuyingCount, data.FridayBuyingCount, data.SaturdayBuyingCount, 
		data.AmBuyingCount, data.PmBuyingCount, data.Buying0To5HourCount, data.Buying6To11HourCount, 
		data.Buying12To17HourCount, data.Buying18To23HourCount, data.BuyingDayAverageBuyingCount, 
		data.LevelBasedAverageBuyingCount, data.Status).Exec(); err != nil {
	
		return err
	}

	return nil
}

func (m *cassBuyingEventDal) GetById(ClientId int64, ProjectId int64) (*model.BuyingEventResponseModel, error) {
	data := &model.BuyingEventResponseModel{}
	if err := m.Client.Query(fmt.Sprintf("SELECT id , client_id , project_id , customer_id , level_index ,total_buying_count ,total_buying_day ,total_buying_hour ,first_buying_year_of_day ,first_buying_year ,first_buying_hour ,first_buying_minute ,first_buying_product_type , second_buying_year_of_day ,second_buying_hour ,second_buying_minute , second_buying_product_type ,third_buying_year_of_day ,third_buying_hour ,third_buying_minute ,third_buying_product_type ,fourth_buying_year_of_day ,fourth_buying_hour ,fourth_buying_minute ,fourth_buying_product_type ,fifth_buying_year_of_day ,fifth_buying_hour ,fifth_buying_minute ,fifth_buying_product_type ,sixth_buying_year_of_day ,sixth_buying_hour ,sixth_buying_minute ,sixth_buying_product_type , seventh_buying_year_of_day , seventh_buying_hour ,seventh_buying_minute ,seventh_buying_product_type ,penultimate_buying_year_of_day , penultimate_buying_hour ,penultimate_buying_minute ,penultimate_buying_product_type ,last_buying_year_of_day ,last_buying_year ,last_buying_hour ,last_buying_minute ,last_buying_product_type ,first_day_buying_count , second_day_buying_count ,third_day_buying_count , fourth_day_buying_count ,fifth_day_buying_count , sixth_day_buying_count , seventh_day_buying_count ,sunday_buying_count ,monday_buying_count ,tuesday_buying_count ,  wednesday_buying_count ,thursday_buying_count , friday_buying_count ,saturday_buying_count , am_buying_count ,pm_buying_count ,buying_0_to_5_hour_count , buying_6_to_11_hour_count ,buying_12_to_17_hour_count ,buying_18_to_23_hour_count ,buying_day_average_buying_count ,level_based_average_buying_count ,status FROM MLDatabase.%s WHERE client_id = ? AND project_id = ? LIMIT 1", m.Table),
		ClientId, ProjectId).Scan(&data.Id, &data.ClientId, &data.ProjectId, &data.CustomerId, &data.LevelIndex, 
			&data.TotalBuyingCount, &data.TotalBuyingDay, &data.TotalBuyingHour, &data.FirstBuyingYearOfDay, 
			&data.FirstBuyingYear,&data.FirstBuyingHour, &data.FirstBuyingMinute, &data.FirstBuyingProductType,
			&data.SecondBuyingYearOfDay, &data.SecondBuyingHour, &data.SecondBuyingMinute, &data.SecondBuyingProductType,
			&data.ThirdBuyingYearOfDay, &data.ThirdBuyingHour, &data.ThirdBuyingMinute, &data.ThirdBuyingProductType,
			&data.FourthBuyingYearOfDay, &data.FourthBuyingHour, &data.FourthBuyingMinute, &data.FourthBuyingProductType,
			&data.FifthBuyingYearOfDay, &data.FifthBuyingHour, &data.FifthBuyingMinute, &data.FifthBuyingProductType,
			&data.SixthBuyingYearOfDay, &data.SixthBuyingHour, &data.SixthBuyingMinute, &data.SixthBuyingProductType,
			&data.SeventhBuyingYearOfDay, &data.SeventhBuyingHour, &data.SeventhBuyingMinute, &data.SeventhBuyingProductType,
			&data.PenultimateBuyingYearOfDay, &data.PenultimateBuyingHour, &data.PenultimateBuyingMinute, 
			&data.PenultimateBuyingProductType, &data.LastBuyingYearOfDay, &data.LastBuyingYear, &data.LastBuyingHour, 
			&data.LastBuyingMinute, &data.LastBuyingProductType, &data.FirstDayBuyingCount, &data.SecondDayBuyingCount, 
			&data.ThirdDayBuyingCount, &data.FourthDayBuyingCount, &data.FifthDayBuyingCount, &data.SixthDayBuyingCount, 
			&data.SeventhDayBuyingCount, &data.SundayBuyingCount, &data.MondayBuyingCount, &data.TuesdayBuyingCount,
			&data.WednesdayBuyingCount, &data.ThursdayBuyingCount, &data.FridayBuyingCount, 
			&data.SaturdayBuyingCount, &data.AmBuyingCount, &data.PmBuyingCount, &data.Buying0To5HourCount, 
			&data.Buying6To11HourCount, &data.Buying12To17HourCount, &data.Buying18To23HourCount, 
			&data.BuyingDayAverageBuyingCount, &data.LevelBasedAverageBuyingCount, &data.Status); err != nil {
	
		return nil, err
	}

	return data, nil
}

func (m *cassBuyingEventDal) UpdateById(ClientId int64, ProjectId int64, data *model.BuyingEventResponseModel) error {
	if err := m.Client.Query(fmt.Sprintf("UPDATE MLDatabase.%s SET id=?, customer_id=?, level_index=?, total_buying_count=? ,total_buying_day=?, total_buying_hour=?, first_buying_year_of_day=?, first_buying_year=?,  first_buying_hour=?,  first_buying_minute=? , first_buying_product_type=?, second_buying_year_of_day=?, second_buying_hour=?,  second_buying_minute=?, second_buying_product_type=?, third_buying_year_of_day=?, third_buying_hour=? ,third_buying_minute=? ,third_buying_product_type=? ,fourth_buying_year_of_day=? ,fourth_buying_hour=? ,fourth_buying_minute=? ,fourth_buying_product_type=? ,fifth_buying_year_of_day=? ,fifth_buying_hour=? ,fifth_buying_minute=? ,fifth_buying_product_type=? ,sixth_buying_year_of_day=? ,sixth_buying_hour=? ,sixth_buying_minute=? ,sixth_buying_product_type=?, seventh_buying_year_of_day=? ,seventh_buying_hour=? ,seventh_buying_minute=? ,seventh_buying_product_type=? ,penultimate_buying_year_of_day=? ,penultimate_buying_hour=? ,penultimate_buying_minute=? ,penultimate_buying_product_type=? ,last_buying_year_of_day=? ,last_buying_year=? ,last_buying_hour=? ,last_buying_minute=? ,last_buying_product_type=? ,first_day_buying_count=? ,second_day_buying_count=? ,third_day_buying_count=? ,fourth_day_buying_count=? ,fifth_day_buying_count=? ,sixth_day_buying_count=? ,seventh_day_buying_count=? ,sunday_buying_count=? ,monday_buying_count=? ,tuesday_buying_count=? ,wednesday_buying_count=? ,thursday_buying_count=? ,friday_buying_count=? ,saturday_buying_count=? ,am_buying_count=? ,pm_buying_count=? ,buying_0_to_5_hour_count=? ,buying_6_to_11_hour_count=? ,buying_12_to_17_hour_count=? ,buying_18_to_23_hour_count=? ,buying_day_average_buying_count=? ,level_based_average_buying_count=?, status=? WHERE client_id = ? AND project_id = ?", m.Table),
	data.Id, data.ClientId, data.ProjectId, data.CustomerId, data.LevelIndex, data.TotalBuyingCount,
	data.TotalBuyingDay, data.TotalBuyingHour, data.FirstBuyingYearOfDay, data.FirstBuyingYear,
	data.FirstBuyingHour, data.FirstBuyingMinute, data.FirstBuyingProductType,
	data.SecondBuyingYearOfDay, data.SecondBuyingHour, data.SecondBuyingMinute, data.SecondBuyingProductType,
	data.ThirdBuyingYearOfDay, data.ThirdBuyingHour, data.ThirdBuyingMinute, data.ThirdBuyingProductType,
	data.FourthBuyingYearOfDay, data.FourthBuyingHour, data.FourthBuyingMinute, data.FourthBuyingProductType,
	data.FifthBuyingYearOfDay, data.FifthBuyingHour, data.FifthBuyingMinute, data.FifthBuyingProductType,
	data.SixthBuyingYearOfDay, data.SixthBuyingHour, data.SixthBuyingMinute, data.SixthBuyingProductType,
	data.SeventhBuyingYearOfDay, data.SeventhBuyingHour, data.SeventhBuyingMinute, data.SeventhBuyingProductType,
	data.PenultimateBuyingYearOfDay, data.PenultimateBuyingHour, data.PenultimateBuyingMinute, 
	data.PenultimateBuyingProductType, data.LastBuyingYearOfDay, data.LastBuyingYear, data.LastBuyingHour, 
	data.LastBuyingMinute, data.LastBuyingProductType, data.FirstDayBuyingCount, data.SecondDayBuyingCount, 
	data.ThirdDayBuyingCount, data.FourthDayBuyingCount, data.FifthDayBuyingCount, data.SixthDayBuyingCount,
	data.SeventhDayBuyingCount, data.SundayBuyingCount,data.MondayBuyingCount, data.TuesdayBuyingCount,
	data.WednesdayBuyingCount, data.ThursdayBuyingCount, data.FridayBuyingCount, data.SaturdayBuyingCount, 
	data.AmBuyingCount, data.PmBuyingCount, data.Buying0To5HourCount, data.Buying6To11HourCount, 
	data.Buying12To17HourCount, data.Buying18To23HourCount, data.BuyingDayAverageBuyingCount, 
	data.LevelBasedAverageBuyingCount, data.Status, ClientId, ProjectId).Exec(); err != nil {

		return err
	}

	return nil
}
