# Lesson 2 — Unit Converter

## Spec

- Accept two CLI arguments: a value (float) and a unit to convert from (e.g. `./converter 100 C`)
- Support at minimum:
  - Temperature: Celsius ↔ Fahrenheit
  - Distance: kilometers ↔ miles
- Print the converted result with the target unit label
- If an unknown unit is passed, print an error and exit
- If arguments are missing or malformed, print an error and exit

---

1. **Define your unit types with `iota`** — use a `const` block with `iota` to enumerate supported units (e.g. `Celsius`, `Fahrenheit`, `Kilometers`, `Miles`). This is Go's idiomatic alternative to string-based enums. Docs: [Go spec: Iota](https://go.dev/ref/spec#Iota) — https://go.dev/ref/spec#Iota

   > `iota` is how almost every language implements enums under the hood — an integer that auto-increments for each constant. Go just makes it visible and lets you do math on it. Classic.

2. **Define a struct for a measurement** — create a struct that holds a numeric value and a unit. This is your core data type. Docs: [Effective Go: Structs](https://go.dev/doc/effective_go#composite_literals) — https://go.dev/doc/effective_go#composite_literals

3. **Add a method to the struct** — attach a `Convert()` method to your struct that returns a new struct with the converted value and target unit. In Go, methods are defined outside the struct body using a receiver. Docs: [Tour of Go: Methods](https://go.dev/tour/methods/1) — https://go.dev/tour/methods/1

   > This is the foundation of object-oriented design without a class system. Go says: any type can have behavior. That's a big idea.

4. **Parse the CLI arguments** — use `os.Args` for the value and unit string. Convert the value to a `float64` with `strconv.ParseFloat`. Docs: [pkg: strconv](https://pkg.go.dev/strconv) — https://pkg.go.dev/strconv

5. **Map the unit string to your const** — write a function that takes the raw string arg (`"C"`, `"km"`, etc.) and returns the corresponding const value. Use a `switch` statement. Docs: [Tour of Go: Switch](https://go.dev/tour/flowcontrol/9) — https://go.dev/tour/flowcontrol/9

   > A `switch` as a dispatch table — mapping input to behavior — is one of the oldest patterns in systems programming. You'll see this everywhere from compilers to game engines.

6. **Handle unknown units** — if the string doesn't match any known unit, print an error and `os.Exit`. Docs: [pkg: os](https://pkg.go.dev/os) — https://pkg.go.dev/os

7. **Print the result** — use `fmt.Printf` to format the output with the numeric result and unit label. Docs: [pkg: fmt](https://pkg.go.dev/fmt) — https://pkg.go.dev/fmt
