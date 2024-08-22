package messaging

import (
	"context"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
)

type ConsumerHandler func(message *kafka.Message) error

func ConsumeTopic(ctx context.Context, consumer *kafka.Consumer, topic string, log *logrus.Logger, handler ConsumerHandler) {
	err := consumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %v", err)
	}

	run := true

	for run {
		select {
		case <-ctx.Done():
			run = false
		default:
			message, err := consumer.ReadMessage(time.Second)
			if err == nil {
				err := handler(message)
				if err != nil {
					log.Errorf("Failed to process message: %v", err)
				} else {
					_, err = consumer.CommitMessage(message)
					if err != nil {
						log.Fatalf("Failed to commit message: %v", err)
					}
				}
			} else if !err.(kafka.Error).IsTimeout() {
				log.Warnf("Consumer error: %v (%v)\n", err, message)
			}
		}
	}

	log.Infof("Closing consumer for topic : %s", topic)
	err = consumer.Close()
	if err != nil {
		panic(err)
	}
}

func ConsumeTopics(ctx context.Context, consumer *kafka.Consumer, topicHandlers map[string]ConsumerHandler, log *logrus.Logger) {
	// Extract topic names from the map keys
	topics := make([]string, 0, len(topicHandlers))
	for topic := range topicHandlers {
		topics = append(topics, topic)
	}

	// Subscribe to all topics in the map
	err := consumer.SubscribeTopics(topics, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topics: %v", err)
	}

	run := true

	for run {
		select {
		case <-ctx.Done():
			run = false
		default:
			// Poll for a message with a timeout
			message, err := consumer.ReadMessage(time.Second)
			if err == nil {
				// Find the handler for the message's topic
				if handler, exists := topicHandlers[*message.TopicPartition.Topic]; exists {
					// Handle the message using the appropriate handler
					err := handler(message)
					if err != nil {
						log.Errorf("Failed to process message on topic %s: %v", *message.TopicPartition.Topic, err)
					} else {
						// Commit the message after successful processing
						_, err = consumer.CommitMessage(message)
						if err != nil {
							log.Fatalf("Failed to commit message on topic %s: %v", *message.TopicPartition.Topic, err)
						}
					}
				} else {
					log.Warnf("No handler found for topic %s", *message.TopicPartition.Topic)
				}
			} else if !err.(kafka.Error).IsTimeout() {
				log.Warnf("Consumer error: %v (%v)\n", err, message)
			}
		}
	}

	log.Infof("Closing consumer for topics: %v", topics)
	err = consumer.Close()
	if err != nil {
		panic(err)
	}
}
