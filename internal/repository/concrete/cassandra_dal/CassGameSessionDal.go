package cassandra_dal

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/model/cassandraTables"

	logger "github.com/appneuroncompany/light-logger"
	"github.com/appneuroncompany/light-logger/clogger"
	"github.com/scylladb/gocqlx/v2"
)

func InsertGameSessionData(session *gocqlx.Session, data *model.GameSessionEveryLoginRespondModel) error {
	q := session.Query(cassandraTables.GameSessionTable.Insert()).BindStruct(data)
	if err := q.ExecRelease(); err != nil {
		clogger.Error(&logger.Messages{
			"InsertGameSessionData err: ": err.Error(),
		})
		return err
	}
	//defer session.Close()
	return nil
}