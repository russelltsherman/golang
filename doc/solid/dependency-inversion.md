# Dependency Inversion Principle (DIP)

High-level code shouldn't depend on low-level implementation details — both
should depend on an abstraction. In Go, that means depending on an interface
rather than a concrete struct.

## The Problem

`Research` (high-level policy: "find all of Alice's children") depends
directly on `Relationships` (low-level detail: how relationships happen to be
stored). If storage ever moves from a slice to a database, `Research` has to
change too, even though its actual logic hasn't:

```go
type Relationship int

const (
    Parent Relationship = iota
    Child
    Sibling
)

type Info struct {
    From, To     string
    Relationship Relationship
}

// Low-level module: owns storage.
type Relationships struct {
    relations []Info
}

func (r *Relationships) AddParentAndChild(parent, child string) {
    r.relations = append(r.relations,
        Info{parent, child, Parent},
        Info{child, parent, Child})
}

// High-level module: depends directly on the concrete Relationships type,
// including its internal []Info slice.
type Research struct {
    relationships *Relationships
}

func (r *Research) FindAllChildrenOf(name string) []string {
    var result []string
    for _, rel := range r.relationships.relations {
        if rel.From == name && rel.Relationship == Parent {
            result = append(result, rel.To)
        }
    }
    return result
}
```

## Applying the Principle in Go

Introduce an interface that describes what `Research` needs — not how
`Relationships` happens to store data — and depend on that instead:

```go
type RelationshipBrowser interface {
    FindAllChildrenOf(name string) []string
}

type Relationships struct {
    relations []Info
}

func (r *Relationships) AddParentAndChild(parent, child string) {
    r.relations = append(r.relations,
        Info{parent, child, Parent},
        Info{child, parent, Child})
}

func (r *Relationships) FindAllChildrenOf(name string) []string {
    var result []string
    for _, rel := range r.relations {
        if rel.From == name && rel.Relationship == Parent {
            result = append(result, rel.To)
        }
    }
    return result
}

// High-level module now depends only on the abstraction.
type Research struct {
    browser RelationshipBrowser
}

func (r *Research) Investigate(name string) []string {
    return r.browser.FindAllChildrenOf(name)
}
```

`Research` no longer knows or cares whether relationships live in a slice, a
map, or a SQL table — and a test can hand it a fake `RelationshipBrowser`
without touching real storage at all.

## Why It Matters in Go

- Go's implicit interface satisfaction means `Relationships` doesn't need to declare it implements `RelationshipBrowser` — the interface can even be defined in the `Research` package itself, right next to the code that consumes it.
- This is the standard shape for testability in Go: high-level code takes an interface parameter, production code passes the real implementation, tests pass a hand-written fake or a mock.
- "Depend on abstractions" doesn't mean every type needs an interface — only introduce one where a real seam (a different implementation, or a test double) is actually needed; an interface with a single implementation and no tests depending on it is often unnecessary indirection.

## Related Principles

- [Open/Closed Principle](./open-closed.md) — depending on an interface is what lets a new implementation be added without modifying the high-level module that consumes it.
- [Interface Segregation Principle](./interface-segregation.md) — `RelationshipBrowser` should stay as narrow as `Research` actually needs, not grow into a general-purpose `Relationships` interface.
