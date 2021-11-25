package IoC

import (
	repository "FilterWorkerService/internal/repository/abstract"
	service "FilterWorkerService/internal/service/abstract"
	cache "FilterWorkerService/pkg/Cache"
	Ijsonparser "FilterWorkerService/pkg/jsonParser"
	IKafka "FilterWorkerService/pkg/kafka"
	"FilterWorkerService/pkg/logger"
)

type IContainer interface {
	Inject()
}

func InjectContainers(container IContainer){
	container.Inject()
}

var AdvEventDal repository.IAdvEventDal
var CacheService service.ICacheService
var AdvEventService service.IAdvEventService

var Kafka IKafka.IKafka
var RedisCache cache.ICache
var Logger logger.ILog
var JsonParser Ijsonparser.IJsonParser