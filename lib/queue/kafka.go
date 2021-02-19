package queue

import (
	"github.com/Shopify/sarama"
)

type Producer struct {
	BrokerServers string               `json:"broker_servers"`
	Config        *sarama.Config       `json:"config"`
	SyncProducer  sarama.SyncProducer  `json:"sync_producer"`
	AsyncProducer sarama.AsyncProducer `json:"async_producer"`
}

type Consumer struct {
	BrokerServers string `json:"broker_servers"`
	Config *sarama.Config `json:"config"`
}

const NamePrefix = "chat_"
