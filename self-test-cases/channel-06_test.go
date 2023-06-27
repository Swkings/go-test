package test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/zeromicro/go-zero/core/threading"
)

type Data struct {
	Flag string
}

func NewData(flag string) *Data {
	return &Data{
		Flag: flag,
	}
}

type Payload struct {
	Name string
	Age  int32
	Data *Data
}

func NewPayload() Payload {
	return Payload{
		Name: "init",
		Age:  100,
	}
}

func (p Payload) SetName(name string) Payload {
	p.Name = name
	return p
}

func (p Payload) SetAge(age int32) Payload {
	p.Age = age
	return p
}

func (p Payload) SetData(data *Data) Payload {
	p.Data = data
	time.Sleep(1 * time.Second)
	return p
}

var Ch chan Payload = make(chan Payload, 999)

func recvChan(ctx context.Context) {
	for {
		select {
		case p := <-Ch:
			fmt.Printf("payload: %+v\n", PrettyMapStruct(p, true))
			panic("panic test")
		case <-ctx.Done():
			fmt.Printf("exit\n")
			return
		}
	}
}

func TestChannelLink(t *testing.T) {
	l := sync.Mutex{}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	threading.GoSafe(func() {
		l.Lock()
		defer l.Unlock()
		recvChan(ctx)
	})

	go func() {
		time.Sleep(time.Second)
		l.Lock()

		fmt.Println("====ok=====")

		l.Unlock()
	}()

	<-ctx.Done()
	// 	Ch <- NewPayload().SetName("test")
	// 	Ch <- NewPayload().SetAge(10)
	// 	Ch <- NewPayload().SetData(NewData("data"))
	// 	time.Sleep(5 * time.Second)
	// cancel()
	// time.Sleep(1 * time.Second)
}
