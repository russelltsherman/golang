# maps

Maps are used to store data values in key:value pairs.

Each element in a map is a key:value pair.

A map is an unordered and changeable collection that does not allow duplicates.

in a map all keys must be the same type
in a map all values must be the same type

The length of a map is the number of its elements.
You can find it using the len() function.

The default value of a map is nil.

Maps hold references to an underlying hash table.

The zero value of a map is nil.
A nil map has no keys, nor can keys be added.

## create

in the following example three different methods are shows for creating maps

```go
package main

import "fmt"

func main() {
  // create a new empty string indexed map of strings
  var colormap1 map[string]string
  colormap1["black"] = "#000000"

  // functionally equivalent to above
  colormap2 := make(map[string]string)
  colormap2["black"] = "#000000"

	colormap3 := map[string]string{
		"black":   "#000000",
	}
}
```

## iterate

```go
package main

import "fmt"

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex for", color, "is", hex)
	}
}

func main() {
	colors := map[string]string{
		"black": "#000000",
		"white": "#ffffff",
		"red":   "#ff0000",
		"blue":  "#00ff00",
		"green": "#0000ff",
	}
	printMap(colors)
}
```
