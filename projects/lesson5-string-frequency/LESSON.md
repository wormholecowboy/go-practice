# Lesson 5 — String Frequency Counter

## Spec

- Accept a filename as a CLI argument
- Read the file and count how many times each word appears
- Print the top 10 most frequent words with their counts, sorted by frequency (descending)
- Normalize words: lowercase, strip punctuation
- If no filename given or file can't be opened, print an error and exit

---

1. **Read the file** — you've done this. `os.Open`, `defer f.Close()`, `bufio.Scanner`. No new concepts here — build the muscle memory. Docs: [pkg: bufio](https://pkg.go.dev/bufio) — https://pkg.go.dev/bufio

2. **Count words with a map** — declare `map[string]int` and increment the count for each word as you scan. Maps are Go's built-in hash table. Docs: [Tour of Go: Maps](https://go.dev/tour/moretypes/19) — https://go.dev/tour/moretypes/19

   > The map-as-counter pattern (`m[key]++`) is one of the most useful idioms in all of programming. In Go, accessing a missing key returns the zero value — so `m["unseen"]++` just works. Elegant.

3. **Normalize with `strings`** — use `strings.ToLower` and `strings.Trim` (or `strings.Map`) to strip punctuation before counting. Docs: [pkg: strings](https://pkg.go.dev/strings) — https://pkg.go.dev/strings

4. **Sort the results** — maps in Go have no order. To sort by frequency you need to convert to a slice of pairs, then sort. Use `sort.Slice` with a custom comparator. Docs: [pkg: sort](https://pkg.go.dev/sort) — https://pkg.go.dev/sort

   > Converting a map to a sorted slice is a classic pattern — you'll do this constantly. Maps give you O(1) lookup, slices give you order. Use both together.

5. **Print the top 10** — loop over the sorted slice and print the first 10 entries with `fmt.Printf`. Use width formatting to align columns cleanly.

6. **Bonus:** try running it against a large text file — grab one from [Project Gutenberg](https://www.gutenberg.org/) — https://www.gutenberg.org/ and see what the most common words are.
