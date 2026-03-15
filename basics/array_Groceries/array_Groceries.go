package main

import "fmt"

func main(){
	groceries_freshness := [...] int {15,67}
	groceries_value := [...]int{10,90}
	freshness_score := 20
	cost := 0
	for i := 0; i < len(groceries_freshness); i++{
		if groceries_freshness[i] >= freshness_score{
			fmt.Printf("Groceries %d is fresh and has value %d\n", i, groceries_value[i])
			cost += groceries_value[i]
		} else {
			fmt.Printf("Groceries %d is not fresh and has value %d\n", i, groceries_value[i])
		}
	}
	fmt.Printf("Total cost of fresh groceries: %d\n", cost)
}