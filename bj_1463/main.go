package main

import "fmt"


func main() {
	var input int
	fmt.Scanln(&input)
	cnt := make([]int, input+1)

	min := func(a int, b int, c int) int {
		if a < b {
			if a < c {
				return a
			} else {
				return c
			}
		} else {
			if b < c {
				return b
			} else {
				return c
			}
		}
	}

	for p := 2; p < input+1; p++ {
		if p%2 == 0 && p%3 == 0 {
			cnt[p] = min(cnt[p-1]+1, cnt[p/2]+1, cnt[p/3]+1)
		} else if p%2 == 0 {
			cnt[p] = min(cnt[p-1]+1, cnt[p/2]+1, 1000001)
		} else if p%3 == 0 {
			cnt[p] = min(cnt[p-1]+1, 1000001, cnt[p/3]+1)
		} else {
			cnt[p] = cnt[p-1] + 1
		}
	}
	fmt.Printf("%d",cnt[input])

}
