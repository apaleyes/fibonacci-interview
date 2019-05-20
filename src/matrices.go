// This approach is based on the realization that the following formula holds
//
// [[0, 1], [1, 1]] * [F_n-1, F_n] = [F_n, F_n+1]
//
// To take this further, we have
//
// [[0, 1], [1, 1]] ^ n = [[F_n-1, F_n], [F_n, F_n+1]]
//
// This allows us to use fast exponentiation (see https://en.wikipedia.org/wiki/Exponentiation_by_squaring#Basic_method)
// Essentially problem comes down to quickly bringing a certain matrix to the given power

package main

import (
    "fmt"
    "strconv"
)

func multiply_matrices(m1 [2][2]int, m2 [2][2]int) [2][2]int {
    result := [2][2]int{{0, 0}, {0, 0}}
    result[0][0] = m1[0][0] * m2[0][0] + m1[0][1] * m2[1][0]
    result[0][1] = m1[0][0] * m2[0][1] + m1[0][1] * m2[1][1]
    result[1][0] = m1[1][0] * m2[0][0] + m1[1][1] * m2[1][0]
    result[1][1] = m1[1][0] * m2[0][1] + m1[1][1] * m2[1][1]

    return result
}

func fast_power(m [2][2]int, n int) [2][2]int {
    // start with identity matrix
    result := [2][2]int{{1, 0}, {0, 1}}
    
    bits := strconv.FormatInt(int64(n), 2)
    
    for _, bit := range bits {
        result = multiply_matrices(result, result)
        if (bit == '1') {
            result = multiply_matrices(result, m)
        }
    }
    
    return result
}

func fib_via_matrix_multiplication(n int) int {
    start := [2][2]int{{0, 1}, {1, 1}}
    end := fast_power(start, n)
    return end[1][0]
}

func main() {
    x := fib_via_matrix_multiplication(8)
    fmt.Printf("This should be 21: %d", x)
}
