package cassandra_dal

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/model/cassandraTables"

	"github.com/scylladb/gocqlx/v2"
)

func InsertLevelBaseSessionData(session *gocqlx.Session, data *model.LevelBaseSessionRespondModel) error {
	q := session.Query(cassandraTables.LevelBaseSessionTable.Insert()).BindStruct(data)
	if err := q.ExecRelease(); err != nil {
		return err
	}
	//defer session.Close()
	return nil
}