package main

import (
	"time"

	"github.com/bjxujiang/gf/os/glog"
)

func main() {
	for i := 0; i < 10; i++ {
		glog.Async().Print("async log", i)
	}
	time.Sleep(time.Second)
}
