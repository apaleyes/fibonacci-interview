// Naive iterative approach, strctly following the definition

package main

import "fmt"

func fib_by_definition(n int) int {
    if (n <= 1) {
        return n
    }

    prev := 1
    next := 1
    i := 2

    for i < n {
        prev, next = next, prev + next
        i += 1
    }

    return next
}

func main() {
    x := fib_by_definition(8)
    fmt.Printf("This should be 21: %d", x)
}