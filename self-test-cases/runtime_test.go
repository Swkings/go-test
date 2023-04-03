package test

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/zeromicro/go-zero/core/threading"
)

func GetFmtFuncName(args ...string) string {
	pc, _, _, _ := runtime.Caller(1)
	function := runtime.FuncForPC(pc)
	nameArr := strings.Split(function.Name(), ".")
	funcName := nameArr[len(nameArr)-1]
	for _, arg := range args {
		if arg == "" {
			continue
		}
		funcName += "-" + arg
	}

	return "[" + funcName + "]"
}

func GetFmtFuncNameInAnonymous(args ...string) string {
	pc, _, _, _ := runtime.Caller(1)
	function := runtime.FuncForPC(pc)
	nameArr := strings.Split(function.Name(), ".")
	funcName := nameArr[len(nameArr)-2]
	for _, arg := range args {
		if arg == "" {
			continue
		}
		funcName += "-" + arg
	}

	return "[" + funcName + "]"
}

func TestFuncName(t *testing.T) {
	signal := make(chan interface{})
	threading.GoSafe(func() {
		fmt.Printf("%v\n", GetFmtFuncNameInAnonymous())
		signal <- "exit"
	})
	<-signal
}
