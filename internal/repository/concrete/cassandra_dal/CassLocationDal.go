package cassandra_dal

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/model/cassandraTables"

	"github.com/scylladb/gocqlx/v2"
)

func InsertLocationData(session *gocqlx.Session, data *model.LocationResponseModel) error {
	q := session.Query(cassandraTables.LocationTable.Insert()).BindStruct(data)
	if err := q.ExecRelease(); err != nil {
		return err
	}
	//defer session.Close()
	return nil
}