package main

import (
	"github.com/bjxujiang/gf/os/gfsnotify"
	"github.com/bjxujiang/gf/os/glog"
)

func main() {
	//path := `D:\temp`
	path := "/Users/john/Temp"
	_, err := gfsnotify.Add(path, func(event *gfsnotify.Event) {
		glog.Println(event)
	})
	if err != nil {
		glog.Fatal(err)
	} else {
		select {}
	}
}
