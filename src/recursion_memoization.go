// Memoization is a technique of keeping record of previously computed results
// This helps combat the poor performance of a naive recursive solution

package main

import "fmt"

func _fib_with_cache(n int, cache map[int]int) int {
    result, found := cache[n]
    if (found) {
        return result
    }

    result = _fib_with_cache(n - 1, cache) +  _fib_with_cache(n - 2, cache)

    cache[n] = result
    return result
}

func fib_memoization(n int) int {
    if (n <= 1) {
        return n
    }

    cache := make(map[int]int)
    cache[1] = 1
    cache[2] = 1

    return _fib_with_cache(n, cache)
}

func main() {
    x := fib_memoization(8)
    fmt.Printf("This should be 21: %d", x)
}