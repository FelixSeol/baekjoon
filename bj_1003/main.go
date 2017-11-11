package main

import (
	"fmt"
)
var cnt0, cnt1 int = 0, 0

func fibonacci(n int) int{
	if n==0{
		cnt0++
		return 0
	}else if n==1{
		cnt1++
		return 1
	}else{
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
func main(){
	var t, x int
	fmt.Scanf("%d",&t)
	for t > 0 {
		fmt.Scanf("%d", &x)
		fibonacci(x)
		fmt.Println(cnt0, cnt1)
		cnt0 =0 
		cnt1= 0
		t--
		}
		
	}
