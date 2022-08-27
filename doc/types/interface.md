# interface

An interface type in Go is kind of like a definition.
It describes the methods that some other type must have.

interfaces facilitate
  code reuse by decoupling
  mocking in testing

One example of an interface type from the standard library is the fmt. Stringer interface, which looks like this:

```go
type Stringer interface {
    String() string
}
```

the following code satisfies or implements the stringer interface
it has a method with the matching signature `String() string`

```go
type Count int

func (c Count) String() string {
    return strconv.Itoa(int(c))
}
```
