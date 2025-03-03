package consumer

import (
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/kinesis/kinesisiface"
)

// Option is used to override defaults when creating a new Consumer
type Option func(*Consumer)

// WithCheckpoint overrides the default checkpoint
func WithCheckpoint(checkpoint Checkpoint) Option {
	return func(c *Consumer) {
		c.checkpoint = checkpoint
	}
}

// WithLogger overrides the default logger
func WithLogger(logger Logger) Option {
	return func(c *Consumer) {
		c.logger = logger
	}
}

// WithCounter overrides the default counter
func WithCounter(counter Counter) Option {
	return func(c *Consumer) {
		c.counter = counter
	}
}

// WithClient overrides the default client
func WithClient(client kinesisiface.KinesisAPI) Option {
	return func(c *Consumer) {
		c.client = client
	}
}

// WithShardIteratorType overrides the starting point for the consumer
func WithShardIteratorType(t string) Option {
	return func(c *Consumer) {
		c.initialShardIteratorType = t
	}
}

// WithGroup overrides the group for the consumer
func WithGroup(listShards func(ksis kinesisiface.KinesisAPI, streamName string) ([]*kinesis.Shard, error)) Option {
	return func(c *Consumer) {
		c.group = NewAllGroup(c.client, c.checkpoint, c.streamName, c.logger, listShards)
	}
}
