package main

import (
	"fmt"
)
func exchangeLowerUpper(s string) string{
	b := make([]byte, len(s))
	for i:= range b {
		c := s[i]
		if c>= 'A' && c<= 'Z'{
			c += 'a' - 'A'
		}else if c>= 'a' && c<='z'{
			c -= 'a' - 'A'
		}else{}
		b[i] = c
	}
	return string(b)
}

func main(){
	var p string
	fmt.Scanln(&p)
	fmt.Print(exchangeLowerUpper(p))
}
