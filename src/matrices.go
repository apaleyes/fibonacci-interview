package main

import "fmt"

func multiply_matrices(m1 [2][2]int, m2 [2][2]int) [2][2]int {
    result := [2][2]int{{0, 0}, {0, 0}}
    result[0][0] = m1[0][0] * m2[0][0] + m1[0][1] * m2[1][0]
    result[0][1] = m1[0][0] * m2[0][1] + m1[0][1] * m2[1][1]
    result[1][0] = m1[1][0] * m2[0][0] + m1[1][1] * m2[1][0]
    result[1][1] = m1[1][0] * m2[0][1] + m1[1][1] * m2[1][1]

    return result
}

func fib_via_matrix_multiplication(n int) {

}

func main() {
    m1 := [2][2]int{{0,1},{1,1}}
    
    fmt.Println(multiply_matrices(m1, m1))
}

