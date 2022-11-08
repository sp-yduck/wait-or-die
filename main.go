package main

import (
	"fmt"
	"time"
)

type waitableFunc func() error

func waitOrDie(fn waitableFunc, interval int, timeout int) {
	for t := 0; t < timeout; t += interval {
		if err := fn(); err != nil {
			time.Sleep(time.Duration(interval) * 1000 * time.Millisecond)
			fmt.Printf("waiting %d\n", t)
			continue
		}
		return
	}
}

func sleep() error {
	time.Sleep(time.Duration(1) * 1000 * time.Millisecond)
	fmt.Println("good morning.")
	return fmt.Errorf("sleeping")
}

func main() {
	waitOrDie(sleep, 1, 8)
}
