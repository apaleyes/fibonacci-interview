// Tail recursion implementation, can also be seen as correct recursion implementation
// As it does not suffer from performance limitations of a naive recursive approach

package main

import "fmt"

func fib_tail_recursion(n int) (int, int) {
    if (n == 1) {
        return 0, 1
    }

    prev, current := fib_tail_recursion(n - 1)
    return current, prev + current
}

func main() {
    _, x := fib_tail_recursion(8)
    fmt.Printf("This should be 21: %d", x)
}