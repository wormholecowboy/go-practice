# Lesson 1 — FizzBuzz CLI

1. **Initialize a module** — run `go mod init` in your project folder. Read [Go modules intro](https://go.dev/doc/code#Organization) — https://go.dev/doc/code#Organization to understand what this creates.

2. **Wire up `main`** — every executable Go program needs `package main` and a `func main()`. Nothing runs without it.

3. **Read a CLI argument** — use `os.Args` to grab the number the user passes in. It's a slice; index 0 is the program name, so your number is at index 1. Docs: [pkg: os](https://pkg.go.dev/os) — https://pkg.go.dev/os

4. **Convert the argument to an integer** — `os.Args` gives you a string. Use `strconv.Atoi` to convert it. It returns two values: the integer and an error. Docs: [pkg: strconv](https://pkg.go.dev/strconv) — https://pkg.go.dev/strconv

5. **Handle the error** — if the user passes something that isn't a number, your program should print a message and exit. Look at `os.Exit` and the `if err != nil` pattern.

6. **Write the loop** — loop from 1 to N. Go has one loop keyword: `for`. [Tour of Go: for](https://go.dev/tour/flowcontrol/1) — https://go.dev/tour/flowcontrol/1

7. **Add the conditionals** — check divisibility with `%`. FizzBuzz before Fizz and Buzz, or you'll get wrong output. [Tour of Go: if](https://go.dev/tour/flowcontrol/5) — https://go.dev/tour/flowcontrol/5

8. **Print with `fmt`** — use `fmt.Println` for each line of output. Docs: [pkg: fmt](https://pkg.go.dev/fmt) — https://pkg.go.dev/fmt
