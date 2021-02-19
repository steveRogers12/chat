package router

import (
	messageController "ChatQueue/controllers/message"
	"github.com/Shopify/sarama"
)



func KafkaRouter (topic string) func(msg *sarama.ConsumerMessage){
	r := map[string]func(msg *sarama.ConsumerMessage){
		"test":messageController.Test,
	}
	return r[topic]
}