package queued

import (
	"context"
	"github.com/Shopify/sarama"
	"go-zero-dandan/common/resd"
)

type Producer struct {
	Pusher *sarama.SyncProducer
}

func NewProducer(addrs []string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V3_3_1_0                          //显示指定协议版本，默认版本好像是1.0.0e很老，用sarama.DefaultVersion查看
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 随机分区partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	config.Producer.Return.Errors = true                      // 失败的消息将在erros channel返回
	// 连接kafka
	client, err := sarama.NewSyncProducer(addrs, config)
	if err != nil {
		return nil, err
	}
	//defer client.Close() //目前不知道应该在哪里主动close
	return &Producer{
		Pusher: &client,
	}, nil
}
func (t *Producer) PushCtx(ctx context.Context, topic string, str string) (partition int32, offset int64, err error) {
	// 发送消息
	if ctx == nil {
		ctx = context.Background()
	}
	partition, offset, err = (*t.Pusher).SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(str),
	})
	if err != nil {
		return 0, 0, resd.ErrorCtx(ctx, err)
	}
	return partition, offset, err
}
