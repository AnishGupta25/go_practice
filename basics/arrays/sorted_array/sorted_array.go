package main 

import "fmt"

func main(){
	arr := [...]int{2,3,4,5,6,7,8,9,1}

	for i := 0; i < len(arr) - 1; i++{
		if arr[i + 1] < arr[i] {
			fmt.Println("array is not sorted")
			return
		}
	}
	fmt.Println("array is sorted")
}