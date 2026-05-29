# Lesson 9 — Rate Limiter (Token Bucket)

## Spec

- Implement a `RateLimiter` with `Allow() bool` and `Wait()` methods
- Configurable: `rate` (tokens per second) and `burst` (max bucket capacity)
- Tokens are refilled by a background goroutine driven by `time.Ticker`
- `Allow()` returns immediately: `true` if a token was available (consumed), `false` otherwise
- `Wait()` blocks until a token is available, then consumes it
- Provide a `Stop()` method to halt the refill goroutine cleanly (no leaks)
- Demo: try to perform 20 operations as fast as possible at rate=5/sec, observe pacing

---

1. **Pick the bucket representation** — a buffered channel of empty structs (`chan struct{}`) of capacity = `burst` is the cleanest token bucket in Go. Each element in the channel is a token. `Allow` is a non-blocking receive; `Wait` is a blocking receive. Docs: [Go blog: Rate limiting](https://gobyexample.com/rate-limiting) — https://gobyexample.com/rate-limiting | [pkg: time](https://pkg.go.dev/time) — https://pkg.go.dev/time

   > `chan struct{}` is the canonical "signal-only" channel — `struct{}` is zero bytes, so it costs nothing per element. The channel is just used for its buffering and blocking semantics. Idiomatic Go.

2. **Declare RateLimiter struct** — fields: `tokens chan struct{}`, `ticker *time.Ticker`, `done chan struct{}`. Constructor `New(rate int, burst int) *RateLimiter` creates the channel with capacity `burst`, pre-fills it with `burst` tokens, starts the ticker at interval `time.Second / time.Duration(rate)`, and launches the refill goroutine. Docs: [pkg: time.Ticker](https://pkg.go.dev/time#Ticker) — https://pkg.go.dev/time#Ticker

3. **Write the refill goroutine** — a `for { select { case <-rl.ticker.C: ... case <-rl.done: return } }` loop. On each tick, attempt a **non-blocking send** of a token into the bucket (using `select` with a `default` clause). If the bucket is full, drop the token — that's the cap. Docs: [Tour: Select](https://go.dev/tour/concurrency/5) — https://go.dev/tour/concurrency/5

   > Non-blocking sends and receives via `select` with `default` are how you "try but don't wait" in Go. This is a fundamental concurrency pattern — once you see it, you'll spot it everywhere in good Go code.

4. **Implement Allow()** — non-blocking receive: `select { case <-rl.tokens: return true; default: return false }`. Immediate decision, no blocking.

5. **Implement Wait()** — blocking receive: `<-rl.tokens`. Simple as that. (For production you'd accept a `context.Context` to allow cancellation — call that out as a follow-up.)

6. **Implement Stop()** — `rl.ticker.Stop()` then `close(rl.done)`. Idempotency is nice to have but not required for the exercise.

7. **Write main** — construct `New(5, 5)`, loop 20 times calling `Wait()` and printing `request i at <elapsed ms>`. You should see the first 5 fire instantly (the initial burst), then one every 200ms.

8. **Bonus** — add `AllowN(n int)` that consumes N tokens atomically. Tricky: you have to peek without committing. Stretch: rewrite using `golang.org/x/time/rate` and compare ergonomics.
