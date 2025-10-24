package queue

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

func (c *Client) Enqueue(ctx context.Context, queueName string, job *Job) error {
	data, err := json.Marshal(job)

	if err != nil {
		c.log.Error(err)
		return err
	}

	if err := c.rdb.LPush(ctx, queueName, data).Err(); err != nil {
		c.log.Error(err)
		return err
	}

	return nil
}

func (c *Client) EnqueueAt(ctx context.Context, queueName string, job *Job, executeAt time.Time) error {
	data, err := json.Marshal(job)

	if err != nil {
		c.log.Error(err)
		return err
	}

	score := float64(executeAt.Unix())
	delayedQueue := queueName + ":delayed"

	if err := c.rdb.ZAdd(ctx, delayedQueue, redis.Z{Score: score, Member: data}).Err(); err != nil {
		c.log.Error(err)
		return err
	}

	return err
}
