package event

import (
	"complaint-service/config"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

func InitKafkaLogger(cfg *config.KafkaConfig) {
	writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{cfg.Brokers[0]},
		Topic:    cfg.Topic,
		Balancer: &kafka.LeastBytes{},
	})
}

func LogLoginEvent(username string) {
	if writer == nil {
		log.Println("⚠️ Kafka writer is not initialized")
		return
	}

	event := map[string]interface{}{
		"event":     "LOGIN",
		"username":  username,
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	}

	data, err := json.Marshal(event)
	if err != nil {
		log.Printf("Failed to marshal login event: %v", err)
		return
	}

	err = writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(username),
			Value: data,
		})
	if err != nil {
		log.Printf("Failed to write Kafka message: %v", err)
	}
}
