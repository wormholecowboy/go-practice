# Lesson 12 — Singly Linked List

## Spec

- Implement a generic `List[T any]` (singly linked, head pointer only)
- Methods:
  - `PushFront(v T)` — O(1)
  - `PushBack(v T)` — O(n) (we only track head; mention this tradeoff)
  - `PopFront() (T, bool)` — O(1), ok-idiom
  - `Len() int` — O(n) (no cached length)
  - `ToSlice() []T` — O(n), useful for tests
  - `Iter(func(T) bool)` — iterate, halt early if the callback returns false
- All methods on a pointer receiver `*List[T]`
- Write table-driven tests for each method, including empty-list edge cases

---

1. **Declare the node and list types** — `node[T any]` has fields `value T` and `next *node[T]`. `List[T any]` has a single field `head *node[T]`. Lowercase `node` because it's an internal implementation detail — callers should never touch it. Docs: [Effective Go: Pointers vs Values](https://go.dev/doc/effective_go#pointers_vs_values) — https://go.dev/doc/effective_go#pointers_vs_values | [Tour: Pointers](https://go.dev/tour/moretypes/1) — https://go.dev/tour/moretypes/1

   > Linked lists are how generations of programmers first met pointers. The mental model — "a struct that points to another struct of the same type" — is foundational. Once linked lists click, trees and graphs follow naturally.

2. **Always use pointer receivers** — `func (l *List[T]) PushFront(v T)`. If you use a value receiver, mutations to `head` won't persist on the caller's list. This is one of the most common Go beginner gotchas. Make it muscle memory: anything that mutates → pointer receiver. Docs: [Go spec: Method sets](https://go.dev/ref/spec#Method_sets) — https://go.dev/ref/spec#Method_sets

3. **Implement PushFront** — create a new node whose `next` is the current head, then point head at the new node. Three-line method. The canonical "prepend" pattern.

4. **Implement PushBack** — walk from head until you find a node whose `next` is nil, then set its `next` to a new node. Handle the empty-list case (set head directly). The O(n) cost is the price of not tracking a tail pointer — call this out in a comment as the tradeoff.

5. **Implement PopFront** — if head is nil, return `zero, false`. Otherwise save head's value, advance `l.head = l.head.next`, return `value, true`. Classic ok-idiom. Docs: [Effective Go: Allocation with new](https://go.dev/doc/effective_go#allocation_new) — https://go.dev/doc/effective_go#allocation_new

   > The "save reference, then mutate" two-step is the universal shape of remove operations on linked structures. You'll see it again in trees and graphs.

6. **Implement Len, ToSlice, Iter** — all variations of "walk from head to nil with a `for n := l.head; n != nil; n = n.next` loop". This is the canonical traversal pattern. `Iter` takes a callback `func(T) bool` and stops if it returns false — Go 1.23+ idiomatically uses `iter.Seq[T]` here, but a callback is simpler and works everywhere. Docs: [pkg: container/list](https://pkg.go.dev/container/list) — https://pkg.go.dev/container/list

7. **Write table-driven tests** — cases per method: empty list, single element, multiple elements, pop-until-empty. Use `ToSlice` to assert state cleanly after each operation. Docs: [Go blog: Table-driven tests](https://go.dev/wiki/TableDrivenTests) — https://go.dev/wiki/TableDrivenTests

8. **Bonus** — make it doubly linked: add `prev *node[T]` and a `tail *node[T]` on the list. Now `PushBack` and `PopBack` both go O(1). Compare with `container/list` from the stdlib — that's a doubly linked list using `interface{}` (pre-generics).
