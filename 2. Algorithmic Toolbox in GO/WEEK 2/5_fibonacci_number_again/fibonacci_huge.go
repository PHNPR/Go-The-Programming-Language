package main

import "fmt"

func main() {
	var n , m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Println(fibonacci_huge(n,m))
}

func fibonacci_huge(n int , m int) int {
	if n < 2 {
		return n
	}
	arr := make([]int , n+1)
	arr[0] , arr[1] = 0 , 1
	time_period := 1
	for i := 2 ; i < n+1 ; i++ {
		arr[i] = ( arr[i-1] + arr[i-2] ) % m
		if arr[i] == 1 && arr[i-1] == 0 {
			time_period = i - 1
			break
		} 
	}
	if time_period > 1 {
		n = n % time_period
	}
	return arr[n]
}