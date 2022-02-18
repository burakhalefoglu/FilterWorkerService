package cassandra_dal

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/pkg/database/cassandra"
	"fmt"

	//logger "github.com/appneuroncompany/light-logger"
	//"github.com/appneuroncompany/light-logger/clogger"
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
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO %s (key, value) VALUES(?,?)", tableName),
		data.Key, data.Value).Exec(); err != nil {
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

func (m *cassTypeStandardizationDal) GetByKey(tableName string, key string) (*model.TypeStandardizationModel, error) {
	data := &model.TypeStandardizationModel{}
	if err := m.Client.Query(fmt.Sprintf("SELECT * FROM %s WHERE key = ? LIMIT 1", tableName),
		key).Scan(&data.Value); err != nil {
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

func (m *cassTypeStandardizationDal) GetAll(tableName string) (*[]model.TypeStandardizationModel, error) {
	var models []model.TypeStandardizationModel
	c := map[string]interface{}{}
 
	iter := m.Client.Query(fmt.Sprintf("SELECT * FROM %s", tableName)).Iter()
	for iter.MapScan(c) {
		models = append(models, model.TypeStandardizationModel{
			Key:   c["key"].(string),
			Value: c["value"].(int16),
		})
		c = map[string]interface{}{}
	}
	return &models, nil
}


func (m *cassTypeStandardizationDal) GetMaxByValue(tableName string) (int16, error) {

	data := &model.TypeStandardizationModel{}
	if err := m.Client.Query(fmt.Sprintf("SELECT MAX(value) FROM MLDatabase.%s LIMIT 1", tableName)).Scan(&data.Key,
		&data.Value); err != nil {
		// clogger.Error(&logger.Messages{
		// 	"Get adv_event_data err: ": err.Error(),
		// })
		return 0, err
	}
	// clogger.Info(&logger.Messages{
	// 	"Get adv_event_data  : ": "SUCCESS",
	// })
	return int16(data.Value), nil


}
