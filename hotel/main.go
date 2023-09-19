package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		time.Sleep(time.Second * 3) // if less than time.After then cancel booking
		cancel()
	}()

	bookHotel(ctx)
}

// context must be the first parameter
func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled")
	case <-time.After(time.Second * 5):
		fmt.Println("Hotel booked")
	}
}
