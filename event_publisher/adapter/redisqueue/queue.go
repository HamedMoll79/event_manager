package redisqueue

import (
	"github.com/adjust/redismq"
)

func New(config Config, name string) *redismq.Queue {
	queue := redismq.CreateQueue(config.Host, config.Port, config.Password, config.DB, name)

	return queue
}

func NewConsumer(queue redismq.Queue, name string) (*redismq.Consumer, error) {
	consumer, err := queue.AddConsumer(name)

	return consumer, err
}
