package main

import (
	"context"
	"os/signal"
	"sync"
	"syscall"

	"edot-monorepo/services/warehouse-service/internal/config"
	"edot-monorepo/services/warehouse-service/internal/delivery/messaging"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func main() {
	// Initialize configurations, logger, and other dependencies
	viperConfig := config.NewViper()
	logger := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, logger)
	validate := config.NewValidator(viperConfig)

	// Start the service
	logger.Info("Starting worker service")

	// Set up context with cancel for graceful shutdown using signal.NotifyContext
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	defer stop() // Ensure stop is called on exit

	// Use a WaitGroup to wait for the consumer to finish
	var wg sync.WaitGroup
	wg.Add(1)

	// Run the product consumer in a separate goroutine
	go func() {
		defer wg.Done()
		RunWarehouseConsumer(ctx, logger, db, validate, viperConfig)
	}()

	// Wait for context cancellation (signal received)
	<-ctx.Done()
	logger.Info("Received shutdown signal, waiting for goroutines to finish")

	// Wait for the consumer to finish processing
	wg.Wait()
	logger.Info("Worker service has shut down gracefully")
}

func RunWarehouseConsumer(ctx context.Context, logger *logrus.Logger, db *gorm.DB, validate *validator.Validate, viperConfig *viper.Viper) {
	logger.Info("setup warehouse consumer")
	consumer := config.NewKafkaConsumer(viperConfig, logger)
	handler := messaging.NewWarehouseConsumer(logger, db, validate)

	topicHandlers := map[string]messaging.ConsumerHandler{
		"shop_assign_warehouse": func(msg *kafka.Message) error {
			return handler.ConsumeShopWarehouseAssigned(msg)
		},
		"shop_created": func(msg *kafka.Message) error {
			return handler.ConsumeShopCreated(msg)
		},
		"stock_changed": func(msg *kafka.Message) error {
			return handler.ConsumeStockChanged(msg)
		},
	}

	messaging.ConsumeTopics(ctx, consumer, topicHandlers, logger)
}
