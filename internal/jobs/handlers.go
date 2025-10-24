package jobs

import (
	"context"
	"golang-api-module/internal/queue"

	"github.com/sirupsen/logrus"
)

func TestJob(ctx context.Context, log *logrus.Logger, job *queue.Job) error {
	log.Infof("job id %s with payload %v", job.ID, job.Payload)
	return nil
}
