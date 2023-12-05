package test

import (
	"fmt"
	"testing"
	"time"
)

type SChan struct {
	Name      string
	Age       int
	ChanGroup ChanGroup
}

type ChanGroup struct {
	C chan string
}

func NewChanGroup() ChanGroup {
	return ChanGroup{
		C: make(chan string, 1),
	}
}

func TestChannel07(t *testing.T) {
	sc := []*SChan{
		{
			Name:      "1",
			Age:       1,
			ChanGroup: NewChanGroup(),
		},
		{
			Name:      "2",
			Age:       2,
			ChanGroup: NewChanGroup(),
		},
		{
			Name:      "3",
			Age:       3,
			ChanGroup: NewChanGroup(),
		},
	}
	for _, ch := range sc {
		if ch.Name == "2" {
			go recv(ch)
			// go func() {
			// 	ev := <-ch.ChanGroup.C
			// 	fmt.Printf("event: %v, s: %+v\n", ev, ch.Name)
			// }()
		}
		time.Sleep(1 * time.Millisecond)
	}
	outTIme := time.Now().UnixNano()
	fmt.Printf("out time: %v\n", outTIme)
	sc[0].ChanGroup.C <- "cancel1"
	sc[1].ChanGroup.C <- "cancel2"
	sc[2].ChanGroup.C <- "cancel3"
	afterTime := time.Now().UnixNano() - outTIme
	fmt.Printf("after time: %v\n", afterTime)
	time.Sleep(3 * time.Second)
	endTime := time.Now().UnixNano() - afterTime
	fmt.Printf("end time: %v\n", endTime)
}

func recv(sc *SChan) {
	ev := <-sc.ChanGroup.C
	fmt.Printf("event: %v, s: %+v\n", ev, sc.Name)
}
