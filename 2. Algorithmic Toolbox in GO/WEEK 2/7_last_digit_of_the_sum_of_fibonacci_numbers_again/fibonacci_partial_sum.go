package main

import "fmt"

func main() {
	var a , b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	result := fibonacci_sum(b) - fibonacci_sum(a-1)
	if result < 0 {
		result += 10
	}
	fmt.Println(result)
}

func fibonacci_sum(n int) int {
	if n < 2 {
		return n
	}
	arr := make([]int , n+1)
	arr[0] , arr[1] = 0 , 1
	time_period := 1
	for i := 2 ; i < n+1 ; i++ {
		arr[i] = ( arr[i-1] + arr[i-2] ) % 10
		if arr[i] == 1 && arr[i-1] == 0 {
			time_period = i - 1
			break
		} 
	}
	if time_period > 1 {
		n = n % time_period
	}
	result := 0
	for i:= 0 ; i < n + 1 ; i++ {
		result += arr[i]
	}
	return result % 10
}