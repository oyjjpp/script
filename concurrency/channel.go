package concurrency

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

// channelConcurrency
// 通过管道实现并发控制
func channelConcurrency() {
	ch := make(chan struct{})

	go func() {
		log.Printf("start working")
		time.Sleep(time.Second * 3)
		ch <- struct{}{}
	}()

	<-ch
	log.Println("finished")
	println()
	fmt.Println()
}

// syncConcurrency
// 通过WaitGroup实现并发控制
func syncConcurrency() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			log.Println(i)
		}(i)
	}
	wg.Wait()
	time.Sleep(5 * time.Second)
	log.Println("finished ", runtime.NumGoroutine())
}

func withContextConcurrency() {
	ctx, cancel := context.WithCancel(context.Background())

	go work(ctx, "one")
	go work(ctx, "two")
	go work(ctx, "three")
	// time.Sleep(2 * time.Second)
	log.Println("stop the goroutine")
	cancel()
	time.Sleep(time.Second * 2)
}

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			log.Println(name, "one stop channel")
			return
		default:
			log.Println(name, "one still workinf")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
