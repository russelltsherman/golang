# channels

a channel:
  is a reference type
  is a typed conduit for sending and receiving values between goroutines
  is the primary way goroutines communicate and synchronize without shared memory

Channels are created with `make` and referenced by the `chan` keyword.
Sending and receiving both use the `<-` operator, with the arrow pointing in
the direction data flows.

```go
ch := make(chan int)

go func() {
  ch <- 42 // send
}()

value := <-ch // receive
fmt.Println(value)
```

By default a channel is **unbuffered**: a send blocks until another goroutine
is ready to receive, and a receive blocks until another goroutine sends. This
is what makes an unbuffered channel act as a synchronization point, not just a
queue.

## Buffered channels

A buffered channel has capacity — sends only block once the buffer is full,
and receives only block once it's empty.

```go
ch := make(chan int, 2) // capacity 2

ch <- 1 // does not block, buffer has room
ch <- 2 // does not block, buffer now full
// ch <- 3 would block until something receives

fmt.Println(<-ch, <-ch) // 1 2
```

## Directional channels

A channel parameter can be restricted to send-only or receive-only, which the
compiler enforces. This documents intent and prevents misuse:

```go
func produce(out chan<- int) {
  out <- 1
}

func consume(in <-chan int) {
  fmt.Println(<-in)
}
```

## Closing a channel

`close(ch)` signals that no more values will be sent. Receiving from a closed
channel returns the zero value immediately, and a second return value reports
whether the channel is still open:

```go
ch := make(chan int, 2)
ch <- 1
ch <- 2
close(ch)

for v := range ch {
  fmt.Println(v) // 1, 2, then the range loop exits
}

v, ok := <-ch
fmt.Println(v, ok) // 0 false
```

Only the sender should close a channel, never the receiver — sending on a
closed channel panics, and closing an already-closed channel panics.

## select

`select` waits on multiple channel operations at once and proceeds with
whichever is ready first, similar to a `switch` for channels:

```go
select {
case v := <-ch1:
  fmt.Println("got", v, "from ch1")
case ch2 <- 10:
  fmt.Println("sent 10 to ch2")
default:
  fmt.Println("neither ready, did not block")
}
```

Without a `default` case, `select` blocks until one of its cases is ready.

## Notes

- A `nil` channel (a `chan` declared with `var` but never `make`'d) blocks forever on both send and receive — this is occasionally used deliberately in a `select` to disable a case.
- Sending or receiving on a channel with no other goroutine able to complete the other side of the operation is a common cause of an "all goroutines are asleep - deadlock!" runtime error.
- Channels are most often paired with `go` (goroutines) and the `sync` package's `WaitGroup`/`Mutex` for coordinating concurrent work.
