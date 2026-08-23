# SOLID Principles in Go

SOLID is a set of five object-oriented design principles that make code easier to
extend, test, and maintain. Go has no classes or implementation inheritance, but
the principles still apply directly to structs, interfaces, and composition —
Go's implicit interface satisfaction and preference for small interfaces make
several of them feel almost native to the language.

## [Single Responsibility Principle](./solid/single-responsibility.md) (SRP)

A type should have only one reason to change. Bundling unrelated concerns —
like an in-memory data structure and its file persistence — into a single type
means a change to either one risks breaking the other, and neither can be
tested in isolation.

## [Open/Closed Principle](./solid/open-closed.md) (OCP)

Types should be open for extension but closed for modification. Instead of
growing a function's branching logic every time a new case appears, model the
varying behavior as an interface, so new cases are added as new types rather
than edits to code that already works.

## [Liskov Substitution Principle](./solid/liskov-substitution.md) (LSP)

Anything that satisfies an interface must be substitutable for it without
surprising the caller. In Go this principle mainly guards against embedding a
type and overriding its behavior in a way that breaks assumptions callers make
about the embedded type's contract.

## [Interface Segregation Principle](./solid/interface-segregation.md) (ISP)

Clients shouldn't be forced to depend on methods they don't use. Prefer
several small, focused interfaces — the way `io.Reader` and `io.Writer` are
composed in the standard library — over one large interface that forces every
implementation to support everything.

## [Dependency Inversion Principle](./solid/dependency-inversion.md) (DIP)

High-level code shouldn't depend on low-level implementation details; both
should depend on an abstraction. In Go, that means a high-level type should
take an interface parameter rather than a concrete struct, so the
implementation can change — or be swapped for a test fake — without touching
the code that depends on it.

Each page shows a Go example that violates the principle, then the same code
reworked to follow it.

See also: [patterns](../patterns) — the Gang-of-Four design patterns, several of
which (Strategy, Decorator, Adapter) are direct applications of these principles.
