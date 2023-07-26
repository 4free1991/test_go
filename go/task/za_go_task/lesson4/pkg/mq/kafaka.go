package mq

import (
	"app/lesson4/config"
	"app/lesson4/pkg/logger"
	"app/lesson4/pkg/shutdown"
	"github.com/Shopify/sarama"
	"time"
)

func InitKafaka() {

	// 生产者初始化
	brokers := config.GetConfig().Stats.Brokers
	//topics := config.GetConfig().Stats.Topics

	kafakaConfig := sarama.NewConfig()
	kafakaConfig.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	kafakaConfig.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	kafakaConfig.Producer.Flush.Frequency = 500 * time.Millisecond //
	kafakaConfig.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, kafakaConfig)
	if err != nil {
		panic(err)
	}

	consumerConfig := sarama.NewConfig()
	consumer, err := sarama.NewConsumer(brokers, consumerConfig)

	if err != nil {
		panic(err)
	}

	shutdown.Add(func() {
		producer.Close()
		logger.Logger.Infof("producer shutdown  success")
		consumer.Close()
		logger.Logger.Info("consumer shutdown success")
	})
}
