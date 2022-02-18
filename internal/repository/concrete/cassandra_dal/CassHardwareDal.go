package cassandra_dal

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/pkg/database/cassandra"
	"fmt"

	//logger "github.com/appneuroncompany/light-logger"
	//"github.com/appneuroncompany/light-logger/clogger"
	"github.com/gocql/gocql"
)

type cassHardwareDal struct {
	Client *gocql.Session
	Table  string
}

func NewCassHardwareDal(Table string) *cassHardwareDal {
	return &cassHardwareDal{Client: cassandra.ConnectDatabase(),
		Table: Table}
}

func (m *cassHardwareDal) Add(data *model.HardwareResponseModel) error {
	if err := m.Client.Query(fmt.Sprintf("INSERT INTO %s (id, client_id, project_id, customer_id, device_type ,graphics_device_type ,graphics_memory_size ,operating_system ,processor_count ,processor_type ,system_memory_size , status) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)", m.Table),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId,
		data.DeviceType,
		data.GraphicsDeviceType,
		data.GraphicsMemorySize,
		data.OperatingSystem,
		data.ProcessorCount,
		data.ProcessorType,
		data.SystemMemorySize,
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

func (m *cassHardwareDal) GetById(ClientId int64, ProjectId int64) (*model.HardwareResponseModel, error) {
	data := &model.HardwareResponseModel{}
	if err := m.Client.Query(fmt.Sprintf("SELECT * FROM %s WHERE client_id = ? AND project_id = ? LIMIT 1", m.Table),
		ClientId, ProjectId).Scan(&data.Id, &data.ClientId, &data.ProjectId, &data.CustomerId,
		&data.DeviceType,
		&data.GraphicsDeviceType,
		&data.GraphicsMemorySize,
		&data.OperatingSystem,
		&data.ProcessorCount,
		&data.ProcessorType,
		&data.SystemMemorySize, &data.Status); err != nil {
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

func (m *cassHardwareDal) UpdateById(ClientId int64, ProjectId int64, data *model.HardwareResponseModel) error {
	if err := m.Client.Query(fmt.Sprintf("UPDATE %s SET id=?, customer_id=?, device_type=? ,graphics_device_type=? ,graphics_memory_size=? ,operating_system=? ,processor_count=? ,processor_type=? ,system_memory_size=?, status=? WHERE client_id = %d AND project_id = %d", m.Table, ClientId, ProjectId),
		data.Id, data.ClientId, data.ProjectId, data.CustomerId,
		data.DeviceType,
		data.GraphicsDeviceType,
		data.GraphicsMemorySize,
		data.OperatingSystem,
		data.ProcessorCount,
		data.ProcessorType,
		data.SystemMemorySize, data.Status).Exec(); err != nil {

		return err
	}

	return nil
}
