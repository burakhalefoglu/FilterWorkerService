package contorller

import (
	"FilterWorkerService/internal/IoC"
	service "FilterWorkerService/internal/service/abstract"
	"FilterWorkerService/pkg/helper"
	Kafka "FilterWorkerService/pkg/kafka"
	"sync"
)

type kafkaController struct {
	Kafka                   *Kafka.IKafka
	AdvEventService         *service.IAdvEventService
	BuyingEventService      *service.IBuyingEventService
	GameSessionService      *service.IGameSessionService
	HardwareService         *service.IHardwareService
	LevelBaseSessionService *service.ILevelBaseSessionService
	LocationService         *service.ILocationService
	ScreenClickService      *service.IScreenClickService
	ScreenSwipeService      *service.IScreenSwipeService
}

func KafkaControllerConstructor() *kafkaController {
	return &kafkaController{
		Kafka:                   &IoC.Kafka,
		AdvEventService:         &IoC.AdvEventService,
		BuyingEventService:      &IoC.BuyingEventService,
		GameSessionService:      &IoC.GameSessionService,
		HardwareService:         &IoC.HardwareService,
		LevelBaseSessionService: &IoC.LevelBaseSessionService,
		LocationService:         &IoC.LocationService,
		ScreenClickService:      &IoC.ScreenClickService,
		ScreenSwipeService:      &IoC.ScreenSwipeService,
	}
}

func (k *kafkaController) StartListen(waitGroup *sync.WaitGroup) {

	waitGroup.Add(8)
	helper.CreateHealthFile()
	go (*k.Kafka).Consume("AdvEventDataModel",
		"AdvEventDataModel_Filter_ConsumerGroup",
		waitGroup,
		(*k.AdvEventService).ConvertRawModelToResponseModel)

	go (*k.Kafka).Consume("BuyingEventDataModel",
		"BuyingEventDataModel_Filter_ConsumerGroup",
		waitGroup,
		(*k.BuyingEventService).ConvertRawModelToResponseModel)

	go (*k.Kafka).Consume("GameSessionDataModel",
		"GameSessionDataModel_Filter_ConsumerGroup",
		waitGroup,
		(*k.GameSessionService).ConvertRawModelToResponseModel)

	go (*k.Kafka).Consume("HardwareModel",
		"HardwareModel_Filter_ConsumerGroup",
		waitGroup,
		(*k.HardwareService).ConvertRawModelToResponseModel)

	go (*k.Kafka).Consume("LevelBaseSessionDataModel",
		"LevelBaseSessionDataModel_Filter_ConsumerGroup",
		waitGroup,
		(*k.LevelBaseSessionService).ConvertRawModelToResponseModel)

	go (*k.Kafka).Consume("LocationDataModel",
		"LocationDataModel_Filter_ConsumerGroup",
		waitGroup,
		(*k.LocationService).ConvertRawModelToResponseModel)

	go (*k.Kafka).Consume("ScreenClickDataModel",
		"ScreenClickDataModel_Filter_ConsumerGroup",
		waitGroup,
		(*k.ScreenClickService).ConvertRawModelToResponseModel)

	go (*k.Kafka).Consume("ScreenSwipeDataModel",
		"ScreenSwipeDataModel_Filter_ConsumerGroup",
		waitGroup,
		(*k.ScreenSwipeService).ConvertRawModelToResponseModel)

}
