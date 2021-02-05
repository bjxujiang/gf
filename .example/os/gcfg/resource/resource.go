package main

import (
	"github.com/bjxujiang/gf/frame/g"
	_ "github.com/bjxujiang/gf/os/gres/testdata"
)

func main() {
	g.Res().Dump()
	g.Dump(g.Config().Get("redis"))

	g.Config().SetFileName("my.ini")
	g.Dump(g.Config().Get("redis"))

	g.Config().SetPath("config-custom")
	g.Config().SetFileName("my.ini")
	g.Dump(g.Config().Get("redis"))

	g.Config().SetFileName("app.conf")
	g.Dump(g.Config().Get("redis"))
}
