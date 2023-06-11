package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

func main() {
	producer() //

}
func client() {
	config := sarama.NewConfig()
	config.Version = sarama.V3_3_1_0
	client, err := sarama.NewConsumerGroup([]string{"127.0.0.1:9092"}, "normal_log_group", config)
	if err != nil {
		fmt.Println("sarama.NewConsumerGroup error")
		panic(err)
	}
	if err := client.Consume(context.Background(), []string{"normal_log"}, &handler{}); err != nil {
		fmt.Println("client.Consume error")
		panic(err)
	}
}

type handler struct {
}

func (t *handler) Setup(session sarama.ConsumerGroupSession) error {
	fmt.Println("set up")
	return nil
}

func (t *handler) Cleanup(session sarama.ConsumerGroupSession) error {
	fmt.Println("clearn up")
	return nil
}

func (t *handler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		time.Sleep(5 * time.Second)
		fmt.Println(string(msg.Value))
	}
	return nil
}

func producer() {
	config := sarama.NewConfig()
	config.Version = sarama.V3_3_1_0                          //显示指定协议版本，默认版本好像是1.0.0e很老，用sarama.DefaultVersion查看
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 随机分区partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	config.Producer.Return.Errors = true                      // 失败的消息将在erros channel返回
	// 构造一个消息
	msg := map[string]string{
		"test": "测试字段",
		"name": "张三",
	}
	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("json Marshal error")
	}
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()
	// 发送消息
	pid, offset, err := client.SendMessage(&sarama.ProducerMessage{
		Topic: "normal_log",
		Value: sarama.StringEncoder(data),
	})
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
