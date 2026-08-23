# Single Responsibility Principle (SRP)

A type should have only one reason to change — one job, one owner.

## The Problem

`Journal` below owns both the in-memory entries *and* how they get persisted to
disk. A change to the storage format (say, switching to JSON, or adding
encryption) now means editing the same type that owns the entry list, and any
other type that wants to persist something differently has to duplicate the
logic:

```go
type Journal struct {
    entries []string
}

func (j *Journal) AddEntry(text string) {
    j.entries = append(j.entries, text)
}

// Persistence is a second, unrelated responsibility bolted onto Journal.
func (j *Journal) Save(filename string) error {
    return os.WriteFile(filename, []byte(strings.Join(j.entries, "\n")), 0644)
}

func (j *Journal) Load(filename string) error {
    data, err := os.ReadFile(filename)
    if err != nil {
        return err
    }
    j.entries = strings.Split(string(data), "\n")
    return nil
}
```

## Applying the Principle in Go

Split the responsibilities into separate types. `Journal` only manages entries;
a `Persister` only knows how to write a string to disk. Neither has to change
because of the other:

```go
type Journal struct {
    entries []string
}

func (j *Journal) AddEntry(text string) {
    j.entries = append(j.entries, text)
}

func (j *Journal) String() string {
    return strings.Join(j.entries, "\n")
}

type Persister struct{}

func (p *Persister) Save(j *Journal, filename string) error {
    return os.WriteFile(filename, []byte(j.String()), 0644)
}
```

Now a `Persister` can save any type with a `String() string` method, and
`Journal` can be tested without ever touching the filesystem.

## Why It Matters in Go

- Small, single-purpose structs are easier to unit test — no filesystem, network, or database mocking required just to test entry management.
- Go's package system reinforces this: a `journal` package and a `persistence` package naturally stay decoupled if each type only does one thing.
- A struct with one responsibility tends to need one constructor and few fields, which keeps `New*` functions simple.

## Related Principles

- [Open/Closed Principle](./open-closed.md) — once responsibilities are separated, each one can be extended independently.
- [Dependency Inversion Principle](./dependency-inversion.md) — `Persister` could depend on a `Stringer` interface rather than `*Journal` directly, decoupling it further.
