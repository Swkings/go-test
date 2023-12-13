package logCompressor

import (
	"context"
	"fmt"
	"testing"
	"time"
)

type Print struct {
	tag string
}

func (p Print) Debugf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	fmt.Println()
}
func (p Print) Infof(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	fmt.Println()
}

func TestLogCompressor(t *testing.T) {
	logger1 := Print{tag: "logger1"}
	logger2 := Print{tag: "logger2"}
	logger3 := Print{tag: "logger3"}
	ctx, cancel := context.WithCancel(context.Background())
	lc := NewCompressorGroup(ctx, 100)
	var iter int = 200
	for i := 1; i <= iter; i++ {
		if i <= 80 {
			lc.Loader().AddMessage(logger1.Debugf, "test log 1 - Debug")
			lc.AddMessage(logger1.Infof, "test log 1 - Info")
			lc.Loader("Common").AddMessage(logger1.Infof, "test log common - Info")
		} else {
			lc.Loader().AddMessage(logger1.Debugf, "test log 1 - Debug - Changed")
			lc.AddMessage(logger1.Infof, "test log 1 - Info - Changed")
			lc.Loader("Common").AddMessage(logger3.Debugf, "test log common - Debug")
		}
		lc.AddMessage(logger2.Debugf, "test log 2 - %v", "Debug")
		lc.AddMessage(logger2.Infof, "test log 2 - Info")
		lc.AddMessage(logger3.Debugf, "test log 3 - Debug")
		lc.AddMessage(logger3.Infof, "test log 3 - Info")
		if i == 190 {
			break
		}
	}
	cancel()
	time.Sleep(1 * time.Second)
}
