package queue

import (
	"github.com/Shopify/sarama"
	"log"
	"strings"
	"time"
)

func (p *Producer) GetSysProducer() error {
	var err error
	p.SyncProducer, err = sarama.NewSyncProducer(strings.Split(p.BrokerServers, ","), p.Config)
	if err != nil {
		log.Printf("Create producer failed %s \n", err.Error())
		return err
	}
	return nil
}

//批量发送数据到kafka
func (p *Producer) BathProduceMsg(topic, key string, messages []string) error {
	if p.SyncProducer == nil {
		if e := p.GetSysProducer(); e != nil {
			log.Fatalf("Create kafka producer failed: %s \n", e.Error())
		}
	}
	defer p.SyncProducer.Close()
	msgs := make([]*sarama.ProducerMessage, len(messages))
	for i, m := range messages {
		msg := sarama.ProducerMessage{}
		msg.Topic = topic
		msg.Key = sarama.StringEncoder(key)
		msg.Value = sarama.StringEncoder(m)
		msg.Timestamp = time.Now()
		msgs[i] = &msg
	}
	return p.SyncProducer.SendMessages(msgs)
}

//单条发送数据到kafka
func (p *Producer) ProducerMsg(topic, key string, msg string) (partition int32, offset int64, err error) {
	if p.SyncProducer == nil {
		if e := p.GetSysProducer(); e != nil {
			log.Fatalf("Create kafka producer failed: %s \n", e.Error())
		}
	}
	defer p.SyncProducer.Close()
	m := &sarama.ProducerMessage{}
	m.Topic = topic
	m.Key = sarama.StringEncoder(key)
	m.Value = sarama.StringEncoder(msg)
	m.Timestamp = time.Now()
	return p.SyncProducer.SendMessage(m)
}