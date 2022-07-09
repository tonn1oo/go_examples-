package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

/// работа с контекстом

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		err := cancelRequest(ctx)
		if err != nil {
			cancel()
		}
	}()
	doRequest(ctx, "https://ya.ru")
}

func cancelRequest(ctx context.Context) error {
	time.Sleep(100 * time.Millisecond)
	return fmt.Errorf("fail request")
}

func doRequest(ctx context.Context, requestStr string) {
	req, _ := http.NewRequest(http.MethodGet, requestStr, nil)
	req = req.WithContext(ctx)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)

	}
	select {
	case <-time.After(500 * time.Microsecond):
		fmt.Printf("response completed, status code=%d", res.StatusCode)
	case <-ctx.Done():
		fmt.Println("request takes too long")

	}
}
