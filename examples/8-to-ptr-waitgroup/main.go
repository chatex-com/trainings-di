package main

import (
	"sync"
	"time"

	"di/di"
)

type A struct {
	Value sync.WaitGroup
}

func main() {
	src := A{}
	_ = di.Set(&src)

	target := &A{}
	_ = di.LoadPtrToPtr(&target)

	src.Value.Add(1)
	go func() {
		<-time.After(5 * time.Second)
		src.Value.Done()
	}()

	target.Value.Wait()
}
