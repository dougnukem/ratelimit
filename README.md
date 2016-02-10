# ratelimit [![Build Status](https://travis-ci.org/dougnukem/ratelimit.png)](https://travis-ci.org/dougnukem/ratelimit) [![Coverage](https://gocover.io/_badge/github.com/dougnukem/ratelimit)](https://gocover.io/github.com/dougnukem/ratelimit)[![Go Report Card](http://goreportcard.com/badge/dougnukem/ratelimit)](http://goreportcard.com/report/dougnukem/ratelimit) [![GoDoc](https://godoc.org/github.com/dougnukem/ratelimit?status.png)](https://godoc.org/github.com/dougnukem/ratelimit)

A [Token Bucket](https://en.wikipedia.org/wiki/Token_bucket) based rate limiter implemented implemented using go-routines.

# Examples

## Blocking RateLimiter.Wait

see [examples/wait_example.go](examples/wait_example.go)
```go
// 2 requests per second refreshing 2 capacity every second
r := ratelimit.NewRateLimiter(time.Second, 2, 2)
start := time.Now()
r.Wait()
fmt.Printf("r.Wait() elapsed[%s] time[%s]\n", time.Now().Sub(start), time.Now())
r.Wait()
fmt.Printf("r.Wait() elapsed[%s] time[%s]\n", time.Now().Sub(start), time.Now())
r.Wait()
fmt.Printf("r.Wait() elapsed[%s] time[%s]\n", time.Now().Sub(start), time.Now())

```

```bash
$ go run examples/wait_example.go
r.Wait() duration[25.783µs] time[2016-02-09 19:01:05.397723074 -0600 CST]
r.Wait() duration[265.595µs] time[2016-02-09 19:01:05.397962943 -0600 CST]
r.Wait() duration[1.002628384s] time[2016-02-09 19:01:06.400325702 -0600 CST]
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
