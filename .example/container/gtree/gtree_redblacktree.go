package main

import (
	"github.com/bjxujiang/gf/container/gtree"
)

func main() {
	tree := gtree.NewRedBlackTree(func(v1, v2 interface{}) int {
		return v1.(int) - v2.(int)
	})
	for i := 0; i < 10; i++ {
		tree.Set(i, i)
	}
	tree.Print()
	tree.Flip()
	tree.Print()
}
