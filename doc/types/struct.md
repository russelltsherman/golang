# structs

A struct (short for structure) is used to create a collection of members of different data types, into a single variable.

While arrays are used to store multiple values of the same data type into a single variable, structs are used to store multiple values of different data types into a single variable.

A struct can be useful for grouping data together to create records.

```go
type struct_name struct {
  member1 datatype;
  member2 datatype;
  member3 datatype;
  ...
}
```

## Access Struct Members

the following example shocases a number of features
Struct composition
Struct received function
Different ways of creating and reading struct values
Pass by reference pointers

```go
import (
	"fmt"
)

type Address struct {
	street string
}

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
}

func (p *Person) setFirstName(name string) {
	(*p).firstName = name
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}

func main() {
	var homer Person
	homer.firstName = "Homer"
	homer.lastName = "Simpson"
	homer.age = 45
	fmt.Println("Name: ", homer.firstName, homer.lastName, "Age: ", homer.age)

	bart := Person{
		firstName: "Bart",
		lastName:  "Simpson",
		age:       13,
		address: Address{
			street: "732 Evergreen Terrace",
		},
	}
	bart.print()

	&bart.setFirstName("El Barto")
	fmt.Println()
  bart.print()
}
```
