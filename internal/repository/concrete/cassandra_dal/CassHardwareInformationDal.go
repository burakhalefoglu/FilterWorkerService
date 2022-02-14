package cassandra_dal

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/model/cassandraTables"

	logger "github.com/appneuroncompany/light-logger"
	"github.com/appneuroncompany/light-logger/clogger"
	"github.com/scylladb/gocqlx/v2"
)

func InsertHardwareInformationData(session *gocqlx.Session, data *model.HardwareInformationResponseModel) error {
	q := session.Query(cassandraTables.HardwareInformationTable.Insert()).BindStruct(data)
	if err := q.ExecRelease(); err != nil {
		clogger.Error(&logger.Messages{
			"InsertHardwareInformationData err: ": err.Error(),
		})
		return err
	}
	//defer session.Close()
	return nil
}