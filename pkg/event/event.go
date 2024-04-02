package event

import (
	"context"
	"github.com/2pgcn/gameim/api/protocol"
	"github.com/segmentio/kafka-go"
)

type EventHeader map[string]string

const eventId = "ID"

type Event interface {
	Header() *EventHeader
	SetId(string)
	GetId() string
	Value() []byte
	String() string
	GetQueueMsg() *QueueMsg
	IsClose() bool
	ToProtocol() (*protocol.Proto, error)
}

func NewHeader(size int) EventHeader {
	return make(EventHeader, size)
}

// todo 改成proto.Marshal
func (eh EventHeader) GetKafkaHead() (res []kafka.Header) {
	for k, v := range eh {
		res = append(res, kafka.Header{
			Key:   k,
			Value: ([]byte)(v),
		})
	}
	return res
}

type Handler func(context.Context, Event) error

type Sender interface {
	Send(ctx context.Context, msg Event) error
	Close() error
}

type Receiver interface {
	Receive(ctx context.Context) (e []chan Event, err error)
	Commit(ctx context.Context, event Event) error
	Close() error
}
