package services

import (
	"context"
	"encoding/json"

	"github.com/kevalsabhani/toll-calculator/client/services"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

const (
	KafkaHost = "localhost:9092"
	GroupId   = "distance-calculator"
	Topic     = "obu-data"
)

type KafkaConsumer struct {
	Reader *kafka.Reader
	Logger *zap.Logger
}

func NewKafkaConsumer(logger *zap.Logger) *KafkaConsumer {
	config := kafka.ReaderConfig{
		Brokers:  []string{KafkaHost},
		GroupID:  GroupId,
		Topic:    Topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	}
	reader := kafka.NewReader(config)

	return &KafkaConsumer{
		Reader: reader,
		Logger: logger,
	}
}

func (kc *KafkaConsumer) Read(ctx context.Context) *services.OBUData {
	message, err := kc.Reader.ReadMessage(ctx)
	if err != nil {
		kc.Logger.Error(err.Error())
	}
	var data services.OBUData
	if err := json.Unmarshal(message.Value, &data); err != nil {
		kc.Logger.Error(err.Error())
	}

	return &data
}
