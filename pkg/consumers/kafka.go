package consumers

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func GetKafkaConsumer(groupId string) *kafka.Consumer {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
	"bootstrap.servers":    "host1:9092,host2:9092",
	"group.id":             groupId,
	"auto.offset.reset":    "smallest"})
	if err == nil {
		return consumer;
	}
	panic(err)
}