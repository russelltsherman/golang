# functions

A function is a block of statements that can be used repeatedly in a program.
Functions are not executed immediately.
They are "saved for later use", and will be executed when they are called.

```go
package main

import ("fmt")

func hello() {
  fmt.Println("Hello World")
}

func main() {
  hello() // call the function
}
```

## Parameters and Arguments

```go
func FunctionName(param1 type, param2 type, param3 type) {
  // code to be executed
}
```

## Function Returns

```go
func FunctionName(param1 type, param2 type) type {
  // code to be executed
  return output
}
```

## Recursion

```go
package main
import ("fmt")

func testcount(x int) int {
  if x == 11 {
    return 0
  }
  fmt.Println(x)
  return testcount(x + 1)
}

func main(){
  testcount(1)
}
```

## Receiver Function

add a function as a method of a type (or struct)

```go
package main

import "fmt"

type deck []string

func (d deck) print() {
  for i, card := range d {
    fmt.Println(i, card)
  }
}

func main() {
  cards := deck{"Ace of Diamonds", "Five of Diamonds", "Six of Spades"}
  cards.print()
}
```
