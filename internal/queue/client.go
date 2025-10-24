package queue

import (
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type Client struct {
	rdb *redis.Client
	log *logrus.Logger
}

type Job struct {
	ID        string                 `json:"id"`
	Type      string                 `json:"type"`
	Payload   map[string]interface{} `json:"payload"`
	CreatedAt time.Time              `json:"created_at"`
	Retry     int                    `json:"retry"`
	MaxRetry  int                    `json:"max_retry"`
}

func NewClient(rdb *redis.Client, log *logrus.Logger) *Client {
	return &Client{
		rdb: rdb,
		log: log,
	}
}

func (c *Client) Close() error {
	return c.rdb.Close()
}
