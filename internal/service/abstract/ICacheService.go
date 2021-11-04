package abstract

type ICacheService interface {
	ManageCache (tableName string, key string) (v int64, s bool, m string)
}
