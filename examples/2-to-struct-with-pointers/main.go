package main

import (
	"fmt"

	"di/di"
)

type A struct {
	Value string
}

type B struct {
	Value *A
}

func main() {
	src := B{Value: &A{Value: "foo"}}
	_ = di.Set(&src)

	var target B
	_ = di.LoadPtr(&target)

	fmt.Println("target =", target.Value.Value)

	fmt.Println("target <- bar")
	target.Value.Value = "bar"

	fmt.Println("target =", target.Value.Value)
	fmt.Println("src =", src.Value.Value)
}
