# Lesson 4 — `wc` Clone

## Spec

- Accept an optional filename as a CLI argument
- If a filename is given, read from that file
- If no filename is given, read from stdin (so it works with pipes: `cat file.txt | ./wc`)
- Count and print: number of lines, words, and bytes
- Output format: `lines words bytes [filename]`
- If the file doesn't exist, print an error and exit

---

1. **Read from a file using `os.Open`** — open the file, get back a file handle. Always `defer` closing it immediately after opening. Docs: [pkg: os](https://pkg.go.dev/os) — https://pkg.go.dev/os

   > `defer` is a Go superpower — it guarantees cleanup runs when the function exits, no matter what. This pattern (open → defer close → use) is everywhere in Go and systems programming generally.

2. **Fall back to stdin** — `os.Stdin` is just another `io.Reader`. If no filename arg is given, use it instead of a file. Your counting logic shouldn't need to know the difference. Docs: [pkg: os#Stdin](https://pkg.go.dev/os#Stdin) — https://pkg.go.dev/os#Stdin

   > This is the Unix philosophy: everything is a file. stdin, stdout, network sockets — all the same interface. Go's `io.Reader` captures this perfectly.

3. **Wrap the reader in a `bufio.Scanner`** — scanning line by line with raw `Read` is painful. `bufio.Scanner` handles it cleanly. Use it to iterate lines, count them, and split each line into words. Docs: [pkg: bufio](https://pkg.go.dev/bufio) — https://pkg.go.dev/bufio

4. **Count words per line** — use `strings.Fields` to split a line into words. It handles multiple spaces and tabs cleanly. Docs: [pkg: strings#Fields](https://pkg.go.dev/strings#Fields) — https://pkg.go.dev/strings#Fields

5. **Count bytes** — `len(string)` in Go returns bytes, not characters. For ASCII text this is fine. Add the byte count of each line as you scan (remember to add 1 for the newline character). Docs: [Go blog: Strings](https://go.dev/blog/strings) — https://go.dev/blog/strings

   > This byte vs character distinction is a gotcha that trips up everyone. Go strings are raw bytes — `len("é")` is 2, not 1. Worth knowing early.

6. **Print the result** — match real `wc` output format: right-aligned columns with the filename if one was given. Use `fmt.Printf` with width formatting e.g. `%8d`. Docs: [pkg: fmt](https://pkg.go.dev/fmt) — https://pkg.go.dev/fmt
