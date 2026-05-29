# Lesson 8 — Ping-Pong Channels

## Spec

- Two goroutines, "pinger" and "ponger", pass a "ball" back and forth N times
- The ball is an `int` representing the current hit count, sent over an unbuffered `chan int`
- Pinger prints `ping N`, ponger prints `pong N`
- Main waits for both goroutines via `sync.WaitGroup`
- N (number of total hits) configurable via `os.Args` or a constant — default 10
- Program exits cleanly when N hits done, no deadlock, no leaked goroutines

---

1. **Set up the channel** — declare `ball := make(chan int)`. Crucially **unbuffered** — sends block until a receiver is ready. That's the synchronization mechanism that makes ping-pong work. Docs: [Tour: Channels](https://go.dev/tour/concurrency/2) — https://go.dev/tour/concurrency/2 | [Effective Go: Channels](https://go.dev/doc/effective_go#channels) — https://go.dev/doc/effective_go#channels

   > Unbuffered channels are Go's most elegant primitive — they're not just data pipes, they're synchronization points. A send and a receive on an unbuffered channel are a *rendezvous*. This is the heart of CSP (Communicating Sequential Processes).

2. **Set up the WaitGroup** — declare `var wg sync.WaitGroup`. Call `wg.Add(2)` before launching the goroutines (Add before Go — important). Each goroutine calls `defer wg.Done()`. Main calls `wg.Wait()` at the end. Docs: [pkg: sync](https://pkg.go.dev/sync) — https://pkg.go.dev/sync | [pkg: sync.WaitGroup](https://pkg.go.dev/sync#WaitGroup) — https://pkg.go.dev/sync#WaitGroup

3. **Write the pinger goroutine** — loops, receives the ball, prints `ping <n>`, then sends `n+1` back. Exit condition: when received `n >= N`, stop. Closing the channel cleanly is part of the puzzle — decide whether pinger or main closes it. Docs: [Go blog: Pipelines and cancellation](https://go.dev/blog/pipelines) — https://go.dev/blog/pipelines

4. **Write the ponger goroutine** — mirror image of pinger but prints `pong <n>`. To kick off the rally, one side needs to send first (or main sends the initial `1`). Decide which.

   > Bootstrapping is half the fun of CSP-style programs. Who starts? Who finishes? Who closes the channel? Get these wrong and you deadlock — Go will tell you with the famous `fatal error: all goroutines are asleep - deadlock!`.

5. **Launch and wait** — `go pinger(...)`, `go ponger(...)`, optionally seed with `ball <- 1`, then `wg.Wait()`. Docs: [Tour: Goroutines](https://go.dev/tour/concurrency/1) — https://go.dev/tour/concurrency/1

6. **Run it** — confirm the output alternates strictly `ping 1`, `pong 2`, `ping 3`, … up to N. If you see two pings in a row, you've made the channel buffered or introduced a race.

7. **Bonus** — run with `go run -race main.go`. The race detector should report nothing because channels are race-free by design. Now try replacing the channel with a shared int and a mutex and watch the race show up.
