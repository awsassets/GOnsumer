package service

import (
	"GOnsumer/service/kafka"

	"github.com/Shopify/sarama"
)

func Kafka() Option {
	return func(o *Options) (err error) {
		config := sarama.NewConfig()
		config.Consumer.Return.Errors = true //TODO debug
		o.Kafka = &kafka.KafkaService{}
		o.Kafka.Worker, err = sarama.NewConsumer(o.Cfg.KafkaConfig.Brokers, config)
		if err != nil {
			return
		}

		o.Kafka.PartitionConsumer, err = o.Kafka.Worker.ConsumePartition(o.Cfg.KafkaConfig.Topic, 0, sarama.OffsetNewest)
		if err != nil {
			return
		}

		return nil
	}
}
