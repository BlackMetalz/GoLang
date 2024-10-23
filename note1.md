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

### Level 9
In Go, the `&` operator is used to get the memory address of a variable. When you use [`fmt.Scanf`] to read input from the user, you need to provide a pointer to the variable where the input should be stored. This is because [`fmt.Scanf`] needs to modify the value of the variable, and it can only do that if it has access to the variable's memory address.

#### Explanation

- **`&name`**: The `&` operator gets the memory address of the variable [`name`]
- **Pointer**: [`fmt.Scanf`] requires a pointer to the variable so it can store the input value at that memory address.

#### Example

Here's a simple example to illustrate this:

```go
package main

import (
    "fmt"
)

func main() {
    var name string
    fmt.Print("Enter your name: ")
    fmt.Scanf("%s", &name) // &name is a pointer to the variable name
    fmt.Println("Hello,", name)
}
```

#### Why a Pointer is Needed

- **Direct Modification**: By passing a pointer,  can directly modify the value of [`name`]
- **Memory Address**: The pointer provides the memory address where the input value should be stored.

Without the `&` operator, [`fmt.Scanf`] would not be able to modify the value of [`name`], and the input would not be stored in the variable.


### Level 11
In Go, the _ (underscore) is known as the blank identifier. It is used to ignore values that you do not need to use. When you call a function that returns multiple values, you can use _ to ignore one or more of those values.

Explanation
In the line:
```
num, _ = strconv.Atoi(str)
```

strconv.Atoi(str) is a function that converts a string to an integer. It returns two values:

The converted integer (num in this case).
An error value (_ in this case).
num is used to store the converted integer.

_ is used to ignore the error value returned by strconv.Atoi.


```
package main

import (
    "fmt"
    "strconv"
)

func main() {
    str := "123"
    num, _ := strconv.Atoi(str) // Ignore the error value
    fmt.Println("Converted number:", num)
}
```

Why Use _?
- Ignoring Unused Values: If you are confident that the conversion will succeed and you do not need to handle the error, you can use _ to ignore the error value.
- Simplifying Code: It simplifies the code by not requiring you to declare a variable for a value you do not intend to use.

Note
- While using _ can simplify the code, it is generally good practice to handle errors to make your code more robust and reliable. In production code, you might want to check the error value to handle any potential issues:

```
num, err := strconv.Atoi(str)
if err != nil {
    fmt.Println("Error converting string to integer:", err)
    return
}
fmt.Println("Converted number:", num)
```


### Level 12 - Constants
- Definition: Constants are immutable values that are known at compile time and do not change during the execution of the program.
- Declaration: Use the const keyword to declare a constant. We need to define constants and it's value at the same time!
- Example for ummitable: like Pi is never changes, so declare Pi value by constant
- Syntax:
```
const <const name> <data type> = <value>
```
#### Untyped constant
- constants are untyped unless they are explicitly given a type at declaration
- allow for flexibility
```
const age = 33
const name, age = "kienlt", 33
```

#### Typed constant
- constants are typed when you explicitly specify the type in the declaration 
- no flexibility
```
const name string = "kien luong trung"
const age int = 33
```
