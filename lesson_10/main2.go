package main

import (
	"fmt"
	"sync"
	"time"
)

/// wait group - примитив синхронизации

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ { //цикл
		wg.Add(1)
		k := i
		go func() { // создаем ананимную горутину
			defer wg.Done() // выхываем метод

			fmt.Println("goroutine working...\n", k)
			time.Sleep(300 * time.Millisecond)
		}()

	}
	wg.Wait()
	fmt.Println("all done")

}
