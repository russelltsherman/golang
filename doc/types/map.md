# maps

Maps are used to store data values in key:value pairs.

Each element in a map is a key:value pair.

A map is an unordered and changeable collection that does not allow duplicates.

The length of a map is the number of its elements.
You can find it using the len() function.

The default value of a map is nil.

Maps hold references to an underlying hash table.

The zero value of a map is nil.
A nil map has no keys, nor can keys be added.

The make function returns a map of the given type, initialized and ready for use.

```go
package main

import "fmt"

type Vertex struct {
 Lat, Long float64
}

var m map[string]Vertex

func main() {
  m = make(map[string]Vertex)

  m["Bell Labs"] = Vertex{
    40.68433, -74.39967,
  }

  fmt.Println(m["Bell Labs"])
}
```
