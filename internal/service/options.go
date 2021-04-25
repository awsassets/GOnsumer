package service

import (
	"GOnsumer/internal/service/kafka"
	"GOnsumer/internal/service/logger"
	"os"

	"github.com/Shopify/sarama"
	"github.com/rs/zerolog"
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

func Logger() Option {
	return func(o *Options) (err error) {
		o.Logger = &logger.LoggerService{
			Logger: zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger(),
		}
		return nil
	}
}
