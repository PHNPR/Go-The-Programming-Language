package main

import (
	"fmt"
	"sort"
)

func main() {
	var len int
	var capacity float32 
	fmt.Scan(&len)
	fmt.Scan(&capacity)

	val_cap := input(len)

	sort.Slice(val_cap , func(i int ,j int) bool {
			return val_cap[i][0]/val_cap[i][1] > val_cap[j][0]/val_cap[j][1] 
	})

	fmt.Println(solver(val_cap , capacity))
}

func input(len int) [][]float32 {
	val_cap := make([][]float32 , len)
	for i := 0 ; i < len ; i++ {
		val_cap[i] = make([]float32 , 2)
		for j := 0 ; j < 2 ; j++ {
			fmt.Scan(&val_cap[i][j]) 
		}
	}
	return val_cap
}

func solver(val_cap [][]float32 , capacity float32) float32 {
	var value float32
	for _ , i := range val_cap {
		if capacity >=  i[1] {
			capacity -= i[1]
			value += i[0]
		} else {
			value += (i[0] / i[1]) * capacity
			return value
		}
	}
	return value
} 