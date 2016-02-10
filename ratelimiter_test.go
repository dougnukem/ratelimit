package ratelimit

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// assertReceive asserts that a bool message is received on the channel before a specified timeout
func assertReceive(t *testing.T, c <-chan struct{}, timeout time.Duration, errorMsg string, msgAndArgs ...interface{}) {
	select {
	case <-c:
		break
	case <-time.After(timeout):
		t.Errorf(errorMsg, msgAndArgs...)
		t.FailNow()
	}
}

var testInterval = time.Millisecond * 10
var testTimeout = testInterval * 4

func TestRateLimiter_Wait(t *testing.T) {

	r := NewRateLimiter(testInterval, 2, 2)
	start := time.Now()
	var end time.Time
	completed := make(chan struct{})
	go func() {
		r.Wait()
		r.Wait()
		r.Wait()
		end = time.Now()
		completed <- struct{}{}
	}()

	assertReceive(t, completed, testTimeout, "Expected rate limiter to complete but timedout")

	expectedEnd := start.Add(testInterval)
	assert.WithinDuration(t, expectedEnd, end, testInterval, "Expected rate limiter waiting interval")
}
