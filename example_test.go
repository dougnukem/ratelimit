package ratelimit

import (
	"fmt"
	"time"
)

func ExampleRateLimiter_Wait() {
	// 2 requests per second refreshing 2 capacity every second
	r := NewRateLimiter(time.Second, 2, 2)
	start := time.Now()
	r.Wait()
	fmt.Printf("r.Wait() elapsed[%s] time[%s]\n", time.Now().Sub(start), time.Now())
	r.Wait()
	fmt.Printf("r.Wait() elapsed[%s] time[%s]\n", time.Now().Sub(start), time.Now())
	r.Wait()
	fmt.Printf("r.Wait() elapsed[%s] time[%s]\n", time.Now().Sub(start), time.Now())
}

func ExampleRateLimiter_WaitMaxDuration() {
	// 2 requests per second refreshing 2 capacity every second
	r := NewRateLimiter(time.Second, 2, 2)
	maxDuration := 500 * time.Millisecond
	start := time.Now()
	s := r.WaitMaxDuration(maxDuration)
	fmt.Printf("r.Wait() success[%s] elapsed[%s] time[%s]\n", s, time.Now().Sub(start), time.Now())
	s = r.WaitMaxDuration(maxDuration)
	fmt.Printf("r.Wait() success[%s] elapsed[%s] time[%s]\n", s, time.Now().Sub(start), time.Now())
	s = r.WaitMaxDuration(maxDuration)
	fmt.Printf("r.Wait() success[%s] elapsed[%s] time[%s]\n", s, time.Now().Sub(start), time.Now())
}

func ExampleRateLimiter_Try() {
	// 2 requests per second refreshing 2 capacity every second
	r := NewRateLimiter(time.Second, 2, 2)
	start := time.Now()
	s := r.Try()
	fmt.Printf("r.Try() success[%t] elapsed[%s] time[%s]\n", s, time.Now().Sub(start), time.Now())
	s = r.Try()
	fmt.Printf("r.Try() success[%t] elapsed[%s] time[%s]\n", s, time.Now().Sub(start), time.Now())
	s = r.Try()
	fmt.Printf("r.Try() success[%t] elapsed[%s] time[%s]\n", s, time.Now().Sub(start), time.Now())
}
