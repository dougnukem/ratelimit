package ratelimit_test

import (
	"fmt"
	"github.com/dougnukem/ratelimit"
	"time"
)

func ExampleRateLimiter_Wait() {
	// 2 requests per second refreshing 2 capacity every second
	r := ratelimit.NewRateLimiter(time.Second, 2, 2)
	start := time.Now()
	r.Wait()
	fmt.Printf("r.Wait() elapsed less than 500 ms [%t]\n", time.Now().Sub(start) < 500*time.Millisecond)
	r.Wait()
	fmt.Printf("r.Wait() elapsed less than 500 ms [%t]\n", time.Now().Sub(start) < 500*time.Millisecond)
	r.Wait()
	fmt.Printf("r.Wait() elapsed greater than 500 ms [%t]\n", time.Now().Sub(start) > 500*time.Millisecond)
	// Output:
	// r.Wait() elapsed less than 500 ms [true]
	// r.Wait() elapsed less than 500 ms [true]
	// r.Wait() elapsed greater than 500 ms [true]
}

func ExampleRateLimiter_WaitMaxDuration() {
	// 2 requests per second refreshing 2 capacity every second
	r := ratelimit.NewRateLimiter(time.Second, 2, 2)
	maxDuration := 500 * time.Millisecond
	s := r.WaitMaxDuration(maxDuration)
	fmt.Printf("r.Wait() success[%t]\n", s)
	s = r.WaitMaxDuration(maxDuration)
	fmt.Printf("r.Wait() success[%t]\n", s)
	s = r.WaitMaxDuration(maxDuration)
	fmt.Printf("r.Wait() success[%t]\n", s)
	// Output:
	// r.Wait() success[true]
	// r.Wait() success[true]
	// r.Wait() success[false]
}

func ExampleRateLimiter_Try() {
	// 2 requests per second refreshing 2 capacity every second
	r := ratelimit.NewRateLimiter(time.Second, 2, 2)
	s := r.Try()
	fmt.Printf("r.Try() success[%t]\n", s)
	s = r.Try()
	fmt.Printf("r.Try() success[%t]\n", s)
	s = r.Try()
	fmt.Printf("r.Try() success[%t]\n", s)
	// Output: r.Try() success[true]
	// r.Try() success[true]
	// r.Try() success[false]
}
