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

type golobbyInjection struct {}

func InjectionConstructor() *golobbyInjection {
	return &golobbyInjection{}
}

func(i *golobbyInjection)Inject(){
	injectLog()
	injectKafka()
	injectCache()
	injectJsonParser()

	injectAdvEventDal()
	injectAdvEventService()
}

func injectAdvEventDal() {
	if err :=container.Singleton(func() repository.IAdvEventDal {
		return mongodb_driver.MdbDAdvEventDalConstructor()
	}); err != nil{
		panic(err)
	}

	if err := container.Resolve(&IoC.AdvEventDal); err != nil{
		panic(err)
	}
}

func injectAdvEventService() {
	if err :=container.Singleton(func() service.IAdvEventService {
		return concrete.AdvEventManagerConstructor()
	}); err != nil{
		panic(err)
	}

	if err := container.Resolve(&IoC.AdvEventService); err != nil{
		panic(err)
	}
}

func injectKafka() {

	if err :=container.Singleton(func() IKafka.IKafka {
		return kafkago.KafkaGoConstructor(/*&IoC.Logger*/)
	}); err != nil{
		panic(err)
	}
	if err := container.Resolve(&IoC.Kafka); err != nil{
		panic(err)
	}
}

func injectCache() {

	if err :=container.Singleton(func() cache.ICache {
		return rediscachev8.RedisCacheConstructor(&IoC.Logger)
	}); err != nil{
		panic(err)
	}
	if err := container.Resolve(&IoC.RedisCache); err != nil{
		panic(err)
	}
}

func injectLog() {
	if err :=container.Singleton(func() logger.ILog {
		return logrus_logstash_hook.LogrusToLogstashLOGConstructor()
	}); err != nil{
		panic(err)
	}

	if err := container.Resolve(&IoC.Logger); err != nil{
		panic(err)
	}
}

func injectJsonParser() {
	if err :=container.Singleton(func() Ijsonparser.IJsonParser {
		return gojson.GoJsonConstructor()
	}); err != nil{
		panic(err)
	}

	if err := container.Resolve(&IoC.JsonParser); err != nil{
		panic(err)
	}
}


