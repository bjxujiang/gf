package main

import (
	"fmt"
	"github.com/bjxujiang/gf/frame/g"
	"github.com/bjxujiang/gf/util/gvalid"
)

func main() {
	g.I18n().SetLanguage("cn")
	err := gvalid.Check("", "required", nil)
	fmt.Println(err.String())
}
