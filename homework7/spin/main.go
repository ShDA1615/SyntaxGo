package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(50 * time.Millisecond)

	time.Sleep(10 * time.Second)

}

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
}

func fibonacci(x int) int {
	if x < 2 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}
