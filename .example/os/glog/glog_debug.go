package main

import (
	"time"

	"github.com/bjxujiang/gf/os/glog"
	"github.com/bjxujiang/gf/os/gtime"
	"github.com/bjxujiang/gf/os/gtimer"
)

func main() {
	gtimer.SetTimeout(3*time.Second, func() {
		glog.SetDebug(false)
	})
	for {
		glog.Debug(gtime.Datetime())
		time.Sleep(time.Second)
	}
}
