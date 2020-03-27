package main

import (
	"di/di"
)

type A struct {
	Value string
}

func main() {
	src := A{Value: "foo"}
	_ = di.Set(&src)

	var target *A
	_ = di.LoadPtrToPtr(&target)
}
