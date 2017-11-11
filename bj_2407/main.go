package main

import "fmt"
import "math/big"

func combination(a int, b int){
	q := new(big.Int).MulRange(int64(a-b+1), int64(a))
	r := new(big.Int).MulRange(1, int64(b))
	fmt.Printf("%d",new(big.Int).Div(q,r))
}

func main(){
	var a, b int
	fmt.Scanf("%d %d",&a, &b)
	combination(a, b)
}

