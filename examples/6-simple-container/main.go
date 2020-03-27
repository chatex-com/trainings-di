package main

import (
	"fmt"

	"di/di"
)

type A struct {
	Value string
}

func main() {
	src := A{Value: "foo"}
	di.SimpleSet(&src)

	var target *A
	e := di.SimpleGet(A{})
	target = e.(*A)

	fmt.Println("target =", target.Value)

	fmt.Println("target <- bar")
	target.Value = "bar"

	fmt.Println("target =", target.Value)
	fmt.Println("src =", src.Value)
}
