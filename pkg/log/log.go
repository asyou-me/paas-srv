package log

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

//异步kafka构建器
func newLogProducer(brokerList []string) *sarama.AsyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms

	producer, err := sarama.NewAsyncProducer(brokerList, config)
	if err != nil {
		panic(err)
	}
	fmt.Println("初始化远程kafka日志系统成功")
	return &producer
}
