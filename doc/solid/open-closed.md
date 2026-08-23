# Open/Closed Principle (OCP)

Types should be open for extension but closed for modification — adding new
behavior shouldn't require editing code that already works and is already
tested.

## The Problem

`FilterByColor` and any future `FilterBySize` each need their own function, and
every new filtering criterion means writing (and testing) another near-identical
function, or growing a single function into a chain of `if`/`switch` branches:

```go
type Product struct {
    Name  string
    Color string
    Size  string
}

func FilterByColor(products []Product, color string) []Product {
    var result []Product
    for _, p := range products {
        if p.Color == color {
            result = append(result, p)
        }
    }
    return result
}

// Adding FilterBySize means writing another near-identical function,
// and combining color AND size means writing yet another.
```

## Applying the Principle in Go

Model a filter criterion as an interface rather than a function per field. New
criteria are added by writing new types that satisfy `Specification` — no
existing code is touched:

```go
type Specification interface {
    IsSatisfied(p Product) bool
}

type ColorSpecification struct {
    Color string
}

func (c ColorSpecification) IsSatisfied(p Product) bool {
    return p.Color == c.Color
}

type SizeSpecification struct {
    Size string
}

func (s SizeSpecification) IsSatisfied(p Product) bool {
    return p.Size == s.Size
}

// AndSpecification composes specifications without modifying them.
type AndSpecification struct {
    specs []Specification
}

func (a AndSpecification) IsSatisfied(p Product) bool {
    for _, spec := range a.specs {
        if !spec.IsSatisfied(p) {
            return false
        }
    }
    return true
}

func Filter(products []Product, spec Specification) []Product {
    var result []Product
    for _, p := range products {
        if spec.IsSatisfied(p) {
            result = append(result, p)
        }
    }
    return result
}
```

`Filter` never changes again. A new criterion — say, filtering by a price
range — is a new `Specification` implementation, and it composes with the
existing ones through `AndSpecification` for free.

## Why It Matters in Go

- Because Go interfaces are satisfied implicitly, adding a new `Specification` requires no changes to `Filter`, `Product`, or any existing spec — true closed-for-modification.
- This is the same shape as the **Strategy** pattern (see [`../../patterns/strategy`](../../patterns/strategy)): behavior is selected by passing in an interface value instead of branching on a type or field.
- Composability (`AndSpecification`, `OrSpecification`) tends to fall out naturally once criteria are values implementing a common interface.

## Related Principles

- [Single Responsibility Principle](./single-responsibility.md) — each `Specification` implementation has exactly one reason to change: its own criterion.
- [Liskov Substitution Principle](./liskov-substitution.md) — every `Specification` must behave consistently with the interface's contract, or composition (`AndSpecification`) breaks.
