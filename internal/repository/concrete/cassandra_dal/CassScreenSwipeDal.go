package cassandra_dal

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/model/cassandraTables"

	"github.com/scylladb/gocqlx/v2"
)

func InsertScreenSwipeData(session *gocqlx.Session, data *model.ScreenSwipeRespondModel) error {
	q := session.Query(cassandraTables.ScreenSwipeTable.Insert()).BindStruct(data)
	if err := q.ExecRelease(); err != nil {
		return err
	}
	//defer session.Close()
	return nil
}