package belajargolanggoroutines

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func Test_Once(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			once.Do(OnlyOnce)
			fmt.Println("perulangan ke-", i)
		}()
	}

	group.Wait()
	fmt.Println("Counter:", counter)
}
