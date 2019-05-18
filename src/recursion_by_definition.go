// Recursive implementation that strictly follows definition of the Fibonacci number
// It has a horrible complexity of O(e^n)
// And therefore should never be used in practice

package main

import "fmt"

func fib_by_definition(n int) int {
    if (n <= 1) {
        return n
    } else {
        return fib_by_definition(n - 1) + fib_by_definition(n - 2)
    }
}

func main() {
    x := fib_by_definition(8)
    fmt.Printf("This should be 21: %d", x)
}