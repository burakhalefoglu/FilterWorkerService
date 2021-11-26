package golobby

import (
	"FilterWorkerService/internal/IoC"
	repository "FilterWorkerService/internal/repository/abstract"
	"FilterWorkerService/internal/repository/concrete/mongodb_driver"
	service "FilterWorkerService/internal/service/abstract"
	"FilterWorkerService/internal/service/concrete"
	cache "FilterWorkerService/pkg/Cache"
	rediscachev8 "FilterWorkerService/pkg/Cache/Redis/RedisV8"
	Ijsonparser "FilterWorkerService/pkg/jsonParser"
	"FilterWorkerService/pkg/jsonParser/gojson"
	IKafka "FilterWorkerService/pkg/kafka"
	"FilterWorkerService/pkg/kafka/kafkago"
	"FilterWorkerService/pkg/logger"
	"FilterWorkerService/pkg/logger/logrus_logstash_hook"

	"github.com/golobby/container/v3"
)

type golobbyInjection struct{}

func InjectionConstructor() *golobbyInjection {
	return &golobbyInjection{}
}

func (i *golobbyInjection) Inject() {
	injectLog()
	injectKafka()
	injectCache()
	injectJsonParser()

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

func injectAdvEventDal() {
	if err := container.Singleton(func() repository.IAdvEventDal {
		return mongodb_driver.MdbDAdvEventDalConstructor()
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
		return mongodb_driver.MdbDBuyingEventDalConstructor()
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
	if err := container.Singleton(func() repository.IGameSessionEveryLoginDal {
		return mongodb_driver.MdbDGameSessionEveryLoginDalConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.GameSessionEveryLoginDal); err != nil {
		panic(err)
	}
}

func injectGameSessionEveryLoginService() {
	if err := container.Singleton(func() service.IGameSessionEveryLoginService {
		return concrete.GameSessionEveryLoginManagerConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.GameSessionEveryLoginService); err != nil {
		panic(err)
	}
}

func injectHardwareInformationDal() {
	if err := container.Singleton(func() repository.IHardwareInformationDal {
		return mongodb_driver.MdbDHardwareInformationDalConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.HardwareInformationDal); err != nil {
		panic(err)
	}
}

func injectHardwareInformationService() {
	if err := container.Singleton(func() service.IHardwareInformationService {
		return concrete.HardwareInformationManagerConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.HardwareInformationService); err != nil {
		panic(err)
	}
}

func injectLevelBaseSessionDal() {
	if err := container.Singleton(func() repository.ILevelBaseSessionDal {
		return mongodb_driver.MdbDLevelBaseSessionDalConstructor()
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
		return mongodb_driver.MdbDLocationDalConstructor()
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
		return mongodb_driver.MdbDScreenClickDalConstructor()
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
		return mongodb_driver.MdbDScreenSwipeDalConstructor()
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

func injectCache() {

	if err := container.Singleton(func() cache.ICache {
		return rediscachev8.RedisCacheConstructor(&IoC.Logger)
	}); err != nil {
		panic(err)
	}
	if err := container.Resolve(&IoC.RedisCache); err != nil {
		panic(err)
	}
}

func injectLog() {
	if err := container.Singleton(func() logger.ILog {
		return logrus_logstash_hook.LogrusToLogstashLOGConstructor()
	}); err != nil {
		panic(err)
	}

	if err := container.Resolve(&IoC.Logger); err != nil {
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
