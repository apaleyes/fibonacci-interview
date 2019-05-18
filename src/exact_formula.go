// Uses exact formula that can be derived from the recursive definition
// Fails for very big inputs due to floating point numbers precision loss

package main

import (
    "fmt"
    "math"
)


func fib_exact_formula(n int) int {
    sqrt_5 := math.Sqrt(5)

    a := (1.0 + sqrt_5) / 2.0
    b := (1.0 - sqrt_5) / 2.0
    res_float := (math.Pow(a, float64(n)) - math.Pow(b, float64(n))) / sqrt_5

    return int(res_float)
}

func main() {
    x := fib_exact_formula(8)
    fmt.Printf("This should be 21: %d", x)
}