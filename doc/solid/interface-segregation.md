# Interface Segregation Principle (ISP)

Clients shouldn't be forced to depend on methods they don't use — prefer several
small, focused interfaces over one large one.

## The Problem

A single `Machine` interface bundles printing, scanning, and faxing. A
simple, print-only device is now forced to implement methods it can't
meaningfully support:

```go
type Machine interface {
    Print(d Document)
    Scan(d Document)
    Fax(d Document)
}

type OldPrinter struct{}

func (o *OldPrinter) Print(d Document) { /* ... */ }

// These exist only to satisfy Machine — callers can invoke them and get
// a runtime failure instead of a compile-time guarantee they're unsupported.
func (o *OldPrinter) Scan(d Document) { panic("not supported") }
func (o *OldPrinter) Fax(d Document)  { panic("not supported") }
```

Any code that accepts a `Machine` now has to know, out of band, whether the
concrete value in hand actually supports `Scan` or `Fax` — the type system
can't tell it.

## Applying the Principle in Go

Split `Machine` into small, single-method interfaces. Types implement only
the ones they can genuinely support, and larger devices compose the small
interfaces back together:

```go
type Printer interface {
    Print(d Document)
}

type Scanner interface {
    Scan(d Document)
}

type Fax interface {
    Fax(d Document)
}

type OldPrinter struct{}

func (o *OldPrinter) Print(d Document) { /* ... */ }
// OldPrinter now only claims to be a Printer — nothing else.

// A capable device composes the small interfaces via embedding.
type MultiFunctionDevice interface {
    Printer
    Scanner
    Fax
}

type ModernMFP struct{}

func (m *ModernMFP) Print(d Document) { /* ... */ }
func (m *ModernMFP) Scan(d Document)  { /* ... */ }
func (m *ModernMFP) Fax(d Document)   { /* ... */ }
```

A function that only needs to print should accept a `Printer`, not a
`Machine` — it then works with both `*OldPrinter` and `*ModernMFP` without
either type over-promising what it supports.

## Why It Matters in Go

- This is idiomatic Go: the standard library's `io.Reader`, `io.Writer`, and `io.Closer` are exactly this pattern — small enough that types compose them (`io.ReadWriter`, `io.ReadCloser`) instead of being forced into one fat interface.
- Because interfaces are satisfied implicitly, a type can freely support `Printer` today and pick up `Scanner` later just by adding a method — no interface declaration needs to change.
- Accepting the narrowest interface a function actually needs (`Printer`, not `Machine`) also makes that function trivially easier to test with a minimal fake.

## Related Principles

- [Single Responsibility Principle](./single-responsibility.md) — small interfaces are the interface-level expression of the same "one reason to change" idea.
- [Dependency Inversion Principle](./dependency-inversion.md) — depending on the narrowest interface is a special case of depending on abstractions rather than concrete types.
