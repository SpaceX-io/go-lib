package cache

import (
	"sync"
	"testing"

	"github.com/robinjoseph08/redisqueue/v2"
)

func TestMemory_Append(t *testing.T) {
	type fields struct {
		items *sync.Map
		queue *sync.Map
		wait  sync.WaitGroup
		mutex sync.RWMutex
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
			fields{},
			args{
				name: "test",
				message: &MemoryMessage{redisqueue.Message{
					ID:     "",
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
			m := &Memory{
				item:  tt.fields.items,
				queue: tt.fields.queue,
				wait:  tt.fields.wait,
				mutex: tt.fields.mutex,
			}
			if err := m.Connect(); err != nil {
				t.Errorf("Connect() error = %v", err)
			}
			if err := m.Append(tt.args.name, tt.args.message); err != nil {
				t.Errorf("Append error = %v", err)
			}
		})
	}
}
