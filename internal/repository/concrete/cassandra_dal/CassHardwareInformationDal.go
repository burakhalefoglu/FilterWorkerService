package cassandra_dal

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/model/cassandraTables"

	"github.com/scylladb/gocqlx/v2"
)

func InsertHardwareInformationData(session *gocqlx.Session, data *model.HardwareInformationResponseModel) error {
	q := session.Query(cassandraTables.HardwareInformationTable.Insert()).BindStruct(data)
	if err := q.ExecRelease(); err != nil {
		return err
	}
	//defer session.Close()
	return nil
}