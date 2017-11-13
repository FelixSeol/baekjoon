package main

import "fmt"
import "errors"

var n, s, d, f, b, k int

func play(land []string) (int, error) {
	defer func() {
		rcv := recover()
		_ = rcv
	}()
	r := 0
	var err error
	if err != nil {
		return -1, errors.New("BUG FOUND")
	}
	if land[s-1] == "C" {
		r = -1
		err = errors.New("BUG FOUND")
	}
	land[s] = "C"
	if land[s-1+f] == "D" {
		land[s-1+f] = "S"
		return r + 1, err
	} else {
		if land[s-1+f] == "K" {
			s -= b
			r, err = play(land)
		} else {
			s += f
			r, err = play(land)
		}
		r++
	}
	return r, err
}

func main() {
	fmt.Scanf("%d%d%d%d%d%d\n", &n, &s, &d, &f, &b, &k)
	fmt.Printf("%d %d %d %d %d %d\n", n, s, d, f, b, k)
	land := make([]string, n)
	police := make([]int, 0)
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
	result, err := play(land)
	fmt.Println(result, err)
	showInGraphic()
}
