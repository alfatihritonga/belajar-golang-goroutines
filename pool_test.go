package belajargolanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			return "Data baru"
		},
	}

	pool.Put("Data 1")
	pool.Put("Data 2")
	pool.Put("Data 3")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(2 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}
