package kafka

import (
	"fmt"
	"os"

	"github.com/Shopify/sarama"
)

type (
	KafkaService struct {
		Worker            sarama.Consumer
		PartitionConsumer sarama.PartitionConsumer
	}
	Config struct {
		Brokers       []string
		Topic         string
		ConsumerGroup string
	}
)

func (k *KafkaService) Consume(i func([]byte, string), sigchan chan os.Signal, doneCh chan struct{}) {
	for {
		select {
		case err := <-k.PartitionConsumer.Errors():
			fmt.Println(err)
		case msg := <-k.PartitionConsumer.Messages():
			i(msg.Value, msg.Topic)
		case <-sigchan:
			fmt.Println("Interrupt is detected")
			doneCh <- struct{}{}
		}
	}
}
