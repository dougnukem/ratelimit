# ratelimit [![Build Status](https://travis-ci.org/dougnukem/ratelimit.png)](https://travis-ci.org/dougnukem/ratelimit) [![Coverage](https://gocover.io/_badge/github.com/dougnukem/ratelimit)](https://gocover.io/github.com/dougnukem/ratelimit)[![Go Report Card](http://goreportcard.com/badge/dougnukem/ratelimit)](http://goreportcard.com/report/dougnukem/ratelimit) [![GoDoc](https://godoc.org/github.com/dougnukem/ratelimit?status.png)](https://godoc.org/github.com/dougnukem/ratelimit)

A [Token Bucket](https://en.wikipedia.org/wiki/Token_bucket) based rate limiter implemented implemented using go-routines.

# Examples

## RateLimiter.Wait - Blocking

see godocs [RateLimiter-WaitMaxDuration](https://godoc.org/github.com/dougnukem/ratelimit#example-RateLimiter-WaitMaxDuration)
```go
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

```

## RateLimiter.Wait - Blocking

see godocs [RateLimiter-Wait](https://godoc.org/github.com/dougnukem/ratelimit#example-RateLimiter-Wait)
```go
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

```



# Contributing
See [Contributing](Contributing.md)

# License
[MIT License](LICENSE)

# TODO
- godoc documentation
- write examples
- benchmark vs non-goroutine implementations like https://github.com/beefsack/go-rate
- configure rate limit policy
  - by default we start the rate limit interval time only when it's actually used (e.g. following twitter api rate limit semantics)
    - more commonly like in a traditionally token bucket algorithm the interval is continuously replenishing tokens see [Token Bucket](https://en.wikipedia.org/wiki/Token_bucket)
