// scan
package main

import (
	"fmt"
)

func main() {
	var v, l, r byte
	re, _ := fmt.Scanf("%c %c %c", &v, &l, &r)
	fmt.Printf("%c %c %c %d\n", v, l, r, re)
}
