package golobby

import (
	"FilterWorkerService/internal/IoC"
	repository "FilterWorkerService/internal/repository/abstract"
	"FilterWorkerService/internal/repository/concrete/cassandra_dal"
	service "FilterWorkerService/internal/service/abstract"
	"FilterWorkerService/internal/service/concrete"
	cache "FilterWorkerService/pkg/Cache"
	rediscachev8 "FilterWorkerService/pkg/Cache/Redis/RedisV8"
	"FilterWorkerService/pkg/database/cassandra"
	Ijsonparser "FilterWorkerService/pkg/jsonParser"
	"FilterWorkerService/pkg/jsonParser/gojson"
	IKafka "FilterWorkerService/pkg/kafka"
	"FilterWorkerService/pkg/kafka/kafkago"

	"github.com/golobby/container/v3"
)

type golobbyInjection struct{}

func InjectionConstructor() *golobbyInjection {
	return &golobbyInjection{}
}

func (i *golobbyInjection) Inject() {
	injectJsonParser()
	injectKafka()
	injectCache()

	injectTypeStandardiseDal()
	injectCacheService()
	injectAdvEventDal()
	injectAdvEventService()
	injectBuyingEventDal()
	injectBuyingEventService()
	injectGameSessionEveryLoginDal()
	injectGameSessionEveryLoginService()
	injectHardwareInformationDal()
	injectHardwareInformationService()
	injectLevelBaseSessionDal()
	injectLevelBaseSessionService()
	injectLocationDal()
	injectLocationManagerService()
	injectScreenClickDal()
	injectScreenClickManagerService()
	injectScreenSwipeDal()
	injectScreenSwipeManagerService()
}

func injectTypeStandardiseDal() {
	if err := container.Singleton(func() repository.ITypeStandardizationDal {
		return cassandra_dal.NewCassTypeStandardizationDal(cassandra.TypeStandardizationModel)
	}); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.TypeStandardizationDal); err != nil {
		panic(err)
	}
}

func injectCacheService() {

	if err := container.Singleton(func() service.ICacheService {
		return concrete.CacheManagerConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.CacheService); err != nil {
		panic(err)
	}
}

func injectAdvEventDal() {
	if err := container.Singleton(func() repository.IAdvEventDal {
		return cassandra_dal.NewCassAdvEventDal(cassandra.AdvEventResponseModel)
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.AdvEventDal); err != nil {
		panic(err)
	}
}

func injectAdvEventService() {
	if err := container.Singleton(func() service.IAdvEventService {
		return concrete.AdvEventManagerConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.AdvEventService); err != nil {
		panic(err)
	}
}

func injectBuyingEventDal() {
	if err := container.Singleton(func() repository.IBuyingEventDal {
		return cassandra_dal.NewCassBuyingEventDal(cassandra.BuyingEventResponseModel)
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.BuyingEventDal); err != nil {
		panic(err)
	}
}

func injectBuyingEventService() {
	if err := container.Singleton(func() service.IBuyingEventService {
		return concrete.BuyingEventManagerConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.BuyingEventService); err != nil {
		panic(err)
	}
}

func injectGameSessionEveryLoginDal() {
	if err := container.Singleton(func() repository.IGameSessionDal {
		return cassandra_dal.NewCassGameSessionDal(cassandra.GameSessionResponseModel)
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.GameSessionDal); err != nil {
		panic(err)
	}
}

func injectGameSessionEveryLoginService() {
	if err := container.Singleton(func() service.IGameSessionService {
		return concrete.GameSessionManagerConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.GameSessionService); err != nil {
		panic(err)
	}
}

func injectHardwareInformationDal() {
	if err := container.Singleton(func() repository.IHardwareDal {
		return cassandra_dal.NewCassHardwareDal(cassandra.HardwareResponseModel)
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.HardwareDal); err != nil {
		panic(err)
	}
}

func injectHardwareInformationService() {
	if err := container.Singleton(func() service.IHardwareService {
		return concrete.HardwareManagerConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.HardwareService); err != nil {
		panic(err)
	}
}

func injectLevelBaseSessionDal() {
	if err := container.Singleton(func() repository.ILevelBaseSessionDal {
		return cassandra_dal.NewCassLevelBaseSessionDal(cassandra.LevelBaseSessionResponseModel)
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.LevelBaseSessionDal); err != nil {
		panic(err)
	}
}

func injectLevelBaseSessionService() {
	if err := container.Singleton(func() service.ILevelBaseSessionService {
		return concrete.LevelBaseSessionManagerConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.LevelBaseSessionService); err != nil {
		panic(err)
	}
}

func injectLocationDal() {
	if err := container.Singleton(func() repository.ILocationDal {
		return cassandra_dal.NewCassLocationDal(cassandra.LocationResponseModel)
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.LocationDal); err != nil {
		panic(err)
	}
}

func injectLocationManagerService() {
	if err := container.Singleton(func() service.ILocationService {
		return concrete.LocationManagerConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.LocationService); err != nil {
		panic(err)
	}
}

func injectScreenClickDal() {
	if err := container.Singleton(func() repository.IScreenClickDal {
		return cassandra_dal.NewCassScreenClickDal(cassandra.ScreenClickResponseModel)
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.ScreenClickDal); err != nil {
		panic(err)
	}
}

func injectScreenClickManagerService() {
	if err := container.Singleton(func() service.IScreenClickService {
		return concrete.ScreenClickManagerConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.ScreenClickService); err != nil {
		panic(err)
	}
}

func injectScreenSwipeDal() {
	if err := container.Singleton(func() repository.IScreenSwipeDal {
		return cassandra_dal.NewCassScreenSwipeDal(cassandra.ScreenSwipeResponseModel)
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.ScreenSwipeDal); err != nil {
		panic(err)
	}
}

func injectScreenSwipeManagerService() {
	if err := container.Singleton(func() service.IScreenSwipeService {
		return concrete.ScreenSwipeManagerConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.ScreenSwipeService); err != nil {
		panic(err)
	}
}

func injectKafka() {

	if err := container.Singleton(func() IKafka.IKafka {
		return kafkago.KafkaGoConstructor( /*&IoC.Logger*/ )
	}); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.Kafka); err != nil {
		panic(err)
	}
}

func injectJsonParser() {
	if err := container.Singleton(func() Ijsonparser.IJsonParser {
		return gojson.GoJsonConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.JsonParser); err != nil {
		panic(err)
	}
}

func injectCache() {

	if err := container.Singleton(func() cache.ICache {
		return rediscachev8.RedisCacheConstructor()
	}); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.RedisCache); err != nil {
		panic(err)
	}
}
