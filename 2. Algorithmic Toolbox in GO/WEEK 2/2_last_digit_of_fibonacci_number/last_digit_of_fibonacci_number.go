package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(fib_last_digit(n))
}

func fib_last_digit(n int) int {
	if n < 2 {
		return 2
	}
	arr := make([]int , n+1)
	arr[0] = 0
	arr[1] = 1
	for i := 2 ; i < n + 1 ; i++ {
		arr[i] = (arr[i-1] + arr[i-2]) % 10
	} 
	return arr[n]
}