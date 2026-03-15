package main
import "fmt"

func main(){
	value := [...]int{15 , 8,  22, 6}
	coupon_price := 30
	coupon_value := 10
	total_price := calculate_value(value[:])
	discount_price := calculate_discount(value[:], coupon_value)
	if total_price > discount_price + coupon_price{
		fmt.Println("Should buy the coupon!")
	} else {
		fmt.Println("Should not buy the coupon!")
	}
}

func calculate_value(value []int) int{
	total_price := 0
	for i := 0; i < len(value); i++{
		total_price += value[i]
	}
	return total_price
}

func calculate_discount(value []int, coupon_value int) int{
	discount_price := 0
	for i := 0; i < len(value); i++{
		if value[i] > coupon_value{
			discount_price = value[i] - coupon_value
		}
	}
	return discount_price
}