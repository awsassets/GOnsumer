package service

import "GOnsumer/service/kafka"

const (
	CONFIG_KAFKA_BROKERS       = "KAFKA_BROKERS"
	CONFIG_KAFKA_TOPICS        = "KAFKA_TOPICS"
	CONFIG_KAFKA_CONSUMERGROUP = "KAFKA_CONSUMER_GROUP"
)

type (
	Config struct {
		KafkaConfig *kafka.Config
	}
)
