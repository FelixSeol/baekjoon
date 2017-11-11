package main

import "fmt"

func main(){
	var total int
	arr := make([]int, 5)
	for i := range arr {
		fmt.Scanf("%d",&arr[i])
		if arr[i] < 40{
			arr[i] = 40
		}
	}

	for _, value:=range arr{
		total += value
	}
	fmt.Print(total/len(arr))
}
