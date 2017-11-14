package main

import (
	"fmt"
	"math"
)

func printMat(mat [][][]int) {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			for k := 0; k < len(mat[0][0]); k++ {
				fmt.Printf("%d\t", mat[i][j][k])
			}
			fmt.Println()
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	var input int
	fmt.Scanln(&input)
	dim1 := input
	dim2 := func() int {
		return int(math.Log2(float64(input)))
	}()
	dim3 := func() int {
		i := 0
		for v := 1; v < input; i++ {
			v *= 3
		}
		return i
	}()
	fmt.Println(dim1, dim2, dim3, input)

	mat := make([][][]int, dim1)
	for i := 0; i < dim1; i++ {
		mat[i] = append(mat[i], make([][]int, dim2)...)
		for j := 0; j < dim2; j++ {
			mat[i][j] = append(mat[i][j], make([]int, dim3)...)
		}
	}
	fmt.Println(mat)
	printMat(mat)
}
