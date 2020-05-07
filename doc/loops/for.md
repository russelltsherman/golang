# Golang For loop

## A simple for loop

It is similar that we use in other programming languages like C, C++, Java, C#, etc.

```go
// This loop starts when i = 0
// executes till i<4 condition is true
// post statement is i++
for i := 0; i < 4; i++{
  fmt.Printf("GeeksforGeeks\n")
}
```

## For loop as Infinite Loop

A for loop is also used as an infinite loop by removing all the three expressions from the for loop.
When the user did not write condition statement in for loop it means the condition statement is true and the loop goes into an infinite loop.

```go
for {
// Statement...
}
```

## Break Statement

The break statement is used to terminate the loop or statement in which it presents.
After that, the control will pass to the statements that present after the break statement, if available.
If the break statement present in the nested loop, then it terminates only those loops which contains break statement.

Flow Chart:

```go
for i := 0; i < 5; i ++ {
  fmt.Println(i)
  // breaks when the value of i = 3
  if i == 3{
    break;
  }
}
```

## Goto Statement

This statement is used to transfer control to the labeled statement in the program.
The label is the valid identifier and placed just before the statement from where the control is transferred.
Generally, goto statement is not used by the programmers because it is difficult to trace the control flow of the program.

```go
var x int = 0
// for loop work as a while loop
Lable1: for x < 8 {
  if x == 5 {
    // using goto statement
    x = x + 1;
    goto Lable1
  }
  fmt.Printf("value is: %d\n", x);
  x++;
}
```

## Continue Statement

This statement is used to skip over the execution part of the loop on a certain condition.
After that, it transfers the control to the beginning of the loop.
Basically, it skips its following statements and continues with the next iteration of the loop.

```go
var x int = 0
// for loop work as a while loop
for x < 8 {
  if x == 5 {
    // skip two iterations
    x = x + 2;
    continue;
  }
  fmt.Printf("value is: %d\n", x);
  x++;
}
```
