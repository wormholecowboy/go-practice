# Lesson 7 — Variadic Logger

## Spec

- Build a small `Logger` type with level-aware methods: `Debug`, `Info`, `Warn`, `Error`
- Each method takes variadic args (`...any`) and formats like `fmt.Println`
- Logger has a minimum `Level` — messages below it are silently dropped
- Output goes to a configurable `io.Writer` (default `os.Stderr`)
- Every line is prefixed: `[LEVEL] 2026-05-29T14:00:00Z message...`
- Write a `main` that constructs a logger at `Info` level and demonstrates all four methods (Debug should be filtered out)

---

1. **Define log levels with iota** — declare a `Level` type as `int` and use a `const ( … iota )` block for `LevelDebug`, `LevelInfo`, `LevelWarn`, `LevelError`. Add a `String()` method so `LevelInfo.String() == "INFO"`. Docs: [Effective Go: Constants](https://go.dev/doc/effective_go#constants) — https://go.dev/doc/effective_go#constants | [Go spec: iota](https://go.dev/ref/spec#Iota) — https://go.dev/ref/spec#Iota

   > `iota` + a `String()` method on a typed int is Go's idiomatic enum. You'll see this pattern in the stdlib (`time.Weekday`, `http.StatusCode` essentially). Worth internalizing.

2. **Declare the Logger struct** — fields: `minLevel Level`, `out io.Writer`. Provide a `New(minLevel Level, out io.Writer) *Logger` constructor. If `out` is nil, default to `os.Stderr`. Docs: [pkg: io](https://pkg.go.dev/io) — https://pkg.go.dev/io | [pkg: log](https://pkg.go.dev/log) — https://pkg.go.dev/log

3. **Write the variadic log core** — a private method `(l *Logger) log(level Level, args ...any)`. Bail early if `level < l.minLevel`. Otherwise format `[LEVEL] timestamp message` using `fmt.Fprintln` to write to `l.out`. Use `time.Now().UTC().Format(time.RFC3339)` for the timestamp. Docs: [pkg: fmt](https://pkg.go.dev/fmt) — https://pkg.go.dev/fmt | [pkg: time](https://pkg.go.dev/time) — https://pkg.go.dev/time

   > Variadic functions (`args ...any`) are how Go handles "any number of arguments" — under the hood it's just a slice. `fmt.Println` itself is variadic, which is why you can pass it any number of values.

4. **Implement the four public methods** — `Debug`, `Info`, `Warn`, `Error` each just forward to `log` with the appropriate level. Tiny wrappers on purpose — keeps the level constant at the call site readable.

5. **Wire up main** — construct a `Logger` at `LevelInfo` writing to `os.Stderr`. Call all four methods with mixed argument types (strings, ints, structs). Confirm Debug is filtered, the other three print with correct prefixes.

6. **Bonus** — add a `SetLevel(Level)` method so the threshold can be changed at runtime. Bonus bonus: make `Error` call `os.Exit(1)` after logging — that's how `log.Fatal` works under the hood.
