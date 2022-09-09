package main

import "fmt"

func main() {
	var n , count int
	var result []int 
	fmt.Scan(&n)
	points := input(n)

	for len(points) > 0 {
		r_m := minima(points)
		points_copy := make([][]int , len(points))
		copy(points_copy , points)
		for 
	}
}

func minima(points [][]int) int {
	mini := points[0][1]
	for _ , i := range points {
		if mini > i[1] {
			mini = i[1]
		}
	}
	return mini
}

func input(n int) [][]int {
	points := make([][]int , n)
	for i := 0 ; i < n ; i++ {
		points[i] = make([]int , 2)
		for j := 0 ; j < 2 ; j++ {
			fmt.Scan(&points[i][j]) 
		}
	}
	return points
}
