package cache

import (
	"testing"
	"time"

	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v7"
	"github.com/robinjoseph08/redisqueue/v2"
)

func TestRedis_Append(t *testing.T) {
	type fields struct {
		ConnectOption   *redis.Options
		ConsumerOptions *redisqueue.ConsumerOptions
		ProducerOptions *redisqueue.ProducerOptions
		client          *redis.Client
		consumer        *redisqueue.Consumer
		producer        *redisqueue.Producer
		mutex           *redislock.Client
	}
	type args struct {
		name    string
		message Message
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"test01",
			fields{
				ConnectOption: &redis.Options{
					Addr: "127.0.0.1:6379",
				},
				ConsumerOptions: &redisqueue.ConsumerOptions{
					VisibilityTimeout: 60 * time.Second,
					BlockingTimeout:   5 * time.Second,
					ReclaimInterval:   1 * time.Second,
					BufferSize:        100,
					Concurrency:       10,
				},
				ProducerOptions: &redisqueue.ProducerOptions{
					StreamMaxLength:      100,
					ApproximateMaxLength: true,
				},
			},
			args{
				name: "test",
				message: &RedisMessage{redisqueue.Message{
					Stream: "test",
					Values: map[string]interface{}{
						"key": "value",
					},
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				ConnectOption: tt.fields.ConnectOption,
			}
			if err := r.Connect(); err != nil {
				t.Errorf("Connect() error = %v", err)
			}
			if err := r.Append(tt.args.name, tt.args.message); err != nil {
				t.Errorf("Append() error = %v", err)
			}
		})
	}
}
