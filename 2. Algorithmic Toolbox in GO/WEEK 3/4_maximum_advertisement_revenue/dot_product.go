package main

import (
	"fmt"
	"sort"
)

func main() {
	var n , value int 
	fmt.Scan(&n)
	prices , clicks := make([]int , n) , make([]int , n)
	for i := 0 ; i < n ; i++ {
		fmt.Scan(&prices[i])
	}
	for i := 0 ; i < n ; i++ {
		fmt.Scan(&clicks[i])
	}
	sort.Slice(prices , func(i, j int) bool {return prices[i] < prices[j]})
	sort.Slice(clicks , func(i, j int) bool {return clicks[i] < clicks[j]})

	for i := 0 ; i < n ; i++ {
		value += prices[i] * clicks[i]
	}
	fmt.Println(value)
}