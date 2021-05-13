package cache

import (
	"time"

	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v7"
	"github.com/robinjoseph08/redisqueue/v2"
)

type Redis struct {
	client          *redis.Client
	ConnectOption   *redis.Options
	ConsumerOptions *redisqueue.ConsumerOptions
	ProducerOptions *redisqueue.ProducerOptions
	consumer        *redisqueue.Consumer
	producer        *redisqueue.Producer
	mutex           *redislock.Client
}

func (r *Redis) Connect() error {
	var err error
	r.client = redis.NewClient(r.ConnectOption)
	_, err = r.client.Ping().Result()
	if err != nil {
		return err
	}
	r.mutex = redislock.New(r.client)
	r.producer, err = r.newProducer(r.client)
	if err != nil {
		return err
	}
	r.consumer, err = r.newConsumer(r.client)
	return err
}

func (r *Redis) newConsumer(client *redis.Client) (*redisqueue.Consumer, error) {
	if r.ConsumerOptions == nil {
		r.ConsumerOptions = &redisqueue.ConsumerOptions{}
	}
	r.ConsumerOptions.RedisClient = client
	return redisqueue.NewConsumerWithOptions(r.ConsumerOptions)
}

func (r *Redis) newProducer(client *redis.Client) (*redisqueue.Producer, error) {
	if r.ProducerOptions == nil {
		r.ProducerOptions = &redisqueue.ProducerOptions{}
	}
	r.ProducerOptions.RedisClient = client
	return redisqueue.NewProducerWithOptions(r.ProducerOptions)
}

func (r *Redis) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}

func (r *Redis) Set(key string, value interface{}, expire int) error {
	return r.client.Set(key, value, time.Duration(expire)*time.Second).Err()
}

func (r *Redis) Del(key string) error {
	return r.client.Del(key).Err()
}

func (r *Redis) HashGet(hashKey, key string) (string, error) {
	return r.client.HGet(hashKey, key).Result()
}

func (r *Redis) HashSet(hashKey string, value ...interface{}) error {
	return r.client.HSet(hashKey, value).Err()
}

func (r *Redis) HashSetNX(hashKey, field string, value interface{}) error {
	return r.client.HSetNX(hashKey, field, value).Err()
}

func (r *Redis) HashMSet(hashKey string, value ...interface{}) error {
	return r.client.HMSet(hashKey, value).Err()
}

func (r *Redis) HashDel(hashKey, key string) error {
	return r.client.HDel(hashKey, key).Err()
}

func (r *Redis) Increase(key string) error {
	return r.client.Incr(key).Err()
}

func (r *Redis) Decrease(key string) error {
	return r.client.Decr(key).Err()
}

func (r *Redis) Expire(key string, duration time.Duration) error {
	return r.client.Expire(key, duration).Err()
}

func (r *Redis) Append(name string, message Message) error {
	err := r.producer.Enqueue(&redisqueue.Message{
		ID:     message.GetID(),
		Stream: name,
		Values: message.GetValues(),
	})
	return err
}

type RedisMessage struct {
	redisqueue.Message
}

func (r *Redis) Register(name string, f ConsumerFunc) {
	r.consumer.Register(name, func(message *redisqueue.Message) error {
		m := new(RedisMessage)
		m.SetID(message.ID)
		m.SetValues(message.Values)
		m.SetStream(message.Stream)
		return f(m)
	})
}

func (r *Redis) Lock(key string, ttl int64, option *redislock.Options) (*redislock.Lock, error) {
	if r.mutex == nil {
		r.mutex = redislock.New(r.client)
	}
	return r.mutex.Obtain(key, time.Duration(ttl)*time.Second, option)
}

func (r *Redis) GetClient() *redis.Client {
	return r.client
}

func (r *Redis) Run() {
	r.consumer.Run()
}

func (r *Redis) ShutDown() {
	r.consumer.Shutdown()
}

func (m *RedisMessage) GetID() string {
	return m.ID
}

func (m *RedisMessage) SetID(id string) {
	m.ID = id
}

func (m *RedisMessage) GetStream() string {
	return m.Stream
}

func (m *RedisMessage) SetStream(stream string) {
	m.Stream = stream
}

func (m *RedisMessage) GetValues() map[string]interface{} {
	return m.Values
}

func (m *RedisMessage) SetValues(value map[string]interface{}) {
	m.Values = value
}
