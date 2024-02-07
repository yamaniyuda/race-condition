package race_condition_case_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSyncPollHandling(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			return "New"
		},
	}
	group := sync.WaitGroup{}

	pool.Put("Ahmad")
	pool.Put("Yuda")
	pool.Put("Zaki")

	for i := 0; i < 10; i++ {
		go func() {
			group.Add(1)
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("selesai")
}
