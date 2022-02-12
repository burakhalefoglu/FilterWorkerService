package abstract

type ICacheService interface {
	ManageCache(tableName string, key string) (v int16, s bool, m string)
}
