package queue

import (
	"ChatQueue/router"
	"github.com/Shopify/sarama"
	"log"
	"strings"
)

//获取符合条件的topic以及对应的分区数
func (c * Consumer)GetTopicsAndPartitions() (map[string]int, error) {
	consumer, e := sarama.NewConsumer(strings.Split(c.BrokerServers, ","), nil)
	defer consumer.Close()
	if e != nil {
		log.Printf("Create kafka consumer failed: %s \n", e.Error())
		return nil, e
	}
	topics, _ := consumer.Topics()//获取所有topic
	whTopic := make(map[string]int)
	for _, t := range topics {
		if strings.HasPrefix(t, NamePrefix) {
			if p, e := consumer.Partitions(t); e == nil {
				whTopic[t] = len(p)
			}
		}
	}
	return whTopic, nil
}

//消费kafka集群下满足条件的所有topic的数据
func (c *Consumer)Consume(done chan bool) {
	whTopic, e := c.GetTopicsAndPartitions()
	if e != nil {
		return
	}
	for t, _ := range whTopic {
		handle := router.KafkaRouter(t)
		consumer, e := sarama.NewConsumer(strings.Split(c.BrokerServers, ","), nil)
		if e != nil {
			log.Printf("Create consumer for %s topic failed: %s \n", t, e.Error())
		}
		partitions, e := consumer.Partitions(t)
		if e != nil {
			log.Printf("Get %s topic's paritions failed: %s \n", t, e.Error())
		}
		for _, p := range partitions {
			pc, e := consumer.ConsumePartition(t, p, sarama.OffsetOldest)
			if e != nil {
				log.Printf("Start consumer for partition failed %d: %s \n", p, t)
				return
			}
			go func(pc sarama.PartitionConsumer) {
				for msg := range pc.Messages() {
					handle(msg)
				}
				done <- true
			}(pc)
		}
	}
}