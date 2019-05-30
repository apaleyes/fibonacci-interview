// This approach can be derived straight from the matrix-base approach (see matrices.go)
// It uses the following formulae:
//
// F[2k] = F[k] * (2 * F[k+1] - F[k])
// F[2k+1] = F[k]^2 + F[k+1]^2
//
// It has the same complexity as the matrix approach
// but less computations on each step

package main

import "fmt"

// For a given n, this method always returns F[n], F[n+1]
func fib_fast_doubling(n int) (int, int) {
    if (n == 0) {
        return 0, 1
    }
    
    if (n == 1) {
        return 1, 1
    }
    
    k := n / 2
    
    // We cannot use + in variable names
    // so k_1 below means k+1
    f_k, f_k_1 := fib_fast_doubling(k)
    
    f_2k := f_k * (2 * f_k_1 - f_k)
    f_2k_1 := f_k * f_k + f_k_1 * f_k_1
    
    if (n == 2 * k) {
        return f_2k, f_2k_1
    } else {
        return f_2k_1, f_2k_1 + f_2k
    }
}

func main() {
    x, _ := fib_fast_doubling(8)
    fmt.Printf("This should be 21: %d", x)
}
