package cassandra_dal

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/model/cassandraTables"

	logger "github.com/appneuroncompany/light-logger"
	"github.com/appneuroncompany/light-logger/clogger"
	"github.com/scylladb/gocqlx/v2"
)

func InsertLevelBaseSessionData(session *gocqlx.Session, data *model.LevelBaseSessionRespondModel) error {
	q := session.Query(cassandraTables.LevelBaseSessionTable.Insert()).BindStruct(data)
	if err := q.ExecRelease(); err != nil {
		clogger.Error(&logger.Messages{
			"InsertLevelBaseSessionData err: ": err.Error(),
		})
		return err
	}
	//defer session.Close()
	return nil
}