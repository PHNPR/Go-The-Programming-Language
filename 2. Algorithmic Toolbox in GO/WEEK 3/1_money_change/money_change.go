package main

import "fmt"

func main() {
	var m int 
	fmt.Scan(&m)
	fmt.Println(change(m))
}

func change(m int) int {
	return m / 10 + ( m % 10 ) / 5 + m % 5
}