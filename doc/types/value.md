# value types

golang is a pass by value language
when a value type value is passed to a function
a copy of that value is created
within the function
changes to a value type value affect the copy
not the original value

be mindful and use pointers when a function is meant to make changes to the original value

* int
* float
* string
* bool
* struct
