package abstract

import "FilterWorkerService/internal/model"

type ITypeStandardizationDal interface {
	Add(tableName string, data *model.TypeStandardizationModel) error
	GetByKey(tableName string, key string) (*model.TypeStandardizationModel, error)
	GetAll(tableName string)(*[]model.TypeStandardizationModel, error)
	GetMaxByValue(tableName string)(int64, error)
}
