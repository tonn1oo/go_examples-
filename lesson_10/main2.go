package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/// wait group - примитив синхронизации

func main() {
	var wg sync.WaitGroup
	var counter uint64

	for i := 0; i < 10; i++ { //цикл
		wg.Add(1)
		k := i
		go func() { // создаем ананимную горутину
			defer wg.Done() // выхываем метод

			for j := 0; j < 1000; j++ {
				atomic.AddUint64(&counter, 1)
			}

			fmt.Println("goroutine working...\n", k)
			time.Sleep(300 * time.Millisecond)
		}()

	}
	wg.Wait()
	fmt.Println("all done")

}
