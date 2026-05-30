# Go Micro-Project Ideas

Each project fits in one sitting. Docs links included — no hand-holding.

---

## Syntax & Fundamentals

- [x] **FizzBuzz CLI** — loops, conditionals, `fmt`, `os.Args`. [Tour of Go](https://go.dev/tour/)
- [x] **Unit converter** — structs, methods, iota/const enums. [Effective Go](https://go.dev/doc/effective_go)
- [x] **Fibonacci (iterative + recursive)** — functions, multiple return values. [Go spec: Functions](https://go.dev/ref/spec#Function_declarations)
- [x] **String frequency counter** — maps, range, `strings` pkg. [pkg: strings](https://pkg.go.dev/strings)
- [ ] **Stack & Queue** — generics (Go 1.18+), type parameters. [Go generics intro](https://go.dev/doc/faq#generics)
- [ ] **Variadic logger** — variadic functions, `log` package, log levels. [pkg: log](https://pkg.go.dev/log)

---

## Concurrency (Go's superpower)

- [ ] **Ping-pong channels** — goroutines, unbuffered channels, `sync.WaitGroup`. [Tour: Concurrency](https://go.dev/tour/concurrency/1)
- [ ] **Fan-out worker pool** — buffered channels, N workers consuming a job queue. [Go blog: Pipelines](https://go.dev/blog/pipelines)
- [ ] **Rate limiter** — `time.Ticker`, channel-based token bucket. [Go blog: Rate limiting](https://gobyexample.com/rate-limiting)
- [ ] **Parallel file hasher** — goroutines, `crypto/md5`, `sync.Mutex`. [pkg: sync](https://pkg.go.dev/sync)
- [ ] **Context timeout demo** — `context.WithTimeout`, `context.WithCancel`, cancellation propagation. [pkg: context](https://pkg.go.dev/context)

---

## Systems & OS

- [x] **`wc` clone** — read stdin/file, count lines/words/bytes, `bufio.Scanner`. [pkg: bufio](https://pkg.go.dev/bufio), [pkg: os](https://pkg.go.dev/os)
- [ ] **File watcher** — poll a dir for changes using `os.Stat`, goroutines, ticker. [pkg: filepath](https://pkg.go.dev/path/filepath)
- [ ] **Process lister** — exec `ps` via `os/exec`, parse output, pretty-print. [pkg: os/exec](https://pkg.go.dev/os/exec)
- [ ] **TCP echo server** — `net.Listen`, `net.Conn`, goroutine per connection. [pkg: net](https://pkg.go.dev/net)
- [ ] **UDP ping** — raw UDP client/server, `net.PacketConn`. [pkg: net](https://pkg.go.dev/net)
- [ ] **Signal handler** — trap SIGINT/SIGTERM, graceful shutdown. [pkg: os/signal](https://pkg.go.dev/os/signal)
- [ ] **Memory usage reporter** — `runtime.MemStats`, `runtime.ReadMemStats`. [pkg: runtime](https://pkg.go.dev/runtime)

---

## Classic Algorithms

- [ ] **Binary search** — slices, zero-value semantics, benchmarking with `testing.B`. [pkg: testing](https://pkg.go.dev/testing)
- [ ] **Merge sort** — recursion, slice manipulation, generics ordering constraint. [Go spec: Comparison operators](https://go.dev/ref/spec#Comparison_operators)
- [ ] **Linked list** — pointer semantics, method sets on pointer receivers. [Effective Go: Pointers](https://go.dev/doc/effective_go#pointers_vs_values)
- [ ] **Ring buffer** — fixed-size circular queue, head/tail index wrapping, O(1) enqueue+dequeue. [Go spec: Arrays](https://go.dev/ref/spec#Array_types) — https://go.dev/ref/spec#Array_types
- [ ] **LRU cache** — maps + doubly linked list, `container/list`. [pkg: container/list](https://pkg.go.dev/container/list)
- [ ] **BFS/DFS on a graph** — adjacency list with maps, recursion vs iteration. [Go spec: Maps](https://go.dev/ref/spec#Map_types)
- [ ] **Trie** — nested structs, rune iteration over strings. [Go blog: Strings](https://go.dev/blog/strings)

---

## Web & HTTP

- [ ] **Hello World HTTP server** — `net/http`, `http.HandleFunc`, query params. [pkg: net/http](https://pkg.go.dev/net/http)
- [ ] **JSON REST API (no framework)** — `encoding/json`, `Decode`/`Encode`, status codes. [pkg: encoding/json](https://pkg.go.dev/encoding/json)
- [ ] **Middleware chain** — `http.Handler` interface, wrapping handlers. [Go blog: HTTP handlers](https://go.dev/blog/http-handler)
- [ ] **Static file server** — `http.FileServer`, `http.StripPrefix`. [pkg: net/http](https://pkg.go.dev/net/http)
- [ ] **HTTP client with retries** — `http.Client`, custom `Transport`, exponential backoff. [pkg: net/http](https://pkg.go.dev/net/http)
- [ ] **Webhook receiver** — validate HMAC signature, `crypto/hmac`. [pkg: crypto/hmac](https://pkg.go.dev/crypto/hmac)

---

## HTML Generation & Templating

- [ ] **HTML report generator** — `html/template`, template inheritance, escaping. [pkg: html/template](https://pkg.go.dev/html/template)
- [ ] **Static site generator (tiny)** — parse Markdown files, render to HTML with `text/template`. [pkg: text/template](https://pkg.go.dev/text/template)
- [ ] **Table renderer** — generate an HTML table from a CSV file. [pkg: encoding/csv](https://pkg.go.dev/encoding/csv)

---

## CLI Tools

- [ ] **Colored CLI output** — `flag` package for args, ANSI escape codes. [pkg: flag](https://pkg.go.dev/flag)
- [ ] **Interactive prompt** — read from stdin line-by-line, REPL loop. [pkg: bufio](https://pkg.go.dev/bufio)
- [ ] **Config file reader** — parse a `.env` or TOML file manually, then with a lib. [pkg: os](https://pkg.go.dev/os)
- [ ] **Progress bar** — `\r` carriage return trick, terminal width via `COLUMNS`. 

---

## Data & Encoding

- [ ] **CSV parser** — `encoding/csv`, struct tags, reflection basics. [pkg: encoding/csv](https://pkg.go.dev/encoding/csv)
- [ ] **JSON → struct round-trip** — marshaling, `omitempty`, custom `MarshalJSON`. [pkg: encoding/json](https://pkg.go.dev/encoding/json)
- [ ] **Base64 encoder/decoder** — `encoding/base64`, URL-safe variant. [pkg: encoding/base64](https://pkg.go.dev/encoding/base64)
- [ ] **Simple cipher** — XOR cipher, then Caesar cipher — byte manipulation. [Go spec: Numeric types](https://go.dev/ref/spec#Numeric_types)

---

## Testing & Tooling

- [ ] **Table-driven tests** — `testing` package, subtests with `t.Run`. [Go blog: Table-driven tests](https://go.dev/wiki/TableDrivenTests)
- [ ] **Benchmark suite** — `testing.B`, `pprof` CPU profile. [pkg: runtime/pprof](https://pkg.go.dev/runtime/pprof)
- [ ] **Fuzz test** — `testing.F`, `f.Add`, `f.Fuzz` for a parser. [Go fuzzing guide](https://go.dev/doc/fuzz/)

---

## Stretch / Fun

- [ ] **Brainfuck interpreter** — byte pointer, array state, classic esoteric lang. 
- [ ] **Mandelbrot PNG** — `image/png`, `image/color`, complex number math. [pkg: image](https://pkg.go.dev/image)
- [ ] **Mini key-value store** — TCP server + simple text protocol, `sync.RWMutex`. [pkg: sync](https://pkg.go.dev/sync)
- [ ] **Port scanner** — concurrent TCP dial with timeout, `sync.WaitGroup`. [pkg: net](https://pkg.go.dev/net)
- [ ] **DNS lookup tool** — `net.LookupHost`, `net.LookupMX`, compare results. [pkg: net](https://pkg.go.dev/net)

---

## Must-Read Docs (bookmark these)

- [A Tour of Go](https://go.dev/tour/) — interactive fundamentals
- [Effective Go](https://go.dev/doc/effective_go) — idiomatic style
- [Go by Example](https://gobyexample.com/) — one-pagers for every concept
- [pkg.go.dev](https://pkg.go.dev/) — stdlib + third-party package docs
- [Go Proverbs](https://go-proverbs.github.io/) — philosophy
- [Go spec](https://go.dev/ref/spec) — authoritative reference
