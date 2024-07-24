# strings

The string data type represents a sequence of Unicode code points.

Or in other words, we can say a string is a sequence of immutable bytes, means once a string is created you cannot change that string.

A string may contain arbitrary data, including bytes with zero value in the human-readable form.

```go
package main
import "fmt"

func main() {
    // str variable which stores strings
   str := "GeeksforGeeks"

   // Display the length of the string
   fmt.Printf("Length of the string is:%d", len(str))

   // Display the string
   fmt.Printf("\nString is: %s", str)

   // Display the type of str variable
   fmt.Printf("\nType of str is: %T", str)
}
```
