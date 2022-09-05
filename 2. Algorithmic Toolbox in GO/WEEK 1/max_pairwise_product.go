package main

import "fmt"

func main() {
	var len int
	fmt.Scan(&len)
	arr := make([]int , len)
	for i := 0 ; i < len ; i++ {
		fmt.Scan(&arr[i])
	} 
	fmt.Println(max_pairwise_product(arr))
}

func max_pairwise_product(a []int) int {
	var m1 , m2 int
    if a[0] > a[1] {
        m1 , m2 = a[0] , a[1]
	} else {
        m1 , m2 = a[1] , a[0]
	}
    for i := 2 ; i < len(a) ; i++ {
        if a[i] >= m1 {
            m1 , m2 = a[i] , m1
		} else if m1 > a[i] && a[i] > m2 {
            m2 = a[i]
		}
	}
    return m1 * m2
}