package main

import "fmt"

func main() {
	var tc, sum, man, cnt int
	var avg, answer float64
	var scores []int
	fmt.Scan(&tc)
	for i := 0; i < tc; i++ {
		fmt.Scan(&man)
		scores = make([]int, man)
		for j := 0; j < man; j++ {
			fmt.Scan(&scores[j])
			sum += scores[j]
		}
		avg = float64(sum) / float64(man)
		for _, k := range scores {
			if float64(k) > avg {
				cnt++
			}
		}
		answer = (float64(cnt) / float64(man)) * 100
		fmt.Printf("%.3f%%\n", answer)
		cnt = 0
		sum = 0
	}

}
