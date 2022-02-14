package cassandra_dal

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/model/cassandraTables"
	logger "github.com/appneuroncompany/light-logger"
	"github.com/appneuroncompany/light-logger/clogger"

	"github.com/scylladb/gocqlx/v2"
)



func InsertAdvEventData(session *gocqlx.Session, data *model.AdvEventRespondModel) error {
	q := session.Query(cassandraTables.AdvEventTable.Insert()).BindStruct(data)
	if err := q.ExecRelease(); err != nil {
		clogger.Error(&logger.Messages{
			"InsertAdvEventData err: ": err.Error(),
		})
		return err
	}
	//defer session.Close()
	return nil
}

func GetAdvEventData(session *gocqlx.Session, data *model.AdvEventRespondModel) (*model.AdvEventRespondModel, error) {
	q := session.Query(cassandraTables.AdvEventTable.Get()).BindStruct(data)
	if err := q.GetRelease(&data); err != nil {
		clogger.Error(&logger.Messages{
			"GetAdvEventData err: ": err.Error(),
		})
		return data, err
	}
	return data, nil
}



