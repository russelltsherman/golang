# Golang Numbers

In Go language, numbers are divided into three sub-categories that are:

## Integers

Both signed and unsigned integers are available in four different sizes.
The signed int is represented by int and the unsigned integer is represented by uint.

| DATA TYPE | DESCRIPTION                                                      |
| --------- | ---------------------------------------------------------------- |
| int8      | 8-bit signed integer                                             |
| int16     | 16-bit signed integer                                            |
| int32     | 32-bit signed integer                                            |
| int64     | 64-bit signed integer                                            |
| uint8     | 8-bit unsigned integer                                           |
| uint16    | 16-bit unsigned integer                                          |
| uint32    | 32-bit unsigned integer                                          |
| uint64    | 64-bit unsigned integer                                          |
| int       | Both in and uint contain same size, either 32 or 64 bit.         |
| uint      | Both in and uint contain same size, either 32 or 64 bit.         |
| rune      | It is a synonym of int32 and also represent Unicode code points. |
| byte      | It is a synonym of int8 .                                        |
| uintptr   | It is an unsigned integer type. Its width is not defined,        |
|           | but its can hold all the bits of a pointer value.                |

## Floating-Point Numbers

Floating-point numbers are divided into two categories as shown in the below table:

| DATA TYPE | DESCRIPTION                           |
| --------- | ------------------------------------- |
| float32   | 32-bit IEEE 754 floating-point number |
| float64   | 64-bit IEEE 754 floating-point number |

## Complex Numbers

The complex numbers are divided into two parts are shown in the below table.
float32 and float64 are also part of these complex numbers.
The in-built function creates a complex number from its imaginary and real part and in-built imaginary and real function extract those parts.

| DATA TYPE  | DESCRIPTION                                                              |
| ---------- | ------------------------------------------------------------------------ |
| complex64  | Complex numbers which contain float32 as a real and imaginary component. |
| complex128 | Complex numbers which contain float64 as a real and imaginary component. |
