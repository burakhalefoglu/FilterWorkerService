package IoC

import (
	repository "FilterWorkerService/internal/repository/abstract"
	service "FilterWorkerService/internal/service/abstract"
	cache "FilterWorkerService/pkg/Cache"
	Ijsonparser "FilterWorkerService/pkg/jsonParser"
	IKafka "FilterWorkerService/pkg/kafka"
)

type IContainer interface {
	Inject()
}

func InjectContainers(container IContainer) {
	container.Inject()
}

var AdvEventDal repository.IAdvEventDal
var AdvEventService service.IAdvEventService
var BuyingEventDal repository.IBuyingEventDal
var BuyingEventService service.IBuyingEventService
var GameSessionDal repository.IGameSessionDal
var GameSessionService service.IGameSessionService
var HardwareDal repository.IHardwareDal
var HardwareService service.IHardwareService
var LevelBaseSessionDal repository.ILevelBaseSessionDal
var LevelBaseSessionService service.ILevelBaseSessionService
var LocationDal repository.ILocationDal
var LocationService service.ILocationService
var ScreenClickDal repository.IScreenClickDal
var ScreenClickService service.IScreenClickService
var ScreenSwipeDal repository.IScreenSwipeDal
var ScreenSwipeService service.IScreenSwipeService
var TypeStandardizationDal repository.ITypeStandardizationDal
var CacheService service.ICacheService
var Kafka IKafka.IKafka
var RedisCache cache.ICache
var JsonParser Ijsonparser.IJsonParser
