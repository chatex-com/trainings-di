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

	var target *A
	err := di.LoadPtrToPtr(&target)

	fmt.Println(err)
}
