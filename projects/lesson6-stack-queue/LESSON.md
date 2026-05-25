# Lesson 6 — Stack & Queue

## Spec

- Implement a generic **Stack** (LIFO) with: `Push`, `Pop`, `Peek`, `IsEmpty`, `Size`
- Implement a generic **Queue** (FIFO) with: `Enqueue`, `Dequeue`, `Peek`, `IsEmpty`, `Size`
- `Pop`/`Dequeue` on an empty structure should return the zero value and `false` (ok-idiom)
- Write a small `main` that demonstrates both with a concrete type (e.g. `int` or `string`)

---

1. **Declare a generic Stack struct** — define a struct with a type parameter `T any` wrapping a `[]T` slice. This is the foundation of generics in Go: parameterizing a type so it works with any element type. Docs: [Go generics intro](https://go.dev/doc/faq#generics) — https://go.dev/doc/faq#generics | [Tour: Type parameters](https://go.dev/tour/generics/1) — https://go.dev/tour/generics/1

2. **Implement Stack methods** — `Push` appends to the slice, `Pop` removes and returns the last element, `Peek` returns the last element without removing, `IsEmpty` checks length, `Size` returns length. All methods are on a pointer receiver `*Stack[T]`. Docs: [Effective Go: Methods](https://go.dev/doc/effective_go#methods) — https://go.dev/doc/effective_go#methods

   > The ok-idiom (`value, ok := stack.Pop()`) is Go's idiomatic way to handle operations that may not produce a result — you'll see it everywhere (maps, type assertions, channel receives). Classic pattern.

3. **Declare a generic Queue struct** — same shape as Stack but FIFO. `Enqueue` appends to the back, `Dequeue` removes from the front (index 0). Docs: [Go spec: Type parameters](https://go.dev/ref/spec#Type_parameter_declarations) — https://go.dev/ref/spec#Type_parameter_declarations

4. **Implement Queue methods** — same set as Stack. For `Dequeue`, slice off the first element with `q.items = q.items[1:]`. Note the tradeoff: this is O(n) — a production queue would use a ring buffer or `container/list`. Docs: [pkg: container/list](https://pkg.go.dev/container/list) — https://pkg.go.dev/container/list

5. **Write main to demo both** — instantiate `Stack[int]` and `Queue[string]`, push/enqueue a few values, pop/dequeue them, print results. Confirm LIFO vs FIFO order visually.

6. **Write tests** — table-driven tests for each method covering the expected case, an edge case (empty structure), and the ok-idiom false return. Docs: [Go blog: Table-driven tests](https://go.dev/wiki/TableDrivenTests) — https://go.dev/wiki/TableDrivenTests
