package redis

import (
	"context"
	"time"

	"github.com/diazjohan98/go-virtual-queue-system/internal/domain"
	"github.com/redis/go-redis/v9"
)

type redisQueueRepository struct {
	client *redis.Client
}

func NewRedisQueueRepository(client *redis.Client) domain.QueueRepository {
	return &redisQueueRepository{
		client: client,
	}
}

func (r *redisQueueRepository) Enqueue(ctx context.Context, eventID string, userID string) error {
	key := "queue:" + eventID

	score := float64(time.Now().UnixNano())

	err := r.client.ZAdd(ctx, key, redis.Z{
		Score:  score,
		Member: userID,
	}).Err()

	return err
}

func (r *redisQueueRepository) GetPosition(ctx context.Context, eventID string, userID string) (int, error) {
	key := "queue:" + eventID

	rank, err := r.client.ZRank(ctx, key, userID).Result()
	if err != nil {
		if err == redis.Nil {
			return -1, nil // User not found in the queue
		}
		return -1, err
	}

	return int(rank) + 1, nil
}

func (r *redisQueueRepository) Dequeue(ctx context.Context, eventID string) (string, error) {
	key := "queue:" + eventID

	results, err := r.client.ZPopMin(ctx, key, 1).Result()
	if err != nil {
		return "", err
	}

	if len(results) == 0 {
		return "", nil
	}

	return results[0].Member.(string), nil
}
