package cache

import "time"

type Adapter interface {
	Connect() error
	Get(key string) (string, error)
	Set(key string, value interface{}, expire int) error
	Del(key string) error
	HashGet(hashKey, key string) error
	HashSet(hashKey string, value ...interface{}) error
	HashSetNX(hashKey, field string, value interface{}) error
	HashMSet(hashKey string, value ...interface{}) error
	HashDel(hashKey, key string) error
	Increase(key string) error
	Decrease(key string) error
	Expire(key string, duration time.Duration) error

	AdapterQueue
}

type AdapterQueue interface {
	Append(name string, message Message) error
	Register(name string, f ConsumerFunc)
	Run()
	ShutDown()
}

type Message interface {
	SetID(string)
	SetStream(string)
	SetValues(map[string]interface{})
	GetID() string
	GetStream() string
	GetValues() map[string]interface{}
}

type ConsumerFunc func(message Message) error
