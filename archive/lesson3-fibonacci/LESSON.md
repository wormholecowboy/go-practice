# Lesson 3 — Fibonacci

## Spec

- Implement Fibonacci **two ways**: iterative and recursive
- Accept a single CLI argument N
- Print the Nth Fibonacci number (0-indexed: F(0)=0, F(1)=1, F(2)=1, F(3)=2…)
- Print the result from both implementations and confirm they match
- If no argument or invalid argument is given, print an error and exit

---

1. **Write the recursive version first** — a function that calls itself with smaller inputs until it hits a base case (`N=0` or `N=1`). Docs: [Tour of Go: Functions](https://go.dev/tour/basics/4) — https://go.dev/tour/basics/4

   > This is one of the most classic CS problems ever. Recursion is the natural mathematical definition of Fibonacci — F(n) = F(n-1) + F(n-2). Elegant, but watch what happens with large N. You'll feel it.

2. **Write the iterative version** — use a `for` loop and two variables to track the last two values, sliding forward each iteration. Docs: [Tour of Go: For](https://go.dev/tour/flowcontrol/1) — https://go.dev/tour/flowcontrol/1

   > This is the classic space/time tradeoff in action. The iterative version runs in O(n) time vs the recursive version's O(2ⁿ). Try both with N=45 and feel the difference.

3. **Use multiple return values** — Go functions can return more than one value. Have each implementation return both the result and the number of steps/calls it took. This is idiomatic Go — no need for out-params or wrapper objects. Docs: [Tour of Go: Multiple return values](https://go.dev/tour/basics/6) — https://go.dev/tour/basics/6

   > Multiple return values are how Go handles errors everywhere (`result, err`). Getting comfortable with this pattern is essential.

4. **Parse CLI args** — same pattern as before: `os.Args`, `strconv.Atoi`, check `err`. Docs: [pkg: strconv](https://pkg.go.dev/strconv) — https://pkg.go.dev/strconv

5. **Print and compare** — use `fmt.Printf` to print both results clearly. Assert they match with a simple `if` — print a warning if they don't.
