package cassandra_dal

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/pkg/database/cassandra"
	"fmt"

	//logger "github.com/appneuroncompany/light-logger"
	//"github.com/appneuroncompany/light-logger/clogger"
	"github.com/gocql/gocql"
)

type cassLocationDal struct {
	Client *gocql.Session
	Table  string
}

func NewCassLocationDal(Table string) *cassLocationDal {
	return &cassLocationDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassLocationDal) Add(data *model.LocationResponseModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO MLDatabase.%s(id, client_id, project_id, customer_id, continent ,country ,city ,region ,org , status) VALUES(?,?,?,?,?,?,?,?,?,?)", m.Table),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId,
		data.Continent,
		data.Country,
		data.City,
		data.Region,
		data.Org,
		data.Status).Exec(); err != nil {
		
		return err
	}
	
	return nil
}

func (m *cassLocationDal) GetById(ClientId int64, ProjectId int64) (*model.LocationResponseModel, error) {
	data := &model.LocationResponseModel{}
	if err := m.Client.Query(fmt.Sprintf("SELECT * FROM MLDatabase.%s WHERE client_id = ? AND project_id = ? LIMIT 1", m.Table),
		ClientId, ProjectId).Scan(&data.Id, &data.ClientId, &data.ProjectId, &data.CustomerId,
		&data.Continent,
		&data.Country,
		&data.City,
		&data.Region,
		&data.Org, &data.Status); err != nil {
		
		return nil, err
	}
	
	return data, nil
}

func (m *cassLocationDal) UpdateById(ClientId int64, ProjectId int64, data *model.LocationResponseModel) error {
	if err := m.Client.Query(fmt.Sprintf("UPDATE MLDatabase.%s SET id=?, customer_id=?, continent=? ,country=? ,city=? ,region=? ,org=? , status=? WHERE client_id = ? AND project_id = ?", m.Table),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId,
		data.Continent,
		data.Country,
		data.City,
		data.Region,
		data.Org, data.Status, ClientId, ProjectId).Exec(); err != nil {

		return err
	}

	return nil
}
