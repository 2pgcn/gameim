package main

import (
	"context"
	"errors"
	"github.com/2pgcn/gameim/api/protocol"
	"github.com/2pgcn/gameim/conf"
	"github.com/2pgcn/gameim/pkg/event"
	"github.com/2pgcn/gameim/pkg/gamelog"
	"strconv"
	"sync/atomic"
	"time"
)

func benchSock(ctx context.Context) {
	consumer, err := event.NewSockReceiver(&conf.Sock{
		Address: c.GetSock().Address,
	})
	if err != nil {
		panic(err)
	}
	e, err := consumer.Receive(ctx)
	if err != nil {
		panic(err)
	}
	producer, err := event.NewSockRender(c.GetSock().Address)
	if err != nil {
		panic(err)
	}
	for _, v1 := range e {
		v := v1
		gopool.GoCtx(func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					return
				case _ = <-v:
					atomic.AddInt64(&countDown, 1)
					//event.PutQueueMsg(msg.GetQueueMsg())
					//err = consumer.Commit(ctx, msg)
					//if err != nil {
					//	fmt.Println(err)
					//}
				}
			}
		})
	}
	var msgId = 0
	gopool.GoCtx(func(ctx context.Context) {
		for {
			for i := 0; i <= 100000; i++ {
				queueMsg := &event.QueueMsg{
					H:    make(map[string]string, 8),
					Data: &protocol.Msg{},
				}
				queueMsg.Data.Msg = []byte("hello world")
				queueMsg.Data.Type = protocol.Type_APP
				queueMsg.Data.ToId = "0"
				queueMsg.Data.SendId = "0"
				msgId++
				queueMsg.SetId(strconv.Itoa(msgId))
				select {
				case <-ctx.Done():
					gamelog.GetGlobalog().Debug("exit producer start stop")
					producer.Close()
					return
				default:
					err = producer.Send(ctx, queueMsg)
					if err != nil {
						gamelog.GetGlobalog().Error(err)
						if errors.Is(err, event.FullError) {
							time.Sleep(time.Second * 1)
							continue
						}
						gamelog.GetGlobalog().Error(err)
						continue
					}
					atomic.AddInt64(&countSend, 1)
				}
			}
		}
	})
	gopool.GoCtx(func(ctx context.Context) {
		select {
		case <-ctx.Done():
			consumer.Close()
		}
	})
	gopool.GoCtx(func(ctx context.Context) {
		result(ctx)
	})
}
