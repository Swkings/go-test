//go:build ignore

package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func main() {
	go func() {
		log.Println("Listen :6060")
		http.ListenAndServe(":6060", nil)
	}()

	runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪，block
	runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪，mutex
	<-time.After(10 * time.Minute)
}

// go tool pprof http://127.0.0.1:6060/debug/pprof/profile
