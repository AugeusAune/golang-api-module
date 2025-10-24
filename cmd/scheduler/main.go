package main

import (
	"context"
	"fmt"
	"golang-api-module/internal/scheduler"
	"golang-api-module/internal/shared/logger"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	log := logger.NewLogger()

	s := scheduler.NewScheduler(log)

	initScheduler(s, log)

	s.Start()

	fmt.Println("Scheduler running with registered jobs")
	fmt.Println("Cron Schedule Format (with seconds):")
	fmt.Println("┌──────────── second (0 - 59)")
	fmt.Println("│ ┌────────── minute (0 - 59)")
	fmt.Println("│ │ ┌──────── hour (0 - 23)")
	fmt.Println("│ │ │ ┌────── day of month (1 - 31)")
	fmt.Println("│ │ │ │ ┌──── month (1 - 12)")
	fmt.Println("│ │ │ │ │ ┌── day of week (0 - 6, Sunday=0)")
	fmt.Println("│ │ │ │ │ │")
	fmt.Println("* * * * * *")

	<-sigChan

	log.Info("shutting down App Scheduler ...")

	s.Stop()
	cancel()
}

func initScheduler(s *scheduler.Scheduler, log *logrus.Logger) {
	s.AddJob("*/5 * * * * *", func() {
		log.Info("Scheduler per 5 seconds")
	})

}
