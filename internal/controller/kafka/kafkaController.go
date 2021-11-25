package kafka

import (
	"FilterWorkerService/internal/IoC"
	Kafka "FilterWorkerService/pkg/kafka"
)

type kafkaController struct {
	Kafka *Kafka.IKafka

}

func KafkaControllerConstructor(
	) *kafkaController {
	return &kafkaController{Kafka: &IoC.Kafka}
}
