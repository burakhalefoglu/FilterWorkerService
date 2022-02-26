package cassandra_dal

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/pkg/database/cassandra"
	"fmt"

	"github.com/gocql/gocql"
)

type cassTypeStandardizationDal struct {
	Client *gocql.Session
	Table  string
}

func NewCassTypeStandardizationDal(Table string) *cassTypeStandardizationDal {
	return &cassTypeStandardizationDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassTypeStandardizationDal) Add(tableName string, data *model.TypeStandardizationModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO MLDatabase.%s(key, value) VALUES(?,?)", tableName),
		data.Key, data.Value).Exec(); err != nil {
		return err
	}
	return nil
}

func (m *cassTypeStandardizationDal) GetByKey(tableName string, key string) (*model.TypeStandardizationModel, error) {
	data := &model.TypeStandardizationModel{}
	if err := m.Client.Query(fmt.Sprintf("SELECT key, value FROM MLDatabase.%s WHERE key = ? LIMIT 1", tableName),
		key).Scan(&data.Key, &data.Value); err != nil {
		return nil, err
	}
	return data, nil
}

func (m *cassTypeStandardizationDal) GetAll(tableName string) (*[]model.TypeStandardizationModel, error) {
	var data model.TypeStandardizationModel
	var models []model.TypeStandardizationModel
	
	iter := m.Client.Query(fmt.Sprintf("SELECT key, value FROM MLDatabase.%s", tableName)).Iter()
	for iter.Scan(&data.Key, &data.Value) {
		models = append(models, data)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}
	return &models, nil
}


func (m *cassTypeStandardizationDal) GetMaxByValue(tableName string) (int16, error) {

	data := &model.TypeStandardizationModel{}
	if err := m.Client.Query(fmt.Sprintf("SELECT MAX(value) FROM MLDatabase.%s LIMIT 1", tableName)).Scan(&data.Value); err != nil {
		return 0, err
	}
	return int16(data.Value), nil

}
