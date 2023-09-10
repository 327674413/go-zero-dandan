package queue

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"go-zero-dandan/common/resd"
	"time"
)

// Producer 生产者
type Producer struct {
	Client sarama.SyncProducer
}

// NewProducer 创建生产者
func NewProducer(addrs []string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V3_3_1_0                          //显示指定协议版本，默认版本好像是1.0.0e很老，用sarama.DefaultVersion查看
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 随机分区partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	config.Producer.Return.Errors = true                      // 失败的消息将在erros channel返回
	client, err := sarama.NewSyncProducer(addrs, config)
	if err != nil {
		return nil, resd.Error(err)
	}
	return &Producer{Client: client}, nil
	//defer client.Close()
}

// Send 发送生产消息
func (t *Producer) Send(topic string, msg string) (partition int32, offset int64, err error) {
	// 发送消息
	return t.Client.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	})
}

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
