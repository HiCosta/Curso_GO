package main

import "fmt"

var f, b int = 1, 0

func main(){
	for i:=1; i<10; i++{
		f+=b
		b+=f
		fmt.Println(f,"\n", b)
	}
}
