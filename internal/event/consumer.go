package event

import (
	"context"
	"encoding/json"
	"log"

	"complaint-service/config"

	"github.com/segmentio/kafka-go"
)

type LoginEvent struct {
	Event     string `json:"event"`
	Username  string `json:"username"`
	Timestamp string `json:"timestamp"`
}

func StartKafkaConsumer(cfg *config.KafkaConfig) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{cfg.Brokers[0]},
		Topic:    cfg.Topic,
		MinBytes: 1,    // 1B
		MaxBytes: 10e6, // 10MB
	})
	defer reader.Close()

	log.Println(" Kafka consumer started...")

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		var event LoginEvent
		if err := json.Unmarshal(m.Value, &event); err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			continue
		}

		log.Printf(" Event received: %+v", event)
	}
}
