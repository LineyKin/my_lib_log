package consumer

import (
	"fmt"
	"my_lib_log/lib/env"

	"github.com/IBM/sarama"
)

const TOPIC = "my_lib_log"

type KafkaConsumer struct {
	consumer sarama.Consumer
}

func connectConsumer(brokers []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	return sarama.NewConsumer(brokers, config)
}

func New() (*KafkaConsumer, error) {
	broker := fmt.Sprintf("localhost:%s", env.GetKafkaPort())
	worker, err := connectConsumer([]string{broker})
	if err != nil {
		return nil, err
	}

	kc := &KafkaConsumer{
		consumer: worker,
	}

	return kc, nil
}

func (kc *KafkaConsumer) Partition() (sarama.PartitionConsumer, error) {
	return kc.consumer.ConsumePartition(TOPIC, 0, sarama.OffsetOldest)
}

func (kc *KafkaConsumer) Close() error {
	return kc.consumer.Close()
}
