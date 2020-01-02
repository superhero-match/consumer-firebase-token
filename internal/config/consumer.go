package config

// Consumer holds the configuration values for the Kafka consumer.
type Consumer struct {
	Brokers []string `env:"KAFKA_BROKERS" default:"[192.168.178.17:9092]"`
	Topic   string   `env:"KAFKA_UPDATE_FIREBASE_TOKEN_TOPIC" default:"update.firebase.token"`
	GroupID string   `env:"KAFKA_CONSUMER_REGISTER_GROUP_ID" default:"consumer-firebase-token"`
}
