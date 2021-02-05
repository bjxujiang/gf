package main

import (
	"fmt"

	"github.com/bjxujiang/gf/frame/g"
	"github.com/bjxujiang/gf/os/gres"
	_ "github.com/bjxujiang/gf/os/gres/testdata"
)

func main() {
	gres.Dump()

	v := g.View()
	s, err := v.Parse("index.html")
	fmt.Println(err)
	fmt.Println(s)
}
