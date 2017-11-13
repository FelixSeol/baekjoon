package main

import "fmt"

func main() {
	var str string
	cups := []int{1, 0, 0, 2}
	fmt.Scanln(&str)
	f := func(a int, b int) {
		tmp := cups[a]
		cups[a] = cups[b]
		cups[b] = tmp
	}
	p := func() {
		var s, b int
		for i := range cups {
			if cups[i] == 1 {
				s = i + 1
			} else if cups[i] == 2 {
				b = i + 1
			} else {

			}
		}
		fmt.Printf("%d\n%d\n", s, b)
	}
	for i := range str {
		c := str[i]
		switch c {
		case 'A':
			f(0, 1)
		case 'B':
			f(0, 2)
		case 'C':
			f(0, 3)
		case 'D':
			f(1, 2)
		case 'E':
			f(1, 3)
		case 'F':
			f(2, 3)
		default:
			fmt.Println("no match character")
		}
	}
	p()
}
