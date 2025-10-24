package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type Handler func(ctx context.Context, log *logrus.Logger, job *Job) error

type Consumer struct {
	client   *Client
	log      *logrus.Logger
	queue    string
	handlers map[string]Handler
}

func NewConsumer(client *Client, queueName string) *Consumer {
	return &Consumer{
		client:   client,
		log:      client.log,
		queue:    queueName,
		handlers: make(map[string]Handler),
	}
}

func (c *Consumer) Register(jobType string, handler Handler) {
	c.handlers[jobType] = handler
}

func (c *Consumer) Start(ctx context.Context, workers int) {
	go c.scheduleDelayedJobs(ctx)

	for i := 0; i < workers; i++ {
		go c.worker(ctx, i)
	}
}

func (c *Consumer) worker(ctx context.Context, id int) {
	c.log.Infof("Worker %d started", id)

	for {
		select {
		case <-ctx.Done():
			c.log.Infof("Worker %d stopped", id)
			return
		default:
			result, err := c.client.rdb.BRPop(ctx, 5*time.Second, c.queue).Result()
			if err == redis.Nil {
				continue
			}

			if err != nil {
				c.log.Infof("Worker %d error: %v", id, err)
				time.Sleep(time.Second)
				continue
			}

			var job Job
			if err := json.Unmarshal([]byte(result[1]), &job); err != nil {
				c.log.Infof("Worker %d unmarshal error: %v", id, err)
				continue
			}

			if err := c.process(ctx, &job); err != nil {
				c.log.Infof("Worker %d job %s failed: %v", id, job.ID, err)
				c.handleFailure(ctx, &job, err)
				continue
			}

			c.log.Infof("Worker %d completed job %s", id, job.ID)
		}
	}
}

func (c *Consumer) process(ctx context.Context, job *Job) error {
	handler, ok := c.handlers[job.Type]
	if !ok {
		c.log.Infof("no handler for job type: %s", job.Type)
		return fmt.Errorf("no handler for job type: %s", job.Type)
	}

	return handler(ctx, c.log, job)
}

func (c *Consumer) handleFailure(ctx context.Context, job *Job, err error) {
	if job.Retry < job.MaxRetry {
		job.Retry++
		delay := time.Duration(job.Retry*job.Retry) * time.Second
		executeAt := time.Now().Add(delay)

		if err := c.client.EnqueueAt(ctx, c.queue, job, executeAt); err != nil {
			c.log.Warnf("Failed to re-enqueue job %s: %v", job.ID, err)
			return
		}

		c.log.Infof("Job %s re-queued (retry %d/%d) after %v", job.ID, job.Retry, job.MaxRetry, delay)
		return
	}

	dlq := c.queue + ":dlq"
	if err := c.client.Enqueue(ctx, dlq, job); err != nil {
		c.log.Infof("Failed to move job %s to DLQ: %v", job.ID, err)
	}
}

func (c *Consumer) scheduleDelayedJobs(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	delayedQueue := c.queue + ":delayed"

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			now := float64(time.Now().Unix())

			result, err := c.client.rdb.ZRangeByScoreWithScores(ctx, delayedQueue, &redis.ZRangeBy{
				Min:   "-inf",
				Max:   fmt.Sprintf("%f", now),
				Count: 100,
			}).Result()

			if err != nil {
				continue
			}

			for _, z := range result {
				jobData := z.Member.(string)
				c.client.rdb.ZRem(ctx, delayedQueue, jobData)
				c.client.rdb.LPush(ctx, c.queue, jobData)
			}
		}
	}
}
