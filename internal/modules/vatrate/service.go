package vatrate

import (
	"context"
	"golang-api-module/internal/queue"
	"time"

	"github.com/sirupsen/logrus"
)

type Service struct {
	ctx         context.Context
	queueClient *queue.Client
	log         *logrus.Logger
}

func NewService(ctx context.Context, queueClient *queue.Client, log *logrus.Logger) *Service {
	return &Service{
		ctx:         ctx,
		queueClient: queueClient,
		log:         log,
	}
}

func (s *Service) AddQueue() {
	s.queueClient.Enqueue(s.ctx, "default", &queue.Job{
		ID:   "test",
		Type: "test_job",
		Payload: map[string]interface{}{
			"test": "Coba",
		},
		CreatedAt: time.Now(),
		MaxRetry:  3,
	})
}
