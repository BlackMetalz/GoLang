package main

import (
    "fmt"
    "unsafe"
)

func main() {
    var a int8 = 1
    var b int16 = 32767
    var c int32 = 2147483647
    var d int64 = 9223372036854775807

    fmt.Printf("int8: %d, size: %d bytes\n", a, unsafe.Sizeof(a))
    fmt.Printf("int16: %d, size: %d bytes\n", b, unsafe.Sizeof(b))
    fmt.Printf("int32: %d, size: %d bytes\n", c, unsafe.Sizeof(c))
    fmt.Printf("int64: %d, size: %d bytes\n", d, unsafe.Sizeof(d))

    var a1 int8 = 1
    var b1 int8 = 10
    var c1 int8 = 100
    // No matter how small the value, the size of the variable is always 1 byte
    fmt.Printf("int8 (value: %d): size: %d bytes\n", a1, unsafe.Sizeof(a1))
    fmt.Printf("int8 (value: %d): size: %d bytes\n", b1, unsafe.Sizeof(b1))
    fmt.Printf("int8 (value: %d): size: %d bytes\n", c1, unsafe.Sizeof(c1))

}