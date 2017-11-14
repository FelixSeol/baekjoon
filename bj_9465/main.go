// main.go
package main

import (
	"fmt"
	"math"
)

func main() {
	var t int
	fmt.Scanln(&t)
	mat := make([][]int, 2)
	penalty := make([][]int, 2)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scanln(&n)
		mat[0] = append(mat[0], make([]int, n)...)
		mat[1] = append(mat[1], make([]int, n)...)
		penalty[0] = append(penalty[0], make([]int, n)...)
		penalty[1] = append(penalty[1], make([]int, n)...)
		for j := 0; j < 2; j++ {
			for k := 0; k < n; k++ {
				fmt.Scan(&mat[j][k])
			}
			fmt.Scanln()
		}
		fmt.Print(mat[0], mat[1])
		fmt.Println()
		fmt.Println()
		getPenal := func(r int, c int) int {
			var p int
			if r == 0 {
				p += mat[1][c]
				if n == 1 {
				} else if c == 0 {
					p += mat[0][1]
				} else if c == n-1 {
					p += mat[0][c-1]
				} else {
					p += mat[0][c+1]
					p += mat[0][c-1]
				}
			} else {
				p += mat[0][c]
				if n == 1 {

				} else if c == 0 {
					p += mat[1][1]
				} else if c == n-1 {
					p += mat[1][c-1]
				} else {
					p += mat[1][c+1]
					p += mat[1][c-1]
				}
			}
			return p
		}
		fmt.Println(getPenal(0, 0), getPenal(0, n-1), getPenal(1, 0), getPenal(1, n-1))

		sum := func() (r int) {
			for _, v := range mat[0] {
				r += v
			}
			for _, v := range mat[1] {
				r += v
			}
			return
		}
		/*dmg := func(r int, c int) int {
			if r == 0 {
				mat[1][c] = -int(math.Abs(float64(mat[1][c])))
				if c == 0 {
					mat[0][1] = -int(math.Abs(float64(mat[0][1])))
				} else if c == n-1 {
					mat[0][c-1] = -int(math.Abs(float64(mat[0][c-1])))
				} else {
					mat[0][c+1] = -int(math.Abs(float64(mat[0][c+1])))
					mat[0][c-1] = -int(math.Abs(float64(mat[0][c-1])))
				}
			} else {
				mat[0][c] = -int(math.Abs(float64(mat[0][c])))
				if c == 0 {
					mat[1][1] = -int(math.Abs(float64(mat[1][1])))
				} else if c == n-1 {
					mat[1][c-1] = -int(math.Abs(float64(mat[1][c-1])))
				} else {
					mat[1][c+1] = -int(math.Abs(float64(mat[1][c+1])))
					mat[1][c-1] = -int(math.Abs(float64(mat[1][c-1])))
				}
			}
			return sum()
		}*/
		min := func(a int, b int) (m int) {
			if a < b {
				m = a
			} else {
				m = b
			}
			return
		}
		_ = min
		score := sum()
		fmt.Println("score", score)
		
		for i:=0;i<2;i++{
			for j:=0;j < n-1; j++{
				penalty[i][j] = getPenal(i, j)
			}
		}
		
		result := make([]int , n)
		
		for i:=-;i<n-1;i++{
			result[i] = min(min(getPenal(0,i),getPenal(1,i)),min(getPenal(0,i+1),getPenal(1,i+1)))
		}
		
		
		fmt.Println(penalty[0])
		fmt.Println(penalty[1])

		mat[0] = nil
		mat[1] = nil
	}
}
