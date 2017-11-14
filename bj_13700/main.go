package main

import "fmt"
import "math"

var n, s, d, f, b, k int
var land []string
var police []int
var cnt []int

func main() {
	fmt.Scanf("%d%d%d%d%d%d\n", &n, &s, &d, &f, &b, &k)
	fmt.Printf("%d %d %d %d %d %d\n", n, s, d, f, b, k)
	land = append(land, make([]string, n)...)
	cnt = append(cnt, make([]int, n)...)
	for i := range cnt {
		cnt[i] = math.MaxInt32
	}
	cnt[s-1] = 0
	fmt.Println(cnt)
	for i := 0; i < k; i++ {
		var p int
		fmt.Scanf("%d", &p)
		police = append(police, p)
	}

	setData := func() {
		for i := range land {
			land[i] = "O"
		}
		land[d-1] = "D"
		land[s-1] = "S"
		for i := range police {
			land[police[i]-1] = "K"
		}
	}
	showInGraphic := func() {
		for i := range land {
			fmt.Printf("%d\t", i+1)
		}
		fmt.Println()
		for i := range land {
			fmt.Printf("%s\t", land[i])
		}
		fmt.Println()
	}

	setData()
	showInGraphic()
	result, _ := play(d - 1)
	fmt.Println(result)
}

func play(pos int) (_cnt int, _err error) {

	min := func(a int, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	if cnt[pos] != math.MaxInt32 {
		return cnt[pos], nil
	}
	var fromfront, fromback int
	if pos-f < 0 || land[pos-f] == "K" {
		fromback = n + 1
	} else {
		fromback, _ = play(pos - f)
	}

	if pos+b > n || land[pos+b] == "K" {
		fromfront = n + 1

	} else {
		fromfront, _ = play(pos + b)
	}

	cnt[pos] = min(fromback, fromfront)
	return _cnt + 1, nil
}
