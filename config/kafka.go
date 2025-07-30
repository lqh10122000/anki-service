package config

type KafkaConfig struct {
	Brokers []string
	Topic   string
}

func LoadKafkaConfig() *KafkaConfig {
	return &KafkaConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "user-logs",
	}
}
