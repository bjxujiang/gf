package main

import (
	"github.com/bjxujiang/gf/frame/g"
	"github.com/bjxujiang/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writef("name: %v, pass: %v", r.Get("name"), r.Get("pass"))
	})
	s.SetPort(8199)
	s.Run()
}
