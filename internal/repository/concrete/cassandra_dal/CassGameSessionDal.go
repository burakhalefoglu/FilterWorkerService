package cassandra_dal

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/model/cassandraTables"

	"github.com/scylladb/gocqlx/v2"
)

func InsertGameSessionData(session *gocqlx.Session, data *model.GameSessionEveryLoginRespondModel) error {
	q := session.Query(cassandraTables.GameSessionTable.Insert()).BindStruct(data)
	if err := q.ExecRelease(); err != nil {
		return err
	}
	//defer session.Close()
	return nil
}