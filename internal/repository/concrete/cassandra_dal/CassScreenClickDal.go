package cassandra_dal

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/model/cassandraTables"

	logger "github.com/appneuroncompany/light-logger"
	"github.com/appneuroncompany/light-logger/clogger"
	"github.com/scylladb/gocqlx/v2"
)

func InsertScreenClickData(session *gocqlx.Session, data *model.ScreenClickRespondModel) error {
	q := session.Query(cassandraTables.ScreenClickTable.Insert()).BindStruct(data)
	if err := q.ExecRelease(); err != nil {
		clogger.Error(&logger.Messages{
			"InsertScreenClickData err: ": err.Error(),
		})
		return err
	}
	//defer session.Close()
	return nil
}