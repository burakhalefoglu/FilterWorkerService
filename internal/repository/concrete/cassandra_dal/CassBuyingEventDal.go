package cassandra_dal

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/model/cassandraTables"

	logger "github.com/appneuroncompany/light-logger"
	"github.com/appneuroncompany/light-logger/clogger"
	"github.com/scylladb/gocqlx/v2"
)

func InsertBuyingEventData(session *gocqlx.Session, data *model.BuyingEventRespondModel) error {
	q := session.Query(cassandraTables.BuyingEventTable.Insert()).BindStruct(data)
	if err := q.ExecRelease(); err != nil {
		clogger.Error(&logger.Messages{
			"InsertBuyingEventData err: ": err.Error(),
		})
		return err
	}
	//defer session.Close()
	return nil
}