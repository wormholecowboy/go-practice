# Lesson 10 — File Watcher (Polling)

## Spec

- Accept a directory path as a CLI argument
- Poll the directory every N ms (default 500ms) and detect:
  - **Created** — new files since last poll
  - **Deleted** — files present last poll, gone now
  - **Modified** — files whose `ModTime` advanced
- Print events to stdout in the form: `CREATE path/to/file`, `MODIFY path/to/file`, `DELETE path/to/file`
- Graceful shutdown on SIGINT (Ctrl-C): print `bye` and exit 0
- Top-level only — don't recurse into subdirectories (mark as a bonus)

---

1. **Read directory entries** — use `os.ReadDir(dir)` to list entries. For each entry, call `entry.Info()` to get a `fs.FileInfo` so you can read `ModTime()`. Skip directories. Docs: [pkg: os](https://pkg.go.dev/os) — https://pkg.go.dev/os | [pkg: io/fs](https://pkg.go.dev/io/fs) — https://pkg.go.dev/io/fs

2. **Build a snapshot map** — `map[string]time.Time` from filename → modtime. This is your "state of the world" at one moment in time. Docs: [pkg: time](https://pkg.go.dev/time) — https://pkg.go.dev/time

   > Diffing two snapshots is the classic poll-based change-detection pattern. It's not as efficient as kernel-level notifications (`inotify` on Linux, `FSEvents` on macOS, `ReadDirectoryChangesW` on Windows — that's what `fsnotify` wraps), but it's simple, portable, and good enough for many tools.

3. **Set up the ticker loop** — `ticker := time.NewTicker(500 * time.Millisecond)`, `defer ticker.Stop()`. Loop with `for { select { case <-ticker.C: ... case <-ctx.Done(): return } }`. Docs: [pkg: time.NewTicker](https://pkg.go.dev/time#NewTicker) — https://pkg.go.dev/time#NewTicker

4. **Diff old vs new snapshots** — three loops:
   - For every key in `new` not in `old` → `CREATE`
   - For every key in `old` not in `new` → `DELETE`
   - For every key in both where `new[k].After(old[k])` → `MODIFY`
   
   After diffing, replace `old = new` for the next iteration.

5. **Handle SIGINT for graceful shutdown** — `sigChan := make(chan os.Signal, 1)`, `signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)`. Use a `context.WithCancel` — when the signal arrives, call `cancel()`. The ticker loop sees `ctx.Done()` and returns. Docs: [pkg: os/signal](https://pkg.go.dev/os/signal) — https://pkg.go.dev/os/signal | [pkg: context](https://pkg.go.dev/context) — https://pkg.go.dev/context

   > Graceful shutdown via signals + context cancellation is the canonical pattern for any long-running Go program (servers, daemons, watchers). Worth practicing until it's reflex.

6. **Run it** — point it at `/tmp/watchdir`, then in another terminal `touch /tmp/watchdir/a`, `echo hi >> /tmp/watchdir/a`, `rm /tmp/watchdir/a`. You should see `CREATE`, `MODIFY`, `DELETE` events.

7. **Bonus** — recurse into subdirectories using `filepath.WalkDir`. Bonus bonus: replace the poll with `github.com/fsnotify/fsnotify` and compare the latency.
