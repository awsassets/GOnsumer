package service

import (
	"GOnsumer/internal/service/kafka"
	"GOnsumer/internal/service/logger"

	"os"
	"strings"

	"github.com/joho/godotenv"
)

type (
	Options struct {
		Name    string
		Version string
		Kafka   *kafka.KafkaService
		Logger  *logger.LoggerService
		Cfg     *Config
	}

	Option func(*Options) error

	Service struct {
		Options
		Sigchan chan os.Signal
		DoneCh  chan struct{}
	}
)

const (
	envFilePath = "../../.env"
)

func (o *Options) connect() (s *Service, err error) {
	return &Service{
		Options: *o,
	}, nil
}

func New(name, version string, options ...Option) (s *Service, err error) {
	cfg, err := loadConfigs()
	if err != nil {
		return
	}
	return newDefaults(name, cfg, options...)
}

func loadConfigs() (c *Config, err error) {
	err = godotenv.Load(envFilePath)
	if err != nil {
		return
	}

	c = &Config{}
	c.KafkaConfig = &kafka.Config{
		Brokers:       strings.Split(os.Getenv(CONFIG_KAFKA_BROKERS), ","),
		Topic:         os.Getenv(CONFIG_KAFKA_TOPICS),
		ConsumerGroup: os.Getenv(CONFIG_KAFKA_CONSUMERGROUP),
	}

	return
}

func newDefaults(name string, cfg *Config, options ...Option) (s *Service, err error) {
	opts := Options{
		Name: name,
		Cfg:  cfg,
	}

	for _, opt := range options {
		if opt != nil {
			if err := opt(&opts); err != nil {
				return nil, err
			}
		}
	}

	s, err = opts.connect()
	s.Sigchan = make(chan os.Signal, 1)
	s.DoneCh = make(chan struct{})

	return
}
