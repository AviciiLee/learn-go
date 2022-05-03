package share_mem

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCount(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(i int) {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}(i)
	}
	wg.Wait()
	t.Logf("counter = %d", counter)
}

func service() string {
	time.Sleep(time.Millisecond * 100)
	return "hello"
}

func AsyncService() chan string {
	retCh := make(chan string)
	go func() {
		ret := service()
		retCh <- ret
	}()
	return retCh
}

func TestChannel(t *testing.T) {
	retCh := AsyncService()
	fmt.Println("hahaha")
	fmt.Println(<-retCh)
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
