package main

import "fmt"
import "errors"

var n, s, d, f, b, k int
var land []string
var police []int

func main() {
	fmt.Scanf("%d%d%d%d%d%d\n", &n, &s, &d, &f, &b, &k)
	fmt.Printf("%d %d %d %d %d %d\n", n, s, d, f, b, k)
	tmp := make([]string, n)
	land = append(land, tmp...)
	for i := 0; i < k; i++ {
		var p int
		fmt.Scanf("%d", &p)
		police = append(police, p)
	}
	fmt.Println(police[0], police[1], police[2], police[3])
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
	r, err := play(s - 1)
	fmt.Println(r, err)
}

func play(_s int) (int, error) {
	var r int
	var err error

	if land[_s-1] == "K" || land[_s-1] == "C" || _s-1 < 0 || _s > n {
		err = errors.New("BUG REPORT")
	} else {
		if land[_s-1+f] == "D" {
			return r + 1, err
		}
		if land[_s-1+f] == "K" {
			r, err = play(_s - b)
		} else {
			r, err = play(_s + f)
		}
		r += 1
	}
	return r, err
}
