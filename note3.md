### Arrays
- len(): the length of the array refers to the number of elements stored in the array


### Slices
- declaring and initializing a slice
```
slice := make([]<data_type>, length, capacity)
```

To copy a slice in Go, you don't need to have two slices with the same length or capacity as you would with arrays. Instead, you can use the built-in `copy` function, which copies elements from a source slice to a destination slice. The `copy` function will copy as many elements as will fit into the destination slice.

#### Example

Here's an example demonstrating how to copy a slice:

```go
package main

import "fmt"

func main() {
    // Source slice
    src := []int{10, 20, 30, 40, 50}

    // Destination slice with the same length as the source slice
    dest := make([]int, len(src))

    // Copy elements from src to dest
    copy(dest, src)

    fmt.Println("Source slice:", src) // Output: [10 20 30 40 50]
    fmt.Println("Destination slice:", dest) // Output: [10 20 30 40 50]
}
```

#### Key Points

- **Source Slice**: The slice from which elements are copied.
- **Destination Slice**: The slice to which elements are copied. It should be initialized with a sufficient length to hold the copied elements.
- **`copy` Function**: The built-in `copy` function is used to copy elements from the source slice to the destination slice.

#### Example with Different Lengths

If the destination slice has a different length, only the elements that fit will be copied:

```go
package main

import "fmt"

func main() {
    // Source slice
    src := []int{10, 20, 30, 40, 50}

    // Destination slice with a smaller length
    dest := make([]int, 3)

    // Copy elements from src to dest
    copy(dest, src)

    fmt.Println("Source slice:", src) // Output: [10 20 30 40 50]
    fmt.Println("Destination slice:", dest) // Output: [10 20 30]
}
```

#### Summary

- **Initialization**: Initialize the destination slice with a sufficient length.
- **Copying**: Use the `copy` function to copy elements from the source slice to the destination slice.
- **Partial Copy**: If the destination slice is smaller, only the elements that fit will be copied.

This approach allows you to copy slices efficiently without needing to manually iterate over the elements.