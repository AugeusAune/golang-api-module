package main

import (
	"context"
	"golang-api-module/config"
	"golang-api-module/internal/jobs"
	"golang-api-module/internal/queue"
	"golang-api-module/internal/shared/logger"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	log := logger.NewLogger()

	if err := godotenv.Load(); err != nil {
		log.Warn("no .env found")
	}

	config := config.Load()

	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: "",
		DB:       0,
	})

	client := queue.NewClient(rdb, log)
	defer client.Close()

	consumer := queue.NewConsumer(client, "default")

	consumer.Register("test_job", jobs.TestJob)

	consumer.Start(ctx, 3)

	<-sigChan

	log.Info("shutting down worker")

	cancel()
}
