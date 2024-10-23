### Dungeon_1 - Level 0
- package main is a special package that is used to create executable program.  It is the entrypoint.
- There is only one `package main` and only one `func main()`
- We use `import` to import package to use in our code that you can use the functions, types, and variables defined in those packages!

### Dungeon_1 - Level 1
- Static Typed: Compilier throws an error when types are used incorrectly (example: C++ and Java).
Example: you for function add (int a,int b). You can not call function like this add(1,"two") --> Error
- Dynamic typed: Compiler does not enforce the type system. Example: Python and Javascript

### Dungeon_1 - level 4:
Here is a table of common `fmt.Printf` format specifiers in Go:

| Verb  | Description |
|-------|-------------|
| `%v`  | Default format |
| `%T`  | Type of value |
| `%d`  | Decimal integer |
| `%b`  | Binary representation |
| `%o`  | Octal representation |
| `%x`  | Hexadecimal representation (lowercase) |
| `%X`  | Hexadecimal representation (uppercase) |
| `%f`  | Floating-point number |
| `%e`  | Scientific notation (lowercase) |
| `%E`  | Scientific notation (uppercase) |
| `%s`  | String |
| `%q`  | Quoted string |
| `%p`  | Pointer address |
| `%t`  | Boolean |

#### Example Usage
Here's an example demonstrating the use of these format specifiers:

```go
package main

import "fmt"

func main() {
    a := 42
    b := 3.14159
    c := "Hello, World!"
    d := true

    fmt.Printf("Default format: %v\n", a)
    fmt.Printf("Type of value: %T\n", a)
    fmt.Printf("Decimal integer: %d\n", a)
    fmt.Printf("Binary representation: %b\n", a)
    fmt.Printf("Octal representation: %o\n", a)
    fmt.Printf("Hexadecimal representation (lowercase): %x\n", a)
    fmt.Printf("Hexadecimal representation (uppercase): %X\n", a)
    fmt.Printf("Floating-point number: %f\n", b)
    fmt.Printf("Scientific notation (lowercase): %e\n", b)
    fmt.Printf("Scientific notation (uppercase): %E\n", b)
    fmt.Printf("String: %s\n", c)
    fmt.Printf("Quoted string: %q\n", c)
    fmt.Printf("Pointer address: %p\n", &a)
    fmt.Printf("Boolean: %t\n", d)
}
```

#### Output
```
Default format: 42
Type of value: int
Decimal integer: 42
Binary representation: 101010
Octal representation: 52
Hexadecimal representation (lowercase): 2a
Hexadecimal representation (uppercase): 2A
Floating-point number: 3.141590
Scientific notation (lowercase): 3.141590e+00
Scientific notation (uppercase): 3.141590E+00
String: Hello, World!
Quoted string: "Hello, World!"
Pointer address: 0xc0000140a0
Boolean: true
```

This example demonstrates how to use various format specifiers with the `fmt.Printf` function in Go.

### Level 7
#### Block
- Inner blocks can access variable declared within outer blocks
- Outer blocks can not access variables declared within inner blocks
#### Local variable
- Declared inside a function or a block
- not accessiable outside the function or the block
- Can be also declared inside looping and conditional statements
### Global variable
- Declared outside of a function or block
- available throughout the lifetime of a program
- declared at the top of the program outside all functions or blocks
- Can be accessed from any part of the program

### Level 8
#### Zero values
- For bool: false
- For int: 0
- For float: 0.0
- For string: ""
- For functions, pointers, interfaces,maps: nil