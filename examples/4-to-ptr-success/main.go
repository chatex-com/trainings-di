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
	_ = di.Set(&src)

	target := &A{}
	_ = di.LoadPtrToPtr(&target)

	fmt.Println("target =", target.Value)

	fmt.Println("target <- bar")
	target.Value = "bar"

	fmt.Println("target =", target.Value)
	fmt.Println("src =", src.Value)
}
