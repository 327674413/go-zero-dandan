package queued

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

// Consumer 消费者
type Consumer struct {
	Client sarama.ConsumerGroup
}

// NewConsumer 创建消费者
func NewConsumer(addrs []string, groupId string, topics []string) {
	config := sarama.NewConfig()
	config.Version = sarama.V3_3_1_0
	client, err := sarama.NewConsumerGroup(addrs, groupId, config)
	if err != nil {
		fmt.Println("sarama.NewConsumerGroup error")
		panic(err)
	}
	if err := client.Consume(context.Background(), topics, &clientHandler{}); err != nil {
		fmt.Println("client.Consume error")
		panic(err)
	}
}

// clientHandler 消费者处理方法
type clientHandler struct {
}

func (t *clientHandler) Setup(session sarama.ConsumerGroupSession) error {
	fmt.Println("set up")
	return nil
}

func (t *clientHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	fmt.Println("clearn up")
	return nil
}

// ConsumeClaim 消费消息
func (t *clientHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		time.Sleep(5 * time.Second)
		fmt.Println(string(msg.Value))
	}
	return nil
}
