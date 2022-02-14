package cassandra_dal

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/model/cassandraTables"

	"github.com/scylladb/gocqlx/v2"
)

func InsertTypeStandardizationData(session *gocqlx.Session, data *model.TypeStandardizationModel) error {
	q := session.Query(cassandraTables.TypeStandardizationTable.Insert()).BindStruct(data)
	if err := q.ExecRelease(); err != nil {
		return err
	}
	//defer session.Close()
	return nil
}