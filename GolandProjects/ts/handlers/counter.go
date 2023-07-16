package handlers

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Counter struct {
	value int
	redis *redis.Client
}

func NewCounter(redisClient *redis.Client) *Counter {
	return &Counter{
		value: 0,
		redis: redisClient,
	}
}

func (c *Counter) Add(ctx context.Context, increment int) error {
	c.value += increment

	// Сохранение значения в Redis
	err := c.redis.Set(ctx, "counter:value", c.value, 0).Err()
	if err != nil {
		// Обработка ошибки сохранения в Redis
		return fmt.Errorf("failed to save counter value to Redis: %v", err)
	}

	return nil
}

func (c *Counter) Sub(ctx context.Context, decrement int) error {
	c.value -= decrement

	// Сохранение значения в Redis
	err := c.redis.Set(ctx, "counter:value", c.value, 0).Err()
	if err != nil {
		// Обработка ошибки сохранения в Redis
		return fmt.Errorf("failed to save counter value to Redis: %v", err)
	}

	return nil
}

func (c *Counter) Value() int {
	return c.value
}
