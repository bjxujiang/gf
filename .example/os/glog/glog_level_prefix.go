package main

import (
	"github.com/bjxujiang/gf/os/glog"
)

func main() {
	l := glog.New()
	l.SetLevelPrefix(glog.LEVEL_DEBU, "debug")
	l.Debug("test")
}
