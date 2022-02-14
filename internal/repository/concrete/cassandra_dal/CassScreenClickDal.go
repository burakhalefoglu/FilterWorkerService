package cassandra_dal

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/model/cassandraTables"

	"github.com/scylladb/gocqlx/v2"
)

func InsertScreenClickData(session *gocqlx.Session, data *model.ScreenClickRespondModel) error {
	q := session.Query(cassandraTables.ScreenClickTable.Insert()).BindStruct(data)
	if err := q.ExecRelease(); err != nil {
		return err
	}
	//defer session.Close()
	return nil
}