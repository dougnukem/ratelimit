package ratelimit

import (
	"time"
)

type RateLimiter struct {
	// config
	resetInterval    time.Duration
	itemsPerInterval int64
	limit            int64
	// state
	remaining     int64
	intervalStart time.Time
	resetAt       time.Time
	// channels
	requestToken chan struct{}
	updateLimits chan updateLimit
}

type updateLimit struct {
	limit     int64
	remaining int64
	resetAt   time.Time
}

func NewRateLimiter(resetInterval time.Duration, itemsPerInterval, limit int64) *RateLimiter {
	requestToken := make(chan struct{})
	updateLimits := make(chan updateLimit)
	r := &RateLimiter{
		resetInterval:    resetInterval,
		itemsPerInterval: itemsPerInterval,
		limit:            limit,
		remaining:        itemsPerInterval,
		requestToken:     requestToken,
		updateLimits:     updateLimits,
	}

	go r.run()

	return r
}

func (r *RateLimiter) Wait() {
	r.requestToken <- struct{}{}
}

func (r *RateLimiter) WaitMaxDuration(duration time.Duration) bool {
	select {
	case r.requestToken <- struct{}{}:
		return true
	case <-time.After(duration):
		return false
	}
}

func (r *RateLimiter) Try() bool {
	select {
	case r.requestToken <- struct{}{}:
		return true
	default:
		return false
	}
}

func (r *RateLimiter) Update(limit, remaining int64, resetAt time.Time) {
	r.updateLimits <- updateLimit{limit: limit, remaining: remaining, resetAt: resetAt}
}

func (r *RateLimiter) run() {
	for {
		durationTilReset := r.runStateUpdate()
		switch {
		case r.limit == r.remaining:
			r.handleInitialState()
		case r.remaining > 0:
			r.handleRateIntervalStarted(durationTilReset)
		default:
			r.handleRateLimitReached(durationTilReset)
		}
	}
}

func (r *RateLimiter) handleInitialState() {
	select {
	case <-r.requestToken:
		now := time.Now()
		r.update(r.limit, r.remaining-1, now.Add(r.resetInterval))
	case u := <-r.updateLimits:
		r.update(u.limit, u.remaining, u.resetAt)
	}
}

func (r *RateLimiter) handleRateLimitReached(durationTilReset time.Duration) {
	select {
	case u := <-r.updateLimits:
		r.update(u.limit, u.remaining, u.resetAt)
	case <-time.After(durationTilReset):
		// will reset when we update state in main run() loop
	}
}

func (r *RateLimiter) handleRateIntervalStarted(durationTilReset time.Duration) {
	select {
	case <-r.requestToken:
		now := time.Now()
		r.update(r.limit, r.remaining-1, now.Add(r.resetInterval))
	case u := <-r.updateLimits:
		r.update(u.limit, u.remaining, u.resetAt)
	}
}

func (r *RateLimiter) update(limit, remaining int64, resetAt time.Time) {
	r.limit = limit
	r.remaining = remaining
	r.resetAt = resetAt
}

func (r *RateLimiter) runStateUpdate() time.Duration {
	if r.resetAt.IsZero() && r.limit == r.remaining {
		return 0
	}

	now := time.Now()
	durationTilReset := r.resetAt.Sub(now)

	if durationTilReset > 0 {
		return durationTilReset
	}

	remaining := r.remaining + r.itemsPerInterval
	if remaining > r.limit {
		remaining = r.limit
	}
	// Reset to new interval (no resetAt time until an item is taken)
	r.update(r.limit, remaining, time.Time{})
	return 0
}
