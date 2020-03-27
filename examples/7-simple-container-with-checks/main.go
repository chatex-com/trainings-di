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
	e, ok := di.SimpleGetWithChecks(A{})
	if !ok {
		fmt.Println("problem with getting from container")
	}
	target, ok = e.(*A)
	if !ok {
		fmt.Println("problem with type casting")
	}

	fmt.Println("target =", target.Value)

	fmt.Println("target <- bar")
	target.Value = "bar"

	fmt.Println("target =", target.Value)
	fmt.Println("src =", src.Value)
}
