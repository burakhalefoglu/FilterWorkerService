package cassandra_dal

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/model/cassandraTables"

	"github.com/scylladb/gocqlx/v2"
)

func InsertAdvEventData(session *gocqlx.Session, data *model.AdvEventRespondModel) error {
	q := session.Query(cassandraTables.AdvEventTable.Insert()).BindStruct(data)
	if err := q.ExecRelease(); err != nil {
		return err
	}
	//defer session.Close()
	return nil
}