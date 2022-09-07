package main

import "fmt"

func main() {
	var a , b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println((a*b)/gcd(a,b))
}

func gcd(a int , b int) int {
	if b == 0 {
		return a
	} 
	return gcd(b , a % b)
}