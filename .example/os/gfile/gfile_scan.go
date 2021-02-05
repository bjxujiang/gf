package main

import (
	"github.com/bjxujiang/gf/os/gfile"
	"github.com/bjxujiang/gf/util/gutil"
)

func main() {
	gutil.Dump(gfile.ScanDir("/Users/john/Documents", "*.*"))
	gutil.Dump(gfile.ScanDir("/home/john/temp/newproject", "*", true))
}
