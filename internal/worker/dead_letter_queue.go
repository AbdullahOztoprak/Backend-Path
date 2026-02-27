package worker

import (
    "context"
    "log"
    "time"

    "github.com/go-redis/redis/v8"
)

type DeadLetterQueue struct {
    client *redis.Client
    ctx    context.Context
}

func NewDeadLetterQueue(redisClient *redis.Client) *DeadLetterQueue {
    return &DeadLetterQueue{
        client: redisClient,
        ctx:    context.Background(),
    }
}

func (dlq *DeadLetterQueue) Push(message string) error {
    err := dlq.client.LPush(dlq.ctx, "dead_letter_queue", message).Err()
    if err != nil {
        log.Printf("Failed to push message to dead letter queue: %v", err)
        return err
    }
    return nil
}

func (dlq *DeadLetterQueue) Process() {
    for {
        message, err := dlq.client.BRPop(dlq.ctx, 5*time.Second, "dead_letter_queue").Result()
        if err != nil {
            log.Printf("Error popping message from dead letter queue: %v", err)
            continue
        }

        // Process the message
        log.Printf("Processing dead letter message: %s", message[1])
        // Add your message processing logic here
    }
}