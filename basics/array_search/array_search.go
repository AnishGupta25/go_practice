package main

import "fmt"

func main(){
	arr := [...]int {1,2,3,4,5,6,7,8,9,10}
	for i:= 0; i < len(arr); i++{
		if arr[i] == 9{
		fmt.Printf("%d : %d\n", i, arr[i])
		}
	}
}