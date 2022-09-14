package main

import "fmt"

func main() {
	var sum int
	fmt.Scan(&sum)
	len := len_calculator(sum)
	fmt.Println(len)

	z := sum - (len*(len-1))/2
	list := make([]int , len)
	for i := 0 ; i < len-1 ; i++ {
		list[i] = i+1
	}
	list[len-1] = z
	fmt.Println(list)
}

func len_calculator(sum int) int {
	for i := 0 ; i <= sum ; i++{
		if i*(i+1) <= sum * 2 && sum * 2 < (i+1)*(i+2) {
			return i
		}
	}
	return -1
}