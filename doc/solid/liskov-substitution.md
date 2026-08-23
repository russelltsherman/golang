# Liskov Substitution Principle (LSP)

Anything that satisfies an interface must be substitutable for that interface
without surprising the caller — it must honor the contract the interface
implies, not just match its method signatures.

## The Problem

Go has no subclassing, so LSP violations usually show up through embedding
instead. `Square` embeds `Rectangle` and overrides `SetWidth`/`SetHeight` to
keep both sides equal — but code written against the general shape of
`Rectangle` now breaks when handed a `Square`:

```go
type Rectangle struct {
    Width, Height int
}

func (r *Rectangle) SetWidth(w int)  { r.Width = w }
func (r *Rectangle) SetHeight(h int) { r.Height = h }
func (r *Rectangle) Area() int       { return r.Width * r.Height }

type Square struct {
    Rectangle
}

func NewSquare(size int) *Square {
    return &Square{Rectangle{size, size}}
}

// Overriding breaks the assumption that setting Width doesn't affect Height.
func (s *Square) SetWidth(w int) {
    s.Width = w
    s.Height = w
}

func (s *Square) SetHeight(h int) {
    s.Width = h
    s.Height = h
}

func UseIt(r *Rectangle) {
    r.SetWidth(5)
    r.SetHeight(10)
    // Assumes Area() == 50. True for Rectangle, false for Square.
    fmt.Println(r.Area())
}
```

Because `UseIt` takes a `*Rectangle`, Go's static dispatch calls
`Rectangle.SetHeight` directly, silently ignoring `Square`'s override — the two
types are not actually substitutable, and the bug is exactly the kind LSP is
meant to catch.

## Applying the Principle in Go

Don't model `Square` as a specialization of `Rectangle` at all — model both as
independent implementations of a shared `Shape` interface, each responsible
only for its own, internally-consistent behavior:

```go
type Shape interface {
    Area() int
}

type Rectangle struct {
    Width, Height int
}

func (r Rectangle) Area() int { return r.Width * r.Height }

type Square struct {
    Size int
}

func (s Square) Area() int { return s.Size * s.Size }

func PrintArea(s Shape) {
    fmt.Println(s.Area())
}
```

Any `Shape` passed to `PrintArea` behaves exactly as the interface promises —
there's no hidden coupling between `Width` and `Height` for a caller to trip
over.

## Why It Matters in Go

- Go's implicit interfaces make it easy to satisfy a method signature without honoring the semantics the interface implies — the compiler won't catch an LSP violation, only tests and careful design will.
- Prefer composition over embedding when the "is-a" relationship doesn't hold for every method — embedding invites exactly the `Square`-via-`Rectangle` trap above.
- If a type's `New*` constructor needs to reject certain field combinations to keep behavior consistent, that's often a sign it shouldn't share an interface/embedding relationship with a more general type at all.

## Related Principles

- [Open/Closed Principle](./open-closed.md) — OCP's composition (e.g. `AndSpecification`) only stays safe if every composed implementation honors LSP.
- [Interface Segregation Principle](./interface-segregation.md) — a narrower interface is less likely to imply behavior a given implementation can't honor.
