package main

import "fmt"

func main() {
	var dist , m , n int
	fmt.Scan(&dist)
	fmt.Scan(&m)
	fmt.Scan(&n)
	stops := input(n)
	fmt.Println(Car_fueling(dist , m , n , stops))
}

func Car_fueling(dist int , m int , n int , stops []int) int {
	num_refills , curr_refill , limit := 0 , 0 , m
	for limit < dist {
		if curr_refill >= n || stops[curr_refill] > limit {
			return -1
		}
		for curr_refill < n - 1 && stops[curr_refill + 1] <= limit {
			curr_refill += 1
		}
		num_refills += 1
		limit = stops[curr_refill] + m
		curr_refill += 1
	}
	return num_refills
}

func input(n int) []int {
	stops := make([]int , n)
	for i := 0 ; i < n ; i++ {
		fmt.Scan(&stops[i])
	}
	return stops
}