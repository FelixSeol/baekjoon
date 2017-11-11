package main

import (
	"bufio"
	"fmt"
	"os"
)
func main(){
	reader := bufio.NewReader(os.Stdin)
	for x:=0; x<100; x++{
		str, _ := reader.ReadString('\n')
		fmt.Print(str)
	}
}
