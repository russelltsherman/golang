# Golang Array

Arrays in Golang or Go programming language is much similar to other programming languages.
In the program, sometimes we need to store a collection of data of the same type, like a list of student marks.
Such type of collection is stored in a program using an Array.
An array is a fixed-length sequence that is used to store homogeneous elements in the memory.
Due to their fixed length array are not much popular like Slice in Go language.
In an array, you are allowed to store zero or more than zero elements in it.
The elements of the array are indexed by using the [] index operator with their zero-based position, means the index of the first element is array[0] and the index of the last element is array[len(array)-1].

Using var keyword: In Go language, an array is created using the var keyword of a particular type with name, size, and elements.
Syntax:

```go
var array_name[length]Typle{item1, item2, item3, ...itemN}
```

Important Points:
In Go language, arrays are mutable, so that you can use array[index] syntax to the left-hand side of the assignment to set the elements of the array at the given index.

```go
// Creating an array of string type
// Using var keyword
var myarr[3]string

// Elements are assigned using index
myarr[0] = "GFG"
myarr[1] = "GeeksforGeeks"
myarr[2] = "Geek"

// Accessing the elements of the array
// Using index value
fmt.Println("Elements of Array:")
fmt.Println("Element 1: ", myarr[0])
fmt.Println("Element 2: ", myarr[1])
fmt.Println("Element 3: ", myarr[2])
```

## Multi-Dimensional Array

As we already know that arrays are 1-D but you are allowed to create a multi-dimensional array. Multi-Dimensional arrays are the arrays of arrays of the same type. In Go language, you can create a multi-dimensional array using the following syntax:

```go
Array_name[Length1][Length2]..[LengthN]Type
```

You can create a multidimensional array using Var keyword or using shorthand declaration.

Note: In multi-dimension array, if a cell is not initialized with some value by the user, then it will initialize with zero by the compiler automatically. There is no uninitialized concept in the Golang.

## Important Observations About Array

In an array, if an array does not initialized explicitly, then the default value of this array is 0.

In an array, you can find the length of the array using len() method as shown below:

In an array, if ellipsis ‘‘…’’ become visible at the place of length, then the length of the array is determined by the initialized elements.

In an array, you are allowed to iterate over the range of the elements of the array.

In Go language, an array is of value type not of reference type.
So when the array is assigned to a new variable, then the changes made in the new variable do not affect the original array.

In an array, if the element type of the array is comparable, then the array type is also comparable.
So we can directly compare two arrays using == operator.
