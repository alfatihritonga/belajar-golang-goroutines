package belajargolanggoroutines

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func AddToMap(data *sync.Map, key string, value int, group *sync.WaitGroup) {
	defer group.Done()

	data.Store(key, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 1; i <= 100; i++ {
		group.Add(1)
		key := "Data ke-" + strconv.Itoa(i)
		go AddToMap(data, key, i, group)
	}

	group.Wait()

	data.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})

	fmt.Println("Selesai")
}
