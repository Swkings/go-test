package test

import (
	"fmt"
	"testing"
)

type RealTimeInfo struct {
	ExitChan chan string
	Vin      string
	Text     string
}

func (r *RealTimeInfo) ClearRealTimeInfo() {
	writeOne := false
	for {
		select {
		case <-r.ExitChan:
			fmt.Printf("Exit %v \n", r.Vin)
			return
		default:
			if !writeOne {
				fmt.Printf("edit %v text %v \n", r.Vin, r.Text)
				writeOne = true
			}
		}
	}
}

func NewRealTimeInfo(vin string, text string) *RealTimeInfo {
	realTimeInfo := &RealTimeInfo{
		Vin:      vin,
		Text:     text,
		ExitChan: make(chan string, 1),
	}

	go realTimeInfo.ClearRealTimeInfo()

	return realTimeInfo
}

func TestCloseChan(t *testing.T) {

	m := map[string]*RealTimeInfo{
		"v1": NewRealTimeInfo("v1", "v1"),
		"v2": NewRealTimeInfo("v2", "v2"),
		"v3": NewRealTimeInfo("v3", "v3"),
	}

	validVinMap := map[string]string{
		"v1": "",
		"v3": "",
	}
	for vin := range m {
		if _, ok := validVinMap[vin]; !ok {
			m[vin].ExitChan <- "exit"
			close(m[vin].ExitChan)
			delete(m, vin)
		}
	}
	fmt.Printf("m: %v\n", PrettyMapStruct(m, true))
	// 保持主进程不退出
	select {}
}
