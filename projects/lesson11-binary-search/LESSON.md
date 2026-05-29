# Lesson 11 — Binary Search

## Spec

- Implement a generic `BinarySearch[T cmp.Ordered](xs []T, target T) (int, bool)`
  - Returns the index and `true` if found
  - Returns `0, false` if not found
- Slice must be assumed sorted ascending — document this in a comment
- Handle edges: empty slice, single element, target at first index, target at last index, target absent
- Write table-driven tests covering each edge case
- Write a benchmark (`testing.B`) comparing binary search vs linear scan on a 100k-element slice

---

1. **Pick the type constraint** — `cmp.Ordered` from the `cmp` package (Go 1.21+) lets the function accept any ordered type (`int`, `float64`, `string`, etc.) without writing your own constraint. Signature: `func BinarySearch[T cmp.Ordered](xs []T, target T) (int, bool)`. Docs: [pkg: cmp](https://pkg.go.dev/cmp) — https://pkg.go.dev/cmp | [Go generics intro](https://go.dev/doc/faq#generics) — https://go.dev/doc/faq#generics

   > `cmp.Ordered` is a beautiful piece of stdlib design — one constraint that covers every primitive type you'd want to compare. Before Go 1.21 you had to declare your own constraint with a big type union. Much cleaner now.

2. **Implement the classic loop** — two pointers `lo, hi := 0, len(xs)-1`. While `lo <= hi`: compute `mid := lo + (hi-lo)/2` (avoids integer overflow vs `(lo+hi)/2`). Compare `xs[mid]` to `target`. Narrow the window. Return when found or when `lo > hi`. Docs: [Go spec: Comparison operators](https://go.dev/ref/spec#Comparison_operators) — https://go.dev/ref/spec#Comparison_operators

   > The `lo + (hi-lo)/2` trick is the famous "binary search overflow bug" fix — `(lo+hi)/2` can overflow for huge arrays. This bug existed in Java's JDK for nine years. Classic CS folklore. Always write the safe version.

3. **Decide on iterative vs recursive** — both work in Go. Iterative is more idiomatic (Go has no tail-call optimization, so recursion costs real stack frames). Implement iterative. Mention the recursive version in a comment as an alternative.

4. **Write table-driven tests** — one `[]struct{ name string; xs []int; target int; wantIdx int; wantOk bool }` and a single `t.Run(tc.name, ...)` loop. Cases: empty, single-found, single-not-found, found-at-start, found-at-end, found-in-middle, not-present-below, not-present-above, not-present-in-gap. Docs: [Go blog: Table-driven tests](https://go.dev/wiki/TableDrivenTests) — https://go.dev/wiki/TableDrivenTests | [pkg: testing](https://pkg.go.dev/testing) — https://pkg.go.dev/testing

   > Table-driven tests are *the* Go testing idiom. Once you internalize it you'll find yourself reaching for it in every language. The combination of struct-of-cases + `t.Run(name)` for subtest naming is exceptionally clean.

5. **Write a benchmark** — `func BenchmarkBinarySearch(b *testing.B)` and `func BenchmarkLinearSearch(b *testing.B)`. Generate a sorted `[]int` of 100,000 elements once (outside the timer with `b.ResetTimer()` after setup). Search for a value near the end (worst case for linear). Run with `go test -bench=. -benchmem`. Docs: [pkg: testing.B](https://pkg.go.dev/testing#B) — https://pkg.go.dev/testing#B

6. **Compare results** — you should see binary search at log₂(100000) ≈ 17 ops vs linear at ~100000. Expect a roughly 1000–10000× speedup. Beautiful demonstration of why algorithmic complexity matters.

7. **Bonus** — compare against `slices.BinarySearch` from the stdlib (Go 1.21+). Read its source on pkg.go.dev to see how the pros write it.
