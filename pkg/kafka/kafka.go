package kafka

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

func writerConfigure(kafkaBrokerUrls []string, clientId string, topic string) (w *kafka.Writer, err error) {
	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: clientId,
	}

	config := kafka.WriterConfig{
		Brokers:          kafkaBrokerUrls,
		Topic:            topic,
		Balancer:         &kafka.LeastBytes{},
		Dialer:           dialer,
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}
	w = kafka.NewWriter(config)
	return w, nil
}

func readerConfigure(kafkaBrokerUrls []string, groupID string, topic string) (r *kafka.Reader, err error) {
	config := kafka.ReaderConfig{
		Brokers:         kafkaBrokerUrls,
		GroupID:         groupID,
		Topic:           topic,
		MinBytes:        10e3,            // 10KB
		MaxBytes:        10e6,            // 10MB
		MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		ReadLagInterval: -1,
	}

	reader := kafka.NewReader(config)
	return reader, nil
}

func Produce(parent context.Context, key, value []byte, topic string) (err error) {
	writer, _ := writerConfigure([]string{"192.168.1.34:9092"}, uuid.New().String(), topic)
	message := kafka.Message{
		Key:   key,
		Value: value,
		Time:  time.Now(),
	}
	err = writer.WriteMessages(parent, message)
	return err
}

func Consume(parent context.Context, topic string, groupId string, callback func(topic string, data []byte) error) {

	reader, _ := readerConfigure([]string{"192.168.1.34:9092"}, groupId, topic)
	defer reader.Close()
	log.Println(reader.Stats().ClientID)
	for {
		m, err := reader.FetchMessage(context.Background())
		if err != nil {
			log.Fatalf("error while receiving message: %s", err.Error())
			continue
		}
		if err != nil {
			log.Fatalf("error while receiving message: %s", err.Error())
			continue
		}
		// fmt.Println("topic: ", topic, " groupId: ", groupId, m.Value)
		callErr := callback(topic, m.Value)

		if callErr == nil {
			if err := reader.CommitMessages(parent, m); err != nil {
				log.Fatal("failed to commit messages:", err)
			}
		}
	}
}
